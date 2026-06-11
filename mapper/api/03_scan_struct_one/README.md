# 03 Scan Struct One

This example shows how to scan a single row using:

```go
mapper.ScanStructOne[T](rows)
```

`ScanStructOne` returns one struct pointer.

If no row is found, it returns `mapper.ErrNoRows`.
If more than one row is found, it returns `mapper.ErrTooManyRows`.