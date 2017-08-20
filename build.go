package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// mkdir migrations/sqls
	if err := os.MkdirAll("migrations/mysql/sqls", 0755); err != nil {
		fmt.Printf("mkdir err: %s\n", err)
		os.Exit(1)
	}
	out, err := exec.Command(
		"/bin/sh",
		"-c",
		"ls -l |tail -n +1 |grep .sql |awk '{print $9}'").
		CombinedOutput()
	if err != nil {
		fmt.Println("err get sql list")
		os.Exit(1)
	}
	fmt.Println(string(out))
	for i, sql := range strings.Split(strings.TrimSuffix(string(out), "\n"), "\n") {
		tableName := strings.Replace(sql, ".sql", "", -1)
		fmt.Printf("%d: %s\n", i+1, tableName)

		// generate js
		if _, err := exec.Command(
			"/bin/sh",
			"-c",
			fmt.Sprintf("sed \"s/TABLE_NAME/%s/g\" template.js > migrations/mysql/%s.js",
				tableName, tableName)).
			CombinedOutput(); err != nil {
			fmt.Printf("err create js for table: %s\n", tableName)
			os.Exit(1)
		}

		// generate up/down sqls
		if _, err := exec.Command(
			"/bin/sh",
			"-c",
			fmt.Sprintf("cp %s migrations/mysql/sqls/%s-up.sql && touch migrations/mysql/sqls/%s-down.sql",
				sql, tableName, tableName)).
			CombinedOutput(); err != nil {
			fmt.Printf("create up/down sqls err: %s\n", err)
			os.Exit(1)
		}
	}
	fmt.Println("\nall done!")
}
