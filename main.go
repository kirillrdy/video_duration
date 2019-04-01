package main

import (
	"github.com/anacrolix/ffprobe"
	"log"
	"path/filepath"
	"time"
)

func main() {
	var sum time.Duration
	files, _ := filepath.Glob("*.mp4")
	for _, file := range files {
		info, _ := ffprobe.Run(file)

		duration, _ := info.Duration()
		log.Print(duration)
		sum += duration
	}

	log.Print(sum)
}
