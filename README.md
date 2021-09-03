# KeySpot CLI Tool

KeySpot CLI Tool allows users to interface with the KeySpot secrets manager from the terminal. The tool is primarily used to take secrets stored with KeySpot and inject them into a process or command as environment variables. This is especially useful for CI/CD pipelines that require access to secret-based credentials, such as AWS or Google Cloud Platform's API keys. To store records with KeySpot, go to [keyspot.app](https://keyspot.app)

# Installation

The KeySpot CLI Tool can be installed on Linux, Mac, and Windows.

## Linux (Ubuntu/Debian)

```bash
curl -s --compressed "https://keyspot.github.io/cli-tool-ppa/KEY.gpg" | sudo apt-key add -
sudo curl -s --compressed -o /etc/apt/sources.list.d/keyspot.list "https://keyspot.github.io/cli-tool-ppa/keyspot.list"
sudo apt update

sudo apt install keyspot
```

## Mac

```bash
brew tap keyspot/cli

brew install keyspot
```

## Windows

```bash
scoop bucket add keyspot https://github.com/keyspot/scoop-bucket

scoop install keyspot
```

# Usage

## run

To inject a record into a process/command as environment variables, use "keyspot run".

```bash
keyspot run "npm start" -k <record-access-key>
```

### run flags

* -c: Command to inject secrets into as environment variables
* -k: Access key of record to be used

## help

At any time you can see all sub-commands keyspot has access to by passing the --help or -h flag.

```bash
keyspot -h
```

Too see the options for a specific sub-command pass the -h flag after that command.

```bash
keyspot run -h
```