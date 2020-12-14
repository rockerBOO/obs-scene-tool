package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/rockerBOO/obs-scene-tool/obs"
	"github.com/spf13/cobra"
)

var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Detect scene files in common user locations.",
	Long: `Detect all scene files in the common user locations. 
			~/.config/obs-studio/basic/scenes`,
	Run: func(cmd *cobra.Command, args []string) {
		directories := []string{filepath.Dir("/home/rockerboo/.config/obs-studio/basic/scenes/")}

		detected_files := Detect(directories)

		for _, file := range detected_files {
			if len(file) == 0 {
				continue
			}

			fmt.Println("Detected this file: ")
			fmt.Println(file)

			x, _ := unwrap([]byte(file))

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
		is_found, file := find_in_directory(dir)

		if !is_found {
			continue
		}

		is_found = obs.HasScenes(file)

		if is_found {
			found_files = append(found_files, file)
		}
	}

	if len(found_files) == 0 {
		log.Printf("No files found")
	} else {
		log.Printf("Found %+v files. %+v", len(found_files), found_files)
	}

	return found_files
}

// Find any scenes files in the directory
func find_in_directory(directory string) (bool, string) {
	dir := filepath.Dir(directory + "/")

	list, err := ioutil.ReadDir(dir)

	if err != nil {
		return false, ""
	}

	for _, file := range list {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".json") {
			return true, filepath.Join(dir, file.Name())
		}
	}

	return false, ""
}

func unwrap(dat []byte) (map[string]interface{}, error) {
	var jsoned map[string]interface{}

	if err := json.Unmarshal(dat, &jsoned); err != nil {
		return nil, err
	}

	return jsoned, nil
}
