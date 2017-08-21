# buildMigrationFiles
generate {table_name}.js and relative up/down sqls

## usage
1. use `MySQLWorkbench` or `mysqldump` dump each table in the current dir with type `.sql`
2. `go run build.go`

## structure
```
├── README.md
├── build.go
├── migrations
│   └── mysql
│       ├── 20170821093300-test_table.js
│       └── sqls
│           ├── 20170821093300-test_table-down.sql
│           └── 20170821093300-test_table-up.sql
├── template.js
└── test_table.sql
```
