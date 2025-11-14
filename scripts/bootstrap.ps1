<#
!WARNING = I.A Generated
Don't ask me about it, i don't understand this part with greater details

This script installs and configures:
- OpenSSH Server
- Firewall rules for SSH
- Git
- Docker Desktop
- Deployment folder
- Deployment SSH key for GitHub Actions
#>

# -----------------------------
# 1. Install OpenSSH Server
# -----------------------------
$sshInstalled = Get-WindowsCapability -Online | Where-Object Name -like 'OpenSSH.Server*'
if ($sshInstalled.State -ne 'Installed') {
    Write-Output "Installing OpenSSH Server..."
    Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
} else {
    Write-Output "OpenSSH Server already installed."
}

# Start and auto-enable SSH service
Start-Service sshd
Set-Service -Name sshd -StartupType 'Automatic'
Write-Output "SSH service started and set to automatic."

# Allow SSH through Firewall
if (-not (Get-NetFirewallRule -DisplayName 'OpenSSH Server (sshd)')) {
    New-NetFirewallRule -Name sshd -DisplayName 'OpenSSH Server (sshd)' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 22
    Write-Output "Firewall rule created for SSH."
} else {
    Write-Output "SSH firewall rule already exists."
}

# -----------------------------
# 2. Install Git
# -----------------------------
if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
    Write-Output "Installing Git..."
    winget install --id Git.Git -e --source winget
} else {
    Write-Output "Git already installed."
}

# -----------------------------
# 3. Install Docker Desktop
# -----------------------------
if (-not (Get-Command docker -ErrorAction SilentlyContinue)) {
    Write-Output "Installing Docker Desktop..."
    winget install --id Docker.DockerDesktop -e --source winget
} else {
    Write-Output "Docker already installed."
}

# -----------------------------
# 4. Create deployment folder
# -----------------------------
$stackPath = "C:\Gole-Certo"
if (-not (Test-Path $stackPath)) { 
    New-Item -Path $stackPath -ItemType Directory
    Write-Output "Deployment folder created at $stackPath"
} else {
    Write-Output "Deployment folder already exists at $stackPath"
}

git clone https://github.com/Mateus-MS/Gole-Certo.git C:\

# -----------------------------
# 5. Create deployment SSH key
# -----------------------------
$sshDir = "$env:USERPROFILE\.ssh"
if (-not (Test-Path $sshDir)) { New-Item -Path $sshDir -ItemType Directory }

$keyPath = "$sshDir\deploy_key"
$pubKeyPath = "${keyPath}.pub"

if (-not (Test-Path $keyPath)) {
    Write-Output "Generating SSH deployment key..."
    # Windows-friendly ssh-keygen call using cmd
    cmd /c "ssh-keygen -t rsa -b 4096 -f `"$keyPath`" -N `` -C github-actions-deploy"
    
    if (Test-Path $pubKeyPath) {
        Write-Output "Deployment SSH key created at $keyPath"
        Write-Output "Public key (add to GitHub Secrets as SSH_KEY):"
        Get-Content $pubKeyPath
    } else {
        Write-Error "Failed to generate SSH key. Check ssh-keygen installation."
    }
} else {
    Write-Output "Deployment SSH key already exists at $keyPath"
}

# Add public key to authorized_keys
$authKeys = "$sshDir\authorized_keys"
if (-not (Test-Path $authKeys)) { New-Item $authKeys -ItemType File }

$pubKey = Get-Content $pubKeyPath

# Read authorized_keys content if it exists
$authContent = @()
if (Test-Path $authKeys) {
    $authContent = Get-Content $authKeys
}

# Check if the public key is already present
if ($pubKey -and -not ($authContent -contains $pubKey)) {
    $pubKey | Out-File -Append $authKeys -Encoding ascii
    Write-Output "Public key added to authorized_keys"
} else {
    Write-Output "Public key already in authorized_keys or missing"
}

Write-Output "`nBootstrap complete! You can now connect via SSH using deploy_key."
Write-Output "Use your Windows username as USERNAME and the IP from 'ipconfig' as HOST in GitHub Actions."
Write-Output "Reboot recommended if Docker Desktop was just installed."
