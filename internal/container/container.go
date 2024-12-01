package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Run starts a new container process
func Run(args []string) {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, args...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // Hostname isolation
			syscall.CLONE_NEWPID | // Process ID isolation
			syscall.CLONE_NEWNS | // Mount namespace
			syscall.CLONE_NEWNET, // Network namespace
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

// RunContainer is called as the init process inside the container
func RunContainer(args []string) {
	if len(args) < 1 {
		fmt.Println("No command specified")
		os.Exit(1)
	}

	// Set hostname
	must(syscall.Sethostname([]byte("container")))

	// Setup root filesystem
	setupRootFS()

	// Mount proc filesystem
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())

	// Cleanup
	must(syscall.Unmount("/proc", 0))
}

// setupRootFS prepares the root filesystem for the container
func setupRootFS() {
	// Make sure /proc exists
	must(os.MkdirAll("/proc", 0755))

	// Mount the root filesystem as private to prevent host mount propagation
	must(syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, ""))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
