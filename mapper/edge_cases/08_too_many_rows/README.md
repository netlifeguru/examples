# 08 Too Many Rows

This example shows how the mapper handles multiple rows when using `ScanStructOne`.

`ScanStructOne` expects exactly one row.

If the query returns more than one row, it returns:

```go
mapper.ErrTooManyRows
```

You can handle it with:

```go
if errors.Is(err, mapper.ErrTooManyRows) {
// expected one row, got multiple rows
}
```

Use `LIMIT 1` in your query when you only want one row.