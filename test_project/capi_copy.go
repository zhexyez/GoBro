package main

// // // // // // // // // // // // // // // //
// THIS IS JUST A COPY OF CLIENT API (CAPI)  //
// TO USE IN FUTURE                          //
// // // // // // // // // // // // // // // //

import (
	capi "./capi"
)

// // Code from here to main() must be in separate file
// type Lets struct {
// }

// func (l *Lets) NewElement(id string, class string, ref string, txt string) {
// 	fmt.Println("0x24\x10"+id+"\x10"+class+"\x10"+ref+"\x10"+txt)
// }

// func (l *Lets) PrintStructure() {
// 	fmt.Println("0x5")
// }

// func (l *Lets) ElementsChangeID(old_id string, new_id string) {
// 	fmt.Println("0x25\x10"+old_id+"\x10"+new_id)
// }

func main() {
	l := capi.Lets{}
	//l.PrintStructure()
	//time.Sleep(time.Second * 3)
	for x:=0; x < 1000000; x++ {
		l.NewElement("d","","","no way")
	}
	for x:=0; x < 1000000; x++ {
		l.NewElement("way","more","params","inside this text")
	}
	for x:=0; x < 1000000; x++ {
		l.NewElement("tobe","presented","with","really amount of text so the file can be tested on latency and other problems")
	}

	l.ElementsChangeID("garry", "bob")
}