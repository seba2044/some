package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

//Tree
type Node struct {
	value []interface{}
	child []*Node
}

//Print tree
func (n *Node) Print(k int) (s string) {
	var tmp string
	tmp = ""
	for i := 0; i < len(n.child); i++ {
		tmp = fmt.Sprintf("%s%s", tmp, n.child[i].Print(k+1))
	}
	if tmp != "" {
		s = fmt.Sprintf("(%s\n%s", n.value, tmp)
		for i := 0; i < k; i++ {
			s = fmt.Sprintf("%s\t", s)
		}
		s = fmt.Sprintf("%s)\n", s)

	} else {
		s = fmt.Sprintf("(%s%s)\n", n.value, tmp)
	}
	for i := 0; i < k; i++ {
		s = fmt.Sprintf("\t%s", s)
	}

	return s
}

//Find tree in key
func (n *Node) Find(key interface{}) (out []*Node) {
	if n.value[0] == key {
		out = append(out, n)
	} else {
		for i := 0; i < len(n.child); i++ {
			out = append(out, n.child[i].Find(key)...)
		}
	}
	return
}

//Prser stack
type stack struct {
	s []interface{}
}

//Stack constructor
func newStack() *stack {
	return &stack{make([]interface{}, 0)}
}

//push value on stack
func (s *stack) Push(v interface{}) {
	s.s = append(s.s, v)
}

//Pop value from stack
func (s *stack) Pop() (out interface{}) {
	l := len(s.s)
	if l == 0 {
		out = errors.New("Empty stack")
	} else {
		out = s.s[l-1]
		s.s = s.s[0 : l-1]
	}
	return
}

//Return index of top value
func (s *stack) top() (size int) {
	return len(s.s)
}

//Open file as str
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

//Parse data, return tree
func sExp_Parse(data string) *Node {

	var token interface{}
	st := newStack()
	var tree *Node
	tree = &Node{make([]interface{}, 0), nil}
	for len(data) > 0 {
		token, data = get_tok(data)
		switch token.(type) {
		case nil: //comment
			//fmt.Println("NIL:", token)
		case error:
			return tree
		case string:
			//fmt.Println("NAPIS:", token)
			if token == "(" {
				st.Push(token)
				node := &Node{make([]interface{}, 0), nil}
				st.Push(node)
				tree.child = append(tree.child, node)
				tree = node
			} else if token == ")" {
				st.Pop()
				st.Pop()
				if (len(st.s)-1) > 0 && st.s[len(st.s)-1] != nil {
					tmp_node, isNode := st.s[len(st.s)-1].(*Node)
					if isNode {
						tree = tmp_node
					}
				}
			} else {
				if (len(st.s)-1) > 0 && st.s[len(st.s)-1] != nil {
					//how to append
					st.s[len(st.s)-1].(*Node).value = append(st.s[len(st.s)-1].(*Node).value, token)
				}
			}
		case int:
			//fmt.Println("INT:", token)
			st.s[len(st.s)-1].(*Node).value = append(st.s[len(st.s)-1].(*Node).value, token)
			//st.Push(token)
		case float64:
			//fmt.Println("F64:", token)
			st.s[len(st.s)-1].(*Node).value = append(st.s[len(st.s)-1].(*Node).value, token)
			//st.Push(token)
		case float32:
			//fmt.Println("F32:", token)
			st.s[len(st.s)-1].(*Node).value = append(st.s[len(st.s)-1].(*Node).value, token)
			//st.Push(token)
		default:
			//fmt.Println("DEF:", token)
			err := errors.New(fmt.Sprintf("Undefined token: %v", token))
			panic(err)
		}
	}
	if st.top() != 0 {
		err := errors.New(fmt.Sprintf("Stack isn't empty, on stack:%s , parse error.", st.top()))
		panic(err)
	}

	return tree
}
func main() {

	data := load_file("key.cfg")
	tree := sExp_Parse(data)
	fmt.Println(tree.Print(0))
	fmt.Println(tree.Find("cos")[0].Print(0))
}
