package main

import (
	"fmt"
	"os"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {
	// given file path, create and return new path to file with fast start encoding

	// new output file filePath
	newPath := fmt.Sprintf("%s.processing", filePath)
	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", newPath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Error executing %s: %v", cmd.String(), err)
	}

	// check that path exists after ffmpeg execution
	if _, err := os.Stat(newPath); err != nil {
		return "", fmt.Errorf("File at path: %s not created: %v", newPath, err)
	}
	return newPath, nil
}
