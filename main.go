package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	core "./core"
)

func Path_Format(path string) string {
	if string(path[:2]) == "E:" {
		return path
	}
	var slash string = "/"
	if runtime.GOOS == "windows" {
		slash = "\\"
	}
	var bytemap []byte = []byte(path)
	for i := range bytemap {
		if bytemap[i] == 47 && slash == "\\" {
			bytemap[i] = 92
		} else if bytemap[i] == 92 && slash == "/" {
			bytemap[i] = 47
		}
	}
	path = string(bytemap)
	if string(path[0]) == "." {
		folding := len(Paths)-1
		j := 0
		for string(path[j]) == "." && folding > 1 {
			folding--
			j++
		}
		i := 0
		making_path := ""
		for ; i < folding; i++ {
			making_path = making_path + string(Paths[i]) + slash
		}
		return making_path + path[j:]
	} else {
		making_path := ""
		for i := 0; i < len(Paths)-1; i++ {
			making_path = making_path + string(Paths[i]) + slash
		}
		return making_path + path
	}
}

func xcute(xcutable *core.Xcutable, api *core.API, chan_completed chan<- bool) {
	var stderr bytes.Buffer
	// This function provides an execution of and connection with
	// xcutables.
	path := Path_Format(xcutable.REF)
	cmd := exec.Command("go", "run", path)
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

var Paths []string
func Path_Parse(given_path string) {
	var charbuf []byte
	if runtime.GOOS == "windows" {
		Paths = append(Paths, string(given_path[0:2]))
		given_path = given_path[3:]
	}
	for i := 0; i < len(given_path); i++ {
		if string(given_path[i]) == "/" ||
			string(given_path[i]) == "\\" {
			Paths = append(Paths, string(charbuf))
			charbuf = []byte{}
		} else {
			charbuf = append(charbuf, given_path[i])
		}
	}
	Paths = append(Paths, string(charbuf))
}

func main() {
	// // // // // // // // // // // //
	//fmt.Println(TotalMemory())     //
	//fmt.Println(ProcessMemory())   //
	// // // // // // // // // // // //

	// MUST DO PATH PROPAGATION HERE
	Path_Parse(os.Args[1])
	fmt.Println("FILE:", os.Args[1])

	i := time.Now()
	page, tree, _, spoterr := core.SPOT(os.Args[1])
	if spoterr != nil {
		fmt.Print(spoterr.Error())
	}
	api := core.API{Page: page, Tree: &tree}
	fmt.Println("Spawned in", time.Since(i).Nanoseconds(), "nanoseconds")
	//core.PrintPOT(api.Page, *api.Tree)
	//fmt.Println(objmap)

	chan_comm := make(chan bool)
	go xcute_man(page.XCUTABLES, &api, chan_comm)
	<-chan_comm
	fmt.Println("Program took", time.Since(i).Milliseconds(), "milliseconds")
	
	// ADD EVERYTHING ON PREBUILD
	core.PrintPOT(api.Page, *api.Tree)
	fmt.Println(core.ObjMap)
	// fmt.Scanln()
}
