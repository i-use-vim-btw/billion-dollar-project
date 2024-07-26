package main

/*
	example deployment command:
	curl -X POST "http://localhost:8080/deploy" -H "Content-Type: application/json" -d '{
    "namespace": "user1",
    "storage": "10Gi"
}'


*/

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

type DeployRequest struct {
	Namespace string `json:"namespace"`
	Storage   string `json:"storage"`
}

func initHelmConfig(namespace string) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	settings := cli.New()

	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, fmt.Errorf("failed to initialize Helm action configuration: %w", err)
	}

	return actionConfig, nil
}

func helmInstall(chartPath, releaseName, namespace string, values map[string]interface{}) (*release.Release, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, fmt.Errorf("failed to initialize Helm action configuration: %w", err)
	}

	client := action.NewInstall(actionConfig)
	client.ReleaseName = releaseName
	client.Namespace = namespace
	client.CreateNamespace = false // already create in templete

	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}

	release, err := client.Run(chart, values)
	if err != nil {
		return nil, fmt.Errorf("failed to install chart: %w", err)
	}

	return release, nil
}

func deployHandler(c *gin.Context) {
	var req DeployRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	namespace := req.Namespace
	releaseName := fmt.Sprintf("rel-%s", namespace)
	chartPath := "../user_manager_chart/"

	values := map[string]interface{}{
		"Namespace": req.Namespace,
		"Storage":   req.Storage,
	}

	_, err := helmInstall(chartPath, releaseName, namespace, values)
	if err != nil {
		log.Printf("Failed to deploy Milvus: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to deploy Milvus: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deployed successfully."})
}

func main() {
	r := gin.Default()
	r.POST("/deploy", deployHandler)
	r.Run(":8080") // Run on port 8080
}
