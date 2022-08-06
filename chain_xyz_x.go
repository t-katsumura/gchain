package gchain

type ChainXYZtoX[P any, Q any, R any] struct {
	fs []func(x P, y Q, z R) P
}

func NewChainXYZtoX[P any, Q any, R any](fs ...func(P, Q, R) P) *ChainXYZtoX[P, Q, R] {
	c := &ChainXYZtoX[P, Q, R]{
		fs: fs,
	}
	return c
}

func (c *ChainXYZtoX[P, Q, R]) Append(f func(P, Q, R) P) *ChainXYZtoX[P, Q, R] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYZtoX[P, Q, R]) Insert(f func(P, Q, R) P, i int) *ChainXYZtoX[P, Q, R] {
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

func (c *ChainXYZtoX[P, Q, R]) Extend(fs ...func(P, Q, R) P) *ChainXYZtoX[P, Q, R] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYZtoX[P, Q, R]) Join(o *ChainXYZtoX[P, Q, R]) *ChainXYZtoX[P, Q, R] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYZtoX[P, Q, R]) Len() int {
	return len(c.fs)
}

func (c *ChainXYZtoX[P, Q, R]) Chain(x P, y Q, z R) P {
	if len(c.fs) == 0 {
		var zero P
		return zero
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret, y, z)
	}
	return ret
}
