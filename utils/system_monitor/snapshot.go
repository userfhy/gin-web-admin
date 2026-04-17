package system_monitor

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

type CPUSnapshot struct {
	Cores      int       `json:"cores"`
	ModelName  string    `json:"modelName"`
	Mhz        float64   `json:"mhz"`
	Usage      float64   `json:"usage"`
	CoreUsages []float64 `json:"coreUsages"`
	Load1      float64   `json:"load1"`
	Load5      float64   `json:"load5"`
	Load15     float64   `json:"load15"`
}

type MemorySnapshot struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	Available   uint64  `json:"available"`
	Cached      uint64  `json:"cached"`
	Buffers     uint64  `json:"buffers"`
	UsedPercent float64 `json:"usedPercent"`
}

type SwapSnapshot struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type DiskPartitionSnapshot struct {
	Path        string  `json:"path"`
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type NetworkSnapshot struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytesSent"`
	BytesRecv   uint64 `json:"bytesRecv"`
	PacketsSent uint64 `json:"packetsSent"`
	PacketsRecv uint64 `json:"packetsRecv"`
	Errin       uint64 `json:"errin"`
	Errout      uint64 `json:"errout"`
	Dropin      uint64 `json:"dropin"`
	Dropout     uint64 `json:"dropout"`
}

type GoRuntimeSnapshot struct {
	Goroutines uint64 `json:"goroutines"`
	HeapAlloc  uint64 `json:"heapAlloc"`
	HeapSys    uint64 `json:"heapSys"`
	HeapIdle   uint64 `json:"heapIdle"`
	HeapInuse  uint64 `json:"heapInuse"`
	StackInuse uint64 `json:"stackInuse"`
	NumGC      uint32 `json:"numGC"`
}

type MonitorSample struct {
	Timestamp         string  `json:"timestamp"`
	CPUUsage          float64 `json:"cpuUsage"`
	MemoryUsedPercent float64 `json:"memoryUsedPercent"`
	SwapUsedPercent   float64 `json:"swapUsedPercent"`
	DiskUsedPercent   float64 `json:"diskUsedPercent"`
	Load1             float64 `json:"load1"`
	Goroutines        uint64  `json:"goroutines"`
	HeapAlloc         uint64  `json:"heapAlloc"`
	NetBytesSent      uint64  `json:"netBytesSent"`
	NetBytesRecv      uint64  `json:"netBytesRecv"`
}

type MonitorSnapshot struct {
	Timestamp string                  `json:"timestamp"`
	Host      *HostStats              `json:"host"`
	CPU       CPUSnapshot             `json:"cpu"`
	Memory    MemorySnapshot          `json:"memory"`
	Swap      SwapSnapshot            `json:"swap"`
	Disks     []DiskPartitionSnapshot `json:"disks"`
	Network   []NetworkSnapshot       `json:"network"`
	GoRuntime GoRuntimeSnapshot       `json:"goRuntime"`
	Current   MonitorSample           `json:"current"`
}

