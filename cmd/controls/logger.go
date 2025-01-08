package controls

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetCommits() {
	var builder []string
	file, _ := os.Open(filepath.Join(GitDir, "heads", "main"))
	hash, _ := io.ReadAll(file)
	hashStr := string(hash)
	commit, _ := ReadObject("p", hashStr)
	var data, parent string
	for {
		data, parent = parseCommit(commit)
		res := fmt.Sprintf("\033[33mCommit: %s\033[0m\n%s", hashStr, data)
		builder = append(builder, res)
		if parent == "" {
			break
		}
		commit, _ = ReadObject("p", parent)
	}
	for i := len(builder) - 1; i >= 0; i-- {
		fmt.Println(builder[i])
	}
}

func parseCommit(commitData string) (string, string) {
	var parent string
	res := strings.Builder{}
	lines := strings.Split(commitData, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}
		if line[:6] == "author" {
			data := strings.Split(line, " ")
			res.WriteString(fmt.Sprintln("Author: ", data[1], data[2]))
			res.WriteString(fmt.Sprintln("Date: ", unixToHuman(data[3]), data[4]))
		}
		if line[:6] == "parent" {
			parent = line[7:]
		}
		if i == len(lines)-1 {
			res.WriteString(fmt.Sprintln("Message: ", lines[i]))
		}
	}
	return res.String(), parent
}

func unixToHuman(unix string) string {
	unixInt, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		panic(err)
	}
	unixTime := time.Unix(unixInt, 0)
	return unixTime.Format("Mon Jan 2 15:04:05 2006 -0700")
}
