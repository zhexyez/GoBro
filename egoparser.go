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

func SPOT(egofile string) {
	ego, err := os.Open(egofile)
	if err != nil {
		log.Fatalln(err)
	}
	defer ego.Close()

	page := NewPage()

	scanner := bufio.NewScanner(ego)
	scanner.Bytes()
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return
		}
		var parent_buffer []*element
		line := scanner.Text()
		if len(line) <= 2 {
			continue
		}
		for i := range line {
			if i >= len(line)-2 {
				continue
			}
			if string(line[i:i+3]) == "###" {
				// comment is one-liner
				continue
			}
			if string(line[i:i+3]) == "<s\x20" {
				// style inclusion is one-liner
				if string(line[i+3:i+6]) == "ref" {
					if string(line[i+7]) == "\"" {
						charbuf := []byte{}
						for j := 8; string(line[j]) != "\""; j++ {
							charbuf = append(charbuf, line[j])
						}
						style := NewStyle(string(charbuf))
						page.AppendStyle(style)
					}
				}
				continue
			}
			if string(line[i:i+3]) == "<x\x20" {
				// xcutable inclusion is one-liner
				if string(line[i+3:i+6]) == "ref" {
					if string(line[i+7]) == "\"" {
						charbuf := []byte{}
						for j := 8; string(line[j]) != "\""; j++ {
							charbuf = append(charbuf, line[j])
						}
						x := NewX(string(charbuf))
						page.AppendX(x)
					}
				}
				continue
			}
			if string(line[i:i+3]) == "<e>" {
				if i >= len(line)-3 {
					continue
				}
				new_element := NewElement(string([]byte{}), string([]byte{}), nil, nil, string([]byte{}))
				parent_buffer = append(parent_buffer, new_element)
				i += 2
			}
			if string(line[i:i+3]) == "</>" {
				if i >= len(line)-3 || len(parent_buffer) == 0 {
					continue
				}
				if len(parent_buffer) == 1 {
					parent_buffer[0].ChangeParent(page)
					//fmt.Println(string(MakeTree_inJSON(parent_buffer[0])))
					fmt.Println(string(MakeTree_inJSON(parent_buffer[0])))
					parent_buffer = nil
				} else {
					parent_buffer[len(parent_buffer)-1].ChangeParent(parent_buffer[len(parent_buffer)-2])
					parent_buffer[len(parent_buffer)-2].AppendChild(parent_buffer[len(parent_buffer)-1])
					parent_buffer = parent_buffer[:len(parent_buffer)-1]
					fmt.Println(parent_buffer)
				}
				i += 2
			}
			if string(line[i:i+3]) == "<e\x20" {
				if i >= len(line)-3 {
					continue
				}
				i += 3
				var charbuf_id, charbuf_class, charbuf_ref []byte
				for string(line[i]) != ">" {
					if string(line[i:i+2]) == "id" {
						if string(line[i+3]) == "\"" {
							i += 4
							for ; string(line[i]) != "\""; i++ {
								charbuf_id = append(charbuf_id, line[i])
							}
							i++
						}
						if string(line[i]) == ">" {
							continue
						} else {
							i++
						}
					}
					if string(line[i:i+3]) == "cls" {
						if string(line[i+4]) == "\"" {
							i += 5
							for ; string(line[i]) != "\""; i++ {
								charbuf_class = append(charbuf_class, line[i])
							}
							i++
						}
						if string(line[i]) == ">" {
							continue
						} else {
							i++
						}
					}
					if string(line[i:i+3]) == "ref" {
						if string(line[i+4]) == "\"" {
							i += 5
							for ; string(line[i]) != "\""; i++ {
								charbuf_ref = append(charbuf_ref, line[i])
							}
							i++
						}
						if string(line[i]) == ">" {
							continue
						} else {
							i++
						}
					}
				}
				//fmt.Println(string(charbuf_id), string(charbuf_class), string(charbuf_ref))
				new_element := NewElement(string(charbuf_id), string(charbuf_class), nil, nil, string(charbuf_ref))
				//fmt.Println(new_element)
				parent_buffer = append(parent_buffer, new_element)
				//fmt.Println(parent_buffer)
			}
		}
	}
	//fmt.Println(string(MakeTree_inJSON(page)))
}
