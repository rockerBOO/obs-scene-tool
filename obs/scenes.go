package obs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ScenesCollection struct {
	CurrentProgramScene string            `json:"current_program_scene"`
	CurrentScene        string            `json:"current_scene"`
	CurrentTransition   string            `json:"current_transition"`
	Groups              []interface{}     `json:"groups"`
	Name                string            `json:"name"`
	PreviewLocked       bool              `json:"preview_locked"`
	QuickTransitions    []QuickTransition `json:"quick_transitions"`
	SavedProjectors     []interface{}     `json:"saved_projectors"`
	ScalingEnabled      bool              `json:"scaling_enabled"`
	ScalingLevel        int               `json:"scaling_level"`
	ScalingOffX         float64           `json:"scaling_off_x"`
	ScalingOffY         float64           `json:"scaling_off_y"`
	SceneOrder          []SceneOrder      `json:"scene_order"`
	Sources             []Source          `json:"sources"`
	TransitionDuration  int               `json:"transition_duration"`
	Transitions         []interface{}     `json:"transitions"`
}

type Filter struct {
	Balance               float64         `json:"balance"`
	DeinterlaceFieldOrder int             `json:"deinterlace_field_order"`
	DeinterlaceMode       int             `json:"deinterlace_mode"`
	Enabled               bool            `json:"enabled"`
	Flags                 int             `json:"flags"`
	Hotkeys               Hotkeys         `json:"hotkeys"`
	ID                    string          `json:"id"`
	Mixers                int             `json:"mixers"`
	MonitoringType        int             `json:"monitoring_type"`
	Muted                 bool            `json:"muted"`
	Name                  string          `json:"name"`
	PrevVer               int             `json:"prev_ver"`
	PrivateSettings       PrivateSettings `json:"private_settings"`
	PushToMute            bool            `json:"push-to-mute"`
	PushToMuteDelay       int             `json:"push-to-mute-delay"`
	PushToTalk            bool            `json:"push-to-talk"`
	PushToTalkDelay       int             `json:"push-to-talk-delay"`
	Settings              Settings        `json:"settings,omitempty"`
	Sync                  int             `json:"sync"`
	VersionedID           string          `json:"versioned_id"`
	Volume                float64         `json:"volume"`
}

type Hotkeys map[string]interface{}
type PrivateSettings map[string]interface{}
type Settings map[string]interface{}

type LibobsMute struct {
	Command bool   `json:"command"`
	Key     string `json:"key"`
}
type LibobsUnmute struct {
	Command bool   `json:"command"`
	Key     string `json:"key"`
}

type AudioDevice struct {
	Balance               float64         `json:"balance"`
	DeinterlaceFieldOrder int             `json:"deinterlace_field_order"`
	DeinterlaceMode       int             `json:"deinterlace_mode"`
	Enabled               bool            `json:"enabled"`
	Flags                 int             `json:"flags"`
	Hotkeys               Hotkeys         `json:"hotkeys"`
	ID                    string          `json:"id"`
	Mixers                int             `json:"mixers"`
	MonitoringType        int             `json:"monitoring_type"`
	Muted                 bool            `json:"muted"`
	Name                  string          `json:"name"`
	PrevVer               int             `json:"prev_ver"`
	PrivateSettings       PrivateSettings `json:"private_settings"`
	PushToMute            bool            `json:"push-to-mute"`
	PushToMuteDelay       int             `json:"push-to-mute-delay"`
	PushToTalk            bool            `json:"push-to-talk"`
	PushToTalkDelay       int             `json:"push-to-talk-delay"`
	Settings              Settings        `json:"settings"`
	Sync                  int             `json:"sync"`
	VersionedID           string          `json:"versioned_id"`
	Volume                float64         `json:"volume"`
}

