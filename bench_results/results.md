# Benchmark results

## Decoding

```
goos: linux
goarch: amd64
pkg: github.com/hchargois/qoi
cpu: AMD Ryzen 7 3700X 8-Core Processor             
                            │ bench_takeyourhatoff │           bench_hchargois           │
                            │        sec/op        │   sec/op     vs base                │
Decode/dice.qoi-16                     8.316m ± 3%   1.598m ± 4%  -80.78% (p=0.000 n=10)
Decode/edgecase.qoi-16                 95.30µ ± 1%   27.60µ ± 1%  -71.04% (p=0.000 n=10)
Decode/kodim10.qoi-16                 18.572m ± 0%   3.797m ± 2%  -79.56% (p=0.000 n=10)
Decode/kodim23.qoi-16                 18.817m ± 1%   3.771m ± 1%  -79.96% (p=0.000 n=10)
Decode/qoi_logo.qoi-16                 492.7µ ± 1%   153.6µ ± 3%  -68.82% (p=0.000 n=10)
Decode/testcard.qoi-16                 701.1µ ± 1%   164.2µ ± 4%  -76.59% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16            714.2µ ± 1%   168.6µ ± 2%  -76.39% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16           42.293m ± 0%   9.315m ± 1%  -77.97% (p=0.000 n=10)
geomean                                2.708m        630.7µ       -76.71%

                            │ bench_takeyourhatoff │            bench_hchargois            │
                            │         B/op         │     B/op       vs base                │
Decode/dice.qoi-16                    3.306Mi ± 0%    4.256Mi ± 0%  +28.73% (p=0.000 n=10)
Decode/edgecase.qoi-16                70.54Ki ± 0%    71.88Ki ± 0%   +1.89% (p=0.000 n=10)
Decode/kodim10.qoi-16                 3.124Mi ± 0%    4.592Mi ± 0%  +46.99% (p=0.000 n=10)
Decode/kodim23.qoi-16                 3.213Mi ± 0%    4.592Mi ± 0%  +42.93% (p=0.000 n=10)
Decode/qoi_logo.qoi-16                432.6Ki ± 0%    474.4Ki ± 0%   +9.67% (p=0.000 n=10)
Decode/testcard.qoi-16                301.6Ki ± 0%    366.4Ki ± 0%  +21.47% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16           309.2Ki ± 0%    366.4Ki ± 0%  +18.50% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16           7.381Mi ± 0%   11.780Mi ± 0%  +59.60% (p=0.000 n=10)
geomean                               967.3Ki         1.203Mi       +27.41%

                            │ bench_takeyourhatoff │           bench_hchargois           │
                            │      allocs/op       │ allocs/op   vs base                 │
Decode/dice.qoi-16                  239023.00 ± 0%   28.00 ± 0%   -99.99% (p=0.000 n=10)
Decode/edgecase.qoi-16               2101.000 ± 0%   9.000 ± 0%   -99.57% (p=0.000 n=10)
Decode/kodim10.qoi-16               618938.00 ± 0%   29.00 ± 0%  -100.00% (p=0.000 n=10)
Decode/kodim23.qoi-16               625324.00 ± 0%   29.00 ± 0%  -100.00% (p=0.000 n=10)
Decode/qoi_logo.qoi-16                9563.00 ± 0%   16.00 ± 0%   -99.83% (p=0.000 n=10)
Decode/testcard.qoi-16               18512.00 ± 0%   17.00 ± 0%   -99.91% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16          18818.00 ± 0%   17.00 ± 0%   -99.91% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16        1384625.00 ± 0%   33.00 ± 0%  -100.00% (p=0.000 n=10)
geomean                                73.97k        20.59        -99.97%
```

