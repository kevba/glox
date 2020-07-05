package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file := flag.String("f", "", "File containing lox source code")
	flag.Parse()

	var err error
	// Execute the file if it is given as a parameter. If no file is given start an interactive interpreter.
	if *file != "" {
		err = runSourceFile(*file)
	} else {
		err = startPrompt()
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Exited without issue")
}

func runSourceFile(file string) error {
	interpreter := loxInterpreter{}
	sourceBytes, err := readSourceFile(file)
	source := string(sourceBytes)

	interpreter.run(source)

	return err
}

func readSourceFile(file string) ([]byte, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return b, fmt.Errorf("Could not read source file %v: %v", err, file)
	}

	return b, nil
}

func startPrompt() error {
	interpreter := loxInterpreter{}
	for {
		reader := bufio.NewReader(os.Stdout)
		fmt.Print(">")

		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		interpreter.run(input)
		if interpreter.hasError {
			log.Println(interpreter.errorText)
		}
	}
}
