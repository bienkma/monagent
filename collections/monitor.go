package collections

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// Function get information interfaces server
func Network(Interfaces string) (InfoNet net.IOCountersStat) {
	v, err := net.IOCounters(true)
	if err != nil {
		Log(err)
		panic(err)
	}
	// almost every return value is a struct
	for i := 0; i < len(v); i++ {
		if v[i].Name == Interfaces {
			InfoNet = v[i]
		}
	}
	return
}

// Function get information Bandwidth
func Bandwidth(Interval uint, Interfaces string) (rx, tx float64) {
	// Check rx, tx is fist time
	t1 := time.Now()
	last_rx := Network(Interfaces).BytesRecv
	last_tx := Network(Interfaces).BytesSent

	// check rx, tx is second time
	time.Sleep(Interval * time.Second)
	t2 := time.Now()
	now_rx := Network(Interfaces).BytesRecv
	now_tx := Network(Interfaces).BytesSent

	// rx, tx = (now_rx - last_rx) * 8 / delta_time(t2-t1)
	delta_time := t2.Sub(t1).Nanoseconds()
	rx = (float64(now_rx-last_rx) * 8) / (float64(delta_time) / 1000000000)
	tx = (float64(now_tx-last_tx) * 8) / (float64(delta_time) / 1000000000)

	return

}

// Function get information Memory
func Memory() (uint64, uint64, uint64) {
	InfoMem, err := mem.VirtualMemory()
	if err != nil {
		Log(err)
		panic(err)
	}
	return InfoMem.Total, InfoMem.Used, InfoMem.Cached
}

// Function get information CPU
func CPU(Interval uint) []float64 {
	CPUPercent, Err := cpu.Percent(Interval*time.Second, false)
	if Err != nil {
		Log(Err)
		panic(Err)
	}
	return CPUPercent
}