```
goos: linux
goarch: amd64
pkg: github.com/hchargois/qoi
cpu: AMD Ryzen 7 3700X 8-Core Processor             
                            │ bench_xfmoulet │           bench_hchargois           │
                            │     sec/op     │   sec/op     vs base                │
Decode/dice.qoi-16               2.546m ± 2%   1.598m ± 4%  -37.22% (p=0.000 n=10)
Decode/edgecase.qoi-16           42.76µ ± 2%   27.60µ ± 1%  -35.45% (p=0.000 n=10)
Decode/kodim10.qoi-16            4.982m ± 2%   3.797m ± 2%  -23.78% (p=0.000 n=10)
Decode/kodim23.qoi-16            4.953m ± 1%   3.771m ± 1%  -23.86% (p=0.000 n=10)
Decode/qoi_logo.qoi-16           234.9µ ± 1%   153.6µ ± 3%  -34.61% (p=0.000 n=10)
Decode/testcard.qoi-16           226.5µ ± 1%   164.2µ ± 4%  -27.51% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16      240.7µ ± 1%   168.6µ ± 2%  -29.95% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16     12.745m ± 1%   9.315m ± 1%  -26.91% (p=0.000 n=10)
geomean                          902.0µ        630.7µ       -30.09%

                            │ bench_xfmoulet │            bench_hchargois             │
                            │      B/op      │     B/op       vs base                 │
Decode/dice.qoi-16              1.840Mi ± 0%    4.256Mi ± 0%  +131.31% (p=0.000 n=10)
Decode/edgecase.qoi-16          68.22Ki ± 0%    71.88Ki ± 0%    +5.35% (p=0.000 n=10)
Decode/kodim10.qoi-16           1.504Mi ± 0%    4.592Mi ± 0%  +205.30% (p=0.000 n=10)
Decode/kodim23.qoi-16           1.504Mi ± 0%    4.592Mi ± 0%  +205.30% (p=0.000 n=10)
Decode/qoi_logo.qoi-16          396.2Ki ± 0%    474.4Ki ± 0%   +19.72% (p=0.000 n=10)
Decode/testcard.qoi-16          260.2Ki ± 0%    366.4Ki ± 0%   +40.79% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16     260.2Ki ± 0%    366.4Ki ± 0%   +40.79% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16     3.778Mi ± 0%   11.780Mi ± 0%  +211.83% (p=0.000 n=10)
geomean                         649.5Ki         1.203Mi        +89.76%

                            │ bench_xfmoulet │           bench_hchargois            │
                            │   allocs/op    │  allocs/op   vs base                 │
Decode/dice.qoi-16                7.000 ± 0%   28.000 ± 0%  +300.00% (p=0.000 n=10)
Decode/edgecase.qoi-16            7.000 ± 0%    9.000 ± 0%   +28.57% (p=0.000 n=10)
Decode/kodim10.qoi-16             7.000 ± 0%   29.000 ± 0%  +314.29% (p=0.000 n=10)
Decode/kodim23.qoi-16             7.000 ± 0%   29.000 ± 0%  +314.29% (p=0.000 n=10)
Decode/qoi_logo.qoi-16            7.000 ± 0%   16.000 ± 0%  +128.57% (p=0.000 n=10)
Decode/testcard.qoi-16            7.000 ± 0%   17.000 ± 0%  +142.86% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16       7.000 ± 0%   17.000 ± 0%  +142.86% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16       7.000 ± 0%   33.000 ± 0%  +371.43% (p=0.000 n=10)
geomean                           7.000         20.59       +194.18%
```

