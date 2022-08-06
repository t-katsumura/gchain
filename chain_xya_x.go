package gchain

type ChainXYAtoX[P any, Q any, S any] struct {
	fs []func(x P, y Q, arr ...S) P
}

func NewChainXYAtoX[P any, Q any, S any](fs ...func(P, Q, ...S) P) *ChainXYAtoX[P, Q, S] {
	c := &ChainXYAtoX[P, Q, S]{
		fs: fs,
	}
	return c
}

func (c *ChainXYAtoX[P, Q, S]) Append(f func(P, Q, ...S) P) *ChainXYAtoX[P, Q, S] {
	c.fs = append(c.fs, f)
	return c
}

func (c *ChainXYAtoX[P, Q, S]) Insert(f func(P, Q, ...S) P, i int) *ChainXYAtoX[P, Q, S] {
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

func (c *ChainXYAtoX[P, Q, S]) Extend(fs ...func(P, Q, ...S) P) *ChainXYAtoX[P, Q, S] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *ChainXYAtoX[P, Q, S]) Join(o *ChainXYAtoX[P, Q, S]) *ChainXYAtoX[P, Q, S] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *ChainXYAtoX[P, Q, S]) Len() int {
	return len(c.fs)
}

func (c *ChainXYAtoX[P, Q, S]) Chain(x P, y Q, a ...S) P {
	if len(c.fs) == 0 {
		var zero P
		return zero
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret, y, a...)
	}
	return ret
}
