package stacks

const Ruby = `name: CI

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
    - name: Setup Ruby
      uses: actions/setup-ruby@v1
      with: 
        ruby-version: '2.7'
    - name: Install Dependencies
      run: bundle install
    - name: Run Tests
      run: bundle exec rspec
`
