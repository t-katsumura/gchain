package gchain

type ChainXYZ[P any, Q any, R any] struct {
	fs []func(x P, y Q, z R)
}

func NewChainXYZ[P any, Q any, R any](fs ...func(P, Q, R)) *ChainXYZ[P, Q, R] {
	c := &ChainXYZ[P, Q, R]{
		fs: fs,
	}
	return c
}

func (c *ChainXYZ[P, Q, R]) Append(f func(P, Q, R)) *ChainXYZ[P, Q, R] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYZ[P, Q, R]) Insert(f func(P, Q, R), i int) *ChainXYZ[P, Q, R] {
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

func (c *ChainXYZ[P, Q, R]) Extend(fs ...func(P, Q, R)) *ChainXYZ[P, Q, R] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYZ[P, Q, R]) Join(o *ChainXYZ[P, Q, R]) *ChainXYZ[P, Q, R] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYZ[P, Q, R]) Len() int {
	return len(c.fs)
}

func (c *ChainXYZ[P, Q, R]) Chain(x P, y Q, z R) {
	for _, f := range c.fs {
		f(x, y, z)
	}
}
