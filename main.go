package main

import (
	//"cs50-romain/terminalfun/box"
	//"cs50-romain/terminalfun/progress"
	"cs50-romain/terminalfun/shell"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"time"
	"unsafe"
)

// Termios represents the terminal settings structure.
type Termios struct {
	Iflag, Oflag, Cflag, Lflag uint32
	Cc                         [20]uint8
	Ispeed, Ospeed             uint32
}

func main() {
	// Get the file descriptor for stdin.
	fd := int(os.Stdin.Fd())

	// Save the original terminal settings.
	origTermios, err := getTermios(fd)
	if err != nil {
		fmt.Println("Error getting terminal attributes:", err)
		return
	}
	defer restoreTermios(fd, origTermios)

	// Set the terminal to raw mode.
	rawTermios := origTermios
	rawTermios.Lflag &^= syscall.ICANON | syscall.ECHO | syscall.ISIG
	rawTermios.Iflag &^= syscall.IXON | syscall.ICRNL
	rawTermios.Oflag &^= syscall.OPOST

	err = setTermios(fd, &rawTermios)
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		return
	}

	// Handle interrupt signal to restore terminal settings.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		restoreTermios(fd, origTermios)
		os.Exit(1)
	}()

	fmt.Println("Terminal is now in raw mode. Press Ctrl+C to exit.")

	/*
	// Example loop to read characters one by one.
	for {
		var buf [1]byte
		os.Stdin.Read(buf[:])
		fmt.Printf("Read character: %q\r\n", buf[0])
	}
	*/
	//box.Render()
	//time.Sleep(5 * time.Second)
	//progress.ProgressRender()
	if err := shell.Run(); err != nil {
		fmt.Println(err)
	}
}

// getTermios retrieves the terminal attributes for the given file descriptor.
func getTermios(fd int) (Termios, error) {
	var termios Termios
	_, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	if errno != 0 {
		return termios, errno
	}
	return termios, nil
}

// setTermios sets the terminal attributes for the given file descriptor.
func setTermios(fd int, termios *Termios) error {
	_, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(termios)), 0, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

// restoreTermios restores the terminal attributes to the original settings.
func restoreTermios(fd int, termios Termios) {
	setTermios(fd, &termios)
}
