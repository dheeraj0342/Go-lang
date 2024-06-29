package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter Username: ")
	//username, _ := reader.ReadString('\n')
	//fmt.Print("Enter Password: ")
	//password, _ := reader.ReadString('\n')
	//fmt.Println("Hello, ", username, " ", password)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a numbers:")
	a, _ := reader.ReadString('\n')
	c, err := strconv.ParseInt(strings.TrimSpace(a), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
