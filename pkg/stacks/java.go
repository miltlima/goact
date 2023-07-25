package stacks

const Java = `
name: CI

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
    - name: Set up JDK
      uses: actions/setup-java@v1
      with: 
        java-version: '11'
    - name: Build with Gradle
      run: ./gradlew build
    - name: Run Tests
      run: ./gradlew test   

`