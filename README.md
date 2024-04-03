<h1 align="center"><i>DoCentre</i></h1>

<div align="center">

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/DoCentre/DoCentre)
[![CI](https://github.com/DoCentre/DoCentre/actions/workflows/ci.yml/badge.svg)](https://github.com/DoCentre/DoCentre/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/DoCentre/DoCentre/graph/badge.svg?token=VE3MI85NDN)](https://codecov.io/gh/DoCentre/DoCentre)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

</div>

---

<p align="center">
    <i>A unified platform streamlining viewing, editing, and approval processes for complex documents across all sectors.</i>
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [API Endpoints](#api)
- [Running the Tests](#tests)
- [Development](#development)
- [License](#license)

## 🧐 About <a name = "about"></a>

_DoCentre_, **Do**cument **Centre**, is a robust document management platform designed to streamline the complexities of document handling within enterprises. From production processes to machine specifications, _DoCentre_ offers a unified solution for managing diverse document types efficiently. With features facilitating viewing, editing, and approval processes all in one place, _DoCentre_ empowers users to seamlessly collaborate and ensure document integrity.

> [!note]
> _DoCentre_ is developed as the course project for the "Cloud Native Development: toward best practice" course at National Yang Ming Chiao Tung University (NYCU).

## 🏁 Getting Started <a name = "getting_started"></a>

### Prerequisites

- The _Go_ programming language (version 1.22 or later is recommended) is required to run the server. You can download and install Go from the [official website](https://go.dev/doc/install).

### Starting the Server

1. Clone the repository (or download the _ZIP_ file):

```console
$ git clone git@github.com:DoCentre/DoCentre.git
# or extract the zip file
$ unzip DoCentre-main.zip
```

2. Change into the project directory:

```console
$ cd DoCentre
# or
$ cd DoCentre-main
```

3. Run the server:

```console
$ go run main.go
```

The server should now be running on `localhost:8080`:

```console
$ curl http://localhost:8080/health
{"message":"health check success"}
```

## 🧾 API Endpoints <a name = "api"></a>

### Health Check

- URL: `/health`
- Method: `GET`
- Description: Check the health of the server.
- Response:
  - `200 OK`: `{"message":"health check success"}`
  - `4xx` or `5xx`

## 🔧 Running the Tests <a name = "tests"></a>

_DoCentre_ uses the `testing` package in _Go_ to write tests.
To run the tests, execute the following command:

```console
$ go test ./...
```

Alternatively, you can run the tests with coverage:

```console
$ go test -cover ./...
```

We also provide _Make_ commands for running tests:

```console
$ make test
# or
$ make test-coverage
```

## 🚀 Development <a name = "development"></a>

### Prerequisites

- [Make](https://www.gnu.org/software/make/#download) is used to gather the required tools and commands for development.
- [pre-commit](https://pre-commit.com/#install) is used to run checks before committing changes.

Install the required tools by running the following command:

```console
$ make tools
```

Several other _Make_ commands are available for development:

```console
$ make help
Usage: make <target>

Targets:
  test            Run tests
  test-coverage   Run tests with coverage
  fmt             Format code
  fmt-check       Check code format
  vet             Run go vet
  lint            Run staticcheck
  misspell-check  Check spelling
  misspell        Fix spelling
  tools           Install tools
  help            Show this help message

```

### Git Hooks

Install _pre-commit_ hooks:

```console
$ pre-commit install
pre-commit installed at .git/hooks/pre-commit
pre-commit installed at .git/hooks/commit-msg
```

We have two _pre-commit_ hooks:
- `pre-commit`: runs checks before committing changes
- `commit-msg`: checks the commit message format

> [!note]
> We follow [Conventional Commits](https://www.conventionalcommits.org) for commit messages.

If the `pre-commit` hook is installed successfully, you can run the following command to check all files:

```console
$ pre-commit run --all-files
```

> [!note]
> Commit hooks can be bypassed by adding the `--no-verify` (`-n`) flag to the `git commit` command.
> You may use it if you need fast and dirty commits. However, do not forget to pass the checks before pushing the changes.

## ✍️ License <a name = "license"></a>

_DoCentre_ is licensed under the [MIT License](LICENSE).
