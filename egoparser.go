package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*var keychars map[byte]byte = map[byte]byte{
	1: 60,  // "<"
	2: 62,  // ">"
	3: 47,  // /
	4: 34,  // \
	5: 61,  // =
	6: 101, // e
	7: 120, // x
	8: 115, // s
	9: 32,  // \x20 (space)
}*/

func ParsEGO(egofile string) {
	ego, err := os.Open(egofile)
	if err != nil {
		log.Fatalln(err)
	}
	defer ego.Close()

	scanner := bufio.NewScanner(ego)
	scanner.Bytes()
	lnnum := 1
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("LINE NUMBER:", lnnum)
		line := scanner.Text()
		if len(line) <= 2 {
			lnnum++ // line cannot be less than 3 characters in size
			continue
		}
		for i := range line {
			if i >= len(line)-2 {
				continue
			}
			if string(line[i:i+3]) == "###" {
				// comment is one-liner
				fmt.Println("Found comment line")
				continue
			}
			if string(line[i:i+3]) == "<s\x20" {
				// style inclusion is one-liner
				fmt.Println("Found style inclusion")
				continue
			}
			if string(line[i:i+3]) == "<x\x20" {
				// xcutable inclusion is one-liner
				fmt.Println("Found xcutable inclusion")
				continue
			}
			if string(line[i:i+3]) == "<e>" {
				fmt.Println("Found <e>", string(line[i:i+3]))
				if i >= len(line)-3 {
					continue
				}
				i += 2
			}
			if string(line[i:i+3]) == "</>" {
				fmt.Println("Found </>", string(line[i:i+3]))
				if i >= len(line)-3 {
					continue
				}
				i += 2
			}
			if string(line[i:i+3]) == "<e\x20" {
				fmt.Println("Found element with details", string(line[i:i+3]))
				if i >= len(line)-3 {
					continue
				}
				i += 2
			}
		}
		lnnum++
	}
}
