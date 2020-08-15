Stopka is a functional, concatenative toy programming language which is clone of [min-lang](https://min-lang.org/)

##### Start REPL

```bash
$ go run cmd/stopka/main.go
```

```
> (1 2 3) (1000 +) map
[[1001 1002 1003]]

> dup
[[1001 1002 1003] [1001 1002 1003]]

> (500 -) map
[[1001 1002 1003] [501 502 503]]

> swap +
[[501 502 503 1001 1002 1003]] 
```

##### Features

- Follows the functional and concatenative programming paradigms
- Simplicity
- No external dependencies, no generators
- Parser in less than 100 lines of code
