package main

import "fmt"

type Formatter interface {
	Format(text string) string
}
type FormatEntity struct {
	str    string
	isBold bool
	isCode bool
	isItal bool
}

func main() {
	var format FormatEntity
	format.str = "Hello go"

	format.Code()
	format.Bold()
	format.Ital()

	fmt.Println(format.Format())
}

func (f *FormatEntity) Format() string {
	if f.isBold {
		f.str = fmt.Sprintf("**%s**", f.str)
	}
	if f.isCode {
		f.str = fmt.Sprintf("`%s`", f.str)
	}
	if f.isItal {
		f.str = fmt.Sprintf("_%s_", f.str)
	}
	return f.str
}

func (f *FormatEntity) Bold() {
	f.isBold = true
}

func (f *FormatEntity) Code() {
	f.isCode = true
}

func (f *FormatEntity) Ital() {
	f.isItal = true
}
