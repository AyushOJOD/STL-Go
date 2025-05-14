package tests

import (
	"testing"

	"github.com/AyushOJOD/stl-go/containers"
)

// Vector Benchmarks
func BenchmarkVectorPushBack(b *testing.B) {
	v := containers.NewVector[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.PushBack(i)
	}
}

func BenchmarkVectorInsert(b *testing.B) {
	v := containers.NewVector[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Insert(0, i) // Insert at beginning
	}
}

func BenchmarkVectorRemoveAt(b *testing.B) {
	v := containers.NewVector[int]()
	for i := 0; i < b.N; i++ {
		v.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.RemoveAt(0) // Remove from beginning
	}
}

// Dequeue Benchmarks
func BenchmarkDequeuePushFront(b *testing.B) {
	d := containers.NewDequeue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}

func BenchmarkDequeuePushBack(b *testing.B) {
	d := containers.NewDequeue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
}

func BenchmarkDequeuePopFront(b *testing.B) {
	d := containers.NewDequeue[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopFront()
	}
}

// Queue Benchmarks
func BenchmarkQueuePush(b *testing.B) {
	q := containers.NewQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkQueuePop(b *testing.B) {
	q := containers.NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

// Stack Benchmarks
func BenchmarkStackPush(b *testing.B) {
	s := containers.NewStack[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	s := containers.NewStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

// List Benchmarks
func BenchmarkListPushBack(b *testing.B) {
	l := containers.NewList[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}
}

func BenchmarkListPushFront(b *testing.B) {
	l := containers.NewList[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.PushFront(i)
	}
}

func BenchmarkListPopFront(b *testing.B) {
	l := containers.NewList[int]()
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.PopFront()
	}
}

// Pair Benchmarks
func BenchmarkPairCreation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containers.NewPair(i, "test")
	}
}

func BenchmarkPairSwap(b *testing.B) {
	p := containers.NewPair(1, "test")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Swap()
	}
} 