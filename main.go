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

	// var int_type_variable []string
	// var int_type_value []string

	length_of_syntax := len(InstructionParser(text))
	// for i := 0; i < length_of_syntax; i++ {
	// 	output := TrimOutFunction((InstructionParser(text)[i]))

	// 	// output function detector
	// 	if strings.Contains(InstructionParser(text)[i], "out") && strings.Contains(InstructionParser(text)[i], "('") && strings.Contains(InstructionParser(text)[i], "')") {
	// 		fmt.Print(output, "\n")
	// 	}

	// 	// type integer detector
	// 	if strings.Contains(InstructionParser(text)[i], "int") {
	// 		int_type_variable = append(int_type_variable, GetIntVar(output)[1])
	// 		int_type_value = append(int_type_value, TrimIntType(output))
	// 	}

	// 	// output integer detector
	// 	if strings.Contains(InstructionParser(text)[i], "out") && !strings.Contains(InstructionParser(text)[i], "'") {
	// 		// fmt.Println(strings.Contains(output, "a"))
	// 		for find := 0; find < len(int_type_variable); find++ {
	// 			if strings.Contains(int_type_variable[find], string(output[1])) { //bug (output)
	// 				fmt.Println(int_type_value[find])
	// 			}
	// 		}
	// 	}

	// }
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
