package main

import (
    "flag"
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
    // Define flags with default values (25 min work, 5 min break)
    workMin := flag.Int("w", 25, "Work duration in minutes")
    breakMin := flag.Int("b", 5, "Break duration in minutes")

    // Parse the flags provided by the user
    flag.Parse()

    fmt.Printf("Pomodoro Started! (Work: %d min, Break: %d min)\n", *workMin, *breakMin)
    fmt.Println("Press Ctrl+C to quit")

    for {
        // Convert the flag valus (int) into time.Duration
        runTimer(time.Duration(*workMin)*time.Minute, "WORK", colorRed)
        runTimer(time.Duration(*breakMin)*time.Minute, "BREAK", colorGreen)
    }
}
