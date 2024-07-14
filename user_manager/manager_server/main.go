package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"log"
	"net/http"
	"os"
)

func helmInstall(chartPath, releaseName, namespace string, values map[string]interface{}) (*release.Release, error) {
	// Initialize Helm settings
	settings := cli.New()

	// Initialize Helm action configuration
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		return nil, fmt.Errorf("failed to initialize Helm action configuration: %w", err)
	}

	// Initialize Helm install client
	client := action.NewInstall(actionConfig)
	client.ReleaseName = releaseName
	client.Namespace = namespace
	client.CreateNamespace = false

	// Load the chart
	chart, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}

	// Install the chart
	release, err := client.Run(chart, values)
	if err != nil {
		return nil, fmt.Errorf("failed to install chart: %w", err)
	}

	return release, nil
}

func main() {
	r := gin.Default()

	r.POST("/install", func(c *gin.Context) {
		var req struct {
			ReleaseName string                 `json:"release_name"`
			Namespace   string                 `json:"namespace"`
			Values      map[string]interface{} `json:"values"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		release, err := helmInstall("../user_manager_chart/", req.ReleaseName, req.Namespace, req.Values)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, release)
	})

	r.Run(":8080") // Listen and serve on port 8080
}
