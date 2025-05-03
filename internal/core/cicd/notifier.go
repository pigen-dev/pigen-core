package cicd

import (
	"fmt"

	shared "github.com/pigen-dev/shared"
)

func PipelineNotifier(pipelineNotification shared.PipelineNotification) error {
	//POST request to backend web to update pipeline status (git url, branch, status)
	fmt.Println("Pipeline repo:", pipelineNotification.RepoUrl)
	fmt.Println("Pipeline branch:", pipelineNotification.Branch)
	fmt.Println("Pipeline status:", pipelineNotification.Status)
	//In case of failure
	//GET request to backend web to get if pipeline AI is enabled (git url, branch)
	//POST request to Pi-Pilot if AI is enabled (cicd type, git url, branch, metadata)
	fmt.Println("Pipeline metadata:", pipelineNotification.Metadata)
	return nil
}