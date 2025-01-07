package controls

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getFileType(fileName string) string {
	nameParts := strings.Split(fileName, ".")
	return strings.TrimSpace(nameParts[len(nameParts)-1])
}

func LoadConfig() map[string]string {
	file, _ := os.Open("credentials.txt")
	if file == nil {
		os.Create("credentials.txt")
		file, _ = os.Open("credentials.txt")
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
	return config
}

func SaveConfig(config map[string]string) {
	var content strings.Builder
	for key, value := range config {
		content.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	}
	file, _ := os.OpenFile("credentials.txt", os.O_RDWR|os.O_CREATE, 0755)
	_, err := io.WriteString(file, content.String())
	if err != nil {
		panic(err)
	}
}
