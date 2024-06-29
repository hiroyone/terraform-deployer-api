// utils/auth.go
package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func InjectAuth(filePath, authToken string) error {
	// Placeholder for actual logic to inject auth token into the file
	// Example: Append auth token to the file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("\n# Auth Token: %s\n", authToken)); err != nil {
		return err
	}

	return nil
}

func DeployResources(filePath string) error {
	// Placeholder for actual deployment logic
	// Example: Run `terraform apply` on the provided file
	cmd := exec.Command("terraform", "apply", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("terraform apply failed: %v", err)
	}

	return nil
}
