package gchain

type ChainX[P any] struct {
	fs []func(x P)
}

func NewChainX[P any](fs ...func(P)) *ChainX[P] {
	c := &ChainX[P]{
		fs: fs,
	}
	return c
}

func (c *ChainX[P]) Append(f func(P)) *ChainX[P] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainX[P]) Insert(f func(P), i int) *ChainX[P] {
	if i >= len(c.fs) {
		c.fs = append(c.fs, f)
	} else {
		if i < 0 {
			i = 0
		}
		c.fs = append(c.fs[:i+1], c.fs[i:]...)
		c.fs[i] = f
	}
	return c
}

func (c *ChainX[P]) Extend(fs ...func(P)) *ChainX[P] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainX[P]) Join(o *ChainX[P]) *ChainX[P] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainX[P]) Len() int {
	return len(c.fs)
}

func (c *ChainX[P]) Chain(x P) {
	for _, f := range c.fs {
		f(x)
	}
}
