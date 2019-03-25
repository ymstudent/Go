package main

import "testing"

func BenchmarkMain(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		main() //10000次，4.096S，每次调用342ms
	}
}
