package monitor

import "github.com/shirou/gopsutil/mem"

func getRam() (stat *mem.VirtualMemoryStat, err error) {
	stat, err = mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	return
}
