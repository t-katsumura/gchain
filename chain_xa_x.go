package gchain

type ChainXAtoX[P any, S any] struct {
	fs []func(x P, arr ...S) P
}

func NewChainXAtoX[P any, S any](fs ...func(P, ...S) P) *ChainXAtoX[P, S] {
	c := &ChainXAtoX[P, S]{
		fs: fs,
	}
	return c
}

func (c *ChainXAtoX[P, S]) Append(f func(P, ...S) P) *ChainXAtoX[P, S] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXAtoX[P, S]) Insert(f func(P, ...S) P, i int) *ChainXAtoX[P, S] {
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

func (c *ChainXAtoX[P, S]) Extend(fs ...func(P, ...S) P) *ChainXAtoX[P, S] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXAtoX[P, S]) Join(o *ChainXAtoX[P, S]) *ChainXAtoX[P, S] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXAtoX[P, S]) Len() int {
	return len(c.fs)
}

func (c *ChainXAtoX[P, S]) Chain(x P, a ...S) P {
	if len(c.fs) == 0 {
		var zero P
		return zero
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret, a...)
	}
	return ret
}
