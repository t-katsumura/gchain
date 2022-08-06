package gchain

type ChainXY[P any, Q any] struct {
	fs []func(x P, y Q)
}

func NewChainXY[P any, Q any](fs ...func(P, Q)) *ChainXY[P, Q] {
	c := &ChainXY[P, Q]{
		fs: fs,
	}
	return c
}

func (c *ChainXY[P, Q]) Append(f func(P, Q)) *ChainXY[P, Q] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXY[P, Q]) Insert(f func(P, Q), i int) *ChainXY[P, Q] {
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

func (c *ChainXY[P, Q]) Extend(fs ...func(P, Q)) *ChainXY[P, Q] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXY[P, Q]) Join(o *ChainXY[P, Q]) *ChainXY[P, Q] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXY[P, Q]) Len() int {
	return len(c.fs)
}

func (c *ChainXY[P, Q]) Chain(x P, y Q) {
	for _, f := range c.fs {
		f(x, y)
	}
}
