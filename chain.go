package gchain

type SimpleChain[T any] struct {
	fs []func(f T) T
}

func NewSimpleChain[T any](fs ...func(T) T) *SimpleChain[T] {
	c := &SimpleChain[T]{
		fs: fs,
	}
	return c
}

func (c *SimpleChain[T]) Append(f func(T) T) *SimpleChain[T] {
	c.fs = append(c.fs, f)
	return c
}

func (c *SimpleChain[T]) Extend(fs ...func(T) T) *SimpleChain[T] {
	c.fs = append(c.fs, fs...)
	return c
}

func (c *SimpleChain[T]) Join(o *SimpleChain[T]) *SimpleChain[T] {
	c.fs = append(c.fs, o.fs...)
	return c
}

func (c *SimpleChain[T]) Len() int {
	return len(c.fs)
}

func (c *SimpleChain[T]) Chain(x T) T {

	if len(c.fs) == 0 {
		panic("gchain.SimpleChain has no method")
	}

	ret := x
	for _, f := range c.fs {
		ret = f(ret)
	}
	return ret
}
