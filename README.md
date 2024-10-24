# Semdapter

The Semdapter binary tool was conceived while I was working on updating workflow files across multiple repositories. Over time, this manual task became impractical, which led to the creation of Semdapter.

With Semdapter, you can specify a list of repositories, run the tool, select which repository you want to update, choose the file to push to the repository, and Semdapter will handle the rest. It will create a new branch from your base branches and generate the pull request automatically.

**NOTE:** Semdapter is currently in beta and developed for linux based systems initially. You may encounter some issues, and parts of the code use static variables, such as branch names, PR body, and title. However, feel free to modify and adapt the project to your needs. New features and updates are coming soon.

Please read the detailed instructions and prerequisites below carefully to ensure proper use of Semdapter.

### Prerequisites

Before using Semdapter, ensure the following tools are installed and configured on your local machine:

- Git

You need to have Git installed and configured. Follow the steps below to set it up:

```
sudo apt-get install git  # For Debian/Ubuntu
sudo yum install git      # For CentOS/Fedora
```

```
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"
```

- Gh

The GitHub CLI allows you to interact with GitHub directly from your terminal. Follow these steps to install and set it up:

```
sudo apt install gh   # For Debian/Ubuntu
sudo dnf install gh   # For Fedora
```

- Log in to GitHub via CLI: After installation, you need to authenticate your GitHub account:

```
gh auth login
```

- Golang

Semdapter is built with Go, so you need to have Go (Golang) installed on your machine. Here's how to install it:

```
sudo apt install golang-go   # For Debian/Ubuntu
sudo dnf install golang      # For Fedora
```

Verify golang installation:

```
go version
```

### How to use?

As you seen, sempdater is a tool built in go, so, you can use the ready binary or build from the source code.

1. Build from the source code, go to the src/ folder and run:

```
go build -o sempdater
```

- Give the necessary permissions:

```
chmod +x sempdater
```

2. Or, run the ready binary located in src/

```
./sempdater
```

Enjoy :)
