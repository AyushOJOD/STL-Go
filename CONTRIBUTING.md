
# Contributing to stl-go

Thanks for your interest in contributing to **stl-go** – a Go-based Standard Template Library inspired by C++ STL! 🚀

We welcome all contributions including bug reports, feature requests, tests, benchmarks, and new containers/algorithms.

---

## 🛠 Project Setup

To get started:

```bash
git clone https://github.com/AyushOJOD/stl-go.git
cd stl-go
go mod tidy
```

---

## 📦 Project Structure

```
stl-go/
├── containers/     # Stack, Queue, PriorityQueue, List, etc.
├── algorithm/      # Algorithms for vector/list manipulation
└── ...
```

- Add new containers in `containers/`
- Add new algorithms in `algorithm/`

---

## ✅ Code Guidelines

- Use Go generics (`type T any`) when appropriate.
- Exported types/functions should use `CamelCase`.
- Keep implementations modular and testable.
- Write unit tests in `*_test.go` files.
- Add benchmarks for performance-critical code using the `testing` package.

---

## 🧪 Running Tests and Benchmarks

Run all unit tests:

```bash
go test ./...
```

Run benchmarks:

```bash
go test -bench=. -benchmem ./...
```

---

## 🚀 Submitting a Pull Request

1. Fork this repo
2. Create a new branch: `git checkout -b feature/your-feature`
3. Make your changes with clear commit messages
4. Run tests to ensure nothing breaks
5. Push: `git push origin feature/your-feature`
6. Open a pull request on GitHub

---

## 🐛 Reporting Issues

If you discover a bug or want to request a feature, please open an [issue](https://github.com/AyushOJOD/stl-go/issues) and provide:

- A clear title and description
- Reproduction steps or use cases
- Relevant logs or test cases

---

## 🙌 Thanks

Your contributions make this project better! Whether it's code, ideas, or feedback — you're helping shape the Go STL ecosystem.

Happy coding! 💙
