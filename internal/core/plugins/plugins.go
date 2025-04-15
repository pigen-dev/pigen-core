package plugins

import (
	"fmt"
	shared "github.com/pigen-dev/shared"
)
func SetupPlugins(pluginStruct shared.PluginStruct)error {
	plugin, client, err := discover(pluginStruct)
	if err != nil {
		return fmt.Errorf("failed to discover plugin: %w", err)
	}
	defer client.Kill()
	err = plugin.SetupPlugin(pluginStruct.Plugin.Config)
	if err != nil {
		return fmt.Errorf("failed to setup plugin: %w", err)
	}
	
	return nil
}

func DestroyPlugin(pluginStruct shared.PluginStruct)error {
	plugin, client, err := discover(pluginStruct)
	if err != nil {
		return fmt.Errorf("failed to discover plugin: %w", err)
	}
	defer client.Kill()
	err = plugin.Destroy(pluginStruct.Plugin.Config)
	if err != nil {
		return fmt.Errorf("failed to setup plugin: %w", err)
	}
	
	return nil
}

func GetOutput(pluginStruct shared.PluginStruct) shared.GetOutputResponse {
	plugin, client, err := discover(pluginStruct)
	if err != nil {
		return shared.GetOutputResponse{Output: nil, Error: err}
	}
	defer client.Kill()
	output := plugin.GetOutput(pluginStruct.Plugin.Config)
	fmt.Println("Output:", output)
	if output.Error != nil {
		return shared.GetOutputResponse{Output: nil, Error: output.Error}
	}
	
	return output
}