package allocation

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func (s S) stack(s1 S) {}

func (s *S) heap(s1 *S) {}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var s S

	f, err := os.Create("stack.out")
	if err != nil {
		b.Failed()
	}
	defer func() {
		_ = f.Close()
	}()

	err = trace.Start(f)
	if err != nil {
		b.Failed()
	}

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	f, err := os.Create("heap.out")
	if err != nil {
		b.Failed()
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		b.Failed()
	}

	for i := 0; i < b.N; i++ {
		s = byPointer()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

func BenchmarkMemoryStackLoop(b *testing.B) {
	var s S
	var s1 S

	s = byCopy()
	s1 = byCopy()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := 0; t < 1000000; t++ {
			s.stack(s1)
		}
	}
}

func BenchmarkMemoryHeapLoop(b *testing.B) {
	var s *S
	var s1 *S

	s = byPointer()
	s1 = byPointer()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := 0; t < 1000000; t++ {
			s.heap(s1)
		}
	}
}
