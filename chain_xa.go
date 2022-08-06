package gchain

type ChainXA[P any, S any] struct {
	fs []func(x P, arr ...S)
}

func NewChainXA[P any, S any](fs ...func(P, ...S)) *ChainXA[P, S] {
	c := &ChainXA[P, S]{
		fs: fs,
	}
	return c
}

func (c *ChainXA[P, S]) Append(f func(P, ...S)) *ChainXA[P, S] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXA[P, S]) Insert(f func(P, ...S), i int) *ChainXA[P, S] {
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

func (c *ChainXA[P, S]) Extend(fs ...func(P, ...S)) *ChainXA[P, S] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXA[P, S]) Join(o *ChainXA[P, S]) *ChainXA[P, S] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXA[P, S]) Len() int {
	return len(c.fs)
}

func (c *ChainXA[P, S]) Chain(x P, a ...S) {
	for _, f := range c.fs {
		f(x, a...)
	}
}
