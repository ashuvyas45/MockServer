package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the mock server")
	_, err := io.Copy(os.Stdin, os.Stdout)
	if err != nil {
		logrus.Fatal(err)
	}
}
