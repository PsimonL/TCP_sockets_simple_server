package connection_stuff

import (
	"bufio"
	"os"
	"strings"
)

type ServerConfig struct {
	Host string
	Port string
	Type string
}

var ServConfObj ServerConfig

func init() {
	file, err := os.Open("connection_params.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var myList []string
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ";")
		for _, word := range words {
			myList = append(myList, word)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	ServConfObj.Host = myList[0]
	ServConfObj.Port = myList[1]
	ServConfObj.Type = myList[2]
}
