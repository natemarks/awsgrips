package testfunctions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FirstAncestorDir given workingDir and targetDir retrun nearest ancestor
// workingDir: /aaa/bbb/ccc/my_target/ggg/hhh/my_target/jjj
// targetDir: my_target
// return : /aaa/bbb/ccc/my_target/ggg/hhh/my_target/
func FirstAncestorDir(workingDir, targetDir string) (result string, err error) {

	sep := string(os.PathSeparator)
	words := strings.Split(workingDir, sep)

	word := words[len(words)-1]
	for word != targetDir {
		words = words[:len(words)-1]
		word = words[len(words)-1]
		if len(words) == 0 {
			break
		}
	}
	if len(words) == 0 {
		msg := fmt.Sprintf("target not found in parent directories: %s", targetDir)
		return "", errors.New(msg)
	}
	return sep + filepath.Join(words...) + sep, err
}

// getProjectPath return the local,  absolute path to awsgrips
// example:
// '/Users/nmarks/go/src/github.com/natemarks/awsgrips'
func getProjectPath() string {
	var result string
	var err error
	wd, _ := os.Getwd()
	result, err = FirstAncestorDir(wd, "awsgrips")
	if err != nil {
		panic(err)
	}
	return result
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
