# Go Graph Algorithms (gga)

[![GitHub Version](https://badge.fury.io/gh/vc-souza%2Fgga.svg)](https://badge.fury.io/gh/vc-souza%2Fgga)
[![CI](https://github.com/vc-souza/gga/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/vc-souza/gga/actions/workflows/ci.yml)
[![Coverage](https://coveralls.io/repos/github/vc-souza/gga/badge.svg?branch=main)](https://coveralls.io/github/vc-souza/gga?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/vc-souza/gga.svg)](https://pkg.go.dev/github.com/vc-souza/gga)
[![Go Report Card](https://goreportcard.com/badge/github.com/vc-souza/gga)](https://goreportcard.com/report/github.com/vc-souza/gga)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Graph algorithms implemented in Go.

The goal of this package is to provide implementations for both graphs and graph algorithms (alongside some common data structures used by them), which can either be used directly or as a reference when implementing your own version. Keep in mind that some design decisions were made with ease of use and formatting support in mind, so you could always write a leaner/faster version (even asymptotically so) with such features stripped away.

At any point, a graph can be exported to a DOT file, which can then be processed by [Graphviz](https://graphviz.org/). This makes it easy to take a snapshot of a graph before and after the execution of an algorithm, formatting the output as desired. Any particular DOT language feature that is not supported by the module can be added/modified in the resulting DOT files.

Examples can be found in the [samples](/internal/samples) folder.

## Algorithms

### [BFS (Breadth-First Search)](/algo/bfs.go)

##### Before
![before](/res/img/bfs/before.svg)

##### After
![after](/res/img/bfs/after.svg)

### [DFS (Depth-First Search)](/algo/dfs.go)

##### Before
![before](/res/img/dfs/before.svg)

##### After
![after](/res/img/dfs/after.svg)

### [Topological Sort](/algo/tsort.go)

##### Before
![before](/res/img/tsort/before.svg)

##### After
![after](/res/img/tsort/after.svg)

### [Connected Components (Undirected)](/algo/cc.go)

##### Before
![before](/res/img/cc/before.svg)

##### After
![after](/res/img/cc/after.svg)

### [Strongly Connected Components (Directed)](/algo/scc.go)

##### Before
![before](/res/img/scc/before.svg)

##### After
![after](/res/img/scc/after.svg)

### [Condensation Graph](/algo/gscc.go)

#### Before
![before](/res/img/gscc/before.svg)

#### Strongly Connected Components (Tarjan's)
![after scc](/res/img/gscc/after-scc.svg)

##### After
![after](/res/img/gscc/after.svg)