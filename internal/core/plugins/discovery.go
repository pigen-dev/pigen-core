package plugins

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	goplugin "github.com/hashicorp/go-plugin"
	shared "github.com/pigen-dev/shared"
)

func discover(pluginStruct shared.PluginStruct) (shared.PluginInterface,*goplugin.Client, error) {
	pluginId := pluginStruct.ID
	pluginVersion := pluginStruct.Version
	repoUrl := pluginStruct.RepoUrl
	pluginFile := fmt.Sprintf("%s-%s", pluginId, pluginVersion)
	currentDir, _ := os.Getwd()
	destDir := filepath.Join(currentDir, "plugins")
	destPath := filepath.Join(destDir, pluginFile)
	log.Println("destination path:", destPath)
	// Check if the binary exists locally
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		log.Printf("plugin binary not found: %s. Downlowading it...", pluginFile)
		err := downloadPlugin(repoUrl, pluginVersion, pluginId, destDir)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to download plugin: %w", err)
		}
	}
	
	// Set up the plugin client
	client := goplugin.NewClient(&goplugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]goplugin.Plugin{"pigenPlugin": &shared.PigenPlugin{}},
		Cmd: exec.Command(destPath),
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create RPC client: %w", err)
	}
	
	raw, err := rpcClient.Dispense("pigenPlugin")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dispense plugin: %w", err)
	}

	plugin := raw.(shared.PluginInterface)
	return plugin, client, nil
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