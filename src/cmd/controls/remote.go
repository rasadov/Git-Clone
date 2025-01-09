package controls

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

var Remotes = map[string]string{}

func SetRemote(name, url string) {
	Remotes[name] = url
}

func RemoveRemote(name string) {
	delete(Remotes, name)
}

func GetRemote(name string) string {
	return Remotes[name]
}

func SaveRemotes() {
	var content strings.Builder
	for key, value := range Remotes {
		content.WriteString("[remote \"" + key + "\"]\n")
		content.WriteString("	url=" + value + "\n")
	}
	file, _ := os.OpenFile(filepath.Join(GitDir, "config"), os.O_RDWR|os.O_CREATE, 0755)
	_, err := file.WriteString(content.String())
	if err != nil {
		panic(err)
	}
}

func init() {
	LoadConfig()
	file, _ := os.Open(filepath.Join(GitDir, "config"))
	data, _ := io.ReadAll(file)
	lines := strings.Split(string(data), "\n")
	for i, v := range lines {
		if strings.Contains(v, "[remote") {
			remote := strings.Split(strings.Split(v, "\"")[1], "\"")[0]
			url := strings.Split(lines[i+1], "=")[1]
			Remotes[remote] = url
		}
	}
}
