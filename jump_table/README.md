# Jump Table Benchmark

Comparing if-else chains vs map-based jump tables in Go with 100 cases.

## Results

| Approach | ns/op | Difference |
|----------|-------|------------|
| If-else | 7381 | baseline |
| JumpTable | 6287 | 15% faster |

## Run

```bash
go test -bench=. -benchmem
```
