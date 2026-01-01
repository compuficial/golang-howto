# Struct Padding Benchmark

Demonstrating struct field ordering impact on memory layout and performance.

## Struct Sizes

| Struct | Size | Difference |
|--------|------|------------|
| Inefficient | 32 bytes | baseline |
| Efficient | 16 bytes | 50% smaller |

## Benchmark Results

| Benchmark | Inefficient | Efficient | Improvement |
|-----------|-------------|-----------|-------------|
| Iterate (1M structs) | 321,597 ns | 270,523 ns | 16% faster |
| Allocate (10K structs) | 23,081 ns | 17,265 ns | 25% faster |

## Run

```bash
go run .           # show struct sizes
go test -bench=.   # run benchmarks
```
