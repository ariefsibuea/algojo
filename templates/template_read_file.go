package template

import (
	"bufio"
	"fmt"
	"os"
)

func temp_main_read_file() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
