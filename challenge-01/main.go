package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	p, _ := cpu.Percent(time.Second, false)
	fmt.Printf("CPU:\t\t%.2f%%\n", p[cpu.CPUser])
	m,_ := mem.VirtualMemory()
	fmt.Printf("Mermory:\t%.2f%%\n", m.UsedPercent)
	d,_ := disk.Usage("/")
	fmt.Printf("Disk: \t\t%.2f%%\n", d.UsedPercent)
}
