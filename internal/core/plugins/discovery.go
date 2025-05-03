package plugins

import (
	"fmt"
	"os/exec"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/pigen-dev/pigen-core/internal/utils"
	shared "github.com/pigen-dev/shared"
)

func discover(pluginStruct shared.PluginStruct) (shared.PluginInterface,*goplugin.Client, error) {
	binaryName := pluginStruct.ID
	pluginVersion := pluginStruct.Version
	repoUrl := pluginStruct.RepoUrl
	pluginBinaryPath, err := utils.PluginGetter(binaryName, repoUrl, pluginVersion)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get plugin binary: %w", err)
	}

	// Set up the plugin client
	client := goplugin.NewClient(&goplugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]goplugin.Plugin{"pigenPlugin": &shared.PigenPlugin{}},
		Cmd: exec.Command(pluginBinaryPath),
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create RPC client: %w", err)
	}

	raw, err := rpcClient.Dispense("pigenPlugin")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dispense plugin: %w", err)
	}

	plugin, ok := raw.(shared.PluginInterface)
	if !ok {
    return nil, nil, fmt.Errorf("dispensed plugin does not implement PluginInterface, got type: %T", raw)
}

	return plugin, client, nil
}