package pigment

import "sync"

var styleRegistry sync.Map

func (c *composer) Compose(styles ...style) style {
	composed := style(c)
	for _, s := range styles {
		if cmp, ok := s.(*composer); ok && cmp.node != nil {
			composed = createComposer(
				composed.(*composer),
				cmp.node.style.code,
				cmp.node.style.reset,
			)
		}
	}
	return composed
}

func Compose(styles ...style) style {
	if len(styles) == 0 {
		return defaultComposer
	}

	composed := styles[0]
	for _, s := range styles[1:] {
		composed = composed.Compose(s)
	}
	return composed
}

func Register(name string, styles ...style) {
	styleRegistry.Store(name, Compose(styles...))
}

func (c *composer) WithCustom(name string) style {
	value, ok := styleRegistry.Load(name)
	if !ok {
		return &composer{}
	}

	cmp, ok := value.(*composer)
	if !ok {
		return &composer{}
	}

	return cmp
}

func WithCustom(name string) style {
	return defaultComposer.WithCustom(name)
}
