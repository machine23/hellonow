# HelloNow

HelloNow is a simple program written in Go. It just prints current / exact time using the NTP package.

## Build

You have to have latest stable `go` toolchain installed.
To build a single binary for direct execution, use `make build` (or `go build -o hellonow ./app`). This command will produce the executable `hellonow` file.

## Run

To run this program - `make run` (or `go run ./...`). It starts the program and displays current time.

```
$ make run
go run ./...
Current time: 2019-05-29 22:12:32.070176951 +0300 EEST m=+0.037493125
```
