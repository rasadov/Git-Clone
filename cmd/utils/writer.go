package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func CreateObject(path, writePath, fileType string) string {
	var contentAndHeader string
	file, _ := os.ReadFile(path)
	stats, _ := os.Stat(path)
	content := string(file)
	if fileType != "" {
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
	z := zlib.NewWriter(&buffer)
	z.Write([]byte(contentAndHeader))
	z.Close()
	os.MkdirAll(filepath.Dir(blobPath), os.ModePerm)
	f, _ := os.Create(blobPath)
	defer f.Close()
	if writePath != "" {
		os.WriteFile(writePath, buffer.Bytes(), 0644)
	}
	return hash
}
