package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type entry struct {
	fileName string
	byteData []byte
}

func getFileType(fileName string) string {
	nameParts := strings.Split(fileName, ".")
	return strings.TrimSpace(nameParts[len(nameParts)-1])
}

func getGitIgnore() ([]string, []string) {
	var (
		ignoreList  []string
		ignoreTypes []string
	)
	file, err := os.Open(".gitignore")
	gitignore, err := io.ReadAll(file)
	//fmt.Println("Gitignore: ", string(gitignore))
	if err == nil {
		ignore := strings.Split(string(gitignore), "\n")
		for i, v := range ignore {
			if v != "" {
				switch v[0] {
				case '*':
					fileType := getFileType(v)
					ignoreTypes = append(ignoreTypes, fileType)
				case '#':
				default:
					ignoreList = append(ignoreList, v)
				}

			}
			if i == len(ignore)-1 {
				break
			}
		}
	}
	return ignoreList, ignoreTypes
}

func validateFile(fileName, fileType string, ignoreList, ignoreTypes []string) bool {
	if len(ignoreTypes) > 0 {
		for _, v := range ignoreTypes {
			if v == fileType {
				return true
			}
		}
	}
	if len(ignoreList) > 0 {
		for _, v := range ignoreList {
			if v == fileName {
				return true
			}
		}
	}
	return false
}

func calcTreeHash(dir string) ([]byte, []byte) {
	var entries []entry
	ignoreList, ignoreTypes := getGitIgnore()
	fileInfos, _ := os.ReadDir(dir)
	contentSize := 0

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		fileType := getFileType(fileName)
		ignore := validateFile(fileName, fileType, ignoreList, ignoreTypes)

		if fileInfo.Name() == ".git" || ignore {
			continue
		}

		if !fileInfo.IsDir() {
			file, _ := os.Open(filepath.Join(dir, fileInfo.Name()))
			byteData, _ := io.ReadAll(file)
			str := fmt.Sprintf("blob %d\u0000%s", len(byteData), string(byteData))
			sha1Val := sha1.New()
			io.WriteString(sha1Val, str)
			str = fmt.Sprintf("100644 %s\u0000", fileInfo.Name())
			byteData = append([]byte(str), sha1Val.Sum(nil)...)
			entries = append(entries, entry{fileInfo.Name(), byteData})
			contentSize += len(byteData)
		} else {
			byteData, _ := calcTreeHash(filepath.Join(dir, fileInfo.Name()))
			str := fmt.Sprintf("40000 %s\u0000", fileInfo.Name())
			b2 := append([]byte(str), byteData...)
			entries = append(entries, entry{fileInfo.Name(), b2})
			contentSize += len(b2)
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].fileName < entries[j].fileName
	})
	str := fmt.Sprintf("tree %d\u0000", contentSize)
	byteData := []byte(str)
	for _, entry := range entries {
		byteData = append(byteData, entry.byteData...)
	}
	sha1Val := sha1.New()
	io.WriteString(sha1Val, string(byteData))
	return sha1Val.Sum(nil), byteData
}

func CreateObject(path, writePath, fileType string) string {
	var contentAndHeader string
	file, _ := os.ReadFile(path)
	stats, _ := os.Stat(path)
	content := string(file)
	if fileType == "" {
		contentAndHeader = fmt.Sprintf("blob %d\x00%s", stats.Size(), content)
	} else {
		contentAndHeader = fmt.Sprintf("%s %d\x00%s", fileType, stats.Size(), content)
	}

	sha := sha1.Sum([]byte(contentAndHeader))
	hash := fmt.Sprintf("%x", sha)
	blobName := []rune(hash)
	blobPath := ".git/objects/"
	for i, v := range blobName {
		blobPath += string(v)
		if i == 1 {
			blobPath += "/"
		}
	}
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	writer.Write([]byte(contentAndHeader))
	writer.Close()
	os.MkdirAll(filepath.Dir(blobPath), os.ModePerm)
	fileCreated, _ := os.Create(blobPath)
	defer fileCreated.Close()
	if writePath != "" {
		os.WriteFile(writePath, buffer.Bytes(), 0644)
	}
	return hash
}

func WriteTree() string {
	currentDir, _ := os.Getwd()
	hash, content := calcTreeHash(currentDir)
	treeHash := hex.EncodeToString(hash)
	os.Mkdir(filepath.Join(".git", "objects", treeHash[:2]), 0755)
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	writer.Write(content)
	writer.Close()
	os.WriteFile(filepath.Join(".git", "objects", treeHash[:2], treeHash[2:]), buffer.Bytes(), 0644)
	return treeHash
}
