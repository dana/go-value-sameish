package valuesameish

import (
	//	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	assert := assert.New(t)
	assert.True(Sameish(1, 1))
}
