package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type BingResponse struct {
	Images []struct {
		URLBase string `json:"urlbase"`
	} `json:"images"`
}

func main() {
	fmt.Println("Fetching Bing Wallpaper of the Day...")

	resp, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US")
	if err != nil {
		fmt.Printf("Error requesting Bing API: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response API: %v\n", err)
		os.Exit(1)
	}

	var bingData BingResponse
	if err := json.Unmarshal(body, &bingData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	if len(bingData.Images) == 0 {
		fmt.Println("Error: No images found from Bing API.")
		os.Exit(1)
	}

	urlBase := bingData.Images[0].URLBase
	fullURL := "https://www.bing.com" + urlBase + "_UHD.jpg"

	// Create Pictures folder silently using UserHomeDir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		os.Exit(1)
	}

	wallpaperDir := filepath.Join(homeDir, "Pictures", "BingWallpapers")
	if err := os.MkdirAll(wallpaperDir, 0755); err != nil {
		fmt.Printf("Error creating wallpaper directory: %v\n", err)
		os.Exit(1)
	}

	// Extract clean filename from the URL payload
	tokens := strings.Split(urlBase, ".")
	filename := "wallpaper"
	if len(tokens) > 1 {
		filename = tokens[1]
	}
	filename += "_UHD.jpg"
	destPath := filepath.Join(wallpaperDir, filename)

	// Download image if it hasn't been fetched yet
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		fmt.Printf("Downloading UHD Wallpaper to %s...\n", destPath)
		if err := downloadFile(destPath, fullURL); err != nil {
			fmt.Printf("Error downloading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Wallpaper already downloaded for today!")
	}

	fmt.Println("Setting desktop background via macOS System Events...")
	script := fmt.Sprintf(`tell application "System Events" to set picture of every desktop to "%s"`, destPath)
	cmd := exec.Command("osascript", "-e", script)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error setting desktop wallpaper: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success! Enjoy your new wallpaper.")
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	return err
}
