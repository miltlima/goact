package stacks

const NodeJS = `
name: Node.js CI

on: 
  push:
    branches:
	  - main
  pull_request:
	branches:
	  - main

jobs:
  build:
    runs-on: ubuntu-latest

	steps:
	- uses: actions/checkout@v2
	- name: Setup Node.js
	  uses: actions/setup-node@v2
	  with:
	    node-version: '14.x'
	- name: Install Dependencies
	  run: npm install
	- name: Run Tests
	  run: npm test
`
