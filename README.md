# Create Express CLI

A Golang-powered CLI tool to instantly scaffold a ready-to-use **Express.js** application — with optional **TypeScript** support — so you can skip repetitive setup and jump straight into coding.

# Motivation behind this project
I’m way too lazy to set up the same Express project every time.  
So instead of doing `npm init` + installing packages + writing the same boilerplate for the millionth time…  
So I made this.

## 🚀 Features
- Generate a **JavaScript** or **TypeScript** Express app in seconds.
- Preconfigured middleware, routes, and project structure.
- Automatic `npm install` after scaffolding.
- Supports multiple templates (`express-basic`, `express-ts`).
- Clean and maintainable Golang CLI architecture.

---

## 📦 Installation

### Clone and build
```bash
git clone https://github.com/yourusername/create-express-cli.git
cd create-express-cli
go build -o create-express-cli
```

## Usage
```bash
./create-express-cli create myapp -t express-basic
# or if you want typescirpt

./create-express-cli create myapp --typescript
# or shorthand if enabled
./create-express-cli create myapp -t express-ts

```


