
## Bit-packing single nucleotide bases

|Checking-ACGT|Bitwise-operation|Method              |Time(ns)|
|:------------|:----------------|:-------------------|:-------|
|             |yes              |base2int_nochecking |20.55   |
|yes          |                 |base2int_array      |24.96   |
|yes          |yes              |base2int_bitwise    |30.77   |
|yes          |yes              |base2int_if         |49.99   |
|yes          |                 |base2int_switch     |51.43   |
|yes          |yes              |base2int_bytes_index|130.8   |

Notes:

1. K-mer encoding: A=0, C=1, T=2, G=3
1. The `base2int` functions check bases, while the benchmark codes skip it.
1. go version go1.19rc2 linux/amd64

Run the bencmark

    # test
    $ go test

    # bench
    $ go test -bench Bench*

## Links

Packages:

- (Go) [kmers](https://github.com/shenwei356/kmers) - Bit-packed k-mers methods for Golang
- (Rust) [nuc2bit](https://github.com/natir/nuc2bit) - A rust crate that provides methods for rapidly encoding and decoding nucleotides in 2-bit representation.
- (C) [kmer-cnt](https://github.com/lh3/kmer-cnt/blob/master/kc-c1.c) - Code examples of fast and simple k-mer counters for tutorial purposes
- ...

Tutorials:

- [how to encode char in 2-bits?](https://stackoverflow.com/questions/39242932/how-to-encode-char-in-2-bits)