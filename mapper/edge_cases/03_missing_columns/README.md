# 03 Missing Columns

This example shows how the mapper handles missing columns.

The destination struct contains more fields than the query returns:

```go
type User struct {
ID     string `db:"id"`
Name   string `db:"name"`
Email  string `db:"email"`
Active bool   `db:"active"`
}
```

But the query only returns:

```sql
SELECT id, name
FROM users
```

The mapper fills the fields it can match and leaves missing fields with their Go zero value.

For example:

```text
Email  -> ""
Active -> false
```