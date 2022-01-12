package main

import "github.com/ImperiumProject/imperium/cmd"

func main() {
	c := cmd.RootCmd()
	c.Execute()
}
