package cicd

import (
	shared "github.com/pigen-dev/shared"
)

func ConnectRepo(pigenSteps shared.PigenStepsFile)shared.ActionRequired {
	cicdPlugin, client, err := discover(pigenSteps)
	if err != nil {
		return shared.ActionRequired{
			ActionUrl: "",
			Error: 	err,
		}
	}
	defer client.Kill()
	resp := cicdPlugin.ConnectRepo(pigenSteps)
	return resp
}

func CreateTrigger(pigenSteps shared.PigenStepsFile)error {
	cicdPlugin, client, err := discover(pigenSteps)
	if err != nil {
		return err
	}
	defer client.Kill()
	err = cicdPlugin.CreateTrigger(pigenSteps)
	return err
}

func GenerateScript(pigenSteps shared.PigenStepsFile)shared.CICDFile {
	cicdPlugin, client, err := discover(pigenSteps)
	if err != nil {
		return shared.CICDFile{
			Error: err,
			FileScript: nil,
		}
	}
	defer client.Kill()
	cicdFile := cicdPlugin.GeneratScript(pigenSteps)
	return cicdFile

}