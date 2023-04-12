package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	core "./core"
)

//const open string = "./"

func xcute(xcutable *core.Xcutable, api *core.API, chan_completed chan<- bool) {
	var stderr bytes.Buffer
	// This function provides an execution of and connection with
	// xcutables.
	cmd := exec.Command("go", "run", xcutable.REF)
	fmt.Println(xcutable.REF)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		params := strings.Split(scanner.Text(), "\x10")
		if params[0] == "0x5" {
			api.PrintStructure()
		}
		if params[0] == "0x24" {
			api.NewElement(params[1:])
		}
		if params[0] == "0x25" {
			api.ElementsChangeID(params[1:])
		}
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(stderr.String())
		fmt.Println(err)
		panic(err)
	}
	chan_completed <- true
}

func xcute_man(xcutables []*core.Xcutable, api *core.API, chan_comm chan<- bool) {
	for x := range xcutables {
		chan_completed := make(chan bool)
		go xcute(xcutables[x], api, chan_completed)
		<-chan_completed
	}
	chan_comm <- true
}

func main() {
	// // // // // // // // // // // //
	//fmt.Println(TotalMemory())     //
	//fmt.Println(ProcessMemory())   //
	// // // // // // // // // // // //

	// MUST DO PATH PROPAGATION

	i := time.Now()
	page, tree, _, spoterr := core.SPOT("E:/GoBro/test_project/testpage.ego")
	if spoterr != nil {
		fmt.Print(spoterr.Error())
	}
	api := core.API{Page: page, Tree: &tree}
	fmt.Println("Spawned in", time.Since(i).Nanoseconds(), "nanoseconds")
	core.PrintPOT(api.Page, *api.Tree)
	//fmt.Println(objmap)

	chan_comm := make(chan bool)
	go xcute_man(page.XCUTABLES, &api, chan_comm)
	<-chan_comm
	fmt.Println("Program took", time.Since(i).Milliseconds(), "milliseconds")
	core.PrintPOT(api.Page, *api.Tree)
	fmt.Println(core.ObjMap)
}
