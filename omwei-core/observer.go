package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("👁️ Willie's Observer is watching history.ow...")
	
	lastSize := int64(0)

	for {
		file, err := os.Open("history.ow")
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		info, _ := file.Stat()
		currentSize := info.Size()

		// Ak súbor narástol, prečítaj nové riadky
		if currentSize > lastSize {
			file.Seek(lastSize, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Printf("🔍 Willie analyzes: %s\n", line)

				// JEDNODUCHÁ LOGIKA: Ak riadok obsahuje anomáliu
				if strings.Contains(line, "anomaly_detected") {
					fmt.Println("⚠️  ALERT: Pattern recognized. Potential instability!")
				}
			}
			lastSize = currentSize
		}
		file.Close()
		time.Sleep(1 * time.Second)
	}
}
