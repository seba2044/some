package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type stack []interface{}

func (s stack) Push(v interface{}) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, interface{}) {
	l := len(s)
	if l == 0 {
		return nil, nil
	} else {
		return s[:l-1], s[l-1]
	}
}

//
func load_file(path string) (out string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	out = string(data)
	return
}

//Get token to parse str
func get_tok(s string) (tok interface{}, remainder string) {
	s = strings.TrimSpace(s)
	if s == "" {
		tok, remainder = errors.New("Empty string."), ""
	} else {
		switch s[0] {
		case ';': //commnets
			if i := strings.Index(s[1:], "\n"); i >= 0 {
				tok, remainder = nil, s[i+1:]
			}

		case '(', ')':
			tok, remainder = s[0:1], s[1:]
		case '"': //str with spaces
			if i := strings.Index(s[1:], "\""); i >= 0 {
				tok, remainder = s[1:i+1], s[i+2:]
			} else {
				tok, remainder = errors.New("Can't find end of string."), ""
			}
		default:
			i := 0
			//get data
			for i < len(s) && s[i] != '(' && s[i] != ')' && s[i] != '"' && !unicode.IsSpace(rune(s[i])) {
				i++
			}
			//check data type
			if is_int, err := strconv.Atoi(s[:i]); err == nil { //int
				tok, remainder = is_int, s[i:]
			} else if is_float, err := strconv.ParseFloat(s[:i], 64); err == nil { //float
				tok, remainder = is_float, s[i:]
			} else {
				tok, remainder = s[:i], s[i:] //string
			}
		}
	}
	return //errors.New("Undefined error."), ""
}
func main() {

	var token interface{}
	data := load_file("key.cfg")
	for len(data) > 0 {
		token, data = get_tok(data)
		switch token.(type) {
		case nil:
			fmt.Println("NIL:", token)
		case error:
			break
		case string:
			fmt.Println("NAPIS:", token)
		case int:
			fmt.Println("INT:", token)
		case float64:
			fmt.Println("F64:", token)
		case float32:
			fmt.Println("F32:", token)
		default:
			fmt.Println("DEF:", token)
		}
	}

}
