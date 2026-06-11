# 02 Scan Struct Slice

This example shows how to scan all rows into a slice using:

```go
mapper.ScanStructSlice[T](rows)
```

`ScanStructSlice` is a convenience wrapper around `ScanStructRows`.

Use it when you want to collect all rows into a `[]T` without writing the callback and append logic yourself.
