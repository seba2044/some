package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	//"unicode/utf8"
)

type my_byte []byte

// Text structure
type text_buffer struct {
	name  string   //name of buffer, do't know if needed.
	path  string   // path to file.
	array [][]byte //array contain text.
}

// Load array frome file.
func (txt *text_buffer) load_array(path string) {
	txt.path = path
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		txt.array = append(txt.array, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

//Get index of byte array as utf8 string.
func index_of_utf8(utf_string string) (utf_vector []int) {
	for index, _ := range string(utf_string) {
		utf_vector = append(utf_vector, index)
	}
	return
}

// Add line to array befor given index.
func (txt *text_buffer) add_line_before(pos_x int) {
	txt.array = append(txt.array[:pos_x], append([][]byte{make([]byte, 1)}, txt.array[pos_x:]...)...)
}

//Insert data into array line.
func (txt *text_buffer) insert_to_line(pos_x, pos_y int, new_text string) {
	vector := index_of_utf8(string(txt.array[pos_x]))
	txt.array[pos_x] = append(txt.array[pos_x][:vector[pos_y]], append([]byte(new_text), txt.array[pos_x][vector[pos_y]:]...)...)
}

//Remove given line.
func (txt *text_buffer) cut_line(pos_x int) (out []byte) {
	out = append(out, txt.array[pos_x])
	txt.array = append(txt.array[:pos_x], txt.array[pos_x+1:]...)
	return
}

//Cut character from line.
func (txt *text_buffer) cut_from_line(line_n, from, to int) (out []byte) {
	vector := index_of_utf8(string(txt.array[line_n]))
	out = append(out, txt.array[line_n][vector[from]:vector[to]]...)
	txt.array[line_n] = append(txt.array[line_n][:from], txt.array[line_n][to:]...)
	return
}

//Insert text array to txt.array.
func (txt *text_buffer) insert_array(pos_x, pos_y int, new_text string) {
	txt.array = append(txt.array[:pos_x], append(bytes.Split([]byte(new_text), []byte("\n")), txt.array[pos_x:]...)...)
}

//Print line with column number
func (txt *text_buffer) Print() {
	for i := 0; i < len(txt.array); i++ {
		fmt.Println(i, ":", string(txt.array[i]))
	}
}
func main() {
	fb := text_buffer{}
	fb.load_array("file.txt")
	fb.add_line_before(0)
	fb.insert_to_line(0, 0, "a")
	//fmt.Println(fb.array)
	fb.insert_to_line(1, 10, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa&*")
	fmt.Println("sd:", string(fb.cut_from_line(1, 0, 100)))
	fb.insert_array(1, 1, "")
	fb.Print()
	//fmt.Println(fb.array)
	return
}
