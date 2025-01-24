param(
    [Parameter(Mandatory=$false)]
    [ValidateSet("x86_64", "arm64", "i386")]
    [string]$Arch = "x86_64"
)

$ErrorActionPreference = 'Stop'

[string]$Version

if (-not $Version) {
    Write-Host "Fetching the latest version..."
    try {
        $response = Invoke-WebRequest -Uri "https://api.github.com/repos/ry0y4n/azurenv/releases/latest" -UseBasicParsing
        $json = $response.Content | ConvertFrom-Json
        $Version = $json.tag_name
        if (-not $Version) {
            Write-Error "Failed to fetch the latest version"
            exit 1
        }
    } catch {
        Write-Error "Failed to fetch the latest version"
        exit 1
    }
}

$url = "https://github.com/ry0y4n/azurenv/releases/download/$Version/azurenv_Windows_$Arch.zip"
$output = "$env:temp\azurenv.zip"
$installDir = "$env:LocalAppData\azurenv"

Write-Host "Downloading azurenv version $Version for architecture: $Arch"
Write-Host "Download URL: $url"

Invoke-WebRequest -Uri $url -OutFile $output

Write-Host "Extracting azurenv"
Expand-Archive -Path $output -DestinationPath $installDir -Force

Write-Host "Adding azurenv to PATH"
$oldPath = [Environment]::GetEnvironmentVariable('Path', [System.EnvironmentVariableTarget]::User)
$newPath = "$oldPath;$installDir"
[Environment]::SetEnvironmentVariable('Path', $newPath, [System.EnvironmentVariableTarget]::User)

Write-Host "Installation complete. Please restart your terminal."