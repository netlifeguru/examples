# 04 Single Row

This example shows how to scan a single row into a struct.

The mapper works with row sets, so for single-row queries you scan normally and return `sql.ErrNoRows` when no row was found.