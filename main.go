package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

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

func xcute(xcutable *xcutable, chan_completed chan<- bool) {
	// This function provides an execution of and connection with
	// xcutables.
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
	// // // // // 
	fmt.Println(TotalMemory())
	fmt.Println(ProcessMemory())
	// // // // //
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
		// fmt.Println(ProcessMemory())
		// if ProcessMemory() > 6 {
		// 	return
		// }
	}
}
