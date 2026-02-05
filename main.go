package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ANSI color codes for terminal output
const (
    colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

func runTimer(duration time.Duration, label string, color string) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    // Channel to listen for interrupt signals (e.g., Ctrl+C)
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

    remaining := int(duration.Seconds())

    for remaining >= 0 {
        select {
        case <-ticker.C:
            minutes := remaining / 60
            seconds := remaining % 60
            // Use \r to overwrite the current line
            fmt.Printf("\r%s[%s] %02d:%02d%s", color, label, minutes, seconds, colorReset)
            remaining--
        case <-sigChan:
            fmt.Printf("\n\nTimer stopped. Have a productive day!")
            os.Exit(0)
        }
    }

	fmt.Printf("\n%s session finished!\n", label)
}

func main() {
    fmt.Println("Pomodoro Timer Start! (Ctrl+C to quit)")

    for {
        runTimer(25*time.Minute, "WORK", colorRed)
        runTimer(5*time.Minute, "BREAK", colorGreen)
    }
}
