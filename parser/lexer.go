package parser

import (
	"io/ioutil"
	"ioutil"
	"log"
)

type Lexer struct {
	File    string
	content string
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func (l *Lexer) Lex() ([]Token, error) {
	l.content = readFile(l.File)
	for i := 0; i < len(l.content); i++ {

	}
}
