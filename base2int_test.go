package kmers

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var seqB = []byte("ACTGactgGTCAgtcaactgGTCAACTGGTCA") // 32-bp

var codeB uint64 // avoid function inline

func Benchmark_base2int_switch(b *testing.B) {
	var code uint64
	var base byte

	for i := 0; i < b.N; i++ {
		for _, base = range seqB {
			code, _ = base2int_switch(base)
		}
		codeB = code // avoid function inline
	}
}

func Benchmark_base2int_nochecking(b *testing.B) {
	var code uint64
	var base byte

	for i := 0; i < b.N; i++ {
		for _, base = range seqB {
			code, _ = base2int_nochecking(base)
		}
		codeB = code // avoid function inline
	}
}

func Benchmark_base2int_if(b *testing.B) {
	var code uint64
	var base byte

	for i := 0; i < b.N; i++ {
		for _, base = range seqB {
			code, _ = base2int_if(base)
		}
		codeB = code // avoid function inline
	}
}

func Benchmark_base2int_index(b *testing.B) {
	var code uint64
	var base byte

	for i := 0; i < b.N; i++ {
		for _, base = range seqB {
			code, _ = base2int_index(base)
		}
		codeB = code // avoid function inline
	}
}

func Benchmark_base2int_bitwise(b *testing.B) {
	var code uint64
	var base byte

	for i := 0; i < b.N; i++ {
		for _, base = range seqB {
			code, _ = base2int_bitwise(base)
		}
		codeB = code // avoid function inline
	}
}

func Benchmark_base2int_array(b *testing.B) {
	var code uint64
	var base byte

	for i := 0; i < b.N; i++ {
		for _, base = range seqB {
			code, _ = base2int_array(base)
		}
		codeB = code // avoid function inline
	}
}

func encodeSeq(seq []byte, f func(b byte) (uint64, bool), checkBases bool) []uint64 {
	vals := make([]uint64, 0, len(seq))
	var ok bool
	var code uint64
	for _, base := range seq {
		code, ok = f(base)
		if checkBases && !ok {
			code = 4
		}

		vals = append(vals, code)
	}
	return vals
}

func equal(a []uint64, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	for i, _a := range a {
		_b := b[i]

		if _a != _b {
			return false
		}
	}

	return true
}

func Test_base2int(t *testing.T) {
	rSwitch := encodeSeq(seqB, base2int_switch, true)
	rNochecking := encodeSeq(seqB, base2int_nochecking, true)
	if !equal(rSwitch, rNochecking) {
		t.Errorf("unequal results between base2int_switch and base2int_nochecking")
	}

	rIf := encodeSeq(seqB, base2int_if, true)
	if !equal(rSwitch, rIf) {
		t.Errorf("unequal results between base2int_switch and base2int_if")
	}

	rIndex := encodeSeq(seqB, base2int_index, true)
	if !equal(rSwitch, rIndex) {
		t.Errorf("unequal results between base2int_switch and base2int_index")
	}

	rBitwise := encodeSeq(seqB, base2int_bitwise, true)
	if !equal(rSwitch, rBitwise) {
		t.Errorf("unequal results between base2int_switch and base2int_bitwise")
	}

	rArray := encodeSeq(seqB, base2int_array, true)
	if !equal(rSwitch, rArray) {
		t.Errorf("unequal results between base2int_switch and base2int_array")
	}

	out := bufio.NewWriter(os.Stderr)
	for _, base := range seqB {
		fmt.Fprintf(out, "%c ", base)
	}
	fmt.Fprintln(out)

	for _, code := range rSwitch {
		fmt.Fprintf(out, "%d ", code)
	}
	fmt.Fprintln(out)

	out.Flush()
}
