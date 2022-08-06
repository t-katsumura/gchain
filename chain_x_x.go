package gchain

type ChainXtoX[P any] struct {
	fs []func(x P) P
}

func NewChainXtoX[P any](fs ...func(P) P) *ChainXtoX[P] {
	c := &ChainXtoX[P]{
		fs: fs,
	}
	return c
}

func (c *ChainXtoX[P]) Append(f func(P) P) *ChainXtoX[P] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXtoX[P]) Insert(f func(P) P, i int) *ChainXtoX[P] {
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

func (c *ChainXtoX[P]) Extend(fs ...func(P) P) *ChainXtoX[P] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXtoX[P]) Join(o *ChainXtoX[P]) *ChainXtoX[P] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXtoX[P]) Len() int {
	return len(c.fs)
}

func (c *ChainXtoX[P]) Chain(x P) P {
	if len(c.fs) == 0 {
		var zero P
		return zero
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret)
	}
	return ret
}
