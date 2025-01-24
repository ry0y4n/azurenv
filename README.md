# azurenv

**azurenv** is a CLI tool that helps you manage environment variables for Azure App Service and Azure Functions without needing to manually copy and paste `.env` content into the Azure Portal. By leveraging this CLI, you can apply your local environment variables to Azure with just a single command, saving time and reducing potential human errors.

> Note: This CLI depends on the Azure CLI. Make sure you have the latest Azure CLI installed and are logged in (`az login`) before using azurenv.

## Installation

> Note: If you encounter an error, you may need to enable script execution. Run the following command in shell as an **administrator**:

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/ry0y4n/azurenv/main/scripts/install.sh | sh
```

### Windows

#### x86_64 (default)

```powershell
powershell -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/ry0y4n/azurenv/main/scripts/install.ps1 -OutFile install.ps1; .\install.ps1"
```

#### arm64

```powershell
powershell -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/ry0y4n/azurenv/main/scripts/install.ps1 -OutFile install.ps1; .\install.ps1 -Arch arm64"
```

#### i386

```powershell
powershell -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/ry0y4n/azurenv/main/scripts/install.ps1 -OutFile install.ps1; .\install.ps1 -Arch i386"
```

## Getting Started

1.  Install **Azure CLI** and ensure you're loggedin:

    ```bash
    az login
    ```

2.  (Optional) Set your subscription if you have multiple:

    ```bash
    az account set --subscription <SubscriptionId>
    ```

3.  Prepare your `.env` file with the environment variables you want to sync.
4.  Run the appropriate subcommand (`webapp apply`, `functionapp apply`) to synchronize.

## Commands

### `azurenv webapp list-remote`

Retrieve the current app settings for an Azure App Service.

```bash
azurenv webapp list-remote --name <WebAppName> --resource-group <ResourceGroup>
```

### `azurenv webapp diff`

Compare local environment variables (from `.env` or any specified file) with the remote Azure App Service's app settings.

```bash
azurenv webapp diff --name <WebAppName> --resource-group <ResourceGroup> --file <.env>
```

### `azurenv webapp apply`

Apply local environment variables to the Azure App Service's app settings.

```bash
azurenv webapp apply --name <WebAppName> --resource-group <ResourceGroup> --file <.env>
```

### `azurenv functionapp list-remote`

Retrieve the current app settings for an Azure Functions.

```bash
azurenv functionapp list-remote --name <FunctionAppName> --resource-group <ResourceGroup>
```

### `azurenv functionapp diff`

Compare local environment variables with the remote Azure Function's app settings.

```bash
azurenv functionapp diff --name <FunctionAppName> --resource-group <ResourceGroup> --file <.env>
```

### `azurenv functionapp apply`

Apply local environment variables to an Azure Functions.

```bash
azurenv functionapp apply --name <FunctionAppName> --resource-group <ResourceGroup> --file <.env>
```

## Other Commands

### `azurenv version`

Display the version of the azurenv CLI tool you’re currently using.

```bash
azurenv version
```

### `azurenv azcheck`

Show the Azure CLI account information you’re logged in with (e.g. subscription name, tenant, etc.). Since azurenv relies on the Azure CLI, this helps you confirm that you’re using the correct subscription and are properly authenticated.

```bash
azurenv azcheck
```

# Contributing

Contributions are welcome! Please open an issue or submit a pull request if you find a bug or have an enhancement request.

# Final Notes

- Make sure you have a working `.env` file with the correct environment variables.
- Always verify your Azure CLI login status with `azurenv azcheck`.
- Use `diff` commands to preview changes before applying them to Azure. This helps prevent accidental overwrites.

Enjoy effortless environment variable management with **azurenv**!
