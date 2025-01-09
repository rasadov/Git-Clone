package controls

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var GitDir string

func getFileType(fileName string) string {
	nameParts := strings.Split(fileName, ".")
	return strings.TrimSpace(nameParts[len(nameParts)-1])
}

func getFileTypeFromContent(content string) string {
	parts := strings.Split(content, " ")
	return parts[0]
}

func LoadConfig() map[string]string {
	file, _ := os.Open(".env")
	if file == nil {
		os.Create(".env")
		file, _ = os.Open(".env")
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	config := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, "=")
			config[parts[0]] = parts[1]
		}
	}
	GitDir = config["gitDir"]
	return config
}

func SaveConfig(config map[string]string) {
	var content strings.Builder
	for key, value := range config {
		content.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	}
	file, _ := os.OpenFile(".env", os.O_RDWR|os.O_CREATE, 0755)
	_, err := io.WriteString(file, content.String())
	if err != nil {
		panic(err)
	}
}

func GetHead(branch string) string {
	file, _ := os.Open(filepath.Join(GitDir, "refs", "heads", branch))
	data, _ := io.ReadAll(file)
	return string(data)
}

func UpdateHead(branch, hash string) {
	file, _ := os.OpenFile(filepath.Join(GitDir, "heads", branch), os.O_RDWR|os.O_CREATE, 0755)
	_, err := io.WriteString(file, hash)
	if err != nil {
		panic(err)
	}
}
