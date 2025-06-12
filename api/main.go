package main

import (
	"basic-service/cmd"
	"os"
)

func main() {
	i := cmd.Execute()
	os.Exit(i)
}
