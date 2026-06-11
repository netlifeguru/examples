# 10 Snake Case

This example shows how the mapper matches `snake_case` database columns to Go struct fields.

For example:

```go
mapper.ToSnakeCase("CreatedAt")
```

returns:

```go
created_at
```

Because of this rule, the mapper can scan the `created_at` column into the `CreatedAt` field even without a `db` tag.

You can still use explicit `db` tags when you want full control over column-to-field mapping.

Output:

```txt
Snake case examples:
CreatedAt -> created_at
UserID    -> user_id

ID: 11111111-1111-1111-1111-111111111111 | Name: John Doe | Email: john@example.com | Active: true | Created: 2026-05-07 11:15:21

```