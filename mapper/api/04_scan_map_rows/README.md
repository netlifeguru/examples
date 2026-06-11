# 04 Scan Map Rows

This example shows how to scan rows into dynamic `map[string]any` values.

```go
mapper.ScanMapRows(rows, callback)
```

Use `ScanMapRows` when the query result shape is dynamic, unknown, or when you are building debugging, admin, export, or
generic query tools.