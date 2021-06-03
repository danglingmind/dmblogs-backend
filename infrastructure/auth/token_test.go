package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatTokenValidReturnsErrorForEmptyTokenHeader(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)

	err = TokenValid(req)
	if err == nil {
		t.Error("unexpected nil error")
	}
	assert.Equal(t, "token contains an invalid number of segments", err.Error())
}