```
goos: linux
goarch: amd64
pkg: github.com/hchargois/qoi
cpu: AMD Ryzen 7 3700X 8-Core Processor             
                            │ bench_arian  │           bench_hchargois           │
                            │    sec/op    │   sec/op     vs base                │
Decode/dice.qoi-16            13.570m ± 1%   1.598m ± 4%  -88.22% (p=0.000 n=10)
Decode/edgecase.qoi-16        400.41µ ± 1%   27.60µ ± 1%  -93.11% (p=0.000 n=10)
Decode/kodim10.qoi-16         13.947m ± 0%   3.797m ± 2%  -72.78% (p=0.000 n=10)
Decode/kodim23.qoi-16         14.023m ± 1%   3.771m ± 1%  -73.11% (p=0.000 n=10)
Decode/qoi_logo.qoi-16        2373.9µ ± 1%   153.6µ ± 3%  -93.53% (p=0.000 n=10)
Decode/testcard.qoi-16        1633.4µ ± 2%   164.2µ ± 4%  -89.95% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16   1682.0µ ± 1%   168.6µ ± 2%  -89.97% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16   35.930m ± 1%   9.315m ± 1%  -74.07% (p=0.000 n=10)
geomean                        4.726m        630.7µ       -86.66%

                            │  bench_arian  │            bench_hchargois            │
                            │     B/op      │     B/op       vs base                │
Decode/dice.qoi-16             3.671Mi ± 0%    4.256Mi ± 0%  +15.94% (p=0.000 n=10)
Decode/edgecase.qoi-16        132.17Ki ± 0%    71.88Ki ± 0%  -45.62% (p=0.000 n=10)
Decode/kodim10.qoi-16          3.004Mi ± 0%    4.592Mi ± 0%  +52.86% (p=0.000 n=10)
Decode/kodim23.qoi-16          3.004Mi ± 0%    4.592Mi ± 0%  +52.86% (p=0.000 n=10)
Decode/qoi_logo.qoi-16         781.2Ki ± 0%    474.4Ki ± 0%  -39.27% (p=0.000 n=10)
Decode/testcard.qoi-16         516.2Ki ± 0%    366.4Ki ± 0%  -29.02% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16    516.2Ki ± 0%    366.4Ki ± 0%  -29.02% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16    7.548Mi ± 0%   11.780Mi ± 0%  +56.06% (p=0.000 n=10)
geomean                        1.258Mi         1.203Mi        -4.30%

                            │  bench_arian   │           bench_hchargois           │
                            │   allocs/op    │ allocs/op   vs base                 │
Decode/dice.qoi-16            480008.00 ± 0%   28.00 ± 0%   -99.99% (p=0.000 n=10)
Decode/edgecase.qoi-16        16392.000 ± 0%   9.000 ± 0%   -99.95% (p=0.000 n=10)
Decode/kodim10.qoi-16         393224.00 ± 0%   29.00 ± 0%   -99.99% (p=0.000 n=10)
Decode/kodim23.qoi-16         393224.00 ± 0%   29.00 ± 0%   -99.99% (p=0.000 n=10)
Decode/qoi_logo.qoi-16         98568.00 ± 0%   16.00 ± 0%   -99.98% (p=0.000 n=10)
Decode/testcard.qoi-16         65544.00 ± 0%   17.00 ± 0%   -99.97% (p=0.000 n=10)
Decode/testcard_rgba.qoi-16    65544.00 ± 0%   17.00 ± 0%   -99.97% (p=0.000 n=10)
Decode/wikipedia_008.qoi-16   988424.00 ± 0%   33.00 ± 0%  -100.00% (p=0.000 n=10)
geomean                          163.4k        20.59        -99.99%
```

## Encoding

