# Mazoid
**Mazoid** is a simple maze generator application written in [Go](https://go.dev/) using the [Raylib](https://github.com/gen2brain/raylib-go) package. It creates new mazes using a recursive depth-first algorithm and solves them using a breadth-first algoritm.

### Building
##### Using Taskfile
If you have [Taskfile](https://taskfile.dev/) installed, simply run `task build` to generate a system specific binary or executable (currently works for Linux and Windows).
##### Using Go
Run `go build -o bin/mazoid main.go` if you are on Linux, `go build -o bin/mazoid.exe main.go` if you are using Windows system.

The runnable files will be placed in the `bin` directory.

### Possible improvements
- Generation settings
- New generator algorithms
- Depth-First search