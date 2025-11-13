package routes

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func DeployRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("New deployment initiating")

		cmd := exec.Command("/bin/bash", "./scripts/deployment.sh")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Error running script: %v\nOutput: %s", err, string(output))
		}

		fmt.Printf("Script output:\n%s", string(output))
	}
}
