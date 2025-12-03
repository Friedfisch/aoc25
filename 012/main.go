package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func loadInput() []int {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex) // this is bad
	data, err := os.ReadFile(path.Join(exPath, "input_test.txt"))
	//data, err := os.ReadFile("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	var r []int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		n, _ := strconv.Atoi(line[1:])
		switch line[0] {
		case 'L':
			r = append(r, -n)

		case 'R':
			r = append(r, n)

		default:
			continue
		}
	}
	return r
}

func main() {
	data := loadInput()
	pos := 50
	cnt := 0

	for _, v := range data {
		//op := pos
		pos += v
		for pos >= 100 {
			pos -= 100
		}
		for pos < 0 {
			pos += 100
		}
		if pos == 0 {
			cnt++
		}
		log.Println(v, pos, cnt)

		// 5395 too low for 1.2
	}
	log.Println("res:", cnt)
}
