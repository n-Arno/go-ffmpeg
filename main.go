package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func AudioConvert(filepath string, format string) string {
	extension := ""
	switch format {
	case "opus":
		extension = ".ogg"
	case "mp3", "wav", "aac", "flac":
		extension = "." + format
	default:
		extension = ".wav"
	}

	if extension == ".wav" {
		return filepath
	}

	target := strings.Replace(filepath, ".wav", extension, -1)
	cmd := exec.Command("ffmpeg", "-y", "-i", filepath, "-vn", target)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return target
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [input.wav]\n", os.Args[0])
		os.Exit(1)
	}

	formats := []string{"mp3", "aac", "opus", "flac"}
	src := os.Args[1]

	for _, format := range formats {
		fmt.Println(AudioConvert(src, format))
	}

}
