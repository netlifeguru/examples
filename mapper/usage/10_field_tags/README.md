# 10 Field Tags

This example shows how to map database columns to differently named Go fields using `db` tags.

The mapper uses the `db` tag to match result columns with struct fields, so your Go field names do not need to match database column names.