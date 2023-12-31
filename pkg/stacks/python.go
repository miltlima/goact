package stacks

const Python = `
name: Python CI

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
    - name: Setup Python
      uses: actions/setup-python@v2
      with: 
        python-version: '3.8'
    - name: Install Dependencies
      run: pip install -r requirements.txt
    - name: Run Tests
      run: pytest
`
