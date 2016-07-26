package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	value []interface{}
	child []*Node
}

func (n *Node) node_print(k int) (s string) {
	var tmp string
	tmp = ""
	for i := 0; i < len(n.child); i++ {
		tmp = fmt.Sprintf("%s%s", tmp, n.child[i].node_print(k+1))
	}
	if tmp != "" {
		s = fmt.Sprintf("(%s\n%s)", n.value, tmp)
	} else {
		s = fmt.Sprintf("(%s%s)\n", n.value, tmp)
	}
	for i := 0; i < k; i++ {
		s = fmt.Sprintf("\t%s", s)
	}

	return s
}
func (n *Node) node_find(key interface{}) (out []*Node) {
	if n.value[0] == key {
		out = append(out, n)
	} else {
		for i := 0; i < len(n.child); i++ {
			out = append(out, n.child[i].node_find(key)...)
		}
	}
	return
}

type stack struct {
	s []interface{}
}

func newStack() *stack {
	return &stack{make([]interface{}, 0)}
}
func (s *stack) Push(v interface{}) {
	s.s = append(s.s, v)
}

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
func (s *stack) top() (size int) {
	return len(s.s)
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
	st := newStack()
	var tree *Node
	tree = &Node{make([]interface{}, 0), nil}
	for len(data) > 0 {
		token, data = get_tok(data)
		switch token.(type) {
		case nil: //comment
			//fmt.Println("NIL:", token)
		case error:
			break
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

	fmt.Println(tree.node_print(0))
	fmt.Println(tree.node_find("cos")[0].node_print(0))
}
