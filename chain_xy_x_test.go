package gchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYtoX1(x string, y string) string {
	return x + "t1" + y
}

func testXYtoX2(x string, y string) string {
	return x + "t2" + y
}

func testXYtoX3(x string, y string) string {
	return x + "t3" + y
}

func TestChainXYtoXNewChainXYtoX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYtoX[string, string]()
		e := &ChainXYtoX[string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYtoX(testXYtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX2))
	}
}

func TestChainXYtoXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYtoX[string, string]()
		c.Append(testXYtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYtoX(testXYtoX1)
		c.Append(testXYtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		c.Append(testXYtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX3))
	}
}

func TestChainXYtoXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYtoX[string, string]()
		c.Insert(testXYtoX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		c.Insert(testXYtoX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		c.Insert(testXYtoX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		c.Insert(testXYtoX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		c.Insert(testXYtoX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		c.Insert(testXYtoX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX3))
	}
}

func TestChainXYtoXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYtoX[string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYtoX[string, string]()
		c.Extend(testXYtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYtoX(testXYtoX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYtoX(testXYtoX1)
		c.Extend(testXYtoX2, testXYtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYtoX3))
	}
}

func TestChainXYtoXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYtoX[string, string]()
		c2 := NewChainXYtoX[string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYtoX(testXYtoX1)
		c2 := NewChainXYtoX[string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYtoX[string, string]()
		c2 := NewChainXYtoX(testXYtoX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYtoX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYtoX(testXYtoX1)
		c2 := NewChainXYtoX(testXYtoX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYtoX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYtoX(testXYtoX1)
		c2 := NewChainXYtoX(testXYtoX2, testXYtoX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYtoX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYtoX3))
	}
}

func TestChainXYtoXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYtoX[string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYtoX(testXYtoX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYtoX(testXYtoX1, testXYtoX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYtoXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYtoX[string, string]()
		str := c.Chain("str:", "-")
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXYtoX(testXYtoX1)
		str := c.Chain("str:", "-")
		assert.Equal(t, "str:t1-", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYtoX(testXYtoX1, testXYtoX2, testXYtoX3)
		str := c.Chain("str:", "-")
		assert.Equal(t, "str:t1-t2-t3-", str)
	}
}
