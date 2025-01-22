package remote

import (
	utilGitConfig "github.com/AmadlaOrg/LibraryUtils/git/config"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func Test_integration_Tags(t *testing.T) {
	gitRemoteService := NewGitRemoteService(
		"https://github.com/AmadlaOrg/QAFixturesEntityMultipleTagVersion",
		&utilGitConfig.Config{})
	tags, err := gitRemoteService.Tags()
	if err != nil {
		t.Errorf("Tags returned an error: %s", err)
	}

	expectedTags := []string{
		"v1.0.0",
		"v1.0.0-alpha.2",
		"v1.0.0-beta.1",
		"v2.0.0",
		"v2.0.1",
		"v2.1.0",
	}

	sort.Strings(tags)
	sort.Strings(expectedTags)

	if !reflect.DeepEqual(tags, expectedTags) {
		t.Errorf("Tags do not match expected values.\nGot: %v\nExpected: %v", tags, expectedTags)
	}

	// Optionally log the sorted tags for debugging purposes
	for _, tag := range tags {
		t.Logf("Retrieved tag: %s", tag)
	}
}

func Test_integration_CommitHeadHash(t *testing.T) {
	gitRemoteService := NewGitRemoteService(
		"https://github.com/AmadlaOrg/QAFixturesEntityMultipleTagVersion",
		&utilGitConfig.Config{})
	hash, err := gitRemoteService.CommitHeadHash()
	if err != nil {
		t.Errorf("CommitHeadHash returned an error: %s", err)
	}

	assert.Equal(t, hash, "8be468562e86eafd0841fe9cfb4a642984c72b87")
}
