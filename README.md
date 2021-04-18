## Golang garbage collection examples

[Garbage Collection(GC)](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) could play a vital effect on golang application performance. This repository demonstrates examples on how garbage collection works.

These examples could serve as a general reference for understanding golang gc phenomena. Different scenarios need explicit analysis to debug the performance related with golang gc though.

Each example has a short description inside its own directory. `cmd/example1` is suggested to be the first one to check, as it describes basic garbage collection process.

### Project Structure

```
github.com/smiletrl/golang_gc
|-- cmd
|   |-- example1
|   |   |-- README.md
|   |   |-- main.go
|   |-- example2
|   |   |-- README.md
|   |   |-- main.go
|-- README.md
```

1. [example1](https://github.com/smiletrl/golang_gc/tree/master/cmd/example1) shows a basic golang garbage collection example, and explains the idea with a simple http server.
2. [example2](https://github.com/smiletrl/golang_gc/tree/master/cmd/example2) shows a memory leak case, where garbage collection can't free the already allocated memory.
3. @todo, add example with gc optimization.

Result of this repository is based on go version: `go version go1.15+ darwin/amd64`

File redis.pdf was downloaded from [redislabs](https://redislabs.com/). Here it works for a file sample.

I'm happy to get feedback if you feel these examples are helpful or not-that helpful :) You are also welcome to contribute more examples to help people to understand Golang garbage collection better.
