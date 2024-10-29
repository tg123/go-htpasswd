package htpasswd

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var contents = `users: user1 user2 user3
admins: user1
`

var contents2 = `users: user1 user2 user3
admins: user1 user2
`

func TestGroups(t *testing.T) {
	// create temp file and write "contents" into it
	var f, err = os.CreateTemp("", "gogroups")
	filename := f.Name()
	assert.NoError(t, err)
	if _, err := f.WriteString(contents); err != nil {
		t.Fatalf("Failed to write temporary file: %s", err.Error())
	}
	f.Close()

	// Read the file and check the contents
	htGroup, err := NewGroups(filename, nil)
	assert.NoError(t, err)
	assert.True(t, htGroup.IsUserInGroup("user1", "users"))
	assert.True(t, htGroup.IsUserInGroup("user1", "admins"))
	assert.True(t, htGroup.IsUserInGroup("user2", "users"))
	assert.False(t, htGroup.IsUserInGroup("user2", "admin"))
	assert.False(t, htGroup.IsUserInGroup("unknownuser", "users"))
	assert.False(t, htGroup.IsUserInGroup("user1", "unknowngroup"))
	assert.False(t, htGroup.IsUserInGroup("unknownuser", "unknowngroup"))
	assert.Len(t, htGroup.GetUserGroups("user1"), 2)
	assert.Len(t, htGroup.GetUserGroups("user2"), 1)
	assert.Len(t, htGroup.GetUserGroups("user3"), 1)
	assert.Len(t, htGroup.GetUserGroups("unknownuser"), 0)

	// Replace temp file with another one (contents2)
	os.Remove(filename)
	f2, errCreate := os.Create(filename)
	if errCreate != nil {
		t.Fatalf("Failed to create temporary file: %s", errCreate.Error())
	}
	if _, err := f2.WriteString(contents2); err != nil {
		t.Fatalf("Failed to write temporary file: %s", err.Error())
	}
	f2.Close()
	defer os.Remove(filename)

	// Reread the file and check the contents again, user2 should now be member of admins too.
	reloadError := htGroup.ReloadGroups(nil)
	assert.NoError(t, reloadError)
	assert.True(t, htGroup.IsUserInGroup("user1", "users"))
	assert.True(t, htGroup.IsUserInGroup("user1", "admins"))
	assert.True(t, htGroup.IsUserInGroup("user2", "users"))
	assert.True(t, htGroup.IsUserInGroup("user2", "admins"))
	assert.False(t, htGroup.IsUserInGroup("unknownuser", "users"))
	assert.False(t, htGroup.IsUserInGroup("user1", "unknowngroup"))
	assert.False(t, htGroup.IsUserInGroup("unknownuser", "unknowngroup"))
	assert.Len(t, htGroup.GetUserGroups("user1"), 2)
	assert.Len(t, htGroup.GetUserGroups("user2"), 2)
	assert.Len(t, htGroup.GetUserGroups("user3"), 1)
	assert.Len(t, htGroup.GetUserGroups("unknownuser"), 0)

	// Test load from reader as well
	r := strings.NewReader(contents2)
	htGroup, err = NewGroupsFromReader(r, nil)
	assert.NoError(t, err)
	assert.True(t, htGroup.IsUserInGroup("user1", "users"))
	assert.True(t, htGroup.IsUserInGroup("user1", "admins"))
	assert.True(t, htGroup.IsUserInGroup("user2", "users"))
	assert.True(t, htGroup.IsUserInGroup("user2", "admins"))
	assert.False(t, htGroup.IsUserInGroup("unknownuser", "users"))
	assert.False(t, htGroup.IsUserInGroup("user1", "unknowngroup"))
	assert.False(t, htGroup.IsUserInGroup("unknownuser", "unknowngroup"))
	assert.Len(t, htGroup.GetUserGroups("user1"), 2)
	assert.Len(t, htGroup.GetUserGroups("user2"), 2)
	assert.Len(t, htGroup.GetUserGroups("user3"), 1)
	assert.Len(t, htGroup.GetUserGroups("unknownuser"), 0)
}
