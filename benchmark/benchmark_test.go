package benchmark

import (
	"testing"
)

func Benchmark_RunTests(b *testing.B) {
	// benchmark tests
	rootChildSendTx(b)
}