package mailinabox

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSystemService_GetStatus(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/status", testHandler("system/getSystemStatus.json", http.MethodPost, http.StatusOK))

	statuses, err := client.System.GetStatus(context.Background())
	require.NoError(t, err)

	expected := []SystemStatus{
		{
			Type:  "heading",
			Text:  "System",
			Extra: []ExtraStatus{},
		},
		{
			Type: "warning",
			Text: "This domain's DNSSEC DS record is not set",
			Extra: []ExtraStatus{
				{
					Monospace: false,
					Text:      "Digest Type: 2 / SHA-25",
				},
			},
		},
	}

	assert.Equal(t, expected, statuses)
}

func TestSystemService_GetVersion(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/version", testHandler("system/getSystemVersion.html", http.MethodGet, http.StatusOK))

	resp, err := client.System.GetVersion(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "v0.46", resp)
}

func TestSystemService_GetUpstreamVersion(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/latest-upstream-version", testHandler("system/getSystemUpstreamVersion.html", http.MethodPost, http.StatusOK))

	resp, err := client.System.GetUpstreamVersion(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "v0.47", resp)
}

func TestSystemService_GetUpdates(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/updates", testHandler("system/getSystemUpdates.html", http.MethodGet, http.StatusOK))

	updates, err := client.System.GetUpdates(context.Background())
	require.NoError(t, err)

	expected := []string{
		"libgnutls30 (3.5.18-1ubuntu1.4)",
		"libxau6 (1:1.0.8-1ubuntu1)",
	}

	assert.Equal(t, expected, updates)
}

func TestSystemService_UpdatePackages(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/update-packages", testHandler("system/updateSystemPackages.html", http.MethodPost, http.StatusOK))

	resp, err := client.System.UpdatePackages(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "Calculating upgrade...\nThe following packages will be upgraded:\n  cloud-init grub-common", resp)
}

func TestSystemService_GetPrivacyStatus(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/privacy", testHandler("system/getSystemPrivacyStatus.json", http.MethodGet, http.StatusOK))

	resp, err := client.System.GetPrivacyStatus(context.Background())
	require.NoError(t, err)

	assert.False(t, resp)
}

func TestSystemService_UpdatePrivacyStatus(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/privacy", testHandler("system/updateSystemPrivacy.html", http.MethodPost, http.StatusOK))

	resp, err := client.System.UpdatePrivacyStatus(context.Background(), "private")
	require.NoError(t, err)

	assert.Equal(t, "OK", resp)
}

func TestSystemService_GetRebootStatus(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/reboot", testHandler("system/getSystemRebootStatus.json", http.MethodGet, http.StatusOK))

	resp, err := client.System.GetRebootStatus(context.Background())
	require.NoError(t, err)

	assert.True(t, resp)
}

func TestSystemService_Reboot(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/reboot", testHandler("system/rebootSystem.html", http.MethodPost, http.StatusOK))

	resp, err := client.System.Reboot(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "No reboot is required, so it is not allowed.", resp)
}

func TestSystemService_GetBackupStatus(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/backup/status", testHandler("system/getSystemBackupStatus.json", http.MethodGet, http.StatusOK))

	status, err := client.System.GetBackupStatus(context.Background())
	require.NoError(t, err)

	expected := &BackupStatus{
		Backups: []Backup{
			{
				Date:      "20200801T023706Z",
				DateDelta: "15 hours, 40 minutes",
				DateStr:   "2020-08-01 03:37:06 BST",
				DeletedIn: "approx. 6 days",
				Full:      false,
				Size:      125332,
				Volumes:   1,
			},
		},
		UnmatchedFileSize: 0,
		Error:             "Something is wrong with the backup",
	}

	assert.Equal(t, expected, status)
}

func TestSystemService_GetBackupConfig(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/backup/config", testHandler("system/getSystemBackupConfig.json", http.MethodGet, http.StatusOK))

	status, err := client.System.GetBackupConfig(context.Background())
	require.NoError(t, err)

	expected := &BackupConfig{
		EncPwFile:           "/home/user-data/backup/secret_key.txt",
		FileTargetDirectory: "/home/user-data/backup/encrypted",
		MinAgeInDays:        3,
		SSHPubKey:           "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDb root@box.example.com\\n",
		Target:              "s3://s3.eu-central-1.amazonaws.com/box-example-com",
		TargetUser:          "string",
		TargetPass:          "string",
	}

	assert.Equal(t, expected, status)
}

func TestSystemService_UpdateBackupConfig(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/system/backup/config", testHandler("system/updateSystemBackupConfig.html", http.MethodPost, http.StatusOK))

	resp, err := client.System.UpdateBackupConfig(context.Background(), "s3://s3.eu-central-1.amazonaws.com/box-example-com", "ACCESS_KEY", "SECRET_ACCESS_KEY", 3)
	require.NoError(t, err)

	assert.Equal(t, "OK", resp)
}
