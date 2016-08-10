package main

import (
	"bufio"
	"fmt"
	"os"
	//"unicode/utf8"
)

type text_buffer struct {
	name  string
	path  string
	array [][]byte
}

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
func (txt *text_buffer) insert(pos_x, pos_y int, new_text string) {
	var vector []int
	for index, runeValue := range string(txt.array[pos_x]) {
		vector = append(vector, index)
		fmt.Printf("%#U index: %d\n", runeValue, index)
	}
	fmt.Println(vector)
	//check why i need thid fucking ...
	txt.array[pos_x] = append(txt.array[pos_x][:vector[pos_y]], append([]byte(new_text), txt.array[pos_x][vector[pos_y]:]...)...)
	fmt.Println(string(txt.array[pos_x]))

}
func main() {
	fb := text_buffer{}
	fb.load_array("file.txt")
	fb.insert(1, 2, "coąśśŋß©↓ə←πœęźżćs")
	/*for i := 0; i < len(fb.array); i++ {*/
	//fmt.Print(string(fb.array[i]), " len: ", utf8.RuneCount(fb.array[i]), "\n")

	/*}*/
	return
}
