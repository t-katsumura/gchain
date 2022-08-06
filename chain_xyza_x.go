package gchain

type ChainXYZAtoX[P any, Q any, R any, S any] struct {
	fs []func(x P, y Q, z R, arr ...S) P
}

func NewChainXYZAtoX[P any, Q any, R any, S any](fs ...func(P, Q, R, ...S) P) *ChainXYZAtoX[P, Q, R, S] {
	c := &ChainXYZAtoX[P, Q, R, S]{
		fs: fs,
	}
	return c
}

func (c *ChainXYZAtoX[P, Q, R, S]) Append(f func(P, Q, R, ...S) P) *ChainXYZAtoX[P, Q, R, S] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYZAtoX[P, Q, R, S]) Insert(f func(P, Q, R, ...S) P, i int) *ChainXYZAtoX[P, Q, R, S] {
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

func (c *ChainXYZAtoX[P, Q, R, S]) Extend(fs ...func(P, Q, R, ...S) P) *ChainXYZAtoX[P, Q, R, S] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYZAtoX[P, Q, R, S]) Join(o *ChainXYZAtoX[P, Q, R, S]) *ChainXYZAtoX[P, Q, R, S] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYZAtoX[P, Q, R, S]) Len() int {
	return len(c.fs)
}

func (c *ChainXYZAtoX[P, Q, R, S]) Chain(x P, y Q, z R, a ...S) P {
	if len(c.fs) == 0 {
		var zero P
		return zero
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret, y, z, a...)
	}
	return ret
}
