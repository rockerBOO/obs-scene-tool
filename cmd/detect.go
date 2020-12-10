package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func init() {
}

func Detect(directories []string) []string {
	found_files := make([]string, 1)

	for _, dir := range directories {
		found, file := find_in_directory(dir)

		if found {
			found_files = append(found_files, file)
		}
	}

	return found_files
}

// Find any scenes files in the directory
func find_in_directory(directory string) (bool, string) {
	list, err := ioutil.ReadDir(directory)

	if err != nil {
		return false, ""
	}

	for _, file := range list {
		if !file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".json") {
			return true, file.Name()
		}
	}

	return false, ""
}

// Find any scenes in the file
// func find(file string) (string, error) {
// 	f, err := os.Open(file)

// 	if err != nil {
// 		return "", err
// 	}

// 	return file.
// }

func unwrap(dat []byte) (map[string]interface{}, error) {
	var jsoned map[string]interface{}

	if err := json.Unmarshal(dat, &jsoned); err != nil {
		return nil, err
	}

	return jsoned, nil
}
