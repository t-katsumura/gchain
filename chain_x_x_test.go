package gchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXtoX1(x string) string {
	return x + "t1"
}

func testXtoX2(x string) string {
	return x + "t2"
}

func testXtoX3(x string) string {
	return x + "t3"
}

func TestChainXtoXNewChainXtoX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXtoX[string]()
		e := &ChainXtoX[string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXtoX(testXtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test creating multiple method chain
		c := NewChainXtoX(testXtoX1, testXtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX2))
	}
}

func TestChainXtoXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXtoX[string]()
		c.Append(testXtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXtoX(testXtoX1)
		c.Append(testXtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXtoX(testXtoX1, testXtoX2)
		c.Append(testXtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX3))
	}
}

func TestChainXtoXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXtoX[string]()
		c.Insert(testXtoX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test inserting a method at -1
		c := NewChainXtoX(testXtoX1, testXtoX2)
		c.Insert(testXtoX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX2))
	}
	{
		// test inserting a method at 0
		c := NewChainXtoX(testXtoX1, testXtoX2)
		c.Insert(testXtoX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX2))
	}
	{
		// test inserting a method at 1
		c := NewChainXtoX(testXtoX1, testXtoX2)
		c.Insert(testXtoX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX2))
	}
	{
		// test inserting a method at 2
		c := NewChainXtoX(testXtoX1, testXtoX2)
		c.Insert(testXtoX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX3))
	}
	{
		// test inserting a method at 3
		c := NewChainXtoX(testXtoX1, testXtoX2)
		c.Insert(testXtoX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX3))
	}
}

func TestChainXtoXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXtoX[string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXtoX[string]()
		c.Extend(testXtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test extending chain with empty list
		c := NewChainXtoX(testXtoX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test extending chain with other methods
		c := NewChainXtoX(testXtoX1)
		c.Extend(testXtoX2, testXtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXtoX3))
	}
}

func TestChainXtoXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXtoX[string]()
		c2 := NewChainXtoX[string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXtoX(testXtoX1)
		c2 := NewChainXtoX[string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXtoX[string]()
		c2 := NewChainXtoX(testXtoX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXtoX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXtoX(testXtoX1)
		c2 := NewChainXtoX(testXtoX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXtoX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXtoX(testXtoX1)
		c2 := NewChainXtoX(testXtoX2, testXtoX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXtoX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXtoX3))
	}
}

func TestChainXtoXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXtoX[string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXtoX(testXtoX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXtoX(testXtoX1, testXtoX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXtoXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXtoX[string]()
		str := c.Chain("str:")
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXtoX(testXtoX1)
		str := c.Chain("str:")
		assert.Equal(t, "str:t1", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXtoX(testXtoX1, testXtoX2, testXtoX3)
		str := c.Chain("str:")
		assert.Equal(t, "str:t1t2t3", str)
	}
}
