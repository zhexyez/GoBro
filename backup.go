package main

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

// if line[i:i+3] == Patterns["style"] &&
// 	line[i+3:i+6] == Patterns["ref"] &&
// 	string(line[i+7]) == Patterns["quotation"] {
// 	// style inclusion is one-liner
// 	charbuf := []byte{}
// 	for j := 8; string(line[j]) != Patterns["quotation"]; j++ {
// 		charbuf = append(charbuf, line[j])
// 	}
// 	style := NewStyle(string(charbuf))
// 	page.AppendStyle(style)
// 	break
// }
// if line[i:i+3] == Patterns["xcutable"] &&
// 	line[i+3:i+6] == Patterns["ref"] &&
// 	string(line[i+7]) == Patterns["quotation"] {
// 	// xcutable inclusion is one-liner
// 	charbuf := []byte{}
// 	for j := 8; string(line[j]) != Patterns["quotation"]; j++ {
// 		charbuf = append(charbuf, line[j])
// 	}
// 	x := NewX(string(charbuf))
// 	page.AppendX(x)
// 	break
// }

// if len(parent_buffer) == 1 {
// 	if len(txt_buffer) > 0 {
// 		parent_buffer[0].AppendTXT(txt_buffer[0])
// 		txt_buffer = [][]byte{}
// 	}
// 	parent_buffer[0].ChangeParent(page)
// 	tree_buffer = append(tree_buffer, parent_buffer[0])
// 	parent_buffer = nil
// } else if len(parent_buffer) >= 2 {
// 	if len(txt_buffer) > 0 {
// 		parent_buffer[len(parent_buffer)-1].AppendTXT(txt_buffer[len(txt_buffer)-1])
// 		txt_buffer = txt_buffer[:len(txt_buffer)-1]
// 	}
// 	parent_buffer[len(parent_buffer)-1].ChangeParent(parent_buffer[len(parent_buffer)-2])
// 	parent_buffer[len(parent_buffer)-2].AppendChild(parent_buffer[len(parent_buffer)-1])
// 	parent_buffer = parent_buffer[:len(parent_buffer)-1]
// }
