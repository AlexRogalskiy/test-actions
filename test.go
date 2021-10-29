package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"unicode/utf8"
)

func main() {
	// fmt.Println("yo")
	cmd := exec.Command("echo", "Hello, world!")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	stdoutReader := bufio.NewReader(stdout)
	// writer := bufio.NewWriter(&bufio.Writer{})

	output, err := readData(stdoutReader)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Wait()

	fmt.Println(utf8.RuneCountInString(output))

	fmt.Println(output)
}

func readData(reader *bufio.Reader) (string, error) {
	var output string = ""
	var line string
	var readErr error
	for {
		line, readErr = reader.ReadString('\n')

		// remove newline, our output is in a slice,
		// one element per line.
		line = strings.TrimSuffix(line, "\n")

		// only return early if the line does not have
		// any contents. We could have a line that does
		// not not have a newline before io.EOF, we still
		// need to add it to the output.
		if len(line) == 0 && readErr == io.EOF {
			break
		}

		// // logger.Logger has a Logf method, but not a Log method.
		// // We have to use the format string indirection to avoid
		// // interpreting any possible formatting characters in
		// // the line.
		// //
		// // See https://github.com/gruntwork-io/terratest/issues/982.
		// log.Logf(t, "%s", line)

		output = output + line
		// if _, err := writer.WriteString(line); err != nil {
		// 	return err
		// }

		if readErr != nil {
			break
		}
	}
	if readErr != io.EOF {
		return output, readErr
	}
	return output, nil
}
