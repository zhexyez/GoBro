package main

import (
	"encoding/json"
	"log"
	"runtime"
)

/*  This file contains interface to interact with
POT (Page Object Tree). TODO: TL;DR
*/

type parent interface {
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
	return &page{}
}

type style struct {
	REF string
}

func NewStyle(ref string) *style {
	// Returns pointer to a new style struct
	return &style{REF: ref}
}

func ChangeStyle(s *style, ref string) {
	s.REF = ref
}

func FlushStyle(s *style) *style {
	return nil
}

type xcutable struct {
	REF string
}

func NewX(ref string) *xcutable {
	// Returns pointer to a new xcutable struct
	return &xcutable{REF: ref}
}

func ChangeX(x *xcutable, ref string) {
	x.REF = ref
}

func FlushX(x *xcutable) *xcutable {
	return nil
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
	ID     string     `json:"ID"`
	CLS    string     `json:"CLASS"`
	PARENT parent     `json:"-"`
	CHILD  []*element `json:"CHILD"`
	REF    string     `json:"REF"`
}

func (e *element) method() {
	// Dummy method to link page and element structs via interface
}

func (e *element) ChangeID(id string) {
	e.ID = id
}

func (e *element) ChangeClass(class string) {
	e.CLS = class
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
	// This method is called by the FlushElement method.
	// Its purpose is to go recursively on each
	// child and remove it (assign to nil).
	e.ID = ""
	e.CLS = ""
	e.PARENT = nil
	e.REF = ""
	if len(e.CHILD) == 0 {
		return
	}
	for i := range e.CHILD {
		e.CHILD[i].element_reset()
		e.CHILD[i] = nil
	}
}

func (e *element) element_reset_keep_children(newparent parent) {
	// This method is called by the Flush method.
	// Its purpose is to change the pointers of
	// first-dimension childs and change their
	// parent to newparent.
	e.ID = ""
	e.CLS = ""
	e.PARENT = nil
	e.REF = ""
	if len(e.CHILD) == 0 {
		return
	}
	for i := range e.CHILD {
		e.CHILD[i].PARENT = newparent
	}
}

func FlushElement(element *element) *element {
	defer runtime.GC()
	// This method resets an element and calls
	// garbage collector to remove it from the memory.
	element.element_reset()
	return nil
}

func FlushElementKeepChildren(element *element, newparent parent) *element {
	defer runtime.GC()
	// This method resets an element and calls
	// garbage collector to remove it from the memory.
	// Also it gives the element's children new parent.
	element.element_reset_keep_children(newparent)
	return nil
}

func NewElement(id string, class string, parent parent, child []*element, ref string) *element {
	// Returns pointer to a new element struct
	return &element{ID: id, CLS: class, PARENT: parent, CHILD: child, REF: ref}
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
