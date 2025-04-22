package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func PluginGetter(pluginId, repoUrl, pluginVersion string)(pluginBinaryPath string, err error){
	binaryName := fmt.Sprintf("%s-%s", pluginId, pluginVersion)
	log.Println("plugin binary name:", binaryName)
	currentDir, _ := os.Getwd()
	destDir := filepath.Join(currentDir, "plugins")
	pluginBinaryPath = filepath.Join(destDir, binaryName)
	log.Println("destination path:", pluginBinaryPath)
	// Check if the binary exists locally
	if _, err := os.Stat(pluginBinaryPath); os.IsNotExist(err) {
		log.Printf("plugin binary not found: %s. Downlowading it...", binaryName)
		err := downloadPlugin(repoUrl, pluginVersion, pluginId, destDir)
		if err != nil {
			return "", fmt.Errorf("failed to download plugin: %w", err)
		}
	}
	return pluginBinaryPath, nil
}

func downloadPlugin(repoUrl, pluginVersion, binaryName, destDir string) error {
	url := fmt.Sprintf("%s/releases/download/%s/%s", repoUrl, pluginVersion, binaryName)
	log.Println("Downloading plugin from:", url)
	resp, err := http.Get(url)
	if err != nil {
			return fmt.Errorf("failed to fetch plugin: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	pluginBinary := fmt.Sprintf("%s-%s", binaryName, pluginVersion)
	destPath := filepath.Join(destDir, pluginBinary)
	out, err := os.Create(destPath)
	if err != nil {
		currentDir, _ := os.Getwd()
		fmt.Println("Current working directory:", currentDir)
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
			return fmt.Errorf("failed to save plugin binary: %w", err)
	}

	// Make it executable
	err = os.Chmod(destPath, 0755)
	if err != nil {
			return fmt.Errorf("failed to make plugin executable: %w", err)
	}

	return nil
}