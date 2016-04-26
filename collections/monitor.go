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
func Network(Interfaces string) (InfoNet net.IOCountersStat) {
	v, _ := net.IOCounters(true)
	// almost every return value is a struct
	for i := 0; i < len(v); i++ {
		if v[i].Name == Interfaces {
			InfoNet = v[i]
		}
	}
	return
}

func Bandwidth(Interfaces string) (rx, tx float64) {
	// Check rx, tx is fist time
	t1 := time.Now()
	last_rx := Network(Interfaces).BytesRecv
	last_tx := Network(Interfaces).BytesSent

	// check rx, tx is second time
	time.Sleep(1 * time.Second)
	t2 := time.Now()
	now_rx := Network(Interfaces).BytesRecv
	now_tx := Network(Interfaces).BytesSent

	// rx, tx = (now_rx - last_rx) * 8 / delta_time(t2-t1)
	delta_time := t2.Sub(t1).Nanoseconds()
	rx = float64(now_rx-last_rx) * 8 / float64(delta_time) / 1000000000
	tx = float64(now_tx-last_tx) * 8 / float64(delta_time) / 1000000000

	return

}

// Function get information Memory
func Memory() (uint64, uint64, uint64) {
	InfoMem, _ := mem.VirtualMemory()
	return InfoMem.Total, InfoMem.Used, InfoMem.Cached
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