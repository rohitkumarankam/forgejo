package templates

import (
	"testing"
	
	"forgejo.org/modules/setting"
)

func TestCommitLink(t *testing.T) {
	// Save original setting
	originalSetting := setting.UI.ForceFileOnlyCommitDiffs
	defer func() {
		setting.UI.ForceFileOnlyCommitDiffs = originalSetting
	}()
	
	repoLink := "/testuser/testRepo"
	commitID := "abc123def456"
	
	// Test with setting disabled (default)
	setting.UI.ForceFileOnlyCommitDiffs = false
	expected := "/testuser/testRepo/commit/abc123def456"
	result := CommitLink(repoLink, commitID)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	
	// Test with setting enabled
	setting.UI.ForceFileOnlyCommitDiffs = true
	expected = "/testuser/testRepo/commit/abc123def456?file-only=true"
	result = CommitLink(repoLink, commitID)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}