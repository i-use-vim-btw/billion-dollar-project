package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"log"
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
	chartPath := "../user_manager_chart"
	releaseName := "u1-rel"
	namespace := "u1-ns"
	storage := "1Gi"
	port := 30000
	values := map[string]interface{}{
		"Namespace":   namespace,
		"releaseName": releaseName,
		"Storage":     storage,
		"Port":        port,
	}

	release, err := helmInstall(chartPath, releaseName, namespace, values)
	if err != nil {
		log.Fatalf("Error installing chart: %v", err)
	}

	fmt.Printf("Chart installed successfully: %s\n", release.Name)
}
