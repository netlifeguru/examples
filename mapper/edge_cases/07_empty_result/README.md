# 07 Empty Result

This example shows how the mapper handles an empty result set when using `ScanStructOne`.

When no row is found, `ScanStructOne` returns:

```go
mapper.ErrNoRows
```

You can handle it with:

```go
if errors.Is(err, mapper.ErrNoRows) {
// not found
}
```