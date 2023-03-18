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

Examples:

- Check is passed ([last commit](https://github.com/AidXylelele/go-lab-2/commit/28a8402df6ae15f90d9f345baebcbf7cc9abeb88))
- Check is not passed ([prelast commit](https://github.com/AidXylelele/go-lab-2/commit/8b6bc1feaab69a2abb56dd2b62dcdcd3ac0ad961))

[Pull Requests](https://github.com/AidXylelele/go-lab-2/pulls?q=is%3Apr+is%3Aclosed)