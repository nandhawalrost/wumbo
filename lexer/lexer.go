package lexer

import (
	"regexp"
	"strings"
)

func InstructionLexer(str string) []string {
	var array []string = strings.Split(str, ";")

	return array
}

func GetIntVar(str string) []string {
	trim_out := regexp.MustCompile("int") // contains out
	trimmed_out := trim_out.ReplaceAllString(str, "")
	trim_opening := regexp.MustCompile("[=\t]") // match either ( or ' or )
	trimmed := trim_opening.ReplaceAllString(trimmed_out, "")

	var array []string = strings.Split(trimmed, " ")

	return array
}

func GetIntType(str string) string {
	trim_out := regexp.MustCompile("int") // contains out
	trimmed_out := trim_out.ReplaceAllString(str, "")
	trim_opening := regexp.MustCompile("[=\t]") // match either ( or ' or )
	trimmed := trim_opening.ReplaceAllString(trimmed_out, "")

	var array []string = strings.Split(trimmed, " ")

	length := len(array)
	return array[length-1]
}
