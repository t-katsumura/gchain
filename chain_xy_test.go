package gchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXY1(x *string, y *string) {
	*x = *x + "t1"
	*y = *y + "e1"
}

func testXY2(x *string, y *string) {
	*x = *x + "t2"
	*y = *y + "e2"
}

func testXY3(x *string, y *string) {
	*x = *x + "t3"
	*y = *y + "e3"
}

func TestChainXYNewChainXY(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXY[*string, *string]()
		e := &ChainXY[*string, *string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXY(testXY1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test creating multiple method chain
		c := NewChainXY(testXY1, testXY2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY2))
	}
}

func TestChainXYAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXY[*string, *string]()
		c.Append(testXY1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXY(testXY1)
		c.Append(testXY2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXY(testXY1, testXY2)
		c.Append(testXY3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY3))
	}
}

func TestChainXYInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXY[*string, *string]()
		c.Insert(testXY1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test inserting a method at -1
		c := NewChainXY(testXY1, testXY2)
		c.Insert(testXY3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY2))
	}
	{
		// test inserting a method at 0
		c := NewChainXY(testXY1, testXY2)
		c.Insert(testXY3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY2))
	}
	{
		// test inserting a method at 1
		c := NewChainXY(testXY1, testXY2)
		c.Insert(testXY3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY2))
	}
	{
		// test inserting a method at 2
		c := NewChainXY(testXY1, testXY2)
		c.Insert(testXY3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY3))
	}
	{
		// test inserting a method at 3
		c := NewChainXY(testXY1, testXY2)
		c.Insert(testXY3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY3))
	}
}

func TestChainXYExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXY[*string, *string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXY[*string, *string]()
		c.Extend(testXY1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test extending chain with empty list
		c := NewChainXY(testXY1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test extending chain with other methods
		c := NewChainXY(testXY1)
		c.Extend(testXY2, testXY3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXY2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXY3))
	}
}

func TestChainXYJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXY[*string, *string]()
		c2 := NewChainXY[*string, *string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXY(testXY1)
		c2 := NewChainXY[*string, *string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXY[*string, *string]()
		c2 := NewChainXY(testXY1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXY1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXY(testXY1)
		c2 := NewChainXY(testXY2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXY2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXY(testXY1)
		c2 := NewChainXY(testXY2, testXY3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXY1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXY2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXY3))
	}
}

func TestChainXYLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXY[*string, *string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXY(testXY1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXY(testXY1, testXY2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXY[*string, *string]()
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		c.Chain(str1Ptr, str2Ptr)
		assert.Equal(t, "", str1)
		assert.Equal(t, "", str2)
	}
	{
		// test executing a single method  chain
		c := NewChainXY(testXY1)
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		c.Chain(str1Ptr, str2Ptr)
		assert.Equal(t, "t1", str1)
		assert.Equal(t, "e1", str2)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXY(testXY1, testXY2, testXY3)
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		c.Chain(str1Ptr, str2Ptr)
		assert.Equal(t, "t1t2t3", str1)
		assert.Equal(t, "e1e2e3", str2)
	}
}
