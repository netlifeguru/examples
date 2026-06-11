# 01 Scan Struct Rows

This example shows the core mapper API:

```go
mapper.ScanStructRows[T](rows, callback)
```

`ScanStructRows` scans database rows into a Go struct using db tags.

The callback is called for every scanned row, which makes this API useful when you want to process rows one by one or
build your own result slice.