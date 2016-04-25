package collections

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// Function get information interfaces server
func Network(Interfaces string) net.IOCountersStat {
	v, _ := net.IOCounters(true)
	var InfoNet net.IOCountersStat
	// almost every return value is a struct
	for i := 0; i < len(v); i++ {
		if v[i].Name == Interfaces {
			InfoNet = v[i]
		}
	}
	return InfoNet
}

// Function get information Memory
func Memory() (uint64, uint64, uint64) {
	InfoMem, _ := mem.VirtualMemory()
	return InfoMem.Used, InfoMem.Cached, InfoMem.Total
}

func CPU() []float64 {
	var Err error
	CPUPercent, Err := cpu.Percent(1*time.Second, false)

	if Err != nil {
		fmt.Printf("Can't read CPU percent!..")
	}

	return CPUPercent
}

func Disk() {
	var Err error
	if Err != nil {
		fmt.Printf("Can't read partiton")
	}

	Disk, Err := disk.Usage("/")

	fmt.Printf("Partiton / Total: %v\n", Disk.Total)
	fmt.Printf("Partition / Free: %v\n", Disk.Free)
}
