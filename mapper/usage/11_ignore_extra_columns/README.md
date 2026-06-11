# 11 Ignore Extra Columns

This example shows that the destination struct does not need to contain every column returned by the query.

The mapper matches known columns by `db` tags and ignores columns that are not present in the struct.