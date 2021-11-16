package main

import (
	"fmt"
	"log"
	"os"
)

func dumpEventLog(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range events {
		fmt.Fprintf(file, "%d %d\n", e.From, e.To)
	}
	defer file.Close()
}

func dumpAgents(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range activeAgents() {
		fmt.Fprintf(file, "%d\n", a.Id)
	}

	defer file.Close()
}
