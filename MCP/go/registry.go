package main

import (
	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	tools_v1 "github.com/mcp-server/mcp-server/tools/v1"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v1.CreateGetcountriesTool(cfg),
		tools_v1.CreateGetcountryTool(cfg),
	}
}
