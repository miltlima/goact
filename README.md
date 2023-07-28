# Goact Documentation

Goact is a command-line tool (CLI) that enables you to create pipeline stacks to automate workflows on GitHub Actions.

## Installation

To use Goact, follow the steps below to install it on your system:

1. Download the Goact source code or clone the GitHub repository.

2. Navigate to the Goact project directory.

3. Compile the Goact code to create the executable:

```bash
go build
```

Move the executable to a folder that is in your system's PATH:

```bash
sudo mv goact /usr/local/bin
```

## Usage

Goact is designed to make it easy to create pipeline stacks in GitHub Actions. With the create command, you can configure GitHub Actions and generate essential files for your actions.

## Command create

The create command is used to create the GitHub Actions configuration for a specific pipeline stack.

```bash
goact create --stack <STACK_NAME>
```

Replace <STACK_NAME> with the name of the stack you want to create, for example, node.js, python, java, golang or ruby.

```bash
goact create --stack node.js
```

## Optional

Optionally, you can create dockerfile with stack desired for you projects per example you choose ruby as a stack add the flag -d will create a dockerfile for ruby

```bash
goact create --stack ruby -d 
```
