package mailinabox

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_Login(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/login", testHandler("user/login.json", http.MethodGet, http.StatusOK))

	names, err := client.User.Login(context.Background())
	require.NoError(t, err)

	expected := &Session{
		APIKey:     "1a2b3c4d5e6f7g8h9i0j",
		Email:      "user@example.com",
		Privileges: []string{"admin"},
		Status:     "ok",
	}

	assert.Equal(t, expected, names)
}

func TestUserService_Login_error(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/login", testHandler("user/login_error.json", http.MethodGet, http.StatusOK))

	_, err := client.User.Login(context.Background())
	require.EqualError(t, err, "invalid: Incorrect username or password")
}

func TestUserService_Logout(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/logout", testHandler("user/logout.json", http.MethodPost, http.StatusOK))

	names, err := client.User.Logout(context.Background())
	require.NoError(t, err)

	expected := &Session{Status: "ok"}

	assert.Equal(t, expected, names)
}
