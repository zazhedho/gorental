package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	result := NewToken("dho", "admin")

	assert.Equal(t, "dho", result.Username, "Expect username = dho")
}

// func TestCheckToken(t *testing.T) {
// 	result, err := CheckToken("abc", "admin")

// 	assert.True(t, true, result, "result must be true")
// 	assert.Equal(t, nil, err, "error must be nil")
// }
