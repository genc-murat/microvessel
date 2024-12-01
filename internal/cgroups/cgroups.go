package cgroups

import (
	"fmt"
	"os"
	"path/filepath"
)

// ResourceConfig holds cgroup resource limits
type ResourceConfig struct {
	MemoryLimit int64
	CPUShares   int64
}

// SetupCgroups creates and configures cgroups for the container
func SetupCgroups(containerID string, config ResourceConfig) error {
	cgroupPath := filepath.Join("/sys/fs/cgroup", containerID)

	// Create cgroup directory
	if err := os.MkdirAll(cgroupPath, 0755); err != nil {
		return fmt.Errorf("failed to create cgroup directory: %v", err)
	}

	// Set memory limit
	if config.MemoryLimit > 0 {
		memoryFile := filepath.Join(cgroupPath, "memory.limit_in_bytes")
		if err := os.WriteFile(memoryFile, []byte(fmt.Sprintf("%d", config.MemoryLimit)), 0644); err != nil {
			return fmt.Errorf("failed to set memory limit: %v", err)
		}
	}

	// Set CPU shares
	if config.CPUShares > 0 {
		cpuFile := filepath.Join(cgroupPath, "cpu.shares")
		if err := os.WriteFile(cpuFile, []byte(fmt.Sprintf("%d", config.CPUShares)), 0644); err != nil {
			return fmt.Errorf("failed to set CPU shares: %v", err)
		}
	}

	return nil
}
