/*
This file contains reader functions
These include Read a blob object and a tree object
*/

package utils

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
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

	ind := 0
	for ind < len(data) {
		modeEnd := ind
		for data[modeEnd] != ' ' {
			modeEnd++
		}
		mode := string(data[ind:modeEnd])
		ind = modeEnd + 1

		nameEnd := ind
		for data[nameEnd] != 0 {
			nameEnd++
		}
		name := string(data[ind:nameEnd])
		ind = nameEnd + 1

		if ind+20 > len(data) {
			break
		}

		sha := data[ind : ind+20]
		ind += 20

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
	path := fmt.Sprintf(".git/objects/%v/%v", hash[0:2], hash[2:])
	file, err := os.Open(path)
	r, _ := zlib.NewReader(io.Reader(file))
	s, _ := io.ReadAll(r)

	switch readerType {
	case "e":
		if err != nil {
			return "", err
		}
		return "Object exists", nil
	case "p":
		fileType := getObjectType(s)
		if fileType == "tree" {
			return getTreeContent(s), nil
		}
		parts, err := getObjectData(s)
		if err != nil {
			return "", err
		}
		result := strings.Join(parts[1:], "\n")
		return result, nil
	case "t":
		if err != nil {
			return "", err
		}
		return getObjectType(s), nil
	case "s":
		stats, err := os.Stat(path)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(int(stats.Size())), nil
	default:
		return "", fmt.Errorf("invalid type: %s", readerType)

	}
}
