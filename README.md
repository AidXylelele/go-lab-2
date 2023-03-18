# Continuous integration and testing automation

This is lab 2 for software architecture subject made by pogromisty

## Download

```bash
git clone https://github.com/MrPaschenko/ci-tests
```

## Run

```bash
go run cmd/example/main.go
```

Specify input by flags:

- -e "expression"
- -f file_path.txt

Specify output by:

- -o file_path.txt

## Test

```bash
go test
```

## GitHub Actions

[Every commit and pull request is checked using GitHub Actions](https://github.com/AidXylelele/go-lab-2/actions)
