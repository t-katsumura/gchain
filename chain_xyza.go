package gchain

type ChainXYZA[P any, Q any, R any, S any] struct {
	fs []func(x P, y Q, z R, arr ...S)
}

func NewChainXYZA[P any, Q any, R any, S any](fs ...func(P, Q, R, ...S)) *ChainXYZA[P, Q, R, S] {
	c := &ChainXYZA[P, Q, R, S]{
		fs: fs,
	}
	return c
}

func (c *ChainXYZA[P, Q, R, S]) Append(f func(P, Q, R, ...S)) *ChainXYZA[P, Q, R, S] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYZA[P, Q, R, S]) Insert(f func(P, Q, R, ...S), i int) *ChainXYZA[P, Q, R, S] {
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

func (c *ChainXYZA[P, Q, R, S]) Extend(fs ...func(P, Q, R, ...S)) *ChainXYZA[P, Q, R, S] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYZA[P, Q, R, S]) Join(o *ChainXYZA[P, Q, R, S]) *ChainXYZA[P, Q, R, S] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYZA[P, Q, R, S]) Len() int {
	return len(c.fs)
}

func (c *ChainXYZA[P, Q, R, S]) Chain(x P, y Q, z R, a ...S) {
	for _, f := range c.fs {
		f(x, y, z, a...)
	}
}
