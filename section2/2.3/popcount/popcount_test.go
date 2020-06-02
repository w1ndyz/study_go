// go test -test.bench=".*"
package popcount

import "testing"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))]) +
		int(pc[byte(x>>(8*8))])
}

func PopCount1(x uint64) int {
	ret := byte(0)
	for i := uint8(0); i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}

func PopCount2(x uint64) int {
	ret := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret += 1
		}
		x = x >> 1
	}
	return ret
}

func PopCount3(x uint64) int {
	ret := 0
	for x != 0 {
		x = (x - 1) & x
		ret += 1
	}
	return ret
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// PopCount(2<<63 - 1)
		// PopCount1(2<<63 - 1)
		// PopCount2(2<<63 - 1)
		PopCount3(2<<63 - 1)
	}
}
