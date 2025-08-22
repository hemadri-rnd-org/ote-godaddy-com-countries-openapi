package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetcountriesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["marketId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("marketId=%v", val))
		}
		if val, ok := args["regionTypeId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("regionTypeId=%v", val))
		}
		if val, ok := args["regionName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("regionName=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/countries%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result []CountrySummary
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetcountriesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_countries",
		mcp.WithDescription("Retrieves summary country information for the provided marketId and filters"),
		mcp.WithString("marketId", mcp.Required(), mcp.Description("MarketId in which the request is being made, and for which responses should be localized")),
		mcp.WithNumber("regionTypeId", mcp.Description("Restrict countries to this region type; required if regionName is supplied")),
		mcp.WithString("regionName", mcp.Description("Restrict countries to this region name; required if regionTypeId is supplied")),
		mcp.WithString("sort", mcp.Description("The term to sort the result countries by.")),
		mcp.WithString("order", mcp.Description("The direction to sort the result countries by.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetcountriesHandler(cfg),
	}
}