type QuickTransition struct {
	Duration    int     `json:"duration"`
	FadeToBlack bool    `json:"fade_to_black"`
	Hotkeys     Hotkeys `json:"hotkeys"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
}

type SceneOrder struct {
	Name string `json:"name"`
}

type Bounds struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Pos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Scale struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Item struct {
	Align           int             `json:"align"`
	Bounds          Bounds          `json:"bounds"`
	BoundsAlign     int             `json:"bounds_align"`
	BoundsType      int             `json:"bounds_type"`
	CropBottom      int             `json:"crop_bottom"`
	CropLeft        int             `json:"crop_left"`
	CropRight       int             `json:"crop_right"`
	CropTop         int             `json:"crop_top"`
	GroupItemBackup bool            `json:"group_item_backup"`
	ID              int             `json:"id"`
	Locked          bool            `json:"locked"`
	Name            string          `json:"name"`
	Pos             Pos             `json:"pos"`
	PrivateSettings PrivateSettings `json:"private_settings"`
	Rot             float64         `json:"rot"`
	Scale           Scale           `json:"scale"`
	ScaleFilter     string          `json:"scale_filter"`
	Visible         bool            `json:"visible"`
}

type Font struct {
	Face  string `json:"face"`
	Flags int    `json:"flags"`
	Size  int    `json:"size"`
	Style string `json:"style"`
}

type Source struct {
	Balance               float64                `json:"balance"`
	DeinterlaceFieldOrder int                    `json:"deinterlace_field_order"`
	DeinterlaceMode       int                    `json:"deinterlace_mode"`
	Enabled               bool                   `json:"enabled"`
	Flags                 int                    `json:"flags"`
	Hotkeys               Hotkeys                `json:"hotkeys,omitempty"`
	ID                    string                 `json:"id"`
	Mixers                int                    `json:"mixers"`
	MonitoringType        int                    `json:"monitoring_type"`
	Muted                 bool                   `json:"muted"`
	Name                  string                 `json:"name"`
	PrevVer               int                    `json:"prev_ver"`
	PrivateSettings       PrivateSettings        `json:"private_settings"`
	PushToMute            bool                   `json:"push-to-mute"`
	PushToMuteDelay       int                    `json:"push-to-mute-delay"`
	PushToTalk            bool                   `json:"push-to-talk"`
	PushToTalkDelay       int                    `json:"push-to-talk-delay"`
	Settings              Settings               `json:"settings,omitempty"`
	Sync                  int                    `json:"sync"`
	VersionedID           string                 `json:"versioned_id"`
	Volume                float64                `json:"volume"`
	Items                 map[string]interface{} `json:"items"`
}

func HasScenes(file string) bool {
	has_found, err := FindScenes(file)

	if err != nil {
		return false
	}

	return len(has_found) > 0
}

func FindScenes(file string) ([]Source, error) {
	scenes, err := open_json(file, ScenesCollection{})

	sources := FindSources(scenes)

	var found_sources []Source

	for _, source := range sources {
		if source.ID == "scene_source" {
			found_sources = append(found_sources, source)
		}

		// items, ok := source["items"].(map[string]interface{})
	}

	return found_sources, err
}

// Find if there are any sounrces in a list
func FindSources(scenes ScenesCollection) []Source {
	var results []Source

	for _, source := range scenes.Sources {
		if source.ID == "scene" {
			fmt.Printf("scene found %+v\n", source.Name)
			items, ok := source.Settings["items"].(map[string]string)

			// TODO items needs to properly create a key:value map

			if ok == false {
				log.Println(items)
				log.Println(source.Settings)
				continue
			}

			if len(items) == 0 {
				continue
			}

			for item := range items {
				log.Printf("%+v\n", item)
			}
		}

		if source.ID == "ffmpeg_source" {
			log.Printf("%+v\n", source.Settings["local_file"])
		}

		if source.ID == "image_source" {
			log.Printf("%+v\n", source.Settings["file"])
		}

		if source.ID == "text_ft2_source" {

		}

		if source.ID == "color_source" {

		}

		if source.ID == "xshm_input" {

		}

		if source.ID == "composite_input" {

		}
	}

	return results
}

func GetSources(file string) ([]Source, error) {
	var scenes ScenesCollection
	scenes, err := open_json(file, scenes)

	if err != nil {
		return []Source{}, err
	}

	return scenes.Sources, nil
}

func open_json(file string, collection ScenesCollection) (ScenesCollection, error) {
	file_source, err := ioutil.ReadFile(file)

	if err != nil {
		return collection, err
	}

	if err := json.Unmarshal(file_source, &collection); err != nil {
		return collection, err
	}

	return collection, nil
}
