package main

import (
	"testing"
)

var result int

func benchmarkfibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1(n)
	}
	result = r
}

func benchmarkfibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2(n)
	}
	result = r
}

func Testfibo1(b *testing.B) {
	benchmarkfibo1(b, 30)
}

func Testfibo2(b *testing.B) {
	benchmarkfibo2(b, 30)
}
