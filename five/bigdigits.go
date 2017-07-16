package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{
		"   000  ",
		"  0   0 ",
		" 0     0 ",
		"0       0",
		"0       0",
		"0       0",
		" 0     0 ",
		"  0   0 ",
		"   000  ",
	},
	{"1", "11", "111"},
	{"2", "22", "222"},
	{"3", "33", "333"},
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("usage: %s \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	// fmt.Println("~~~", os.Args)
	// fmt.Println("$$$", stringOfDigits)
	for row := range bigDigits[0] {
		line := ""

		for column := range stringOfDigits {

			// fmt.Println("...", stringOfDigits[column])
			// fmt.Println("!!!", stringOfDigits[column]-'0')
			digit := stringOfDigits[column] - '0' //??
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
