# 01 Null Values

This example shows how the mapper handles nullable database values.

The destination struct uses standard Go SQL nullable types:

```go
sql.NullString
sql.NullTime
```

When a database column is `NULL`, the mapper keeps the value invalid:

```go
sql.NullString{Valid: false}
sql.NullTime{Valid: false}
```

When a database column contains a value, the mapper fills the nullable type and sets Valid to true.