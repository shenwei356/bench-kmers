package kmers

import "bytes"

func base2int_switch(b byte) (uint64, bool) {
	switch b {
	case 'a', 'A':
		return 0, true
	case 'c', 'C':
		return 1, true
	case 'g', 'G':
		return 3, true
	case 't', 'T':
		return 2, true
	}
	return 0, false
}

func base2int_nochecking(b byte) (uint64, bool) {
	return (uint64(b) >> 1) & 3, true
}

func base2int_if(b byte) (uint64, bool) {
	if b == 'a' || b == 'c' || b == 'g' || b == 't' || b == 'A' || b == 'C' || b == 'G' || b == 'T' {
		return (uint64(b) >> 1) & 3, true
	}
	return 0, false
}

var bases = []byte{'a', 'c', 'g', 't', 'A', 'C', 'G', 'T'}

func base2int_index(b byte) (uint64, bool) {
	if bytes.IndexByte(bases, b) >= 0 {
		return (uint64(b) >> 1) & 3, true
	}
	return 0, false
}

func base2int_bitwise(b byte) (uint64, bool) {
	d := b - 65
	if d <= 51 && 0x8004500080045&(1<<d) > 0 {
		return (uint64(b) >> 1) & 3, true
	}
	return 0, false
}

var base2int = [256]uint64{
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 0, 4, 1, 4, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 2, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 0, 4, 1, 4, 4, 4, 3, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 2, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
}

func base2int_array(b byte) (uint64, bool) {
	v := base2int[b]
	if v < 4 {
		return v, true
	}
	return 0, false
}
