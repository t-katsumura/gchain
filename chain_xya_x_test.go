package gchain

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYAtoX1(x string, y string, arr ...string) string {
	return x + "t1" + y + strings.Join(arr, "")
}

func testXYAtoX2(x string, y string, arr ...string) string {
	return x + "t2" + y + strings.Join(arr, "")
}

func testXYAtoX3(x string, y string, arr ...string) string {
	return x + "t3" + y + strings.Join(arr, "")
}

func TestChainXYAtoXNewChainXYAtoX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYAtoX[string, string, string]()
		e := &ChainXYAtoX[string, string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYAtoX(testXYAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX2))
	}
}

func TestChainXYAtoXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYAtoX[string, string, string]()
		c.Append(testXYAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYAtoX(testXYAtoX1)
		c.Append(testXYAtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		c.Append(testXYAtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX3))
	}
}

func TestChainXYAtoXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYAtoX[string, string, string]()
		c.Insert(testXYAtoX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		c.Insert(testXYAtoX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		c.Insert(testXYAtoX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		c.Insert(testXYAtoX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		c.Insert(testXYAtoX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		c.Insert(testXYAtoX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX3))
	}
}

func TestChainXYAtoXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYAtoX[string, string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYAtoX[string, string, string]()
		c.Extend(testXYAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYAtoX(testXYAtoX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYAtoX(testXYAtoX1)
		c.Extend(testXYAtoX2, testXYAtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYAtoX3))
	}
}

func TestChainXYAtoXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYAtoX[string, string, string]()
		c2 := NewChainXYAtoX[string, string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYAtoX(testXYAtoX1)
		c2 := NewChainXYAtoX[string, string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYAtoX[string, string, string]()
		c2 := NewChainXYAtoX(testXYAtoX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYAtoX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYAtoX(testXYAtoX1)
		c2 := NewChainXYAtoX(testXYAtoX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYAtoX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYAtoX(testXYAtoX1)
		c2 := NewChainXYAtoX(testXYAtoX2, testXYAtoX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYAtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYAtoX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYAtoX3))
	}
}

func TestChainXYAtoXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYAtoX[string, string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYAtoX(testXYAtoX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYAtoXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYAtoX[string, string, string]()
		arr := []string{}
		str := c.Chain("str:", "-", arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXYAtoX(testXYAtoX1)
		arr := []string{}
		str := c.Chain("str:", "-", arr...)
		assert.Equal(t, "str:t1-", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2, testXYAtoX3)
		arr := []string{}
		str := c.Chain("str:", "-", arr...)
		assert.Equal(t, "str:t1-t2-t3-", str)
	}
	{
		// test executing a empty chain with extra args
		c := NewChainXYAtoX[string, string, string]()
		arr := []string{"a", "b"}
		str := c.Chain("str:", "-", arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain with extra args
		c := NewChainXYAtoX(testXYAtoX1)
		arr := []string{"a", "b"}
		str := c.Chain("str:", "-", arr...)
		assert.Equal(t, "str:t1-ab", str)
	}
	{
		// test executing a multiple method  chain with extra args
		c := NewChainXYAtoX(testXYAtoX1, testXYAtoX2, testXYAtoX3)
		arr := []string{"a", "b"}
		str := c.Chain("str:", "-", arr...)
		assert.Equal(t, "str:t1-abt2-abt3-ab", str)
	}
}
