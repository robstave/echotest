package main

import (
	"echotest/cmd"
	"log"
)

// see rybit/seltzer for refernces on how this was done

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
