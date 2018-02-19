package duplicate_file

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DuplicateFileInfo struct {
	FileName             string
	MatchContentFileName []string
}

/*
* "Given a directory with lots of files, find the files that have the same content"
 */
func SameContentFiles(dir string) {
	// it will return a map where all file with same file sizes are divided
	fileListMap := getFilesInDir(dir)

	if len(fileListMap) == 0 {
		return
	}
	fileMap := make(map[string][]string)
	for size, fileList := range fileListMap {
		fmt.Println(size, fileList)
		if len(fileList) <= 1 {
			continue
		}
		tempMd5Map := make(map[string][]string)
		//tempMd5List := make([]string, 0)
		for _, file := range fileList {
			md5sum := fmt.Sprintf("%x", getMd5(file))
			entry, exists := tempMd5Map[md5sum]
			if exists {
				// compare byte-by-byte for the file & the entry first file
				if compareFiles(file, entry[0]) {
					entry = append(entry, file)
				}
			} else {
				entry = append(entry, file)
			}
			tempMd5Map[md5sum] = entry
		}
	}
	/*
		for _, file := range fileList {
			md5sum := fmt.Sprintf("%x", getMd5(file))

			entry, exists := fileMap[md5sum]
			if exists {
				entry.MatchContentFileName = append(entry.MatchContentFileName, file)
			} else {
				entry.FileName = file
			}
			fileMap[md5sum] = entry
		}
	*/
	//printDuplicateFileContent(fileMap)
}

func compareFiles(a, b string) bool {

	f1, err1 := ioutil.ReadFile(a)

	if err1 != nil {
		return false
	}

	f2, err2 := ioutil.ReadFile(b)
	if err2 != nil {
		return false
	}

	return bytes.Equal(f1, f2)

}

func getMd5(file string) []byte {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil
	}
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		h = nil
		return nil
	}
	return h.Sum(nil)
}

// returns a map of [file size][]files
func getFilesInDir(dir string) map[int64][]string {
	fileMap := make(map[int64][]string)
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		//fmt.Println(f.Size())
		if !f.IsDir() {
			fileList := fileMap[f.Size()]
			fileList = append(fileList, path)
			fileMap[f.Size()] = fileList
		}
		return err
	})
	//printFileNames(fileList)
	return fileMap
}

func printFileNames(files []string) {
	for _, file := range files {
		fmt.Println(file)
	}
}

func printDuplicateFileContent(fileMap map[string]DuplicateFileInfo) {
	for _, fileInfo := range fileMap {
		fmt.Println("file:", fileInfo.FileName, "has duplicate files as follows:")
		//fmt.Println(strings.Trim(fmt.Sprintf(fileInfo.MatchContentFileName), "[]"))
		fmt.Println(fileInfo.MatchContentFileName)
	}
}
