# fastabs

![seven minute abs](https://memecrunch.com/meme/629QU/seven-minute-abs/image.png)

So I did a thing that required absolute value and I just did it by checking if the value was negative and multiplying. It seems like some bit flipping and no branching could be cool, so I stole [code from StackOverflow](https://stackoverflow.com/a/2074403/14101698), ported it to Go, and tried it out. Turns out, the compiler figured out my little ruse and the difference between the naive "invert if negative" and fancy-pants "twiddle the bits" version is lower than the noise of other stuff running on my laptop. We're talking sub-nanosecond anyway.

```text
goos: darwin
goarch: amd64
pkg: github.com/dangermike/fastabs
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkFastAbs/round_0/slowabs-8   1000000000           0.2851 ns/op
BenchmarkFastAbs/round_0/fastabs-8   1000000000           0.2804 ns/op
BenchmarkFastAbs/round_1/slowabs-8   1000000000           0.2811 ns/op
BenchmarkFastAbs/round_1/fastabs-8   1000000000           0.2823 ns/op
BenchmarkFastAbs/round_2/slowabs-8   1000000000           0.2806 ns/op
BenchmarkFastAbs/round_2/fastabs-8   1000000000           0.2804 ns/op
BenchmarkFastAbs/round_3/slowabs-8   1000000000           0.2792 ns/op
BenchmarkFastAbs/round_3/fastabs-8   1000000000           0.2806 ns/op
PASS
ok    github.com/dangermike/fastabs  2.745s
```
