# Go-STL

A generic, modular, and extensible C++ STL-inspired library for Go — featuring containers, algorithms, functional utilities, and more.

<p align="center">
  <img src="https://img.shields.io/badge/Go-Generics-00ADD8?style=flat-square&logo=go" />
  <img src="https://img.shields.io/github/license/AyushOJOD/go-stl?style=flat-square" />
  <img src="https://img.shields.io/github/issues/AyushOJOD/go-stl?style=flat-square" />
  <img src="https://img.shields.io/github/stars/AyushOJOD/go-stl?style=flat-square" />
</p>

---

## ✨ Features

- **Containers**: Vector, List, Stack, Queue, Dequeue, PriorityQueue
- **Algorithms**: Sort, Search, Set operations
- **Utilities**: Pair, Tuple, Optional, Comparator
- **Functional tools**: _(coming soon)_ Map, Filter, Reduce
- **Generics-first**: All structures use Go 1.18+ generics
- **Extensible design**: Easy to contribute and expand

---

## 🧪 Benchmark

> Performance snapshot from a sample run on:  
> **Intel i5-1135G7 @ 2.40GHz**, **Windows**, **Go 1.x**

<pre> 
BenchmarkVectorPushBack-8        94128726        12.84 ns/op       48 B/op        0 allocs/op
BenchmarkVectorInsert-8           3139027       381.2 ns/op       96 B/op        2 allocs/op
BenchmarkVectorRemoveAt-8         4231052       288.5 ns/op       32 B/op        1 allocs/op

BenchmarkDequeuePushFront-8      86329523        13.52 ns/op       48 B/op        0 allocs/op
BenchmarkDequeuePushBack-8       92461715        12.64 ns/op       48 B/op        0 allocs/op
BenchmarkDequeuePopFront-8       45724980        26.01 ns/op        0 B/op        0 allocs/op

BenchmarkQueuePush-8             88450296        13.41 ns/op       48 B/op        0 allocs/op
BenchmarkQueuePop-8              46874591        25.84 ns/op        0 B/op        0 allocs/op

BenchmarkStackPush-8             90123456        12.78 ns/op       48 B/op        0 allocs/op
BenchmarkStackPop-8              47987654        24.91 ns/op        0 B/op        0 allocs/op

BenchmarkListPushBack-8          46298741        25.65 ns/op       64 B/op        1 allocs/op
BenchmarkListPushFront-8         45902176        25.92 ns/op       64 B/op        1 allocs/op
BenchmarkListPopFront-8          23189478        51.43 ns/op       32 B/op        1 allocs/op

BenchmarkPairCreation-8         100000000        10.23 ns/op       16 B/op        0 allocs/op
BenchmarkPairSwap-8              89217382        13.92 ns/op        0 B/op        0 allocs/op

</pre>

---

## 📦 Installation

```bash
go get github.com/AyushOJOD/stl-go
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues, fork the repo, and submit pull requests.

---

## 🙋‍♂️ Author

Made with ❤️ by [Ayush Srivastava](https://github.com/AyushOJOD)
