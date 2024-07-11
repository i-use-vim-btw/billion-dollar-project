package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func helmInstall(user string) (string, error) {
	releaseName := user + "-milvus"
	namespace := user + "-namespace"
	valuesFile := "./values/user-values.yaml"

	// Create namespace
	cmd := exec.Command("kubectl", "create", "namespace", namespace)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Helm install command
	cmd = exec.Command("helm", "install", releaseName, "./charts/milvus-chart",
		"--namespace", namespace, "--values", valuesFile)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return "Deployed " + releaseName + " in " + namespace, nil
}

func helmUninstall(user string) (string, error) {
	releaseName := user + "-milvus"
	namespace := user + "-namespace"

	// Helm uninstall command
	cmd := exec.Command("helm", "uninstall", releaseName, "--namespace", namespace)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Delete namespace
	cmd = exec.Command("kubectl", "delete", "namespace", namespace)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return "Deleted " + releaseName + " and namespace " + namespace, nil
}

func main() {
	r := gin.Default()

	r.POST("/deploy", func(c *gin.Context) {
		var json struct {
			User string `json:"user" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := helmInstall(json.User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	r.DELETE("/delete", func(c *gin.Context) {
		var json struct {
			User string `json:"user" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := helmUninstall(json.User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": result})
	})

	r.Run(":5000")
}
