package duplicate_file

import "testing"

var (
	testDir = "/home/vagrant/temp_dir"
)

func TestDuplicateFileContent(t *testing.T) {
	SameContentFiles(testDir)
}
