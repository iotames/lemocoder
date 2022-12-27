package status

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// https://blog.csdn.net/whatday/article/details/109620192

func byteToGB(b uint64) float64 {
	return float64(b / (1024 * 1024 * 1024))
}

type CpuInfo struct {
	UsedPercent float64
	Num         int
	Name        string
}

type DiskPartInfo struct {
	Device, MountPoint string
	UsedPercent        float64
	Total, Free        uint64
}

func (m DiskPartInfo) TotalText() string {
	gb := byteToGB(m.Total)
	return fmt.Sprintf("%.2f GB", gb)
}

func (m DiskPartInfo) FreeText() string {
	gb := byteToGB(m.Free)
	return fmt.Sprintf("%.2f GB", gb)
}

type MemoryInfo struct {
	UsedPercent float64
	Total, Free uint64
}

func (m MemoryInfo) TotalText() string {
	gb := byteToGB(m.Total)
	return fmt.Sprintf("%.2f GB", gb)
}

func (m MemoryInfo) FreeText() string {
	gb := byteToGB(m.Free)
	return fmt.Sprintf("%.2f GB", gb)
}

func GetCpuInfo() CpuInfo {
	cpuNum, _ := cpu.Counts(false)
	cpup, _ := cpu.Percent(time.Second, false)
	cpuPercent := 0.0
	if len(cpup) > 0 {
		cpuPercent = cpup[0]
	}
	cpuInfo, _ := cpu.Info()
	cpuName := ""
	if len(cpuInfo) > 0 {
		cpuName = cpuInfo[0].ModelName
	}
	return CpuInfo{Num: cpuNum, UsedPercent: cpuPercent, Name: cpuName}
}

func GetDiskInfo() []DiskPartInfo {
	var dps []DiskPartInfo
	parts, _ := disk.Partitions(true)
	for _, part := range parts {
		info, _ := disk.Usage(part.Mountpoint)
		dps = append(dps, DiskPartInfo{UsedPercent: info.UsedPercent, Total: info.Total, Free: info.Free, Device: part.Device, MountPoint: part.Mountpoint})
	}
	return dps
}

func GetMemoryInfo() MemoryInfo {
	vm, _ := mem.VirtualMemory()
	return MemoryInfo{UsedPercent: vm.UsedPercent, Total: vm.Total, Free: vm.Free}
}
