package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func Vite(asset string) string {
	port := os.Getenv("VITE_PORT")
	if len(port) == 0 {
		port = "5173"
	}

	u := fmt.Sprintf("http://localhost:%s", port)

	fmt.Println(u)

	return getViteEntry(asset)

	// return "<script type=\"module\" src=\"http://localhost:5173/js/main.js\"></script>"
}

type viteFileManifest struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css"`
}

func getViteEntry(asset string) string {
	if os.Getenv("APP_MODE") == "dev" {
		return fmt.Sprintf("<script type=\"module\" src=\"http://localhost:5173/%s\"></script>", asset)
	}

	// TODO: Handle error
	data, _ := getViteEntries()

	var metaTags string

	metaTags += fmt.Sprintf("<script src=\"/%s\"></script>", data[asset].File)
	for _, cssFile := range data[asset].CSS {
		metaTags += fmt.Sprintf("<link rel=\"stylesheet\" href=\"/%s\" />", cssFile)
	}

	return metaTags
}

func getViteEntries() (map[string]viteFileManifest, error) {
	wd, err := os.Getwd()

	manifest, err := os.ReadFile(path.Join(wd, "manifest.json"))
	if err != nil {
		fmt.Println(err)
	}

	var data map[string]viteFileManifest

	err = json.Unmarshal(manifest, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
