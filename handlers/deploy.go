// handlers/deploy.go
package handlers

import (
	"encoding/json"
	"net/http"
	"terraform-deployer-api/models"
	"terraform-deployer-api/utils"
)

func DeployHandler(w http.ResponseWriter, r *http.Request) {
	var req models.DeployRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = utils.SaveFile(req.ModuleFile, "./uploaded_modules/main.tf")
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	err = utils.InjectAuth("./uploaded_modules/main.tf", req.AuthToken)
	if err != nil {
		http.Error(w, "Failed to inject auth", http.StatusInternalServerError)
		return
	}

	err = utils.DeployResources("./uploaded_modules/main.tf")
	if err != nil {
		http.Error(w, "Failed to deploy resources", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Module deployed successfully"))
}
