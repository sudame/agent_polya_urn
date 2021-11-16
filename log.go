package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// dump event log (raw)
func dumpEventLog(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range events {
		fmt.Fprintf(file, "%d %d\n", e.From, e.To)
	}
	defer file.Close()
}

// dump agents (raw)
func dumpAgents(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range activeAgents() {
		fmt.Fprintf(file, "%d\n", a.Id)
	}

	defer file.Close()
}

// aggregate network edges using SQL with sqlite3
func aggEdges(path string) {

	os.Remove("./temp.db")

	db, err := sql.Open("sqlite3", "./temp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `create table event (id integer not null primary key, source integer, target integer, weight integer);
	delete from event;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`insert into event(source, target, weight) values (?, ?, 1)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, e := range events {
		_, err = stmt.Exec(e.From, e.To)
		if err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()

	rows, err := db.Query("select source, target, sum(weight) from event group by source, target")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{"source", "target", "weight"})

	for rows.Next() {
		var source int
		var target int
		var weight int
		err := rows.Scan(&source, &target, &weight)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]string{
			strconv.Itoa(source),
			strconv.Itoa(target),
			strconv.Itoa(weight),
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	os.Remove("./temp.db")
}

// aggregate network nodes
func aggNodes(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{"id", "label"})

	for _, a := range agents {
		w.Write([]string{
			strconv.Itoa(a.Id),
			strconv.Itoa(a.Id),
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
