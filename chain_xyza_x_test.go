package gchain

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYZAtoX1(x string, y string, z string, arr ...string) string {
	return x + "t1" + y + z + strings.Join(arr, "")
}

func testXYZAtoX2(x string, y string, z string, arr ...string) string {
	return x + "t2" + y + z + strings.Join(arr, "")
}

func testXYZAtoX3(x string, y string, z string, arr ...string) string {
	return x + "t3" + y + z + strings.Join(arr, "")
}

func TestChainXYZAtoXNewChainXYZAtoX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYZAtoX[string, string, string, string]()
		e := &ChainXYZAtoX[string, string, string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYZAtoX(testXYZAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX2))
	}
}

func TestChainXYZAtoXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYZAtoX[string, string, string, string]()
		c.Append(testXYZAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYZAtoX(testXYZAtoX1)
		c.Append(testXYZAtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		c.Append(testXYZAtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX3))
	}
}

func TestChainXYZAtoXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYZAtoX[string, string, string, string]()
		c.Insert(testXYZAtoX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		c.Insert(testXYZAtoX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		c.Insert(testXYZAtoX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		c.Insert(testXYZAtoX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		c.Insert(testXYZAtoX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		c.Insert(testXYZAtoX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX3))
	}
}

func TestChainXYZAtoXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYZAtoX[string, string, string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYZAtoX[string, string, string, string]()
		c.Extend(testXYZAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYZAtoX(testXYZAtoX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYZAtoX(testXYZAtoX1)
		c.Extend(testXYZAtoX2, testXYZAtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZAtoX3))
	}
}

func TestChainXYZAtoXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYZAtoX[string, string, string, string]()
		c2 := NewChainXYZAtoX[string, string, string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYZAtoX(testXYZAtoX1)
		c2 := NewChainXYZAtoX[string, string, string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYZAtoX[string, string, string, string]()
		c2 := NewChainXYZAtoX(testXYZAtoX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZAtoX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYZAtoX(testXYZAtoX1)
		c2 := NewChainXYZAtoX(testXYZAtoX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZAtoX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYZAtoX(testXYZAtoX1)
		c2 := NewChainXYZAtoX(testXYZAtoX2, testXYZAtoX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZAtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZAtoX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYZAtoX3))
	}
}

func TestChainXYZAtoXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYZAtoX[string, string, string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYZAtoX(testXYZAtoX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYZAtoXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYZAtoX[string, string, string, string]()
		arr := []string{}
		str := c.Chain("str:", "-", "*", arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXYZAtoX(testXYZAtoX1)
		arr := []string{}
		str := c.Chain("str:", "-", "*", arr...)
		assert.Equal(t, "str:t1-*", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2, testXYZAtoX3)
		arr := []string{}
		str := c.Chain("str:", "-", "*", arr...)
		assert.Equal(t, "str:t1-*t2-*t3-*", str)
	}
	{
		// test executing a empty chain with extra args
		c := NewChainXYZAtoX[string, string, string, string]()
		arr := []string{"a", "b"}
		str := c.Chain("str:", "-", "*", arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain with extra args
		c := NewChainXYZAtoX(testXYZAtoX1)
		arr := []string{"a", "b"}
		str := c.Chain("str:", "-", "*", arr...)
		assert.Equal(t, "str:t1-*ab", str)
	}
	{
		// test executing a multiple method  chain with extra args
		c := NewChainXYZAtoX(testXYZAtoX1, testXYZAtoX2, testXYZAtoX3)
		arr := []string{"a", "b"}
		str := c.Chain("str:", "-", "*", arr...)
		assert.Equal(t, "str:t1-*abt2-*abt3-*ab", str)
	}
}
