package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const open string = "../"

type MEMORYSTATUSEX struct {
    dwLength                uint32
    dwMemoryLoad            uint32
    ullTotalPhys            uint64
    ullAvailPhys            uint64
    ullTotalPageFile        uint64
    ullAvailPageFile        uint64
    ullTotalVirtual         uint64
    ullAvailVirtual         uint64
    ullAvailExtendedVirtual uint64
}

type PROCESS_MEMORY_COUNTERS struct {
    cb                         uint32
    PageFaultCount             uint32
    PeakWorkingSetSize         uint64
    WorkingSetSize             uint64
    QuotaPeakPagedPoolUsage    uint64
    QuotaPagedPoolUsage        uint64
    QuotaPeakNonPagedPoolUsage uint64
    QuotaNonPagedPoolUsage     uint64
    PagefileUsage              uint64
    PeakPagefileUsage          uint64
}

func TotalMemory () uint64 {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
    globalMemoryStatusEx := kernel32.NewProc("GlobalMemoryStatusEx")

    var memInfo MEMORYSTATUSEX
    memInfo.dwLength = uint32(unsafe.Sizeof(memInfo))

    ret, _, err := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
    if ret == 0 {
        panic(fmt.Sprintf("Call to GlobalMemoryStatusEx failed: %v", err))
    }

    // fmt.Printf("Total memory: %d GB\n", memInfo.ullTotalPhys / 1073741824)
	return memInfo.ullTotalPhys
}

func ProcessMemory() uint64 {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
    psapi := syscall.NewLazyDLL("psapi.dll")

    getCurrentProcess := kernel32.NewProc("GetCurrentProcess")
    getProcessMemoryInfo := psapi.NewProc("GetProcessMemoryInfo")

    var memCounters PROCESS_MEMORY_COUNTERS

    currentProcess, _, _ := getCurrentProcess.Call()
    ret, _, err := getProcessMemoryInfo.Call(currentProcess, uintptr(unsafe.Pointer(&memCounters)), uintptr(unsafe.Sizeof(memCounters)))
    if ret == 0 {
        panic(fmt.Sprintf("Call to GetProcessMemoryInfo failed: %v", err))
    }

    //fmt.Printf("Working set size: %d MB\n", memCounters.WorkingSetSize / 1048576)
	return memCounters.WorkingSetSize
}

func xcute(xcutable *xcutable, api *API, chan_completed chan<- bool) {
	// This function provides an execution of and connection with
	// xcutables.
	cmd := exec.Command("go", "run", open+xcutable.REF)
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
		panic(err)
	}
	chan_completed <- true
}

func xcute_man(xcutables []*xcutable, api *API, chan_comm chan<- bool) {
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

	i := time.Now()
	page, tree, _, spoterr := SPOT("testpage.ego")
	if spoterr != nil {
		fmt.Print(spoterr.Error())
	}
	api := API{Page: page, Tree: &tree}
	fmt.Println("Spawned in", time.Since(i).Nanoseconds(), "nanoseconds")

	//fmt.Println(objmap)

	chan_comm := make(chan bool)
	go xcute_man(page.XCUTABLES, &api, chan_comm)
	<-chan_comm
	fmt.Println("Program took", time.Since(i).Milliseconds(), "milliseconds")
	PrintPOT(api.Page, *api.Tree)
	fmt.Println(ObjMap)
}
