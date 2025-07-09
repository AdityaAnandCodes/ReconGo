# 🔎 ReconGo — Advanced Network Reconnaissance CLI Tool

**ReconGo** is a fast, extensible, and user-friendly network reconnaissance tool written in Go. Designed for security researchers, penetration testers, and network engineers, ReconGo offers subdomain discovery, TCP/UDP port scanning, banner grabbing, and local IP discovery—all from a single command-line interface.



## ✨ Features

- 🔍 Subdomain Discovery (using wordlists)
- 🔐 TCP Port Scanning with optional banner grabbing
- 📡 UDP Port Scanning
- 🧭 Common Port Auto-Scan
- 🌐 Local IP Address Enumeration
- 💻 Cross-platform Support (Windows, Linux, macOS)
- ⚡ Highly Concurrent & Fast
- 📦 Modular Codebase with Cobra CLI

## 📁 Project Structure

```
ReconGo/
├── cmd/         # Cobra CLI commands
├── finder/      # Subdomain discovery logic
├── network/     # Network utilities like local IP enumeration
├── scanner/     # TCP & UDP scanning logic
├── helpers/     # Wordlists and misc resources
├── builds/      # Compiled binaries (ignored by Git)
├── main.go      # CLI entrypoint
├── go.mod
└── README.md
```

## 🚀 Installation

### 🔧 Prerequisites

- [Go 1.20+](https://golang.org/dl/)
- Git
- (Optional) [Make](https://www.gnu.org/software/make/) or PowerShell for simplified builds

### 🏗️ Build from Source

Clone the repository:

```bash
git clone https://github.com/AdityaAnandCodes/ReconGo.git
cd ReconGo
```

**Option 1: Manual Build**
```bash
go build -o builds/recongo main.go
```

**Option 2: Using Makefile (Linux/macOS)**
```bash
make build       # Builds for current system
make build-linux # Cross-compile for Linux
make build-mac   # Cross-compile for macOS
```

**Option 3: Using PowerShell (Windows)**
```powershell
mkdir builds
go build -o builds/recongo.exe main.go
```

## 🛠️ Usage

All functionality is modular via subcommands. Run the binary from `builds/` or add it to your PATH.

```bash
./recongo --help
```

## 🧪 Examples

**🔐 TCP Port Scan**
```bash
recongo tcp --host 192.168.1.1 --range 20-80 --banner
recongo tcp --host scanme.nmap.org --ports 22,80,443
```

**📡 UDP Port Scan**
```bash
recongo udp --host 192.168.1.1 --ports 53,161,123
```

**🔍 Subdomain Discovery**
```bash
recongo subdomain --domain example.com --wordlist helpers/subdomains.txt
```

**🌐 Discover Local Network IPs**
```bash
recongo ipscan
```

## ⚙️ Configuration

All flags are CLI-based (no config files required). Common flags include:

| Flag        | Description                              |
|-------------|------------------------------------------|
| `--host`    | Target IP or hostname                    |
| `--range`   | Port range (e.g., 20-443)                |
| `--ports`   | Specific ports (e.g., 80,443,22)         |
| `--banner`  | Enable TCP banner grabbing               |
| `--domain`  | Target domain for subdomain scanning     |
| `--wordlist`| Path to wordlist for subdomains          |

## 📦 Packaging & Distribution

Compiled binaries are saved in the `builds/` folder. You can manually:

- Copy to a location in your PATH
- Or add `builds/` to your environment path

**Example (Windows PowerShell):**
```powershell
[Environment]::SetEnvironmentVariable(
  "Path",
  $env:Path + ";C:\path\to\ReconGo\builds",
  [EnvironmentVariableTarget]::User
)
```

## 🤝 Contributing

Pull requests are welcome! Please open an issue first to discuss your changes.

**Dev Setup**
```bash
go mod tidy
go run main.go
```

## 📜 License

This project is licensed under the MIT License. See LICENSE for details.

## 🙌 Acknowledgments

- Cobra CLI
- Go net package
- Inspired by tools like nmap, Amass, and masscan

## 🔗 Connect  

Built by Aditya Anand
📧 Email: adityaanandatwork276@gmail.com