```
goos: linux
goarch: amd64
pkg: github.com/hchargois/qoi
cpu: AMD Ryzen 7 3700X 8-Core Processor             
                            │ bench_takeyourhatoff │           bench_hchargois           │
                            │        sec/op        │   sec/op     vs base                │
Encode/dice.qoi-16                    26.951m ± 0%   2.127m ± 0%  -92.11% (p=0.000 n=10)
Encode/edgecase.qoi-16                752.33µ ± 0%   45.21µ ± 0%  -93.99% (p=0.000 n=10)
Encode/kodim10.qoi-16                 39.133m ± 0%   3.858m ± 1%  -90.14% (p=0.000 n=10)
Encode/kodim23.qoi-16                 39.651m ± 0%   3.898m ± 2%  -90.17% (p=0.000 n=10)
Encode/qoi_logo.qoi-16                4445.5µ ± 0%   312.6µ ± 0%  -92.97% (p=0.000 n=10)
Encode/testcard.qoi-16                3440.5µ ± 0%   232.8µ ± 1%  -93.23% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16           3487.5µ ± 0%   247.8µ ± 0%  -92.89% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16            97.84m ± 1%   11.59m ± 0%  -88.15% (p=0.000 n=10)
geomean                                10.64m        861.1µ       -91.91%

                            │ bench_takeyourhatoff │           bench_hchargois            │
                            │         B/op         │     B/op      vs base                │
Encode/dice.qoi-16                 4759.360Ki ± 0%   6.287Ki ± 0%  -99.87% (p=0.000 n=10)
Encode/edgecase.qoi-16              140.338Ki ± 0%   4.016Ki ± 0%  -97.14% (p=0.000 n=10)
Encode/kodim10.qoi-16              5763.117Ki ± 0%   8.132Ki ± 1%  -99.86% (p=0.000 n=10)
Encode/kodim23.qoi-16               5830.15Ki ± 0%   12.36Ki ± 2%  -99.79% (p=0.000 n=10)
Encode/qoi_logo.qoi-16              813.074Ki ± 0%   4.024Ki ± 0%  -99.51% (p=0.000 n=10)
Encode/testcard.qoi-16              603.240Ki ± 0%   4.029Ki ± 0%  -99.33% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16         603.586Ki ± 0%   4.030Ki ± 0%  -99.33% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16        14374.50Ki ± 0%   59.01Ki ± 9%  -99.59% (p=0.000 n=10)
geomean                               1.727Mi        7.479Ki       -99.58%

                            │ bench_takeyourhatoff │           bench_hchargois           │
                            │      allocs/op       │ allocs/op   vs base                 │
Encode/dice.qoi-16                1228832.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/edgecase.qoi-16              36143.000 ± 0%   2.000 ± 0%   -99.99% (p=0.000 n=10)
Encode/kodim10.qoi-16             1794550.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/kodim23.qoi-16             1808800.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/qoi_logo.qoi-16             209180.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/testcard.qoi-16             163163.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16        162697.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16       4448864.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                494.3k        2.000       -100.00%
```

```
goos: linux
goarch: amd64
pkg: github.com/hchargois/qoi
cpu: AMD Ryzen 7 3700X 8-Core Processor             
                            │ bench_xfmoulet │           bench_hchargois           │
                            │     sec/op     │   sec/op     vs base                │
Encode/dice.qoi-16              17.038m ± 0%   2.127m ± 0%  -87.52% (p=0.000 n=10)
Encode/edgecase.qoi-16          531.75µ ± 0%   45.21µ ± 0%  -91.50% (p=0.000 n=10)
Encode/kodim10.qoi-16           17.319m ± 0%   3.858m ± 1%  -77.73% (p=0.000 n=10)
Encode/kodim23.qoi-16           17.630m ± 0%   3.898m ± 2%  -77.89% (p=0.000 n=10)
Encode/qoi_logo.qoi-16          3200.2µ ± 0%   312.6µ ± 0%  -90.23% (p=0.000 n=10)
Encode/testcard.qoi-16          2184.2µ ± 0%   232.8µ ± 1%  -89.34% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16     2205.3µ ± 0%   247.8µ ± 0%  -88.76% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16      43.47m ± 0%   11.59m ± 0%  -73.34% (p=0.000 n=10)
geomean                          6.074m        861.1µ       -85.82%

                            │ bench_xfmoulet  │           bench_hchargois            │
                            │      B/op       │     B/op      vs base                │
Encode/dice.qoi-16            1894.105Ki ± 0%   6.287Ki ± 0%  -99.67% (p=0.000 n=10)
Encode/edgecase.qoi-16          68.087Ki ± 0%   4.016Ki ± 0%  -94.10% (p=0.000 n=10)
Encode/kodim10.qoi-16         1570.607Ki ± 0%   8.132Ki ± 1%  -99.48% (p=0.000 n=10)
Encode/kodim23.qoi-16          1571.07Ki ± 0%   12.36Ki ± 2%  -99.21% (p=0.000 n=10)
Encode/qoi_logo.qoi-16         389.164Ki ± 0%   4.024Ki ± 0%  -98.97% (p=0.000 n=10)
Encode/testcard.qoi-16         260.197Ki ± 0%   4.029Ki ± 0%  -98.45% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16    260.198Ki ± 0%   4.030Ki ± 0%  -98.45% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16    4022.48Ki ± 0%   59.01Ki ± 9%  -98.53% (p=0.000 n=10)
geomean                          654.6Ki        7.479Ki       -98.86%

                            │ bench_xfmoulet  │           bench_hchargois           │
                            │    allocs/op    │ allocs/op   vs base                 │
Encode/dice.qoi-16            480009.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/edgecase.qoi-16         16393.000 ± 0%   2.000 ± 0%   -99.99% (p=0.000 n=10)
Encode/kodim10.qoi-16         393225.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/kodim23.qoi-16         393225.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/qoi_logo.qoi-16         98569.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/testcard.qoi-16         65545.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16    65545.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16   988425.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                           163.4k        2.000       -100.00%
```

