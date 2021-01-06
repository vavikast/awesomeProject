package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"runtime"
)

func main() {

	fmt.Println(*cpu.ProcInfo())
}
