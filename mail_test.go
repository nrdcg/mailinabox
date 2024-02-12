package mailinabox

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_GetUsers(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users", testHandler("mail/mailusers.json", http.MethodGet, http.StatusOK))

	records, err := client.Mail.GetUsers(context.Background())
	require.NoError(t, err)

	expected := []MailUsers{{
		Domain: "example.com",
		Users: []User{
			{
				Email:      "user@example.com",
				Privileges: []string{"admin"},
				Status:     "active",
				Mailbox:    "/home/user-data/mail/mailboxes/example.com/user",
			},
		},
	}}

	assert.Equal(t, expected, records)
}

func TestUserService_AddUser(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users/add", testHandler("mail/adduser.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.AddUser(context.Background(), "user@example.com", "secret", "admin")
	require.NoError(t, err)

	assert.Equal(t, "mail user added\nupdated DNS: OpenDKIM configuration", resp)
}

func TestUserService_RemoveUser(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users/remove", testHandler("mail/removeuser.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.RemoveUser(context.Background(), "user@example.com")
	require.NoError(t, err)

	assert.Equal(t, "OK", resp)
}

func TestUserService_AddUserPrivilege(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users/privileges/add", testHandler("mail/adduserprivilege.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.AddUserPrivilege(context.Background(), "user@example.com", "admin")
	require.NoError(t, err)

	assert.Equal(t, "OK", resp)
}

func TestUserService_RemoveUserPrivilege(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users/privileges/remove", testHandler("mail/removeuserprivilege.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.RemoveUserPrivilege(context.Background(), "user@example.com", "admin")
	require.NoError(t, err)

	assert.Equal(t, "OK", resp)
}

func TestUserService_SetUserPassword(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users/password", testHandler("mail/setuserpassword.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.SetUserPassword(context.Background(), "user@example.com", "secret")
	require.NoError(t, err)

	assert.Equal(t, "OK", resp)
}

func TestUserService_GetUserPrivileges(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/users/privileges", testHandler("mail/getMailUserPrivileges.html", http.MethodGet, http.StatusOK))

	domains, err := client.Mail.GetUserPrivileges(context.Background(), "user@example.com")
	require.NoError(t, err)

	assert.Equal(t, "admin", domains)
}

func TestUserService_GetDomains(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/domains", testHandler("mail/getMailDomains.html", http.MethodGet, http.StatusOK))

	domains, err := client.Mail.GetDomains(context.Background())
	require.NoError(t, err)

	expected := []string{"example1.com", "example2.com"}

	assert.Equal(t, expected, domains)
}

func TestUserService_GetAliases(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/aliases", testHandler("mail/getMailAliases.json", http.MethodGet, http.StatusOK))

	aliases, err := client.Mail.GetAliases(context.Background())
	require.NoError(t, err)

	expected := []MailAliases{
		{
			Domain: "example.com",
			Aliases: []Alias{
				{
					Address:          "user@example.com",
					AddressDisplay:   "user@example.com",
					ForwardsTo:       []string{"user@example.com"},
					PermittedSenders: []string{"user@example.com"},
					Required:         true,
				},
			},
		},
	}

	assert.Equal(t, expected, aliases)
}

func TestUserService_UpsertAlias(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/aliases/add", testHandler("mail/upsertMailAlias.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.UpsertAlias(context.Background(), true, "user@example.com", []string{"user@example.com"}, []string{"user@example.com"})
	require.NoError(t, err)

	assert.Equal(t, "alias updated", resp)
}

func TestUserService_RemoveAliases(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/mail/aliases/remove", testHandler("mail/removeMailAlias.html", http.MethodPost, http.StatusOK))

	resp, err := client.Mail.RemoveAliases(context.Background(), "user@example.com")
	require.NoError(t, err)

	assert.Equal(t, "alias updated", resp)
}
