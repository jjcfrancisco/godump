package main

import (
    "os/exec"
    "os"
	"database/sql"
	"fmt"
    "log"

	_ "github.com/lib/pq"
)

var dumpsDir string = fmt.Sprintf("%s/dumps", os.Getenv("HOME"))

type PgConn struct {
	db *sql.DB
}

func NewConn(c *inputConf) (*PgConn, error) {

	conn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		c.user, c.database, c.password, c.hostname, c.port)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PgConn{
		db: db,
	}, nil
}

func dump(s string, m model) error {

    err := createDir(dumpsDir)
    if err != nil {
        log.Fatal(err)
    }

    dumpFile := dumpsDir + "/" + m.search.Value()
    uri := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`, m.creds.user,
                                                    m.creds.password,
                                                    m.creds.hostname,
                                                    m.creds.port,
                                                    m.creds.database) 

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal("Error when opening postgres connection.")
	}

	if err := db.Ping(); err != nil {
		return err
	}

    cmd := exec.Command("pg_dump", uri, "-Fc", "-f", dumpFile)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    return nil
}
