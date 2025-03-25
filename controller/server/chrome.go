package server

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// TakeScreenshot takes a screenshot of the given URL or file path with the specified dimensions.
// The binary parameter selects which headless browser to use. If binary equals "chromium-browser",
// an extra flag ("--headless=old") is added; otherwise, "chrome-headless-shell" is used.
// It returns the screenshot image as a byte slice or an error.
func TakeScreenshot(urlOrPath string, dim [2]int, binary string) ([]byte, error) {
	if binary == "" {
		binary = "chrome-headless-shell"
	}

	// Create a temporary file for the screenshot.
	tmpFile, err := os.CreateTemp("", "*.png")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %w", err)
	}
	tmpFilePath := tmpFile.Name()
	tmpFile.Close() // Close immediately since we only need the file path.
	// Ensure the temporary file is removed.
	defer func() {
		if err := os.Remove(tmpFilePath); err != nil {
			log.Printf("Failed to remove temporary file: %v", err)
		}
	}()

	log.Printf("Taking screenshot: url=%s, path=%s", urlOrPath, tmpFilePath)

	// Build the command arguments.
	var args []string
	if binary == "chromium-browser" {
		args = append(args, "--headless=old")
	}

	args = append(args,
		fmt.Sprintf("--screenshot=%s", tmpFilePath),
		fmt.Sprintf("--window-size=%d,%d", dim[0], dim[1]),
		"--force-device-scale-factor=1",
		"--no-sandbox",
		"--disable-gpu",
		"--disable-software-rasterizer",
		"--disable-dev-shm-usage",
		"--hide-scrollbars",
		urlOrPath,
	)

	// Execute the command.
	cmd := exec.Command(binary, args...)
	log.Printf("Running command: %s %v", binary, args)
	output, err := cmd.CombinedOutput()
	log.Printf("Running command: %s", output)
	if err != nil || !fileExists(tmpFilePath) {
		log.Printf("Failed to take screenshot: error=%s", err)
		return nil, fmt.Errorf("failed to take screenshot: %s", err)
	}

	// Read the screenshot file.
	imageBytes, err := os.ReadFile(tmpFilePath)
	if err != nil {
		log.Printf("Error reading screenshot: error=%v, file=%s", err, tmpFilePath)
		return nil, err
	}

	return imageBytes, nil
}

// fileExists checks if a file exists at the given path.
func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

