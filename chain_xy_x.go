package gchain

type ChainXYtoX[P any, Q any] struct {
	fs []func(x P, y Q) P
}

func NewChainXYtoX[P any, Q any](fs ...func(P, Q) P) *ChainXYtoX[P, Q] {
	c := &ChainXYtoX[P, Q]{
		fs: fs,
	}
	return c
}

func (c *ChainXYtoX[P, Q]) Append(f func(P, Q) P) *ChainXYtoX[P, Q] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYtoX[P, Q]) Insert(f func(P, Q) P, i int) *ChainXYtoX[P, Q] {
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

func (c *ChainXYtoX[P, Q]) Extend(fs ...func(P, Q) P) *ChainXYtoX[P, Q] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYtoX[P, Q]) Join(o *ChainXYtoX[P, Q]) *ChainXYtoX[P, Q] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYtoX[P, Q]) Len() int {
	return len(c.fs)
}

func (c *ChainXYtoX[P, Q]) Chain(x P, y Q) P {
	if len(c.fs) == 0 {
		var zero P
		return zero
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret, y)
	}
	return ret
}
