package gchain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testXYZ1(x *string, y *string, z *string) {
	*x = *x + "t1"
	*y = *y + "e1"
	*z = *z + "s1"
}

func testXYZ2(x *string, y *string, z *string) {
	*x = *x + "t2"
	*y = *y + "e2"
	*z = *z + "s2"
}

func testXYZ3(x *string, y *string, z *string) {
	*x = *x + "t3"
	*y = *y + "e3"
	*z = *z + "s3"
}

func TestChainXYZNewChainXYZ(t *testing.T) {
	{
		// test creating empty chain
		c := NewChainXYZ[*string, *string, *string]()
		e := &ChainXYZ[*string, *string, *string]{}
		assert.Equal(t, e, c)
	}
	{
		// test creating 1 method chain
		c := NewChainXYZ(testXYZ1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test creating multiple method chain
		c := NewChainXYZ(testXYZ1, testXYZ2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ2))
	}
}

func TestChainXYZAppend(t *testing.T) {
	{
		// test appending a method to empty chain
		c := NewChainXYZ[*string, *string, *string]()
		c.Append(testXYZ1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test appending a method to a single method chain
		c := NewChainXYZ(testXYZ1)
		c.Append(testXYZ2)
		assert.Equal(t, 2, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ2))
	}
	{
		// test appending a method to a two method chain
		c := NewChainXYZ(testXYZ1, testXYZ2)
		c.Append(testXYZ3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ3))
	}
}

func TestChainXYZInsert(t *testing.T) {
	{
		// test inserting a method into empty chain
		c := NewChainXYZ[*string, *string, *string]()
		c.Insert(testXYZ1, 0)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test inserting a method at -1
		c := NewChainXYZ(testXYZ1, testXYZ2)
		c.Insert(testXYZ3, -1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ2))
	}
	{
		// test inserting a method at 0
		c := NewChainXYZ(testXYZ1, testXYZ2)
		c.Insert(testXYZ3, 0)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ3))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ2))
	}
	{
		// test inserting a method at 1
		c := NewChainXYZ(testXYZ1, testXYZ2)
		c.Insert(testXYZ3, 1)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ3))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ2))
	}
	{
		// test inserting a method at 2
		c := NewChainXYZ(testXYZ1, testXYZ2)
		c.Insert(testXYZ3, 2)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ3))
	}
	{
		// test inserting a method at 3
		c := NewChainXYZ(testXYZ1, testXYZ2)
		c.Insert(testXYZ3, 3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ3))
	}
}

func TestChainXYZExtend(t *testing.T) {
	{
		// test extending empty chain with empty list
		c := NewChainXYZ[*string, *string, *string]()
		c.Extend()
		assert.Equal(t, 0, len(c.fs))
	}
	{
		// test extending empty chain with another method
		c := NewChainXYZ[*string, *string, *string]()
		c.Extend(testXYZ1)
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test extending chain with empty list
		c := NewChainXYZ(testXYZ1)
		c.Extend()
		assert.Equal(t, 1, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test extending chain with other methods
		c := NewChainXYZ(testXYZ1)
		c.Extend(testXYZ2, testXYZ3)
		assert.Equal(t, 3, len(c.fs))
		assert.Equal(t, reflect.ValueOf(c.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c.fs[1]), reflect.ValueOf(testXYZ2))
		assert.Equal(t, reflect.ValueOf(c.fs[2]), reflect.ValueOf(testXYZ3))
	}
}

func TestChainXYZJoin(t *testing.T) {
	{
		// test joining empty chain and empty chain
		c1 := NewChainXYZ[*string, *string, *string]()
		c2 := NewChainXYZ[*string, *string, *string]()
		c1.Join(c2)
		assert.Equal(t, 0, len(c1.fs))
	}
	{
		// test joining chain and empty chain
		c1 := NewChainXYZ(testXYZ1)
		c2 := NewChainXYZ[*string, *string, *string]()
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test joining empty chain and another chain
		c1 := NewChainXYZ[*string, *string, *string]()
		c2 := NewChainXYZ(testXYZ1)
		c1.Join(c2)
		assert.Equal(t, 1, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZ1))
	}
	{
		// test joining chain and another chain
		c1 := NewChainXYZ(testXYZ1)
		c2 := NewChainXYZ(testXYZ2)
		c1.Join(c2)
		assert.Equal(t, 2, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZ2))
	}
	{
		// test joining chain and another chain with multiple methods
		c1 := NewChainXYZ(testXYZ1)
		c2 := NewChainXYZ(testXYZ2, testXYZ3)
		c1.Join(c2)
		assert.Equal(t, 3, len(c1.fs))
		assert.Equal(t, reflect.ValueOf(c1.fs[0]), reflect.ValueOf(testXYZ1))
		assert.Equal(t, reflect.ValueOf(c1.fs[1]), reflect.ValueOf(testXYZ2))
		assert.Equal(t, reflect.ValueOf(c1.fs[2]), reflect.ValueOf(testXYZ3))
	}
}

func TestChainXYZLen(t *testing.T) {
	{
		// test getting length of empty chain
		c := NewChainXYZ[*string, *string, *string]()
		assert.Equal(t, 0, c.Len())
	}
	{
		// test getting length single method chain
		c := NewChainXYZ(testXYZ1)
		assert.Equal(t, 1, c.Len())
	}
	{
		// test getting length of multiple method chain
		c := NewChainXYZ(testXYZ1, testXYZ2)
		assert.Equal(t, 2, c.Len())
	}
}

func TestChainXYZChain(t *testing.T) {
	{
		// test executing a empty chain
		c := NewChainXYZ[*string, *string, *string]()
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		c.Chain(str1Ptr, str2Ptr, str3Ptr)
		assert.Equal(t, "", str1)
		assert.Equal(t, "", str2)
		assert.Equal(t, "", str3)
	}
	{
		// test executing a single method  chain
		c := NewChainXYZ(testXYZ1)
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		c.Chain(str1Ptr, str2Ptr, str3Ptr)
		assert.Equal(t, "t1", str1)
		assert.Equal(t, "e1", str2)
		assert.Equal(t, "s1", str3)
	}
	{
		// test executing a multiple method  chain
		c := NewChainXYZ(testXYZ1, testXYZ2, testXYZ3)
		var str1, str2, str3 string
		str1Ptr := &str1
		str2Ptr := &str2
		str3Ptr := &str3
		c.Chain(str1Ptr, str2Ptr, str3Ptr)
		assert.Equal(t, "t1t2t3", str1)
		assert.Equal(t, "e1e2e3", str2)
		assert.Equal(t, "s1s2s3", str3)
	}
}