func GetMonitorSnapshot() (*MonitorSnapshot, error) {
	hostInfo, err := GetHostInfo()
	if err != nil {
		return nil, err
	}

	cpuInfo, _ := cpu.Info()
	cpuUsages, err := cpu.Percent(200*time.Millisecond, true)
	if err != nil {
		return nil, err
	}
	totalUsage := 0.0
	if len(cpuUsages) > 0 {
		for _, v := range cpuUsages {
			totalUsage += v
		}
		totalUsage /= float64(len(cpuUsages))
	}
	loadAvg, _ := load.Avg()

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	swapStat, _ := mem.SwapMemory()

	disks := make([]DiskPartitionSnapshot, 0)
	partitions, _ := disk.Partitions(false)
	for _, part := range partitions {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			continue
		}
		disks = append(disks, DiskPartitionSnapshot{
			Path:        part.Mountpoint,
			Fstype:      part.Fstype,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
		})
	}
	if len(disks) == 0 {
		root, err := disk.Usage("/")
		if err == nil {
			disks = append(disks, DiskPartitionSnapshot{
				Path:        "/",
				Fstype:      "",
				Total:       root.Total,
				Used:        root.Used,
				Free:        root.Free,
				UsedPercent: root.UsedPercent,
			})
		}
	}

	netStats, _ := net.IOCounters(true)
	networks := make([]NetworkSnapshot, 0, len(netStats))
	var totalSent uint64
	var totalRecv uint64
	for _, item := range netStats {
		networks = append(networks, NetworkSnapshot{
			Name:        item.Name,
			BytesSent:   item.BytesSent,
			BytesRecv:   item.BytesRecv,
			PacketsSent: item.PacketsSent,
			PacketsRecv: item.PacketsRecv,
			Errin:       item.Errin,
			Errout:      item.Errout,
			Dropin:      item.Dropin,
			Dropout:     item.Dropout,
		})
		totalSent += item.BytesSent
		totalRecv += item.BytesRecv
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	now := time.Now()

	snapshot := &MonitorSnapshot{
		Timestamp: now.Format(time.RFC3339),
		Host:      hostInfo,
		CPU: CPUSnapshot{
			Cores:      len(cpuUsages),
			Usage:      totalUsage,
			CoreUsages: cpuUsages,
			Load1:      loadAvg.Load1,
			Load5:      loadAvg.Load5,
			Load15:     loadAvg.Load15,
		},
		Memory: MemorySnapshot{
			Total:       vmStat.Total,
			Used:        vmStat.Used,
			Free:        vmStat.Free,
			Available:   vmStat.Available,
			Cached:      vmStat.Cached,
			Buffers:     vmStat.Buffers,
			UsedPercent: vmStat.UsedPercent,
		},
		Swap: SwapSnapshot{
			Total:       swapStat.Total,
			Used:        swapStat.Used,
			Free:        swapStat.Free,
			UsedPercent: swapStat.UsedPercent,
		},
		Disks:   disks,
		Network: networks,
		GoRuntime: GoRuntimeSnapshot{
			Goroutines: uint64(runtime.NumGoroutine()),
			HeapAlloc:  memStats.HeapAlloc,
			HeapSys:    memStats.HeapSys,
			HeapIdle:   memStats.HeapIdle,
			HeapInuse:  memStats.HeapInuse,
			StackInuse: memStats.StackInuse,
			NumGC:      memStats.NumGC,
		},
	}
	if len(cpuInfo) > 0 {
		snapshot.CPU.ModelName = cpuInfo[0].ModelName
		snapshot.CPU.Mhz = cpuInfo[0].Mhz
	}
	diskUsedPercent := 0.0
	if len(disks) > 0 {
		diskUsedPercent = disks[0].UsedPercent
	}
	snapshot.Current = MonitorSample{
		Timestamp:         snapshot.Timestamp,
		CPUUsage:          snapshot.CPU.Usage,
		MemoryUsedPercent: snapshot.Memory.UsedPercent,
		SwapUsedPercent:   snapshot.Swap.UsedPercent,
		DiskUsedPercent:   diskUsedPercent,
		Load1:             snapshot.CPU.Load1,
		Goroutines:        snapshot.GoRuntime.Goroutines,
		HeapAlloc:         snapshot.GoRuntime.HeapAlloc,
		NetBytesSent:      totalSent,
		NetBytesRecv:      totalRecv,
	}
	return snapshot, nil
}

func (m MonitorSnapshot) Summary() string {
	return fmt.Sprintf(
		"CPU %.2f%% | MEM %.2f%% | SWAP %.2f%% | Load %.2f | Goroutines %d",
		m.CPU.Usage,
		m.Memory.UsedPercent,
		m.Swap.UsedPercent,
		m.CPU.Load1,
		m.GoRuntime.Goroutines,
	)
}
