package main

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
)

func installHelmChart(namespace, releaseName, chartPath string, valuesFile string) (*release.Release, error) {
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(cli.New().RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Printf(format, v...)
	}); err != nil {
		return nil, err
	}

	install := action.NewInstall(actionConfig)
	install.Namespace = namespace
	install.ReleaseName = releaseName

	valueOpts := &values.Options{
		ValueFiles: []string{valuesFile},
	}

	p := getter.All(cli.New())
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		return nil, err
	}

	cp, err := install.ChartPathOptions.LocateChart(chartPath, cli.New())
	if err != nil {
		return nil, err
	}

	chartRequested, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}

	release, err := install.Run(chartRequested, vals)
	if err != nil {
		return nil, err
	}

	return release, nil
}

func main() {
	namespace := "u1"
	releaseName := "u1-rel"
	chartPath := "../user_manager_chart"
	valuesFile := "../user_manager_chart/values.yaml"
	rel, err := installHelmChart(namespace, releaseName, chartPath, valuesFile)
	if err != nil {
		fmt.Printf("Failed to install Helm chart: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Helm chart installed successfully: %s\n", rel.Name)
}
