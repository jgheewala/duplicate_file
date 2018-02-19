package duplicate_file

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	testDir = "/home/vagrant/temp_dir"
)

func TestDuplicateFileContent(t *testing.T) {
	rv := SameContentFiles(testDir)
	expectedMap := make(map[string][]string)
	expectedMap["322d8361849d9f0de2bdffac9f4050b2"] = []string{
		"/home/vagrant/temp_dir/b.txt",
		"/home/vagrant/temp_dir/d.txt",
		"/home/vagrant/temp_dir/d1/b_dup.txt",
	}
	expectedMap["fc78c7e108a64f944ff2ee8f51c23c55"] = []string{
		"/home/vagrant/temp_dir/f.txt",
		"/home/vagrant/temp_dir/f1.txt",
	}
	expectedMap["7f0f3cf2d246125248a6a74c261ec89e"] = []string{
		"/home/vagrant/temp_dir/a.txt",
		"/home/vagrant/temp_dir/c.txt",
		"/home/vagrant/temp_dir/d1/a_dup.txt",
		"/home/vagrant/temp_dir/e.txt",
	}
	if !reflect.DeepEqual(rv, expectedMap) {
		fmt.Println("")
		fmt.Println("===============RETURN INFORMATION=============")
		fmt.Println("")
		printDuplicateFileContent(rv)
		fmt.Println("")
		fmt.Println("===============EXPECTED INFORMATION===========")
		fmt.Println("")
		printDuplicateFileContent(expectedMap)
		t.Error()
	}
}
