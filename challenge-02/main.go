package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Stats struct {
	CPU float64 `json:cpu`
	Disk float64 `json:disk`
	Memory float64 `json:memory`
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	p,_ := cpu.Percent(time.Second, false)
	m,_ := mem.VirtualMemory()
	d,_ := disk.Usage("/")
	s := Stats{CPU: p[cpu.CPUser], Disk: d.UsedPercent, Memory: m.UsedPercent}
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/stats", handleStats)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
