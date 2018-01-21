# Terraform Lambda Go

## Introduction

Use Terraform to deploy AWS Lambda functions in Go

## Contents

- [Install](#install)

## Install

```bash
brew install terraform
terraform init
govendor sync
```

## Usage

```bash
GOOS=linux go build -o bin/main
terraform apply
```
