package testfunctions

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// getProjectPath return the absolut path to awsgrips
// example:
// sep + filepath.Join(result...) + sep
// '/Users/nmarks/go/src/github.com/natemarks/awsgrips'
func getProjectPath() string {
	var found bool = false
	var result []string
	sep := string(os.PathSeparator)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	words := strings.Split(wd, sep)

	for _, word := range words {
		result = append(result, word)
		if word == "awsgrips" {
			found = true
			break
		}
	}
	if !found {
		return ""
	}
	return sep + filepath.Join(result...) + sep
}

// JSONFileToByteArray Given a filePath, return the contents as a byte array
// json.Unmarshall takes the byte array as input
func JSONFileToByteArray(filePath string) (byteArray []byte, err error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return byteArray, err
	}
	defer jsonFile.Close()
	byteArray, err = ioutil.ReadAll(jsonFile)
	return byteArray, err
}
