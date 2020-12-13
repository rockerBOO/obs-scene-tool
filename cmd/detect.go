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
		directories := []string{filepath.Dir("/home/rockerboo/.config/obs-studio/basic/scenes/")}

		fmt.Printf("directories: %+v\n", directories)

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
		fmt.Printf("find in directory %+v\n", dir)
		is_found, file := find_in_directory(dir)

		if !is_found {
			fmt.Printf("no file found %+v\n", dir)
			continue
		}

		sources, err := find_sources(file)

		if err != nil {
			fmt.Printf("no sources found %+v\n", err)
		}

		for _, source := range sources {
			fmt.Printf("source found %+v\n", source.Name)
		}

		if is_found {
			found_files = append(found_files, file)
		}
	}

	fmt.Printf("%+v\n", found_files)

	return found_files
}

type Scenes struct {
	Sources []Source
}

type Source struct {
	Id      string
	Name    string
	Enabled bool
}

func open_json(file string, unpack_json Scenes) (*Scenes, error) {
	file_source, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("%v %+v\n", file, err)
		return &Scenes{}, err
	}

	if err := json.Unmarshal(file_source, &unpack_json); err != nil {
		fmt.Printf("invalid unmarshal for %+v %+v\n", file, err)
		return &Scenes{}, err
	}

	return &unpack_json, nil
}

func find_sources(file string) ([]Source, error) {
	scenes, err := open_json(file, Scenes{})

	if err != nil {
		return []Source{}, err
	}

	return scenes.Sources, nil
}

// Find any scenes files in the directory
func find_in_directory(directory string) (bool, string) {
	dir := filepath.Dir(directory + "/")

	println(dir)

	list, err := ioutil.ReadDir(dir)

	fmt.Printf("find_in_directory %+v\n", err)

	if err != nil {
		return false, ""
	}

	for _, file := range list {
		fmt.Printf("file in list %+v\n", dir+file.Name())
		if file.IsDir() {
			fmt.Printf("is a dir %+v\n", dir+file.Name())
			continue
		}

		if strings.HasSuffix(file.Name(), ".json") {
			fmt.Printf("Found file %+v\n", dir+file.Name())
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
