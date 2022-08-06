package gchain

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXA1(x *string, arr ...string) {
	*x = *x + "t1" + strings.Join(arr, "")
}

func testXA2(x *string, arr ...string) {
	*x = *x + "t2" + strings.Join(arr, "")
}

func testXA3(x *string, arr ...string) {
	*x = *x + "t3" + strings.Join(arr, "")
}

func TestChainXANewChainXA(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXA[*string, string]()
		e := &ChainXA[*string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXA(testXA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test creating multiple method chain
		c := NewChainXA(testXA1, testXA2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA2))
	}
}

func TestChainXAAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXA[*string, string]()
		c.Append(testXA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXA(testXA1)
		c.Append(testXA2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXA(testXA1, testXA2)
		c.Append(testXA3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA3))
	}
}

func TestChainXAInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXA[*string, string]()
		c.Insert(testXA1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test inserting a method at -1
		c := NewChainXA(testXA1, testXA2)
		c.Insert(testXA3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA2))
	}
	{
		// test inserting a method at 0
		c := NewChainXA(testXA1, testXA2)
		c.Insert(testXA3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA2))
	}
	{
		// test inserting a method at 1
		c := NewChainXA(testXA1, testXA2)
		c.Insert(testXA3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA2))
	}
	{
		// test inserting a method at 2
		c := NewChainXA(testXA1, testXA2)
		c.Insert(testXA3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA3))
	}
	{
		// test inserting a method at 3
		c := NewChainXA(testXA1, testXA2)
		c.Insert(testXA3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA3))
	}
}

func TestChainXAExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXA[*string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXA[*string, string]()
		c.Extend(testXA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test extending chain with empty list
		c := NewChainXA(testXA1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test extending chain with other methods
		c := NewChainXA(testXA1)
		c.Extend(testXA2, testXA3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXA3))
	}
}

func TestChainXAJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXA[*string, string]()
		c2 := NewChainXA[*string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXA(testXA1)
		c2 := NewChainXA[*string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXA[*string, string]()
		c2 := NewChainXA(testXA1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXA1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXA(testXA1)
		c2 := NewChainXA(testXA2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXA2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXA(testXA1)
		c2 := NewChainXA(testXA2, testXA3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXA1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXA2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXA3))
	}
}

func TestChainXALen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXA[*string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXA(testXA1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXA(testXA1, testXA2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXAChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXA[*string, string]()
		var str string
		strPtr := &str
		c.Chain(strPtr)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain
		c := NewChainXA(testXA1)
		var str string
		strPtr := &str
		c.Chain(strPtr)
		assert.Equal(t, "t1", str)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXA(testXA1, testXA2, testXA3)
		var str string
		strPtr := &str
		c.Chain(strPtr)
		assert.Equal(t, "t1t2t3", str)
	}
	{
		// test executing a empty chain with extra args
		c := NewChainXA[*string, string]()
		var str string
		strPtr := &str
		arr := []string{"a", "b"}
		c.Chain(strPtr, arr...)
		assert.Equal(t, "", str)
	}
	{
		// test executing a single method  chain with extra args
		c := NewChainXA(testXA1)
		var str string
		strPtr := &str
		arr := []string{"a", "b"}
		c.Chain(strPtr, arr...)
		assert.Equal(t, "t1ab", str)
	}
	{
		// test executing a multiple method  chain with extra args
		c := NewChainXA(testXA1, testXA2, testXA3)
		var str string
		strPtr := &str
		arr := []string{"a", "b"}
		c.Chain(strPtr, arr...)
		assert.Equal(t, "t1abt2abt3ab", str)
	}
}
