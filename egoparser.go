package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Keychars map[byte]string = map[byte]string{
	60:  "<",
	62:  ">",
	47:  "/",
	34:  "\"",
	61:  "=",
	101: "e",
	120: "x",
	115: "s",
	114: "r",
	102: "f",
	105: "i",
	100: "d",
	99:  "c",
	108: "l",
	97:  "a",
	32:  "\x20",
	35:  "#",
}

var Patterns map[string]string = map[string]string{
	"style":      Keychars[60] + Keychars[115] + Keychars[32],
	"xcutable":   Keychars[60] + Keychars[120] + Keychars[32],
	"elementNoP": Keychars[60] + Keychars[101] + Keychars[62],
	"elementWiP": Keychars[60] + Keychars[101] + Keychars[32],
	"elementCLS": Keychars[60] + Keychars[47] + Keychars[62],
	"quotation":  Keychars[34],
	"propEND":    Keychars[62],
	"ref":        Keychars[114] + Keychars[101] + Keychars[102],
	"id":         Keychars[105] + Keychars[100],
	"class":      Keychars[99] + Keychars[108] + Keychars[97] + Keychars[115] + Keychars[115],
	"comment":    Keychars[35] + Keychars[35] + Keychars[35],
}

func SPOT(egofile string) (*page, []*element) {
	ego, err := os.Open(egofile)
	if err != nil {
		log.Fatalln(err)
		return nil, nil
	}
	defer ego.Close()

	page := NewPage()
	var parent parent
	var child_buffer, tree_buffer, parent_buffer []*element

	scanner := bufio.NewScanner(ego)
	scanner.Bytes()

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return nil, nil
		}

		line := scanner.Text()

		if len(line) <= 2 {
			continue
		}

		for i := range line {
			if i >= len(line)-2 {
				continue
			} else if string(line[i:i+3]) == Patterns["comment"] {
				// comment is one-liner
				continue
			}
			if string(line[i:i+3]) == Patterns["style"] &&
				string(line[i+3:i+6]) == Patterns["ref"] &&
				string(line[i+7]) == Patterns["quotation"] {
				// style inclusion is one-liner
				charbuf := []byte{}
				for j := 8; string(line[j]) != Patterns["quotation"]; j++ {
					charbuf = append(charbuf, line[j])
				}
				style := NewStyle(string(charbuf))
				page.AppendStyle(style)
				continue
			}
			if string(line[i:i+3]) == Patterns["xcutable"] &&
				string(line[i+3:i+6]) == Patterns["ref"] &&
				string(line[i+7]) == Patterns["quotation"] {
				// xcutable inclusion is one-liner
				charbuf := []byte{}
				for j := 8; string(line[j]) != Patterns["quotation"]; j++ {
					charbuf = append(charbuf, line[j])
				}
				x := NewX(string(charbuf))
				page.AppendX(x)
				continue
			}
			if string(line[i:i+3]) == Patterns["elementNoP"] {
				parent = page
				if len(parent_buffer) >= 1 {
					// we choose between element and a page (nil) as a parent
					parent = parent_buffer[len(parent_buffer)-1]
				}
				new_element := NewElement("std", "", parent, child_buffer, "")
				parent_buffer = append(parent_buffer, new_element)
				if i >= len(line)-3 {
					continue
				}
				i += 2
			}
			if string(line[i:i+3]) == Patterns["elementCLS"] {
				if len(parent_buffer) == 0 {
					continue
				} else if len(parent_buffer) == 1 {
					parent_buffer[0].ChangeParent(page)
					tree_buffer = append(tree_buffer, parent_buffer[0])
					parent_buffer = nil
				} else if len(parent_buffer) >= 2 {
					parent_buffer[len(parent_buffer)-1].ChangeParent(parent_buffer[len(parent_buffer)-2])
					parent_buffer[len(parent_buffer)-2].AppendChild(parent_buffer[len(parent_buffer)-1])
					parent_buffer = parent_buffer[:len(parent_buffer)-1]
				}
				if i >= len(line)-3 {
					continue
				}
				i += 2
			}
			if string(line[i:i+3]) == Patterns["elementWiP"] {
				// something seems not right. can make it more compact
				parent = page
				if i >= len(line)-3 {
					continue
				}
				i += 3
				var charbuf_id, charbuf_class, charbuf_ref []byte
				for string(line[i]) != Patterns["propEND"] {
					if string(line[i:i+2]) == Patterns["id"] {
						if string(line[i+3]) == Patterns["quotation"] {
							i += 4
							for ; string(line[i]) != Patterns["quotation"]; i++ {
								charbuf_id = append(charbuf_id, line[i])
							}
							i++
						}
						if string(line[i]) == Patterns["propEND"] {
							break
						} else {
							i++
						}
					}
					if string(line[i:i+5]) == Patterns["class"] {
						if string(line[i+6]) == Patterns["quotation"] {
							i += 7
							for ; string(line[i]) != Patterns["quotation"]; i++ {
								charbuf_class = append(charbuf_class, line[i])
							}
							i++
						}
						if string(line[i]) == Patterns["propEND"] {
							break
						} else {
							i++
						}
					}
					if string(line[i:i+3]) == Patterns["ref"] {
						if string(line[i+4]) == Patterns["quotation"] {
							i += 5
							for ; string(line[i]) != Patterns["quotation"]; i++ {
								charbuf_ref = append(charbuf_ref, line[i])
							}
							i++
						}
						if string(line[i]) == Patterns["propEND"] {
							break
						} else {
							i++
						}
					}
				}
				if len(charbuf_id) == 0 {
					charbuf_id = []byte("std")
				}
				if len(parent_buffer) >= 1 {
					parent = parent_buffer[len(parent_buffer)-1]
				}
				new_element := NewElement(string(charbuf_id), string(charbuf_class), parent, child_buffer, string(charbuf_ref))
				parent_buffer = append(parent_buffer, new_element)
			}
		}
	}
	for len(parent_buffer) > 0 {
		if len(parent_buffer) == 1 {
			parent_buffer[0].ChangeParent(page)
			tree_buffer = append(tree_buffer, parent_buffer[0])
			parent_buffer = nil
		} else if len(parent_buffer) >= 2 {
			parent_buffer[len(parent_buffer)-1].ChangeParent(parent_buffer[len(parent_buffer)-2])
			parent_buffer[len(parent_buffer)-2].AppendChild(parent_buffer[len(parent_buffer)-1])
			parent_buffer = parent_buffer[:len(parent_buffer)-1]
		}
	}
	if page != nil {
		return page, tree_buffer
	}
	log.Fatal("Page does not exist")
	return nil, nil
}

func PrintPOT(page parent, tree []*element) {
	fmt.Println("page        :", string(MakeTree_inJSON(page)))
	fmt.Println("tree length :", len(tree))
	for el := range tree {
		fmt.Println("element", el+1, "  :", string(MakeTree_inJSON(tree[el])))
	}
}
