package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This file contains is the core to create
// POT (Page Object Tree). TODO: TL;DR

// Hashmap of objects. Gets updated with any change
var ObjMap map[string][]*Element = make(map[string][]*Element)

var keychars map[byte]string = map[byte]string{
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
	9:   "\x09",
}

// Can be extended to contain "context" keyword
var Patterns map[string]string = map[string]string{
	"style":      keychars[60] + keychars[115] + keychars[32],
	"xcutable":   keychars[60] + keychars[120] + keychars[32],
	"elementNoP": keychars[60] + keychars[101] + keychars[62],
	"elementWiP": keychars[60] + keychars[101] + keychars[32],
	"elementCLS": keychars[60] + keychars[47] + keychars[62],
	"quotation":  keychars[34],
	"propEND":    keychars[62],
	"ref":        keychars[114] + keychars[101] + keychars[102],
	"id":         keychars[105] + keychars[100],
	"class":      keychars[99] + keychars[108] + keychars[97] + keychars[115] + keychars[115],
	"comment":    keychars[35] + keychars[35] + keychars[35],
}

type POTError struct {
	TXT string
}

func (e *POTError) Error() string {
	return e.TXT
}

// This function reads a property of an element with parameters
func checkElementWiP(line *string, i *int, n int, charbuf *[]byte, EOL *bool, propend *bool) {
	if string((*line)[(*i)+n]) == Patterns["quotation"] {
		(*i) += (n+1)
		for ; string((*line)[(*i)]) != Patterns["quotation"]; (*i)++ {
			(*charbuf) = append((*charbuf), (*line)[(*i)])
		}
		(*i)++
	}
	if string((*line)[(*i)]) == Patterns["propEND"] {
		if (*i) == len((*line))-1 {
			(*EOL) = true
			(*propend) = true
			return
		} else {
			(*i)++
			(*propend) = true
			return
		}
	} else {
		(*i)++
		(*propend) = true
		return
	}
}

// This function creates new element and adds it to the buffer
func makeNewElement(parent_buffer *[]*Element, txt_buffer *[][]byte, charbuf_id *[]byte, charbuf_class *[]byte, parent *Parent, child_buffer *[]*Element, charbuf_ref *[]byte) {
	(*txt_buffer) = append((*txt_buffer), []byte{})
	new_element := NewElement(string((*charbuf_id)), string((*charbuf_class)), *parent, *child_buffer, string((*charbuf_ref)), "")
	ObjMap[new_element.ID] = append(ObjMap[new_element.ID], new_element)
	(*parent_buffer) = append((*parent_buffer), new_element)
}

// This function reads the ref field of styles and executables
func fillRef(line *string, i *int) []byte {
	if (*line)[(*i)+3:(*i)+6] == Patterns["ref"] && string((*line)[(*i)+7]) == Patterns["quotation"] {
		charbuf := []byte{}
		for j := 8; string((*line)[j]) != Patterns["quotation"]; j++ {
			charbuf = append(charbuf, (*line)[j])
		}
		return charbuf
	} else {
		return []byte{}
	}
}

// This function updates tree_buffer
func editTreeBuffer(parent_buffer *[]*Element, tree_buffer *[]*Element, txt_buffer *[][]byte, page *Page) {
	if len((*parent_buffer)) == 1 {
		if len((*txt_buffer)) > 0 {
			(*parent_buffer)[0].AppendTXT((*txt_buffer)[0])
			(*txt_buffer) = [][]byte{}
		}
		(*parent_buffer)[0].ChangeParent(page)
		(*tree_buffer) = append((*tree_buffer), (*parent_buffer)[0])
		(*parent_buffer) = nil
	} else if len((*parent_buffer)) >= 2 {
		if len((*txt_buffer)) > 0 {
			(*parent_buffer)[len((*parent_buffer))-1].AppendTXT((*txt_buffer)[len((*txt_buffer))-1])
			(*txt_buffer) = (*txt_buffer)[:len((*txt_buffer))-1]
		}
		(*parent_buffer)[len((*parent_buffer))-1].ChangeParent((*parent_buffer)[len((*parent_buffer))-2])
		(*parent_buffer)[len((*parent_buffer))-2].AppendChild((*parent_buffer)[len((*parent_buffer))-1])
		(*parent_buffer) = (*parent_buffer)[:len((*parent_buffer))-1]
	}
}

