package gchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYZtoX1(x string, y string, z string) string {
	return x + "t1" + y + z
}

func testXYZtoX2(x string, y string, z string) string {
	return x + "t2" + y + z
}

func testXYZtoX3(x string, y string, z string) string {
	return x + "t3" + y + z
}

func TestChainXYZtoXNewChainXYZtoX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYZtoX[string, string, string]()
		e := &ChainXYZtoX[string, string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYZtoX(testXYZtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX2))
	}
}

func TestChainXYZtoXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYZtoX[string, string, string]()
		c.Append(testXYZtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYZtoX(testXYZtoX1)
		c.Append(testXYZtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		c.Append(testXYZtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX3))
	}
}

func TestChainXYZtoXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYZtoX[string, string, string]()
		c.Insert(testXYZtoX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		c.Insert(testXYZtoX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		c.Insert(testXYZtoX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		c.Insert(testXYZtoX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		c.Insert(testXYZtoX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		c.Insert(testXYZtoX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX3))
	}
}

func TestChainXYZtoXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYZtoX[string, string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYZtoX[string, string, string]()
		c.Extend(testXYZtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYZtoX(testXYZtoX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYZtoX(testXYZtoX1)
		c.Extend(testXYZtoX2, testXYZtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZtoX3))
	}
}

func TestChainXYZtoXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYZtoX[string, string, string]()
		c2 := NewChainXYZtoX[string, string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYZtoX(testXYZtoX1)
		c2 := NewChainXYZtoX[string, string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYZtoX[string, string, string]()
		c2 := NewChainXYZtoX(testXYZtoX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZtoX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYZtoX(testXYZtoX1)
		c2 := NewChainXYZtoX(testXYZtoX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZtoX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYZtoX(testXYZtoX1)
		c2 := NewChainXYZtoX(testXYZtoX2, testXYZtoX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZtoX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYZtoX3))
	}
}

func TestChainXYZtoXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYZtoX[string, string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYZtoX(testXYZtoX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYZtoXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYZtoX[string, string, string]()
		str := c.Chain("str:", "-", "*")
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXYZtoX(testXYZtoX1)
		str := c.Chain("str:", "-", "*")
		assert.Equal(t, "str:t1-*", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYZtoX(testXYZtoX1, testXYZtoX2, testXYZtoX3)
		str := c.Chain("str:", "-", "*")
		assert.Equal(t, "str:t1-*t2-*t3-*", str)
	}
}
