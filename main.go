package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"
)

// This function provides an execution of and connection with
// xcutables.
func xcute(xcutable *xcutable, chan_completed chan<- bool) {
	cmd := exec.Command("go", "run", xcutable.REF)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "you are cutie" {
			fmt.Println("thanks")
		}
		if scanner.Text() == "and me?" {
			fmt.Println("you too!")
		}
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	chan_completed <- true
}

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
	page, tree, spoterr := SPOT("testpage.ego")
	if spoterr != nil {
		fmt.Print(spoterr.Error())
	}
	fmt.Println("Spawned in", time.Since(i).Nanoseconds(), "nanoseconds")
	PrintPOT(page, tree)

	// xcutables are executed 1 by 1 in mentioned order
	for x := range page.XCUTABLES {
		chan_completed := make(chan bool)
		go xcute(page.XCUTABLES[x], chan_completed)
		<-chan_completed
	}
}
