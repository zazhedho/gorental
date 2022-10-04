package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	result, err := HashPassword("mypassword")

	fmt.Println(result)
	assert.NotEqual(t, "mypassword", result, "password harus sudah di-hash dan tidak sama dengan 'mypassword'")
	assert.Equal(t, nil, err, "error must be nil")
}

func TestCheckPassword(t *testing.T) {
	passInput := "mypassword"
	hashedPassword := "$2a$10$GHjsp91ocdZlHxq4HRvoo./I52BZ4fkqjMebBufPvhQZeiB.5CyDe"

	assert.True(t, true, CheckPassword(hashedPassword, passInput), "actual must be true")
}
