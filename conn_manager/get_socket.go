package conn_manager

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Credentials struct {
	Username string
	Password string
}

type ServerConfig struct {
	Host        string
	Port        string
	Type        string
	Credentials Credentials
}

var ServConfObj ServerConfig

func scanFile(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var fileInfo []string
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ";")
		for _, word := range words {
			fileInfo = append(fileInfo, word)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	return fileInfo
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
		return
	}

	filePath := filepath.Join(dir, "../conn_manager")

	socketFile, err := os.Open(filepath.Join(filePath, "socket_details.txt"))
	if err != nil {
		panic(err.Error())
	}
	defer socketFile.Close()
	socketInfo := scanFile(socketFile)

	credFile, err := os.Open(filepath.Join(filePath, "credentials.txt"))
	if err != nil {
		panic(err.Error())
	}
	defer credFile.Close()
	credInfo := scanFile(credFile)

	ServConfObj.Host = socketInfo[0]
	ServConfObj.Port = socketInfo[1]
	ServConfObj.Type = socketInfo[2]

	ServConfObj.Credentials.Username = credInfo[0]
	ServConfObj.Credentials.Password = credInfo[1]
}
