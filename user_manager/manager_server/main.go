package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/install", func(c *gin.Context) {
		var req struct {
			ReleaseName string `json:"release_name"`
			Namespace   string `json:"namespace"`
			Port        int    `json:"port"`
			Storage     string `json:"storage"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		values := map[string]interface{}{
			"port":    req.Port,
			"storage": req.Storage,
		}

		release, err := helmInstall("../user_manager_chart/", req.ReleaseName, req.Namespace, values)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, release)
	})

	r.DELETE("/uninstall/:release_name/:namespace", func(c *gin.Context) {
		releaseName := c.Param("release_name")
		namespace := c.Param("namespace")

		if err := helmUninstall(releaseName, namespace); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "uninstalled"})
	})

	r.Run(":8080") // Listen and serve on port 8080
}
