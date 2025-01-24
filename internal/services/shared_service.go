package services

import (
	"azurenv/internal/appconfig"
	"azurenv/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func ListRemote(config appconfig.AppConfig) {
	azCmd := exec.Command("az", "webapp", "config", "appsettings", "list",
		"--name", config.AppName,
		"--resource-group", config.ResourceGroup,
	)

	out, err := azCmd.Output()
	if err != nil {
		log.Fatalf("[ERROR] Failed to run 'az webapp config appsettings list: %v\n", err)
	}

	// Parse JSON output
	var settings []appconfig.AppSettings
	if err := json.Unmarshal(out, &settings); err != nil {
		log.Fatalf("[ERROR] Failed to parse JSON: %v\n", err)
	}

	// Display key-value pairs
	fmt.Printf("App Settings for App Service '%s':\n", config.AppName)
	for i, s := range settings {
		fmt.Printf("[%d] %s=%s\n", i+1, s.Name, s.Value)
	}
}

func Apply(config appconfig.AppConfig) {
	// Load local .env file
	keys, envMap, err := utils.LoadEnvFile(config.FilePath)
	if err != nil {
		log.Fatalf("[ERROR] Failed to load %s: %v\n", config.FilePath, err)
	}

	if len(keys) == 0 {
		log.Println("[WARN] No environment variables found in", config.FilePath)
		return
	}

	// Prepare arguments
	var settingsArgs []string
	for _, k := range keys {
		v := envMap[k]
		if k == "" {
			continue
		}
		arg := fmt.Sprintf("%s=%s", k, v)
		settingsArgs = append(settingsArgs, arg)
	}

	// Build the Azure CLI command
	azArgs := []string{
		"webapp", "config", "appsettings", "set",
		"--name", config.AppName,
		"--resource-group", config.ResourceGroup,
		"--settings",
	}
	azArgs = append(azArgs, settingsArgs...)

	cmdAz := exec.Command("az", azArgs...)
	out, err := cmdAz.Output()
	if err != nil {
		log.Fatalf("[ERROR] Failed to run az command:\n%v\nOutput:%s", err, string(out))
	}

	// Parse JSON output
	var updated []appconfig.AppSettings
	if err := json.Unmarshal(out, &updated); err != nil {
		log.Fatalf("[ERROR] Failed to parse JSON output: %v\nOutput: %s", err, string(out))
	}

	// Display success message
	fmt.Printf("Successfully applied %d keys to App Service '%s':\n", len(settingsArgs), config.AppName)
	for i, pair := range settingsArgs {
		fmt.Printf("[%d] %s\n", i+1, pair)
	}
}
