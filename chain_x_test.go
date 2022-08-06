package gchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testX1(x *string) {
	*x = *x + "t1"
}

func testX2(x *string) {
	*x = *x + "t2"
}

func testX3(x *string) {
	*x = *x + "t3"
}

func TestChainXNewChainX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainX[*string]()
		e := &ChainX[*string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainX(testX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test creating multiple method chain
		c := NewChainX(testX1, testX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX2))
	}
}

func TestChainXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainX[*string]()
		c.Append(testX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainX(testX1)
		c.Append(testX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainX(testX1, testX2)
		c.Append(testX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX3))
	}
}

func TestChainXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainX[*string]()
		c.Insert(testX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test inserting a method at -1
		c := NewChainX(testX1, testX2)
		c.Insert(testX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX2))
	}
	{
		// test inserting a method at 0
		c := NewChainX(testX1, testX2)
		c.Insert(testX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX2))
	}
	{
		// test inserting a method at 1
		c := NewChainX(testX1, testX2)
		c.Insert(testX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX2))
	}
	{
		// test inserting a method at 2
		c := NewChainX(testX1, testX2)
		c.Insert(testX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX3))
	}
	{
		// test inserting a method at 3
		c := NewChainX(testX1, testX2)
		c.Insert(testX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX3))
	}
}

func TestChainXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainX[*string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainX[*string]()
		c.Extend(testX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test extending chain with empty list
		c := NewChainX(testX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test extending chain with other methods
		c := NewChainX(testX1)
		c.Extend(testX2, testX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testX3))
	}
}

func TestChainXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainX[*string]()
		c2 := NewChainX[*string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainX(testX1)
		c2 := NewChainX[*string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainX[*string]()
		c2 := NewChainX(testX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainX(testX1)
		c2 := NewChainX(testX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainX(testX1)
		c2 := NewChainX(testX2, testX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testX3))
	}
}

func TestChainXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainX[*string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainX(testX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainX(testX1, testX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainX[*string]()
		var str string
		strPtr := &str
		c.Chain(strPtr)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainX(testX1)
		var str string
		strPtr := &str
		c.Chain(strPtr)
		assert.Equal(t, "t1", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainX(testX1, testX2, testX3)
		var str string
		strPtr := &str
		c.Chain(strPtr)
		assert.Equal(t, "t1t2t3", str)
	}
}
