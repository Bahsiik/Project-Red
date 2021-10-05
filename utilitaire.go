package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	return text
}
