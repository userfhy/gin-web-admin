package system_monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryStats struct {
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent float64
}

type DiskStats struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

type SystemStats struct {
	HostInfo    *HostStats
	CPUUsage    float64
	MemoryUsage *MemoryStats
	DiskUsage   *DiskStats
}

// 新增 HostStats 结构体
type HostStats struct {
	Hostname     string
	Uptime       uint64
	BootTime     string
	OS           string
	Platform     string
	KernelArch   string
	ProcessCount uint64
	Users        []UserInfo
}

type UserInfo struct {
	User     string
	Terminal string
	Host     string
	Started  int
}

// 其他原有结构体保持不变...

// GetHostInfo 获取主机信息
func GetHostInfo() (*HostStats, error) {
	// 获取基础主机信息
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("获取主机信息失败: %v", err)
	}

	// 获取用户登录信息
	users, _ := host.Users() // 允许用户信息获取失败

	// 转换用户信息
	var userList []UserInfo
	for _, u := range users {
		userList = append(userList, UserInfo{
			User:     u.User,
			Terminal: u.Terminal,
			Host:     u.Host,
			Started:  u.Started,
		})
	}

	return &HostStats{
		Hostname:     hostInfo.Hostname,
		Uptime:       hostInfo.Uptime,
		BootTime:     formatBootTime(hostInfo.BootTime),
		OS:           hostInfo.OS,
		Platform:     fmt.Sprintf("%s %s", hostInfo.Platform, hostInfo.PlatformVersion),
		KernelArch:   fmt.Sprintf("%s %s", hostInfo.KernelVersion, hostInfo.KernelArch),
		ProcessCount: hostInfo.Procs,
		Users:        userList,
	}, nil
}

// 辅助函数：格式化启动时间
func formatBootTime(bootTime uint64) string {
	t := time.Unix(int64(bootTime), 0)
	return t.Format("2006-01-02 15:04:05")
}

// GetCPUUsage 获取CPU使用率（百分比）
func GetCPUUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percent) == 0 {
		return 0, fmt.Errorf("no CPU data available")
	}
	return percent[0], nil
}

// GetMemoryUsage 获取内存使用情况
func GetMemoryUsage() (*MemoryStats, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &MemoryStats{
		Total:       memInfo.Total,
		Used:        memInfo.Used,
		Free:        memInfo.Free,
		UsedPercent: memInfo.UsedPercent,
	}, nil
}

// GetDiskUsage 获取磁盘使用情况
func GetDiskUsage(path string) (*DiskStats, error) {
	diskInfo, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}

	return &DiskStats{
		Total:       diskInfo.Total,
		Free:        diskInfo.Free,
		Used:        diskInfo.Used,
		UsedPercent: diskInfo.UsedPercent,
	}, nil
}

// 修改后的 GetSystemStats 包含主机信息
func GetSystemStats() (*SystemStats, error) {
	hostInfo, err := GetHostInfo()
	if err != nil {
		return nil, err
	}

	cpuUsage, err := GetCPUUsage()
	if err != nil {
		return nil, err
	}

	memStats, err := GetMemoryUsage()
	if err != nil {
		return nil, err
	}

	diskStats, err := GetDiskUsage("/")
	if err != nil {
		return nil, err
	}

	return &SystemStats{
		HostInfo:    hostInfo,
		CPUUsage:    cpuUsage,
		MemoryUsage: memStats,
		DiskUsage:   diskStats,
	}, nil
}

// 更新后的格式化输出
func (s *SystemStats) String() string {
	hostInfo := fmt.Sprintf(
		"主机名称: %s\n"+
			"系统运行: %s (启动时间: %s)\n"+
			"操作系统: %s\n"+
			"内核架构: %s\n"+
			"进程数量: %d\n"+
			"登录用户: %d",
		s.HostInfo.Hostname,
		formatUptime(s.HostInfo.Uptime),
		s.HostInfo.BootTime,
		s.HostInfo.OS,
		s.HostInfo.KernelArch,
		s.HostInfo.ProcessCount,
		len(s.HostInfo.Users),
	)

	return fmt.Sprintf(
		"=== 主机信息 ===\n%s\n\n"+
			"=== 性能指标 ===\n"+
			"CPU使用率: %.2f%%\n"+
			"内存使用: %.2f%% (总: %s, 已用: %s, 剩余: %s)\n"+
			"磁盘使用: %.2f%% (总: %s, 已用: %s, 剩余: %s)",
		hostInfo,
		s.CPUUsage,
		s.MemoryUsage.UsedPercent,
		formatBytes(s.MemoryUsage.Total),
		formatBytes(s.MemoryUsage.Used),
		formatBytes(s.MemoryUsage.Free),
		s.DiskUsage.UsedPercent,
		formatBytes(s.DiskUsage.Total),
		formatBytes(s.DiskUsage.Used),
		formatBytes(s.DiskUsage.Free),
	)
}

// 新增辅助函数：格式化运行时间
func formatUptime(seconds uint64) string {
	duration := time.Duration(seconds) * time.Second
	return fmt.Sprintf("%v", duration-(duration%time.Second))
}

// 字节单位转换工具函数
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
