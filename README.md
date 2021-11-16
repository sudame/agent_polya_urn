## Requirements
- Go (1.17 or later)
- gcc (code depends on [cgo](https://go.dev/blog/cgo), which require gcc)
  - memo: on Windows, we use [mingw-w64](https://winlibs.com/) for gcc

## How to run
### TL;DR
```sh
go run .
```

### parameters
```
-iter int
      iteration (default 1000)
-nu int
      ν (default 1)
-rho int
      ρ (default 1)
```

example

```sh
# ν=5, ρ=15, iterations=10000 steps
go run . -nu=5 -rho=15 -iter=10000
```
