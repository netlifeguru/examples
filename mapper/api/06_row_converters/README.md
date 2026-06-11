# 06 Row Converters

This example shows how to use `mapper.Row` and converter helpers.

```go
row.String("id")
row.Bool("active")
row.Time("created_at")
```

`mapper.Row` is a small wrapper around `map[string]any` that provides typed accessors.

This is useful when you scan dynamic rows with `ScanMapRows`, or when you want to implement custom mapping logic
manually.