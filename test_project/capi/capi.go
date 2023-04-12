package capi

import "fmt"

// Code from here to main() must be in separate file
type Lets struct {
}

func (l *Lets) NewElement(id string, class string, ref string, txt string) {
	fmt.Println("0x24\x10"+id+"\x10"+class+"\x10"+ref+"\x10"+txt)
}

func (l *Lets) PrintStructure() {
	fmt.Println("0x5")
}

func (l *Lets) ElementsChangeID(old_id string, new_id string) {
	fmt.Println("0x25\x10"+old_id+"\x10"+new_id)
}