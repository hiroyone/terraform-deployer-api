// models/request.go
package models

type DeployRequest struct {
	ModuleFile string `json:"module_file"`
	AuthToken  string `json:"auth_token"`
}
