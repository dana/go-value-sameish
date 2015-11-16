package valuesameish

import (
	//	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicIntIntSame(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish(1, 1))
}
func TestBasicIntIntNotSame(t *testing.T) {
	assert := assert.New(t)
	assert.False(Sameish(1, 2))
}
func TestBasicIntStringSame1(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish(1, "1"))
}
func TestBasicIntStringSame2(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish(1, "1.0"))
}
func TestBasicIntFloatSame(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish(1, 1.0))
}
func TestBasicStringStringSame(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish("same", "same"))
}
func TestBasicStringStringDifferent(t *testing.T) {
	assert := assert.New(t)
	assert.False(Sameish("not", "same"))
}
func TestBasicStringIntSame1(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish("1", 1))
}
func SkipTestBasicStringFloatSame(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish("1", 1.0))
}

func TestFirstMapFail(t *testing.T) {
	assert := assert.New(t)
	thing := map[string]interface{}{
		"a": "b",
	}
	assert.False(Sameish(thing, 2))
}
