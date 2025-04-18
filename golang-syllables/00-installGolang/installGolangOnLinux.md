# 🐹 Installing Golang on Linux + Setting PATH Permanently

This guide helps you install Golang on any Linux distribution and set the `PATH` variable so the `go` command is available globally.

---

## 📦 Step 1: Download Golang

Visit the official Go downloads page:  
👉 https://go.dev/dl/

Or download directly using `wget`:

```bash
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
```

> ⚠️ Replace the version with the latest one from the website if needed.

---

## 🧹 Step 2: Remove Previous Go Installation (if any)

If you've installed Go before, remove it to avoid conflicts:

```bash
sudo rm -rf /usr/local/go
```

---

## 📂 Step 3: Extract Go to `/usr/local`

Extract the downloaded archive:

```bash
sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
```

This installs Go to `/usr/local/go`.

---

## 🔧 Step 4: Set PATH Permanently

You need to add Go’s binary path to your shell’s `PATH`.

### For Bash users:

Edit your `.bashrc` file:

```bash
nano ~/.bashrc
```

Add this line at the end:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Then apply the changes:

```bash
source ~/.bashrc
```



## ✅ Step 5: Verify Go Installation

Close and reopen your terminal, then run:

```bash
go version
```

Expected output:

```bash
go version go1.22.2 linux/amd64
```

---

## 📁 Optional: Set Up Go Workspace (GOPATH)

Modern Go uses modules, so this is optional. But if needed:

```bash
mkdir -p $HOME/go_projects
echo 'export GOPATH=$HOME/go_projects' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
```

You can now store your projects in `~/go_projects`.

---

## 🧠 Bonus Tips

- Check Go environment variables:

```bash
go env
```

- Check Go install path:

```bash
which go
```

---

## 🎉 You're Ready!

Golang is installed and ready to use. Happy coding!

