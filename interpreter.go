package main

import "log"

type loxInterpreter struct {
	hasError  bool
	errorText error
}

func (l loxInterpreter) run(line string) {
	log.Println(line)
}
