package core

import (
	"encoding/json"
	"log"
	"runtime"
)

/*  This file contains interface to interact with
POT (Page Object Tree). TODO: TL;DR
*/

type Parent interface {
	// Dummy interface to link Page and Element structs via interface
	method()
}

type Page struct {
	// Root Element in the tree
	STYLES    []*Style
	XCUTABLES []*Xcutable
}

func (p *Page) method() {
	// Dummy method to link Page and Element structs via interface
}

func (p *Page) AppendStyle(s *Style) {
	p.STYLES = append(p.STYLES, s)
}

func (p *Page) AppendX(x *Xcutable) {
	p.XCUTABLES = append(p.XCUTABLES, x)
}

func NewPage() *Page {
	// Returns pointer to a new Page struct
	return &Page{}
}

type Style struct {
	REF string
}

func NewStyle(ref string) *Style {
	// Returns pointer to a new style struct
	return &Style{REF: ref}
}

func (s *Style) ChangeStyle(ref string) {
	s.REF = ref
}

func FlushStyle(s *Style) *Style {
	return nil
}

type Xcutable struct {
	REF string
}

func NewX(ref string) *Xcutable {
	// Returns pointer to a new xcutable struct
	return &Xcutable{REF: ref}
}

func (x *Xcutable) ChangeX(ref string) {
	x.REF = ref
}

func FlushX(x *Xcutable) *Xcutable {
	return nil
}

type Element struct {
	// An Element definition.
	//
	// ID and Class links Element with the style
	// and used in selection of Elements
	//
	// Parent points the Element to its Parent Element.
	// The top-level Element belongs to the Page struct,
	// so pass nil on it.
	// Each Element !!!MUST!!! have Parent, otherwise crash
	//
	// CHILD holds a slice of pointers to its child Elements.
	// The Element without childs holds nil
	//
	// REF holds reference link. Can be used as HREF in HTML
	ID     string     `json:"ID"`
	CLS    string     `json:"CLASS"`
	PARENT Parent     `json:"-"`
	CHILD  []*Element `json:"CHILD"`
	REF    string     `json:"REF"`
	TXT    string     `json:"text"`
}

func (e *Element) method() {
	// Dummy method to link Page and Element structs via interface
}

func (e *Element) ChangeID(id string) {
	e.ID = id
}

func (e *Element) ChangeClass(class string) {
	e.CLS = class
}

func (e *Element) ChangeParent(Parent Parent) {
	e.PARENT = Parent
}

func (e *Element) AppendChild(child *Element) {
	e.CHILD = append(e.CHILD, child)
}

func (e *Element) ChangeRef(ref string) {
	e.REF = ref
}

func (e *Element) ChangeTXT(txt []byte) {
	e.TXT = string(txt)
}

func (e *Element) AppendTXT(txt []byte) {
	e.TXT += string(txt)
}

func (e *Element) RemoveOneFromEnd() {
	e.TXT = e.TXT[:len(e.TXT)-1]
}

func (e *Element) element_reset() {
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

func (e *Element) element_reset_keep_children(newParent Parent) {
	// This method is called by the Flush method.
	// Its purpose is to change the pointers of
	// first-dimension childs and change their
	// Parent to newParent.
	e.ID = ""
	e.CLS = ""
	e.PARENT = nil
	e.REF = ""
	if len(e.CHILD) == 0 {
		return
	}
	for i := range e.CHILD {
		e.CHILD[i].PARENT = newParent
	}
}

func FlushElement(Element *Element) *Element {
	defer runtime.GC()
	// This method resets an Element and calls
	// garbage collector to remove it from the memory.
	Element.element_reset()
	return nil
}

func FlushElementKeepChildren(Element *Element, newParent Parent) *Element {
	defer runtime.GC()
	// This method resets an element and calls
	// garbage collector to remove it from the memory.
	// Also it gives the Element's children new Parent.
	Element.element_reset_keep_children(newParent)
	return nil
}

func NewElement(id string, class string, Parent Parent, child []*Element, ref string, txt string) *Element {
	// Returns pointer to a new Element struct
	return &Element{ID: id, CLS: class, PARENT: Parent, CHILD: child, REF: ref, TXT: txt}
}

func MakeTree_inJSON(Element Parent) []byte {
	// Can be used for visual representation
	tree, err := json.Marshal(Element)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return tree
}
