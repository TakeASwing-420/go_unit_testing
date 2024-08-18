package main

import (
    "testing"
)

func BenchmarkConcatStrings(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ConcatStrings("hello", "world")
    }
}

func BenchmarkConcatStringsWithAllocations(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        ConcatStrings("hello", "world")
    }
}

func BenchmarkSub(b *testing.B) {
    b.Run("ShortStrings", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            ConcatStrings("a", "b")
        }
    })

    b.Run("LongStrings", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            ConcatStrings("this is a long string", "and another long string")
        }
    })
}
