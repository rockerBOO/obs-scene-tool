package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detect scene files in common user locations.",
	Long: `Detect all scene files in the common user locations. 
			~/.config/obs-studio/basic/scenes`,
	Run: func(cmd *cobra.Command, args []string) {
		directories := []string{"~/.config/obs-studio/basic/scenes"}
		for _, detect := range Detect(directories) {
			if len(detect) == 0 {
				continue
			}

			fmt.Println("Detected this file: ")
			fmt.Println(detect)

			x, _ := unwrap([]byte(detect))

			fmt.Println(x)
		}
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)
}

// Detect if any Scene files on the machine
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
	dir := filepath.Dir(directory)
	list, err := ioutil.ReadDir(dir)

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
