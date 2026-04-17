package sysService

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/system_monitor"
)

type ServerMonitorVO struct {
	Snapshot *system_monitor.MonitorSnapshot `json:"snapshot"`
	History  []system_monitor.MonitorSample  `json:"history"`
	Warnings []MonitorWarning                `json:"warnings"`
}

type MonitorWarning struct {
	Code      string  `json:"code"`
	Level     string  `json:"level"`
	Message   string  `json:"message"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
}

type ServerMonitorStreamVO struct {
	Seq              uint64                          `json:"seq"`
	Mode             string                          `json:"mode"`
	SampleIntervalMs int64                           `json:"sampleIntervalMs"`
	ServerRetryMs    int64                           `json:"serverRetryMs"`
	Snapshot         *system_monitor.MonitorSnapshot `json:"snapshot,omitempty"`
	History          []system_monitor.MonitorSample  `json:"history,omitempty"`
	AppendSample     *system_monitor.MonitorSample   `json:"appendSample,omitempty"`
	Warnings         []MonitorWarning                `json:"warnings"`
}

var (
	monitorOnce    sync.Once
	monitorMu      sync.RWMutex
	monitorHistory = make([]system_monitor.MonitorSample, 0, 90)
	monitorSeq     uint64
)

const (
	monitorSampleInterval = 10 * time.Second
	monitorHistoryLimit   = 90
	cpuWarnThreshold      = 85
	memoryWarnThreshold   = 90
	swapWarnThreshold     = 80
	diskWarnThreshold     = 90
)

func (s *Service) GetServerMonitor() (*ServerMonitorVO, error) {
	s.startMonitorSampler()

	snapshot, err := system_monitor.GetMonitorSnapshot()
	if err != nil {
		return nil, err
	}

	history := s.recordMonitorSample(snapshot.Current)

	return &ServerMonitorVO{
		Snapshot: snapshot,
		History:  history,
		Warnings: buildMonitorWarnings(snapshot),
	}, nil
}

func (s *Service) GetServerMonitorStreamInit(serverRetry time.Duration) (*ServerMonitorStreamVO, error) {
	data, err := s.GetServerMonitor()
	if err != nil {
		return nil, err
	}

	return &ServerMonitorStreamVO{
		Seq:              nextMonitorSeq(),
		Mode:             "init",
		SampleIntervalMs: monitorSampleInterval.Milliseconds(),
		ServerRetryMs:    serverRetry.Milliseconds(),
		Snapshot:         data.Snapshot,
		History:          data.History,
		Warnings:         data.Warnings,
	}, nil
}

func (s *Service) GetServerMonitorStreamDelta(serverRetry time.Duration) (*ServerMonitorStreamVO, error) {
	s.startMonitorSampler()

	snapshot, err := system_monitor.GetMonitorSnapshot()
	if err != nil {
		return nil, err
	}

	history := s.recordMonitorSample(snapshot.Current)
	appendSample := snapshot.Current
	if len(history) > 0 {
		appendSample = history[len(history)-1]
	}

	return &ServerMonitorStreamVO{
		Seq:              nextMonitorSeq(),
		Mode:             "append",
		SampleIntervalMs: monitorSampleInterval.Milliseconds(),
		ServerRetryMs:    serverRetry.Milliseconds(),
		Snapshot:         snapshot,
		AppendSample:     &appendSample,
		Warnings:         buildMonitorWarnings(snapshot),
	}, nil
}

func (s *Service) startMonitorSampler() {
	monitorOnce.Do(func() {
		go func() {
			s.collectMonitorSample()

			ticker := time.NewTicker(monitorSampleInterval)
			defer ticker.Stop()

			for {
				<-ticker.C
				s.collectMonitorSample()
			}
		}()
	})
}

func (s *Service) collectMonitorSample() {
	snapshot, err := system_monitor.GetMonitorSnapshot()
	if err != nil {
		logging.Warnf("collect system monitor sample failed: %v", err)
		return
	}

	monitorMu.Lock()
	appendMonitorSample(snapshot.Current)
	monitorMu.Unlock()
}

func (s *Service) recordMonitorSample(sample system_monitor.MonitorSample) []system_monitor.MonitorSample {
	monitorMu.Lock()
	defer monitorMu.Unlock()

	appendMonitorSample(sample)

	return append([]system_monitor.MonitorSample(nil), monitorHistory...)
}

func appendMonitorSample(sample system_monitor.MonitorSample) {
	if sample.Timestamp == "" {
		sample.Timestamp = time.Now().Format(time.RFC3339)
	}
	if len(monitorHistory) > 0 && monitorHistory[len(monitorHistory)-1].Timestamp == sample.Timestamp {
		monitorHistory[len(monitorHistory)-1] = sample
		return
	}

	monitorHistory = append(monitorHistory, sample)
	if len(monitorHistory) > monitorHistoryLimit {
		monitorHistory = append([]system_monitor.MonitorSample(nil), monitorHistory[len(monitorHistory)-monitorHistoryLimit:]...)
	}
}

func (s *Service) GetServerMonitorSummary() (string, error) {
	data, err := s.GetServerMonitor()
	if err != nil {
		return "", err
	}
	if data == nil || data.Snapshot == nil {
		return "", fmt.Errorf("empty monitor snapshot")
	}
	return data.Snapshot.Summary(), nil
}

func buildMonitorWarnings(snapshot *system_monitor.MonitorSnapshot) []MonitorWarning {
	if snapshot == nil {
		return nil
	}

	warnings := make([]MonitorWarning, 0, 4)
	if snapshot.CPU.Usage >= cpuWarnThreshold {
		warnings = append(warnings, MonitorWarning{
			Code:      "cpu_high",
			Level:     "warning",
			Message:   "CPU 使用率过高",
			Value:     snapshot.CPU.Usage,
			Threshold: cpuWarnThreshold,
		})
	}
	if snapshot.Memory.UsedPercent >= memoryWarnThreshold {
		warnings = append(warnings, MonitorWarning{
			Code:      "memory_high",
			Level:     "warning",
			Message:   "内存使用率过高",
			Value:     snapshot.Memory.UsedPercent,
			Threshold: memoryWarnThreshold,
		})
	}
	if snapshot.Swap.UsedPercent >= swapWarnThreshold {
		warnings = append(warnings, MonitorWarning{
			Code:      "swap_high",
			Level:     "warning",
			Message:   "交换分区使用率过高",
			Value:     snapshot.Swap.UsedPercent,
			Threshold: swapWarnThreshold,
		})
	}

	maxDiskUsage := 0.0
	for _, item := range snapshot.Disks {
		if item.UsedPercent > maxDiskUsage {
			maxDiskUsage = item.UsedPercent
		}
	}
	if maxDiskUsage >= diskWarnThreshold {
		warnings = append(warnings, MonitorWarning{
			Code:      "disk_high",
			Level:     "warning",
			Message:   "磁盘使用率过高",
			Value:     maxDiskUsage,
			Threshold: diskWarnThreshold,
		})
	}

	return warnings
}

func nextMonitorSeq() uint64 {
	return atomic.AddUint64(&monitorSeq, 1)
}
