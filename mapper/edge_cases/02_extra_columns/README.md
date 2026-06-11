# 02 Extra Columns

This example shows how the mapper handles extra columns returned by a query.

The query returns:

```sql
SELECT id, name, email, active, created_at
FROM users
```

But the destination struct only contains:

```go
type UserSummary struct {
ID   string `db:"id"`
Name string `db:"name"`
}
```

The mapper maps known fields and ignores columns that are not present in the destination struct.