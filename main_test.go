package main

import "testing"
func BenchmarkConvert(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
			convert("./color.jpg","dest.jpg")
	}
}
func BenchmarkConvert2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
			convert2("./color2.jpg","dest2.jpg")
	}
}

