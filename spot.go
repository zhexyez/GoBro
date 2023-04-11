package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var ObjMap map[string][]*element = make(map[string][]*element)

var charbuf_id, charbuf_class, charbuf_ref []byte

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

// // // // // // // // //
//  SECTION DEPRECATED	//
// // // // // // // // //

// txt_buffer = append(txt_buffer, []byte{})
// new_element := NewElement("std", "", parent, child_buffer, "", "")
// ObjMap[new_element.ID] = append(ObjMap[new_element.ID], new_element)
// parent_buffer = append(parent_buffer, new_element)

// txt_buffer = append(txt_buffer, []byte{})
// new_element := NewElement(string(charbuf_id), string(charbuf_class), parent, child_buffer, string(charbuf_ref), "")
// ObjMap[new_element.ID] = append(ObjMap[new_element.ID], new_element)
// parent_buffer = append(parent_buffer, new_element)

// func makehash(element *element, objmap map[string][]*element) {
// 	objmap[element.ID] = append(objmap[element.ID], element)
// 	if len(element.CHILD) > 0 {
// 		for el := range element.CHILD {
// 			makehash(element.CHILD[el], objmap)
// 		}
// 	}
// }

// func pothash(tree []*element, objmap map[string][]*element) {
// 	for el := range tree {
// 		makehash(tree[el], objmap)
// 	}
// }

// if line[i:i+2] == Patterns["id"] {
// 	if string(line[i+3]) == Patterns["quotation"] {
// 		i += 4
// 		for ; string(line[i]) != Patterns["quotation"]; i++ {
// 			charbuf_id = append(charbuf_id, line[i])
// 		}
// 		i++
// 	}
// 	if string(line[i]) == Patterns["propEND"] {
// 		if i == len(line)-1 {
// 			EOL = true
// 			break
// 		} else {
// 			i++
// 			break
// 		}
// 	} else {
// 		i++
// 	}

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

func makeNewElement(parent_buffer *[]*element, txt_buffer *[][]byte, charbuf_id *[]byte, charbuf_class *[]byte, parent *parent, child_buffer *[]*element, charbuf_ref *[]byte) {
	(*txt_buffer) = append((*txt_buffer), []byte{})
	new_element := NewElement(string((*charbuf_id)), string((*charbuf_class)), *parent, *child_buffer, string((*charbuf_ref)), "")
	ObjMap[new_element.ID] = append(ObjMap[new_element.ID], new_element)
	(*parent_buffer) = append((*parent_buffer), new_element)
}

func checkRefStyle() {

}

func SPOT(egofile string) (*page, []*element, *map[string][]*element, *POTError) {
	ego, err := os.Open(egofile)
	if err != nil {
		return nil, nil, nil, &POTError{TXT: "Unable to open the file. Possible error:" + err.Error()}
	}
	defer ego.Close()

	// Create instance of page and intermediate buffers
	page := NewPage()
	var parent parent
	var child_buffer, tree_buffer, parent_buffer []*element
	var txt_buffer [][]byte

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
				// TODO: make one function. They are almost the same
				checkRefStyle()
				if line[i:i+3] == Patterns["style"] &&
					line[i+3:i+6] == Patterns["ref"] &&
					string(line[i+7]) == Patterns["quotation"] {
					// style inclusion is one-liner
					charbuf := []byte{}
					for j := 8; string(line[j]) != Patterns["quotation"]; j++ {
						charbuf = append(charbuf, line[j])
					}
					style := NewStyle(string(charbuf))
					page.AppendStyle(style)
					break
				}
				if line[i:i+3] == Patterns["xcutable"] &&
					line[i+3:i+6] == Patterns["ref"] &&
					string(line[i+7]) == Patterns["quotation"] {
					// xcutable inclusion is one-liner
					charbuf := []byte{}
					for j := 8; string(line[j]) != Patterns["quotation"]; j++ {
						charbuf = append(charbuf, line[j])
					}
					x := NewX(string(charbuf))
					page.AppendX(x)
					break
				}
				if line[i:i+3] == Patterns["elementNoP"] {
					parent = page
					if len(parent_buffer) >= 1 {
						// We choose between element and a page (nil) as a parent
						parent = parent_buffer[len(parent_buffer)-1]
					}
					makeNewElement(&parent_buffer, &txt_buffer, &charbuf_id, &charbuf_class, &parent, &child_buffer, &charbuf_ref)
					charbuf_id, charbuf_class, charbuf_ref = []byte{}, []byte{}, []byte{}
					if i >= len(line)-3 {
						break
					}
					i += 3
				}
				// TODO: rewrite so it makes one function with the last
				// before page != nil statement
				if line[i:i+3] == Patterns["elementCLS"] {
					if len(parent_buffer) == 0 {
						break
					} else if len(parent_buffer) == 1 {
						if len(txt_buffer) > 0 {
							parent_buffer[0].AppendTXT(txt_buffer[0])
							txt_buffer = [][]byte{}
						}
						parent_buffer[0].ChangeParent(page)
						tree_buffer = append(tree_buffer, parent_buffer[0])
						parent_buffer = nil
					} else if len(parent_buffer) >= 2 {
						if len(txt_buffer) > 0 {
							parent_buffer[len(parent_buffer)-1].AppendTXT(txt_buffer[len(txt_buffer)-1])
							txt_buffer = txt_buffer[:len(txt_buffer)-1]
						}
						parent_buffer[len(parent_buffer)-1].ChangeParent(parent_buffer[len(parent_buffer)-2])
						parent_buffer[len(parent_buffer)-2].AppendChild(parent_buffer[len(parent_buffer)-1])
						parent_buffer = parent_buffer[:len(parent_buffer)-1]
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
		if len(parent_buffer) == 1 {
			if len(txt_buffer) > 0 {
				parent_buffer[0].AppendTXT(txt_buffer[0])
				txt_buffer = [][]byte{}
			}
			parent_buffer[0].ChangeParent(page)
			tree_buffer = append(tree_buffer, parent_buffer[0])
			parent_buffer = nil
		} else if len(parent_buffer) >= 2 {
			if len(txt_buffer) > 0 {
				parent_buffer[len(parent_buffer)-1].AppendTXT(txt_buffer[len(txt_buffer)-1])
				txt_buffer = txt_buffer[:len(txt_buffer)-1]
			}
			parent_buffer[len(parent_buffer)-1].ChangeParent(parent_buffer[len(parent_buffer)-2])
			parent_buffer[len(parent_buffer)-2].AppendChild(parent_buffer[len(parent_buffer)-1])
			parent_buffer = parent_buffer[:len(parent_buffer)-1]
		}
	}
	if page != nil {
		return page, tree_buffer, &ObjMap, nil
	}
	log.Fatal("Page does not exist")
	return nil, nil, nil, &POTError{TXT: "Page does not exist!"}
}

func PrintPOT(page parent, tree []*element) {
	fmt.Println("page        :", string(MakeTree_inJSON(page)))
	fmt.Println("tree length :", len(tree))
	for el := range tree {
		fmt.Println("element", el+1, "  :", string(MakeTree_inJSON(tree[el])))
	}
}
