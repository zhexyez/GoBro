package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
)

type pageorparent interface {
	// Dummy interface to link page and element structs via interface
	method()
}

type page struct {
	// Dummy struct to link the top-level element
}

func (p *page) method() {
	// Dummy method to link page and element structs via interface
}

func NewPage() *page {
	// Returns pointer to a new page struct
	page := page{}
	return &page
}

type element struct {
	// An element definition.
	//
	// ID and Class links element with the style
	// and used in selection of elements
	//
	// PARENT points the element to its parent element.
	// The top-level element belongs to the page struct,
	// so pass nil on it.
	// Each element !!!MUST!!! have parent, otherwise crash
	//
	// CHILD holds a slice of pointers to its child elements.
	// The element without childs holds nil
	//
	// REF holds reference link. Can be used as HREF in HTML
	ID     string       `json:"ID"`
	CLASS  string       `json:"CLASS"`
	PARENT pageorparent `json:"-"`
	CHILD  []*element   `json:"CHILD"`
	REF    string       `json:"REF"`
}

func (e *element) method() {
	// Dummy method to link page and element structs via interface
}

func (e *element) ChangeID(id string) {
	e.ID = id
}

func (e *element) ChangeClass(class string) {
	e.CLASS = class
}

func (e *element) ChangeParent(parent *element) {
	e.PARENT = parent
}

func (e *element) AppendChild(child *element) {
	e.CHILD = append(e.CHILD, child)
}

func (e *element) ChangeRef(ref string) {
	e.REF = ref
}

func (e *element) element_reset() {
	// This method is called by the Flush method.
	// Its purpose is to go recursively on each
	// child and remove it (assign to nil),
	// then it explicitly runs garbage collector
	// so the memory is freed at the time.
	e.ID = ""
	e.CLASS = ""
	e.PARENT = nil
	e.REF = ""
	if len(e.CHILD) == 0 {
		return
	}
	for i := range e.CHILD {
		e.CHILD[i].element_reset()
		e.CHILD[i] = nil
		runtime.GC()
	}
	return
}

func FlushElement(element *element) *element {
	// This method resets an element and calls
	// garbage collector to remove it from the memory.
	element.element_reset()
	element = nil
	runtime.GC()
	return element
}

func NewElement(id string, class string, parent pageorparent, child []*element, ref string) *element {
	// Returns pointer to a new element struct
	element := element{id, class, parent, child, ref}
	return &element
}

func MakeTree_inJSON(element *element) []byte {
	// Can be used for visual representation
	tree, err := json.Marshal(element)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return tree
}

func main() {
	fmt.Println("to be >")

	// test #1> creating
	page := NewPage()
	elem := NewElement("baby", "boom", page, nil, "")
	elem2 := NewElement("baby2", "boom2", elem, nil, "")
	elem3 := NewElement("baby3", "boom3", elem, nil, "")
	elem.AppendChild(elem2)
	elem.AppendChild(elem3)

	// test #2> changing
	elem.ChangeClass("boogie")
	elem.ChangeID("rasta")

	// test #3> printing info
	fmt.Println("element>", elem)
	fmt.Println(elem.PARENT)
	fmt.Println(elem.CHILD)
	fmt.Println("element2>", elem2)
	fmt.Println("element3>", elem3)

	// test #4> printing in JSON format
	fmt.Println(string(MakeTree_inJSON(elem)))

	// test #5> removing
	elem = FlushElement(elem)
	fmt.Println("element>", elem)
	fmt.Println("element2>", elem2)
	fmt.Println("element3>", elem3)
}
