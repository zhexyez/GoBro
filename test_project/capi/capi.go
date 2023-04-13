package capi

import "fmt"

func NewElement(id string, class string, ref string, txt string) {
	fmt.Println("0x24\x10"+id+"\x10"+class+"\x10"+ref+"\x10"+txt)
}

func PrintStructure() {
	fmt.Println("0x5")
}

func ElementsChangeID(old_id string, new_id string) {
	fmt.Println("0x25\x10"+old_id+"\x10"+new_id)
}

func CapiLog(s ...string) {
	var printbuf string
	for i := range s {
		printbuf = printbuf + s[i] + "\x20"
	}
	fmt.Println("0xF9\x10" + printbuf)
}