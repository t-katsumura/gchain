package gchain

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYZA1(x *string, y *string, z *string, arr ...string) {
	*x = *x + "t1" + strings.Join(arr, "")
	*y = *y + "e1" + strings.Join(arr, "")
	*z = *z + "s1" + strings.Join(arr, "")
}

func testXYZA2(x *string, y *string, z *string, arr ...string) {
	*x = *x + "t2" + strings.Join(arr, "")
	*y = *y + "e2" + strings.Join(arr, "")
	*z = *z + "s2" + strings.Join(arr, "")
}

func testXYZA3(x *string, y *string, z *string, arr ...string) {
	*x = *x + "t3" + strings.Join(arr, "")
	*y = *y + "e3" + strings.Join(arr, "")
	*z = *z + "s3" + strings.Join(arr, "")
}

func TestChainXYZANewChainXYZA(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYZA[*string, *string, *string, string]()
		e := &ChainXYZA[*string, *string, *string, string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYZA(testXYZA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYZA(testXYZA1, testXYZA2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA2))
	}
}

func TestChainXYZAAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYZA[*string, *string, *string, string]()
		c.Append(testXYZA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYZA(testXYZA1)
		c.Append(testXYZA2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYZA(testXYZA1, testXYZA2)
		c.Append(testXYZA3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA3))
	}
}

func TestChainXYZAInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYZA[*string, *string, *string, string]()
		c.Insert(testXYZA1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYZA(testXYZA1, testXYZA2)
		c.Insert(testXYZA3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYZA(testXYZA1, testXYZA2)
		c.Insert(testXYZA3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYZA(testXYZA1, testXYZA2)
		c.Insert(testXYZA3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYZA(testXYZA1, testXYZA2)
		c.Insert(testXYZA3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYZA(testXYZA1, testXYZA2)
		c.Insert(testXYZA3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA3))
	}
}

func TestChainXYZAExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYZA[*string, *string, *string, string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYZA[*string, *string, *string, string]()
		c.Extend(testXYZA1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYZA(testXYZA1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYZA(testXYZA1)
		c.Extend(testXYZA2, testXYZA3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZA2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZA3))
	}
}

func TestChainXYZAJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYZA[*string, *string, *string, string]()
		c2 := NewChainXYZA[*string, *string, *string, string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYZA(testXYZA1)
		c2 := NewChainXYZA[*string, *string, *string, string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYZA[*string, *string, *string, string]()
		c2 := NewChainXYZA(testXYZA1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZA1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYZA(testXYZA1)
		c2 := NewChainXYZA(testXYZA2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZA2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYZA(testXYZA1)
		c2 := NewChainXYZA(testXYZA2, testXYZA3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZA1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZA2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYZA3))
	}
}

func TestChainXYZALen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYZA[*string, *string, *string, string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYZA(testXYZA1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYZA(testXYZA1, testXYZA2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYZAChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYZA[*string, *string, *string, string]()
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		arr := []string{}
		c.Chain(str1Ptr, str2Ptr, str3Ptr, arr...)
		assert.Equal(t, "", str1)
		assert.Equal(t, "", str2)
		assert.Equal(t, "", str3)
	}
	{
		// test executing a single method  chain
		c := NewChainXYZA(testXYZA1)
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		arr := []string{}
		c.Chain(str1Ptr, str2Ptr, str3Ptr, arr...)
		assert.Equal(t, "t1", str1)
		assert.Equal(t, "e1", str2)
		assert.Equal(t, "s1", str3)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYZA(testXYZA1, testXYZA2, testXYZA3)
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		arr := []string{}
		c.Chain(str1Ptr, str2Ptr, str3Ptr, arr...)
		assert.Equal(t, "t1t2t3", str1)
		assert.Equal(t, "e1e2e3", str2)
		assert.Equal(t, "s1s2s3", str3)
	}
	{
		// test executing a empty chain with extra args
		c := NewChainXYZA[*string, *string, *string, string]()
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, str3Ptr, arr...)
		assert.Equal(t, "", str1)
		assert.Equal(t, "", str2)
		assert.Equal(t, "", str3)
	}
	{
		// test executing a single method  chain with extra args
		c := NewChainXYZA(testXYZA1)
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, str3Ptr, arr...)
		assert.Equal(t, "t1ab", str1)
		assert.Equal(t, "e1ab", str2)
		assert.Equal(t, "s1ab", str3)
	}
	{
		// test executing a multiple method  chain with extra args
		c := NewChainXYZA(testXYZA1, testXYZA2, testXYZA3)
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		arr := []string{"a", "b"}
		c.Chain(str1Ptr, str2Ptr, str3Ptr, arr...)
		assert.Equal(t, "t1abt2abt3ab", str1)
		assert.Equal(t, "e1abe2abe3ab", str2)
		assert.Equal(t, "s1abs2abs3ab", str3)
	}
}
