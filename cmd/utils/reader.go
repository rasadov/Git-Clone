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
	"strings"
)

func getObjectData(path string) ([]string, error) {
	file, _ := os.Open(path)
	r, _ := zlib.NewReader(io.Reader(file))
	s, _ := io.ReadAll(r)
	parts := strings.Split(string(s), "\x00")
	if err := r.Close(); err != nil {
		return nil, err
	}
	return parts, nil
}

func getObjectHeaderPart(ind int, path string) (string, error) {
	parts, err := getObjectData(path)
	if err != nil {
		return "", err
	}
	headerParts := strings.Split(parts[0], " ")
	if ind >= len(headerParts) {
		return "", fmt.Errorf("index out of range")
	}
	return headerParts[ind], nil
}

func ReadObject(readerType, hash string) (string, error) {
	// Get the type of the object
	path := fmt.Sprintf(".git/objects/%v/%v", hash[0:2], hash[2:])
	switch readerType {
	case "e":
		_, err := os.Open(path)
		if err != nil {
			return "", err
		}
		return "Object exists", nil
	case "p":
		parts, err := getObjectData(path)
		if err != nil {
			return "", err
		}
		return parts[1], nil
	case "t":
		res, err := getObjectHeaderPart(0, path)
		if err != nil {
			return "", err
		}
		return res, nil
	case "s":
		res, err := getObjectHeaderPart(1, path)
		if err != nil {
			return "", err
		}
		return res, nil
	default:
		return "", fmt.Errorf("invalid type: %s", readerType)

	}
}
