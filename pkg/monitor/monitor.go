package monitor

import (
	"fmt"
	"math"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/viper"
)

type SystemInfo struct {
	CPU         []cpu.InfoStat
	CPU_PERCENT []float64
	RAM         *mem.VirtualMemoryStat
	RAM_PERCENT float64
}

func Get() SystemInfo {
	cpu, _ := getCpu()
	cpu_percent, _ := getCpuPercent()
	ram, _ := getRam()
	ram_percent := ram.UsedPercent
	d := SystemInfo{
		CPU:         cpu,
		CPU_PERCENT: cpu_percent,
		RAM:         ram,
		RAM_PERCENT: ram_percent,
	}
	log(d)
	return d
}

func log(d SystemInfo) {
	fmt.Println(d.CPU)
	for i, v := range d.CPU_PERCENT {
		fmt.Printf("✅ CPU %v : %v\r\n", i, v)
	}
	fmt.Println(d.RAM)
	fmt.Println("✅ RAM_PERCENT ", d.RAM_PERCENT)
}

func Check(d SystemInfo) (message string, status bool) {
	cpuRisk := viper.GetViper().GetInt(`risk.cpu.percent`)
	ramRisk := viper.GetViper().GetInt(`risk.ram.percent`)
	ip := viper.GetViper().GetString(`host.ip`)
	hostName := viper.GetViper().GetString(`host.name`)
	message = "\r\n"
	message += fmt.Sprintf("💻 Host Name : %v\r\n", hostName)
	message += fmt.Sprintf("💻 Host IP : %v\r\n", ip)
	message += fmt.Sprintf("🔧 SETTING RISK CPU :%v%%\r\n", cpuRisk)
	message += fmt.Sprintf("🔧 SETTING RISK RAM :%v%%\r\n", ramRisk)
	for i, v := range d.CPU_PERCENT {
		message += fmt.Sprintf("♨️  CPU%v : %v%% \r\n", i, math.Floor(v))

		if v > float64(cpuRisk) {
			status = true
		}
	}
	message += fmt.Sprintf("♨️  RAM : %vGB\r\n", d.RAM.Total/1024/1024/1024)
	message += fmt.Sprintf("♨️  USED RAM : %v%%", math.Floor(d.RAM_PERCENT))

	if d.RAM_PERCENT > float64(ramRisk) {
		status = true
	}
	return
}
