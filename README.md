# buildMigrationFiles
generate {table_name}.js and relative up/down sqls

## usage
1. use `MySQLWorkbench` or `mysqldump` dump each table in the current dir with type `.sql`
2. `go run build.go`

## structure
```
├── build.go
├── migrations
│   └── mysql
│       ├── sqls
│       │   ├── test_table-down.sql
│       │   └── test_table-up.sql
│       └── test_table.js
├── template.js
└── test_table.sql

```
