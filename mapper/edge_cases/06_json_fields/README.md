# 06 JSON Fields

This example shows how the mapper handles JSON-like fields.

The database may return a JSON column as `[]byte` or `string`.

The mapper can unmarshal that value into a Go map:

```go
Payload map[string]any `db:"payload"`
```

This is useful for event payloads, metadata, settings, audit logs, and other semi-structured data.

Ak `events.payload` v demo DB nie je valid JSON, treba ho upraviť napríklad na:

```json
{
  "action": "login",
  "ip": "127.0.0.1"
}
```