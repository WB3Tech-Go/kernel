package strings

import "testing"

func TestCleanStringEmptyLeading(t *testing.T) {
	inputString :="  hello there!"
	outputString := CleanString(inputString)

	if outputString != "hello there!" { t.Error("'  hello there!' did not default to 'hello there!'; leading spaces not truncated")}
}

func TestCleanStringEmptyTrailing(t *testing.T) {
	inputString :="hello there!  "
	outputString := CleanString(inputString)

	if outputString != "hello there!" { t.Error("'hello there!  ' did not default to 'hello there!'; leading spaces not truncated")}
}

func TestCleanStringEmptyLeadingAndTrailing(t *testing.T) {
	inputString :="  hello there!  "
	outputString := CleanString(inputString)

	if outputString != "hello there!" { t.Error("'  hello there!  ' did not default to 'hello there!'; leading spaces not truncated")}
}

func TestCleanStringEmptyLeadingAndTrailingWithSpaceInMiddle(t *testing.T) {
	inputString :="  hello \n there!  "
	outputString := CleanString(inputString)

	if outputString != "hello \n there!" { t.Error("'  hello \n there!  ' did not default to 'hello there!'; leading spaces not truncated")}
}

func TestCleanStringLeadingLineBreak(t *testing.T) {
	inputString :="\nhello there!"
	outputString := CleanString(inputString)

	if outputString != "hello there!" { t.Error("'\nhello there!' did not default to 'hello there!'; leading spaces not truncated")}
}

func TestCleanStringTrailingLineBreak(t *testing.T) {
	inputString :="hello there!\n"
	outputString := CleanString(inputString)

	if outputString != "hello there!" { t.Error("'hello there!\n' did not default to 'hello there!'; leading spaces not truncated")}
}
