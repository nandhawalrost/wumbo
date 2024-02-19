package detector

import (
	"fmt"
	"regexp"
	"strings"
	"wlang/lexer"
)

func TrimOutFunction(str string) string {
	trim_out := regexp.MustCompile("(out)") // contains out
	trimmed_out := trim_out.ReplaceAllString(str, "")
	trim_opening := regexp.MustCompile("[(')\n]") // match either ( or ' or )
	trimmed := trim_opening.ReplaceAllString(trimmed_out, "")

	return trimmed
}

func DetectLine(text string, length int, instruction []string) {
	var int_type_variable []string
	var int_type_value []string

	instruction_lexer := lexer.InstructionLexer(text)
	length_of_syntax := len(instruction_lexer)
	for i := 0; i < length_of_syntax; i++ {
		output := TrimOutFunction((instruction_lexer[i]))

		// output function detector
		if strings.Contains(instruction_lexer[i], "out") && strings.Contains(instruction_lexer[i], "('") && strings.Contains(instruction_lexer[i], "')") {
			fmt.Print(output, "\n")
		}

		// type integer detector
		if strings.Contains(instruction_lexer[i], "int") {
			int_type_variable = append(int_type_variable, lexer.GetIntVar(output)[1])
			int_type_value = append(int_type_value, lexer.GetIntType(output))
		}

		// output integer detector
		if strings.Contains(instruction_lexer[i], "out") && !strings.Contains(instruction_lexer[i], "'") {
			// fmt.Println(strings.Contains(output, "a"))
			for find := 0; find < len(int_type_variable); find++ {
				if strings.Contains(int_type_variable[find], string(output[1])) { //bug (output)
					fmt.Println(int_type_value[find])
				}
			}
		}

	}
}
