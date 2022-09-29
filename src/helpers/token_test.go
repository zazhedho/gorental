package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	result := NewToken("dho", "admin")

	assert.Equal(t, "dho", result.Username, "Expect username = dho")
}

func BenchmarkNewToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewToken("dho", "admin")
	}

}
