![Go Build](https://github.com/fglo/particles-rules-of-attraction/actions/workflows/go-build.yml/badge.svg)

# Particles' Rules of Attraction Simulator

Program simulating particles' rules of attraction. 

It's heavily inspired by [hunar4321's life_code](https://github.com/hunar4321/life_code) (it started as a port of his javascript code).

## Gif

500 particles of each color:

![gif](./img/gif.gif)

1000 particles of each color:

![gif3](./img/gif3.gif)

Different amounts of each color of particles:

* 100 red particles
* 1500 green particles
* 500 blue particles
* 1200 yellow particles

![gif4](./img/gif4.gif)

Other random results:

![gif5](./img/gif5.gif)

![gif6](./img/gif6.gif)

## Roadmap (if I have time and energy to put into it)

- [ ] User Interface: changing rules live
- [ ] User Interface: adding/removing new types (colors) of particles
- [ ] User Interface: rules import/export
- [x] computation parallelization
- [ ] rule mutations
- [ ] optimalizations :v (always ongoing)

## Running simulation

### Desktop

Run the program using `go run` command:

```bash
go run cmd/particlelifesim/main.go
```

### Web (bad performance tho)

[simulation visualization (bad performance tho)](https://fglo.github.io/particles-rules-of-attraction/index.html)

Compile the program to webasm and run a http server:

```bash
GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/particlelifesim/main.go
python3 -m http.server --cgi 8080 --directory static/ 
```

or run the program using wasmserve

```bash
go install github.com/hajimehoshi/wasmserve@latest
wasmserve ./cmd/particlelifesim/
```
