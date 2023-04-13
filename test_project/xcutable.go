package main

import (
	"time"

	capi "./capi"
)

func main() {
	//l := capi.Lets{}
	capi.PrintStructure()
	time.Sleep(time.Second * 3)
	// for x:=0; x < 1000000; x++ {
	// 	capi.NewElement("d","","","no way")
	// }
	// for x:=0; x < 1000000; x++ {
	// 	capi.NewElement("way","more","params","inside this text")
	// }
	// for x:=0; x < 1000000; x++ {
	// 	capi.NewElement("tobe","presented","with","really amount of text so the file can be tested on latency and other problems")
	// }
	capi.ElementsChangeID("tobe", "bob")
	//capi.PrintStructure()
	//capi.CapiLog("Hi there!")
}