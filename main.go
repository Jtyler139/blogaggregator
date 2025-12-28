package main

import _ "github.com/lib/pq"

import (
	"github.com/jtyler139/blogaggregator/internal/config"
	"os"
	"log"
	"database/sql"
	"github.com/jtyler139/blogaggregator/internal/database"
)

type state struct {
	db	*database.Queries
	cfg *config.Config
}

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/gator"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(db)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
	
}