```
goos: linux
goarch: amd64
pkg: github.com/hchargois/qoi
cpu: AMD Ryzen 7 3700X 8-Core Processor             
                            │ bench_arian  │           bench_hchargois           │
                            │    sec/op    │   sec/op     vs base                │
Encode/dice.qoi-16            17.052m ± 1%   2.127m ± 0%  -87.53% (p=0.000 n=10)
Encode/edgecase.qoi-16        507.70µ ± 0%   45.21µ ± 0%  -91.10% (p=0.000 n=10)
Encode/kodim10.qoi-16         21.202m ± 0%   3.858m ± 1%  -81.80% (p=0.000 n=10)
Encode/kodim23.qoi-16         21.440m ± 0%   3.898m ± 2%  -81.82% (p=0.000 n=10)
Encode/qoi_logo.qoi-16        2927.4µ ± 0%   312.6µ ± 0%  -89.32% (p=0.000 n=10)
Encode/testcard.qoi-16        2233.4µ ± 0%   232.8µ ± 1%  -89.58% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16   2222.5µ ± 0%   247.8µ ± 0%  -88.85% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16    53.88m ± 0%   11.59m ± 0%  -78.49% (p=0.000 n=10)
geomean                        6.473m        861.1µ       -86.70%

                            │   bench_arian   │           bench_hchargois            │
                            │      B/op       │     B/op      vs base                │
Encode/dice.qoi-16            3025.899Ki ± 0%   6.287Ki ± 0%  -99.79% (p=0.000 n=10)
Encode/edgecase.qoi-16          69.781Ki ± 0%   4.016Ki ± 0%  -94.25% (p=0.000 n=10)
Encode/kodim10.qoi-16         2966.109Ki ± 0%   8.132Ki ± 1%  -99.73% (p=0.000 n=10)
Encode/kodim23.qoi-16          2933.17Ki ± 0%   12.36Ki ± 2%  -99.58% (p=0.000 n=10)
Encode/qoi_logo.qoi-16         418.848Ki ± 0%   4.024Ki ± 0%  -99.04% (p=0.000 n=10)
Encode/testcard.qoi-16         299.399Ki ± 0%   4.029Ki ± 0%  -98.65% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16    304.528Ki ± 0%   4.030Ki ± 0%  -98.68% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16    7780.67Ki ± 0%   59.01Ki ± 9%  -99.24% (p=0.000 n=10)
geomean                          927.1Ki        7.479Ki       -99.19%

                            │   bench_arian    │           bench_hchargois           │
                            │    allocs/op     │ allocs/op   vs base                 │
Encode/dice.qoi-16             616215.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/edgecase.qoi-16          18481.000 ± 0%   2.000 ± 0%   -99.99% (p=0.000 n=10)
Encode/kodim10.qoi-16          778553.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/kodim23.qoi-16          775640.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/qoi_logo.qoi-16         105815.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/testcard.qoi-16          80504.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/testcard_rgba.qoi-16     80144.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
Encode/wikipedia_008.qoi-16   1974394.000 ± 0%   2.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                            234.9k        2.000       -100.00%
```
