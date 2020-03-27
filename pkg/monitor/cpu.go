package monitor

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func getCpu() (stat []cpu.InfoStat, err error) {
	stat, err = cpu.Info()
	if err != nil {
		panic(err)
	}
	return
}

func getCpuPercent() (p []float64, err error) {
	p, err = cpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}
	return
}
