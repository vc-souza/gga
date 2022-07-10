# Go Graph Algorithms (gga)

[![GitHub Version](https://badge.fury.io/gh/vc-souza%2Fgga.svg)](https://badge.fury.io/gh/vc-souza%2Fgga)
[![CI](https://github.com/vc-souza/gga/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/vc-souza/gga/actions/workflows/ci.yml)
[![Coverage](https://coveralls.io/repos/github/vc-souza/gga/badge.svg?branch=main)](https://coveralls.io/github/vc-souza/gga?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/vc-souza/gga.svg)](https://pkg.go.dev/github.com/vc-souza/gga)
[![Go Report Card](https://goreportcard.com/badge/github.com/vc-souza/gga)](https://goreportcard.com/report/github.com/vc-souza/gga)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Graph algorithms implemented in Go.

The goal of this package is to provide implementations for both graphs and graph algorithms (alongside some common data structures used by them), which can either be used directly or as a reference when implementing your own version.

Keep in mind that some design decisions were made with ease of use and formatting support in mind, so you could always write a leaner version with such features stripped away. The asymptotic time complexity of the algorithms won't change, but you might get better constant factors. The exception would be using a better auxiliary data structure for some algorithms (like using a [Fibonacci Heap](https://en.wikipedia.org/wiki/Fibonacci_heap) in [Prim's](https://en.wikipedia.org/wiki/Prim%27s_algorithm) algorithm).

At any point, a graph can be exported to a DOT file, which can then be processed by [Graphviz](https://graphviz.org/). This makes it easy to take a snapshot of a graph before and after the execution of an algorithm, formatting the output as desired. Any particular DOT language feature that is not supported by the package can be added/modified in the resulting DOT files.

Examples can be found in the [docs](https://pkg.go.dev/github.com/vc-souza/gga).