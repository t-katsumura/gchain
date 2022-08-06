package gchain

type ChainXYA[P any, Q any, S any] struct {
	fs []func(x P, y Q, arr ...S)
}

func NewChainXYA[P any, Q any, S any](fs ...func(P, Q, ...S)) *ChainXYA[P, Q, S] {
	c := &ChainXYA[P, Q, S]{
		fs: fs,
	}
	return c
}

func (c *ChainXYA[P, Q, S]) Append(f func(P, Q, ...S)) *ChainXYA[P, Q, S] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYA[P, Q, S]) Insert(f func(P, Q, ...S), i int) *ChainXYA[P, Q, S] {
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

func (c *ChainXYA[P, Q, S]) Extend(fs ...func(P, Q, ...S)) *ChainXYA[P, Q, S] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYA[P, Q, S]) Join(o *ChainXYA[P, Q, S]) *ChainXYA[P, Q, S] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYA[P, Q, S]) Len() int {
	return len(c.fs)
}

func (c *ChainXYA[P, Q, S]) Chain(x P, y Q, a ...S) {
	for _, f := range c.fs {
		f(x, y, a...)
	}
}
