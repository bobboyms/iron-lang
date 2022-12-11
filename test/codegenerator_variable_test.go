package test

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestCodeGeneratorVariable(t *testing.T) {

	f, err := os.Open("variable.ir")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)

	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Compiler(builder.String())
	RunMake()

}
