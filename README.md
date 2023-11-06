# Dependency Report - 🚧 wip

> ## 👷🏗️ Under development
> This project is under development and is not ready for use.
---
![GitHub](https://img.shields.io/github/license/giovannymassuia/dependency-report)
![GitHub/v/tag/giovannymassuia/minimalist-java](https://img.shields.io/github/v/tag/giovannymassuia/dependency-report?label=version)
![GitHub issues](https://img.shields.io/github/issues/giovannymassuia/dependency-report)

This CLI tool aims to provide a simple way to generate a dependency report for a projects using
Maven, Gradle and Node.

## How to use

### Install

```shell
go install github.com/ivanpirog/dependency-report@latest
```

- to update: `go install -a github.com/ivanpirog/dependency-report@latest`

#### Test

```shell
dependency-report --version
# or
dependency-report --help
```

### Requirements

- Go 1.21+ (https://golang.org/dl/) **required*
    - *Make sure the `GOPATH` is set correctly,
      see [here](https://golang.org/doc/gopath_code.html#GOPATH)*
- Maven
- Gradle
- Node

### Commands

- `repo`
- `scan`

## Providers

- Bitbucket
    - [Auth](https://developer.atlassian.com/bitbucket/api/2/reference/meta/authentication)
        - Supported methods:
            - OAuth 2.0 - `https://bitbucket.org/<workspace_id>/workspace/settings/api`
                - scope required: `repository:read`
            - App Passwords - `https://bitbucket.org/account/settings/app-passwords`
        - [App Password x OAuth](https://community.atlassian.com/t5/Bitbucket-questions/Bitbucket-API-OAuth-vs-App-password/qaq-p/2193984#:~:text=The%20difference%20is%2C%20OAuth%20is,access%20is%20based%20on%20the)
        - Which one to use:
            - If you are wanting simplicity, AppPassword would be the suggested choice - if you are
              wishing to favour security at the cost of code complexity, OAuth is the suggested
              choice. [Source](https://community.atlassian.com/t5/Bitbucket-questions/Bitbucket-API-OAuth-vs-App-password/qaq-p/2193984#:~:text=If%20you%20are%20wanting%20simplicity%2C%20AppPassword%20would%20be%20the%20suggested%20choice%20%2D%20if%20you%20are%20wishing%20to%20favour%20security%20at%20the%20cost%20of%20code%20complexity%2C%20OAuth%20is%20the%20suggested%20choice.)

- Github

## Project Dependencies

- [Go](https://golang.org/)
    - [cobra](https://github.com/spf13/cobra)
    - [colored-cobra](github.com/ivanpirog/coloredcobra)
    - [go-git](https://github.com/go-git/go-git)
- See more in [go.mod](./go.mod)
