package gchain

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYA1(x *string, y *string, arr ...string) {
	*x = *x + "t1" + strings.Join(arr, "")
	*y = *y + "e1" + strings.Join(arr, "")
}

func testXYA2(x *string, y *string, arr ...string) {
	*x = *x + "t2" + strings.Join(arr, "")
	*y = *y + "e2" + strings.Join(arr, "")
}

func testXYA3(x *string, y *string, arr ...string) {
	*x = *x + "t3" + strings.Join(arr, "")
	*y = *y + "e3" + strings.Join(arr, "")
}

func TestChainXYANewChainXYA(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYA[*string, *string, string]()
		e := &ChainXYA[*string, *string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYA(testXYA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYA(testXYA1, testXYA2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA2))
	}
}

func TestChainXYAAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYA[*string, *string, string]()
		c.Append(testXYA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYA(testXYA1)
		c.Append(testXYA2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYA(testXYA1, testXYA2)
		c.Append(testXYA3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA3))
	}
}

func TestChainXYAInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYA[*string, *string, string]()
		c.Insert(testXYA1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYA(testXYA1, testXYA2)
		c.Insert(testXYA3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYA(testXYA1, testXYA2)
		c.Insert(testXYA3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYA(testXYA1, testXYA2)
		c.Insert(testXYA3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYA(testXYA1, testXYA2)
		c.Insert(testXYA3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYA(testXYA1, testXYA2)
		c.Insert(testXYA3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA3))
	}
}

func TestChainXYAExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYA[*string, *string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYA[*string, *string, string]()
		c.Extend(testXYA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYA(testXYA1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYA(testXYA1)
		c.Extend(testXYA2, testXYA3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYA3))
	}
}

func TestChainXYAJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYA[*string, *string, string]()
		c2 := NewChainXYA[*string, *string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYA(testXYA1)
		c2 := NewChainXYA[*string, *string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYA[*string, *string, string]()
		c2 := NewChainXYA(testXYA1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYA1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYA(testXYA1)
		c2 := NewChainXYA(testXYA2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYA2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYA(testXYA1)
		c2 := NewChainXYA(testXYA2, testXYA3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYA1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYA2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYA3))
	}
}

func TestChainXYALen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYA[*string, *string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYA(testXYA1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYA(testXYA1, testXYA2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYAChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYA[*string, *string, string]()
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, arr...)
		assert.Equal(t, "", str1)
		assert.Equal(t, "", str2)
	}
	{
		// test executing a single method  chain
		c := NewChainXYA(testXYA1)
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		arr := []string{}
		c.Chain(str1Ptr, str2Ptr, arr...)
		assert.Equal(t, "t1", str1)
		assert.Equal(t, "e1", str2)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYA(testXYA1, testXYA2, testXYA3)
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		arr := []string{}
		c.Chain(str1Ptr, str2Ptr, arr...)
		assert.Equal(t, "t1t2t3", str1)
		assert.Equal(t, "e1e2e3", str2)
	}
	{
		// test executing a empty chain with extra args
		c := NewChainXYA[*string, *string, string]()
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, arr...)
		assert.Equal(t, "", str1)
		assert.Equal(t, "", str2)
	}
	{
		// test executing a single method  chain with extra args
		c := NewChainXYA(testXYA1)
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, arr...)
		assert.Equal(t, "t1ab", str1)
		assert.Equal(t, "e1ab", str2)
	}
	{
		// test executing a multiple method  chain with extra args
		c := NewChainXYA(testXYA1, testXYA2, testXYA3)
		var str1, str2 string
		str1Ptr := &str1
		str2Ptr := &str2
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, arr...)
		assert.Equal(t, "t1abt2abt3ab", str1)
		assert.Equal(t, "e1abe2abe3ab", str2)
	}
}
