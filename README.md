# htmlcoder-go

Blogpost: <https://www.kaper.com/software/first-try-at-golang-htmlcoder-utility/>

## Description

Written by me as first attempt at GoLang.

This is a simple command line tool, which you can use to html encode or decode some data.
You can either use this as a shell (pipe) filter, or you can process a file, or you can
pass in the data to be encoded/decoded as command line arguments.

## Usage

Example use:

```shell
# show help:
htmlcoder -h

# decode command line data:
htmlcoder -d -c -- "&#34;hello&#34;" test

# decode a file:
htmlcoder -d -f /tmp/somefile.html

# encode output from another command:
echo '"hello" test' | htmlcoder
```

## Build

Build using:

```shell
# compile
go build htmlcoder.go

# optional: put the tool in your path for ease of use
sudo install -o root -g root -m 755 htmlcoder /usr/local/bin/
```

## Used resources which got me started

I am running Linux Mint, which is based on Ubuntu.

GoLang installation:

I installed GoLang using the installation instructions from: <https://golang.org/doc/install>
And I added these lines to my ~/.profile:

```shell
export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin"
export GOPATH="$HOME/go"
export GOROOT="/usr/local/go"
```

Development environment / "editor": my first pitfall - do NOT install Visual Studio Code as "flatpak",
as that will complicate getting GoLang support working. Just install it as a "normal" package.
I used this instruction: <https://techviewleo.com/install-visual-studio-code-on-linux-mint/>

```shell
# Used on Linux Mint 20 (Find proper manual for your OS!)

# Update packages, and make sure you have apt-transport-https installed
sudo apt update
sudo apt install apt-transport-https

# Add microsoft as package source
curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > /tmp/microsoft.gpg
sudo install -o root -g root -m 644 /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'

# Update package list, and install "code" (= Visual Studio Code)
sudo apt update
sudo apt install code
```

After installing Visual Studio, it did ask at some point in time if I wanted GoLang support (I think when opening a *.go file). I just clicked yes, and "download all" to get that up and running.

Note: I also tried using IntelliJ, as that is the development environment I use daily, but I could not get the GoLang plugin to work 100%. Probably some setting I can not find yet. VSC does work without issues for now.

Internet resources:

- <https://tour.golang.org/welcome/1> -> introduction to the language

- <https://gobyexample.com/> -> another introduction, using some examples

- <https://golang.org/doc/> -> documentation

- <https://pkg.go.dev/> -> package documentation (the stuff you "import")

- <https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code> -> the manual I used to setup Visual Studio Code debugging for my first example

From the above links, I completely went through the first two links to get a basic understanding of the language.
And after that, I went on to write my first little go tool, as show in this project.
I am well aware that I'm not using the module stuff yet, and not using fancy tools for dependency
management, but that is maybe something to look for in the next project.

Thijs Kaper, November 6, 2021.
