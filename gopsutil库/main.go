package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	// for {
	percent, _ := cpu.Percent(0, false)
	fmt.Printf("cpu percent:%T\n", percent)
	// }
}

func main() {
	getCpuInfo()
}
