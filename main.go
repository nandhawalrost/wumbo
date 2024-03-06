package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"wlang/detector"
)

func ReadFile(name string, extension string) {
	fileName := name + extension

	data, err := os.ReadFile(name + extension)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\n Running: %s", fileName)
	fmt.Printf("\n Size: %d bytes", len(data))
	fmt.Printf("\n Raw: %s", data)
	fmt.Printf("\n")
	fmt.Printf("\n Converted to String:")
	text := string(data)
	fmt.Println("\n", text)
	fmt.Printf("\n")
	fmt.Printf("Output:")
	fmt.Println()

	length_of_syntax := len(InstructionParser(text))
	detector.DetectLine(text, length_of_syntax, InstructionParser(text))
}

// output() function trimmer
func TrimOutFunction(str string) string {
	trim_out := regexp.MustCompile("(out)") // contains out
	trimmed_out := trim_out.ReplaceAllString(str, "")
	trim_opening := regexp.MustCompile("[(')\n]") // match either ( or ' or )
	trimmed := trim_opening.ReplaceAllString(trimmed_out, "")

	return trimmed
}

// type integer function trimmer
func TrimIntType(str string) string {
	trim_out := regexp.MustCompile("int") // contains out
	trimmed_out := trim_out.ReplaceAllString(str, "")
	trim_opening := regexp.MustCompile("[=\t]") // match either ( or ' or )
	trimmed := trim_opening.ReplaceAllString(trimmed_out, "")

	var array []string = strings.Split(trimmed, " ")

	length := len(array)
	return array[length-1]
}

// get type int value and address
func GetIntVar(str string) []string {
	trim_out := regexp.MustCompile("int") // contains out
	trimmed_out := trim_out.ReplaceAllString(str, "")
	trim_opening := regexp.MustCompile("[=\t]") // match either ( or ' or )
	trimmed := trim_opening.ReplaceAllString(trimmed_out, "")

	var array []string = strings.Split(trimmed, " ")

	return array
}

// every instruction parsed by ';' (semicolon)
func InstructionParser(str string) []string {
	var array []string = strings.Split(str, ";")

	return array
}

func main() {
	ReadFile("output", ".w")
}
