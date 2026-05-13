# Go Workshop — Setup Guide

---

## Windows

### 1. Install VS Code
1. Go to [code.visualstudio.com](https://code.visualstudio.com) and click **Download for Windows**
2. Run the downloaded `.exe` installer
3. Accept the licence agreement and click through the installer — the default options are fine
4. Launch VS Code when the installer finishes

### 2. Install Go
1. Go to [go.dev/dl](https://go.dev/dl) in your browser and download the latest `.msi` installer
2. Run the `.msi` installer — the default options are fine

### 3. Verify the installation
Open the terminal inside VS Code with **Ctrl + `** (backtick) or via **Terminal → New Terminal** in the menu, then run:
```
go version
```
You should see something like `go version go1.22.0 windows/amd64`. If you see an error, close and reopen the terminal and try again.

### 4. Install the Go extension
1. Open the Extensions panel with **Ctrl + Shift + X**
2. Search for **Go** and install the extension published by **Go Team at Google**
3. Open any `.go` file (or create a new one) — VS Code will show a prompt saying Go tools are missing. Click **Install All**
4. Wait for the tools to finish installing — you will see a notification when done

---

## macOS

### 1. Install VS Code
1. Go to [code.visualstudio.com](https://code.visualstudio.com) and click **Download for Mac**
2. Open the downloaded `.zip` file — this extracts the VS Code app
3. Drag **Visual Studio Code.app** into your **Applications** folder
4. Launch VS Code from Applications

### 2. Install Go
1. Go to [go.dev/dl](https://go.dev/dl) in your browser and download the latest `.pkg` installer for your Mac (choose **ARM64** if you have an M1/M2/M3 Mac, **x86-64** if you have an older Intel Mac — if unsure, click the Apple menu → About This Mac)
2. Run the `.pkg` installer — the default options are fine

### 3. Verify the installation
Open the terminal inside VS Code with **Ctrl + `** (backtick) or via **Terminal → New Terminal** in the menu, then run:
```
go version
```
You should see something like `go version go1.22.0 darwin/arm64`. If you see `command not found`, close and reopen the terminal and try again.

### 4. Install the Go extension
1. Open the Extensions panel with **Cmd + Shift + X**
2. Search for **Go** and install the extension published by **Go Team at Google**
3. Open any `.go` file (or create a new one) — VS Code will show a prompt saying Go tools are missing. Click **Install All**
4. Wait for the tools to finish installing — you will see a notification when done

---

## Linux

### 1. Install VS Code
1. Go to [code.visualstudio.com](https://code.visualstudio.com) and download the package for your distribution — `.deb` for Ubuntu/Debian, `.rpm` for Fedora/RHEL
2. For `.deb` based systems, open a terminal and run:
```
sudo dpkg -i ~/Downloads/code_*.deb
```
3. For `.rpm` based systems, run:
```
sudo rpm -i ~/Downloads/code_*.rpm
```
4. Launch VS Code from your applications menu or by running `code` in a terminal

### 2. Install Go
1. Open the terminal inside VS Code with **Ctrl + `** (backtick) or via **Terminal → New Terminal** in the menu
2. Go to [go.dev/dl](https://go.dev/dl) in your browser and copy the download URL for the latest Linux `.tar.gz`
3. In the VS Code terminal, run the following commands one at a time:
```
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```
> Replace the filename with the exact version you downloaded from go.dev/dl

4. Add Go to your PATH by running:
```
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
> If you are using zsh instead of bash, replace `.bashrc` with `.zshrc`

### 3. Verify the installation
In the VS Code terminal, run:
```
go version
```
You should see something like `go version go1.22.0 linux/amd64`. If you see `command not found`, run `source ~/.bashrc` and try again.

### 4. Install the Go extension
1. Open the Extensions panel with **Ctrl + Shift + X**
2. Search for **Go** and install the extension published by **Go Team at Google**
3. Open any `.go` file (or create a new one) — VS Code will show a prompt saying Go tools are missing. Click **Install All**
4. Wait for the tools to finish installing — you will see a notification when done
> If VS Code warns that Go tools could not be found after installation, run `source ~/.bashrc` in the terminal and restart VS Code