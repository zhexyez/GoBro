package main

import (
	"fmt"
	"time"
)

func main() {
	// test #1
	/*fmt.Println("to be >")

	// test #1.1> creating
	page := NewPage()
	elem := NewElement("baby", "boom", page, nil, "")
	elem2 := NewElement("baby2", "boom2", elem, nil, "")
	elem3 := NewElement("baby3", "boom3", elem, nil, "")
	elem.AppendChild(elem2)
	elem.AppendChild(elem3)

	// test #1.2> changing
	elem.ChangeClass("boogie")
	elem.ChangeID("rasta")

	// test #1.3> printing info
	fmt.Println("element>", elem)
	fmt.Println(elem.PARENT)
	fmt.Println(elem.CHILD)
	fmt.Println("element2>", elem2)
	fmt.Println("element3>", elem3)

	// test #1.4> printing in JSON format
	fmt.Println(string(MakeTree_inJSON(elem)))

	// test #1.5> removing keeping children
	elem = FlushElementKeepChildren(elem, page)
	fmt.Println("element>", elem)
	fmt.Println("element2>", elem2)
	fmt.Println("element3>", elem3) */

	i := time.Now()
	page, tree := SPOT("testpage.ego")
	fmt.Println("Spawned in", time.Since(i).Nanoseconds(), "nanoseconds")
	PrintPOT(page, tree)

}
