package main

import (
	"flag"
	"fmt"
	"github.com/anacrolix/ffprobe"
	"os"
	"path/filepath"
	"time"
)

func main() {

	remove := flag.Bool("rm", false, "Removes videos longer than 1h")
	flag.Parse()

	var sum time.Duration
	files, _ := filepath.Glob("*.mp4")
	for _, file := range files {
		info, err := ffprobe.Run(file)
		if err == nil {
			duration, _ := info.Duration()
			if *remove == true && duration > time.Hour {
				os.Remove(file)
			}
			fmt.Printf("%s %s\n", file, duration)
			sum += duration
		}
	}

	fmt.Printf("Total: %s\n", sum)
}
