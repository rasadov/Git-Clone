package controls

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func skipTreeHeader(data []byte) []byte {
	headerEnd := 0
	for data[headerEnd] != 0 {
		headerEnd++
	}
	return data[headerEnd+1:]
}

func getTreeContent(data []byte) string {
	var res strings.Builder
	data = skipTreeHeader(data)

	i := 0
	for i < len(data) {
		modeEnd := i
		for data[modeEnd] != ' ' {
			modeEnd++
		}
		mode := string(data[i:modeEnd])
		i = modeEnd + 1

		nameEnd := i
		for data[nameEnd] != 0 {
			nameEnd++
		}
		name := string(data[i:nameEnd])
		i = nameEnd + 1

		if i+20 > len(data) {
			break
		}

		sha := data[i : i+20]
		i += 20

		res.WriteString(fmt.Sprintf("%s %s %x\n", mode, name, sha))
	}
	return res.String()
}

func getObjectType(data []byte) string {
	return strings.Split(string(data), " ")[0]
}

func getObjectData(data []byte) ([]string, error) {
	parts := strings.Split(string(data), "\x00")
	return parts, nil
}

func ReadObject(readerType, hash string) (string, error) {
	path := filepath.Join(GitDir, "objects", hash[:2], hash[2:])
	file, err := os.Open(path)
	reader, _ := zlib.NewReader(io.Reader(file))
	byteData, _ := io.ReadAll(reader)

	switch readerType {
	case "e":
		if err != nil {
			return "", err
		}
		return "Object exists", nil
	case "p":
		fileType := getObjectType(byteData)
		if fileType == "tree" {
			return getTreeContent(byteData), nil
		}
		return string(byteData), nil
	case "t":
		if err != nil {
			return "", err
		}
		return getObjectType(byteData), nil
	case "s":
		stats, err := os.Stat(path)
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(stats.Size(), 10), nil
	default:
		return "", fmt.Errorf("invalid type: %s", readerType)

	}
}
