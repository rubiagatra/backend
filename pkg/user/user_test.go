package user_test

import (
	"testing"

	"github.com/rubiagatra/backend/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	assert.Equal(t, "user", user.GetUser(), "they should be equal")
}
