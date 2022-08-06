package gchain

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXAtoX1(x string, arr ...string) string {
	return x + "t1" + strings.Join(arr, "")
}

func testXAtoX2(x string, arr ...string) string {
	return x + "t2" + strings.Join(arr, "")
}

func testXAtoX3(x string, arr ...string) string {
	return x + "t3" + strings.Join(arr, "")
}

func TestChainXAtoXNewChainXAtoX(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXAtoX[string, string]()
		e := &ChainXAtoX[string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXAtoX(testXAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test creating multiple method chain
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX2))
	}
}

func TestChainXAtoXAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXAtoX[string, string]()
		c.Append(testXAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXAtoX(testXAtoX1)
		c.Append(testXAtoX2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		c.Append(testXAtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX3))
	}
}

func TestChainXAtoXInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXAtoX[string, string]()
		c.Insert(testXAtoX1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test inserting a method at -1
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		c.Insert(testXAtoX3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX2))
	}
	{
		// test inserting a method at 0
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		c.Insert(testXAtoX3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX2))
	}
	{
		// test inserting a method at 1
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		c.Insert(testXAtoX3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX2))
	}
	{
		// test inserting a method at 2
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		c.Insert(testXAtoX3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX3))
	}
	{
		// test inserting a method at 3
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		c.Insert(testXAtoX3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX3))
	}
}

func TestChainXAtoXExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXAtoX[string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXAtoX[string, string]()
		c.Extend(testXAtoX1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test extending chain with empty list
		c := NewChainXAtoX(testXAtoX1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test extending chain with other methods
		c := NewChainXAtoX(testXAtoX1)
		c.Extend(testXAtoX2, testXAtoX3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXAtoX2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXAtoX3))
	}
}

func TestChainXAtoXJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXAtoX[string, string]()
		c2 := NewChainXAtoX[string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXAtoX(testXAtoX1)
		c2 := NewChainXAtoX[string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXAtoX[string, string]()
		c2 := NewChainXAtoX(testXAtoX1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXAtoX1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXAtoX(testXAtoX1)
		c2 := NewChainXAtoX(testXAtoX2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXAtoX2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXAtoX(testXAtoX1)
		c2 := NewChainXAtoX(testXAtoX2, testXAtoX3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXAtoX1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXAtoX2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXAtoX3))
	}
}

func TestChainXAtoXLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXAtoX[string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXAtoX(testXAtoX1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXAtoX(testXAtoX1, testXAtoX2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXAtoXChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXAtoX[string, string]()
		arr := []string{}
		str := c.Chain("str:", arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXAtoX(testXAtoX1)
		arr := []string{}
		str := c.Chain("str:", arr...)
		assert.Equal(t, "str:t1", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXAtoX(testXAtoX1, testXAtoX2, testXAtoX3)
		arr := []string{}
		str := c.Chain("str:", arr...)
		assert.Equal(t, "str:t1t2t3", str)
	}
	{
		// test executing a empty chain with extra args
		c := NewChainXAtoX[string, string]()
		arr := []string{"a", "b"}
		str := c.Chain("str:", arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain with extra args
		c := NewChainXAtoX(testXAtoX1)
		arr := []string{"a", "b"}
		str := c.Chain("str:", arr...)
		assert.Equal(t, "str:t1ab", str)
	}
	{
		// test executing a multiple method  chain with extra args
		c := NewChainXAtoX(testXAtoX1, testXAtoX2, testXAtoX3)
		arr := []string{"a", "b"}
		str := c.Chain("str:", arr...)
		assert.Equal(t, "str:t1abt2abt3ab", str)
	}
}
