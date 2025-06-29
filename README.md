![QOI Logo](https://qoiformat.org/qoi-logo.svg)

# QOI Go

[![Go Reference](https://pkg.go.dev/badge/github.com/hchargois/qoi.svg)](https://pkg.go.dev/github.com/hchargois/qoi)

Implementation of QOI encoding and decoding in Go.

For more info about QOI, see: 

- the specification at https://qoiformat.org/qoi-specification.pdf
- the reference implementation at https://github.com/phoboslab/qoi

## Performance

Even though this implementation is very straightforward and not particularly
optimised (there is not much _to_ optimise with such a simple format anyway),
I found that it is much faster than the existing Go implementations:

| compared to... | decoding | encoding |
| - | - | - |
| [takeyourhatoff/qoi](https://github.com/takeyourhatoff/qoi) | 4.3x faster | 12.4x faster |
| [xfmoulet/qoi](https://github.com/xfmoulet/qoi) | 1.43x faster | 7.1x faster |
| [arian/go-qoi](https://github.com/arian/go-qoi) | 7.5x faster | 7.5x faster |

Full benchmark comparisons can be found [here](/bench_results/results.md).

I think a lot of the difference is caused by these libraries making huge
amounts of unneeded allocations.

Compared to the reference C implem, I didn't make a full benchmark, but I have
some evidence that this package would be about 20-50% slower.
Encoding especially takes a hit because it needs to compare pixels a lot and Go
doesn't have (C-style) union types, so 4 bytes cannot be cast to an uint32 for
free to make a single comparison, it needs to do 4 compares.

The only thing I did that I would call "optimisation" (i.e. apart from doing
sensible allocations) simply consists of replacing the check for whether we can
use an OP_DIFF in the encoder (and similarly for OP_LUMA):

```
if dr >= -2 && dr <= 1 && dg >= -2 && dg <= 1 && db >= -2 && db <= 1 { ...
```

with this equivalent version:

```
if uint8((dr+2)|(dg+2)|(db+2))&0xfc == 0 { ...
```

This avoids having too many branches and for photographic images where OP_DIFF
and OP_LUMA are plentiful, this can speed up encoding by more than 10 %.
