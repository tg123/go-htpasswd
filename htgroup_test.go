package htpasswd

import (
	"io/ioutil"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

var contents = 
`users: user1 user2 user3
admins: user1
`

var contents2 = 
`users: user1 user2 user3
admins: user1 user2
`

func TestGroups(t *testing.T) { 
	// create temp file and write "contents" into it
	var f, err = ioutil.TempFile("", "gogroups")
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
	f, err = os.Create(f.Name())
	if _, err := f.WriteString(contents2); err != nil {
		t.Fatalf("Failed to write temporary file: %s", err.Error())
	}
	f.Close()
	defer os.Remove(filename)

	// Reread the file and check the contents again, user2 should now be member of admins too.
	reloadError := htGroup.ReloadGroups(nil);
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
}