// This function spawns the POT (Page Object Tree)
// and returns pointers as well as hashes
func SPOT(egofile string) (*Page, []*Element, *map[string][]*Element, *POTError) {
	ego, err := os.Open(egofile)
	if err != nil {
		return nil, nil, nil, &POTError{TXT: "Unable to open the file. Possible error:" + err.Error()}
	}
	defer ego.Close()

	// Create instance of page and intermediate buffers
	page := NewPage()
	var parent Parent
	var child_buffer, tree_buffer, parent_buffer []*Element
	var charbuf_id, charbuf_class, charbuf_ref []byte
	var txt_buffer [][]byte
	var charbuf []byte

	var EOL bool = false

	scanner := bufio.NewScanner(ego)
	scanner.Bytes()

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, nil, nil, &POTError{TXT: "Unable to scan the file. Possible error:" + err.Error()}
		}

		line := scanner.Text()

		// Remove all zeros from the beginning of the line
		zeros_counter := 0
		for line[zeros_counter] == 32 || line[zeros_counter] == 9 {
			zeros_counter++
		}
		line = line[zeros_counter:]

		if len(line) <= 2 {
			continue
		}

		for i := 0; i < len(line); {
			if i < len(line)-2 {
				if line[i:i+3] == Patterns["comment"] {
					// Comment is one-liner. We don't care about saving comments
					break
				}
				// style and xcutable inclusions are one-liners
				if line[i:i+3] == Patterns["style"] {
					charbuf = fillRef(&line, &i)
					style := NewStyle(string(charbuf))
					page.AppendStyle(style)
					break
				}
				if line[i:i+3] == Patterns["xcutable"] {
					charbuf = fillRef(&line, &i)
					x := NewX(string(charbuf))
					page.AppendX(x)
					break
				}
				if line[i:i+3] == Patterns["elementNoP"] {
					parent = page
					if len(parent_buffer) >= 1 {
						parent = parent_buffer[len(parent_buffer)-1]
					}
					makeNewElement(&parent_buffer, &txt_buffer, &charbuf_id, &charbuf_class, &parent, &child_buffer, &charbuf_ref)
					charbuf_id, charbuf_class, charbuf_ref = []byte{}, []byte{}, []byte{}
					if i >= len(line)-3 {
						break
					}
					i += 3
				}
				if line[i:i+3] == Patterns["elementCLS"] {
					if len(parent_buffer) == 0 {
						break
					} else {
						editTreeBuffer(&parent_buffer, &tree_buffer, &txt_buffer, page)
					}
					if i >= len(line)-3 {
						break
					}
					i += 3
					continue
				}
				if line[i:i+3] == Patterns["elementWiP"] {
					// something seems not right. can make it more compact
					if i >= len(line)-3 {
						break
					}
					parent = page
					i += 3
					// var charbuf_id, charbuf_class, charbuf_ref []byte
					var occurencies_check [3]byte = [3]byte{0,0,0}
					var propend bool = false
					for !propend {
						if line[i:i+2] == Patterns["id"] && occurencies_check[0] == 0 {
							checkElementWiP(&line, &i, 3, &charbuf_id, &EOL, &propend)
							occurencies_check[0]++
						} else if line[i:i+5] == Patterns["class"] && occurencies_check[1] == 0 {
							checkElementWiP(&line, &i, 6, &charbuf_class, &EOL, &propend)
							occurencies_check[0]++
						} else if line[i:i+3] == Patterns["ref"] && occurencies_check[2] == 0 {
							checkElementWiP(&line, &i, 4, &charbuf_class, &EOL, &propend)
							occurencies_check[0]++
						}
					}
					if len(charbuf_id) == 0 {
						charbuf_id = []byte("std")
					}
					if len(parent_buffer) >= 1 {
						parent = parent_buffer[len(parent_buffer)-1]
					}
					makeNewElement(&parent_buffer, &txt_buffer, &charbuf_id, &charbuf_class, &parent, &child_buffer, &charbuf_ref)
					charbuf_id, charbuf_class, charbuf_ref = []byte{}, []byte{}, []byte{}
				}
			}
			if len(parent_buffer) == 0 {
				break
			}
			if len(txt_buffer) == 0 {
				txt_buffer = append(txt_buffer, []byte{})
			}
			if EOL {
				EOL = false
				break
			}
			txt_buffer[len(txt_buffer)-1] = append(txt_buffer[len(txt_buffer)-1], line[i])
			i++
		}
	}
	for len(parent_buffer) > 0 {
		editTreeBuffer(&parent_buffer, &tree_buffer, &txt_buffer, page)
	}
	if page != nil {
		return page, tree_buffer, &ObjMap, nil
	}
	log.Fatal("Page does not exist")
	return nil, nil, nil, &POTError{TXT: "Page does not exist!"}
}

// This function prints the structure of POT
// to the stdout
func PrintPOT(page Parent, tree []*Element) {
	fmt.Println("page        :", string(MakeTree_inJSON(page)))
	fmt.Println("tree length :", len(tree))
	for el := range tree {
		fmt.Println("element", el+1, "  :", string(MakeTree_inJSON(tree[el])))
	}
}
