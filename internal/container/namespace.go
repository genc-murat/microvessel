package container

import (
	"syscall"
)

// NamespaceConfig holds configuration for container namespaces
type NamespaceConfig struct {
	Hostname string
	Mount    bool
	PID      bool
	Network  bool
	UTS      bool
}

// GetNamespaceFlags returns syscall flags based on config
func GetNamespaceFlags(config NamespaceConfig) uintptr {
	var flags uintptr

	if config.UTS {
		flags |= syscall.CLONE_NEWUTS
	}
	if config.PID {
		flags |= syscall.CLONE_NEWPID
	}
	if config.Mount {
		flags |= syscall.CLONE_NEWNS
	}
	if config.Network {
		flags |= syscall.CLONE_NEWNET
	}

	return flags
}
