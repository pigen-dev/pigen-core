package cicd

import (
	"fmt"
	"os/exec"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/pigen-dev/pigen-core/internal/utils"
	shared "github.com/pigen-dev/shared"
)

func discover(cicdStruct shared.PigenStepsFile) (shared.CicdInterface,*goplugin.Client, error) {
	cicdType := cicdStruct.Type
	pluginVersion := cicdStruct.Version
	repoUrl := cicdStruct.RepoUrl
	pluginBinaryPath, err := utils.PluginGetter(cicdType, repoUrl, pluginVersion)
	// Set up the plugin client
	client := goplugin.NewClient(&goplugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]goplugin.Plugin{"cicdPlugin": &shared.CicdPlugin{}},
		Cmd: exec.Command(pluginBinaryPath),
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create RPC client: %w", err)
	}
	
	raw, err := rpcClient.Dispense("cicdPlugin")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dispense plugin: %w", err)
	}

	cicdPlugin := raw.(shared.CicdInterface)
	return cicdPlugin, client, nil
}