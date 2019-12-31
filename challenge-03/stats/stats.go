package stats

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type Stats struct {
	CPU float64 `json:cpu`
	Disk float64 `json:disk`
	Memory float64 `json:memory`
}

func FetchStats() Stats {
	p,_ := cpu.Percent(time.Second, false)
	m,_ := mem.VirtualMemory()
	d,_ := disk.Usage("/")
	return Stats{CPU: p[cpu.CPUser], Disk: d.UsedPercent, Memory: m.UsedPercent}
}