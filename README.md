# todolist-CLI
A simple command-line to-do list built to get more familiar with Go. It stores tasks in a local SQLite database and exposes add, list, complete, and delete commands through a Cobra-based CLI.
## What it does
Tasks are stored with an ID, the task text, a creation timestamp, and a completed flag. Each command is a subcommand of the `todo` binary.
```bash
todo add "buy milk"
todo list
```
```
 ID | Task     | Date Created        | Completed |
 1  | buy milk | 2026-06-23 19:38:22 | false     |
```
By default `list` hides completed tasks. Pass `--all` (or `-a`) to show everything:
```bash
todo list --all
```
Mark a task done or remove it by its ID:
```bash
todo complete 1
todo delete 1
```
## Concepts covered
* **CLI structure with Cobra**
    Each command (`add`, `list`, `complete`, `delete`) lives in its own file under `cmd/` and registers itself onto the root command in its `init()`.
* **Flags**
    A local `--all`/`-a` bool flag on the `list` command, read back at run time via `cmd.Flags().GetBool`.
* **Pure-Go SQLite**
    Storage uses *modernc.org/sqlite*, a CGo-free SQLite driver, so the binary builds and cross-compiles without a C toolchain.
* **Separation of concerns**
    All SQL lives in the *database* package behind a *Database* type. The command files never touch SQL directly, they just call methods like *AddTodo*, *ListTodo*, *CompleteTodo*, and *DeleteTodo*.
* **Shared connection**
    A single package-level *Database* is connected once in *Execute()* before any command runs, then reused across every subcommand.
* **Stable IDs**
    The displayed ID is the database primary key, so the number shown in `list` is the same one `complete` and `delete` operate on.
* **Tabular output**
    Aligned columns rendered with the standard library *text/tabwriter*.
## Project structure
```
cmd/                Command definitions (root, add, list, complete, delete)
internal/
  database/         Database type, connection, and SQL methods
main.go             Application entry point
list_database.db    SQLite database file
```
## Running
```
go build -o todo
./todo list
```
## Schema
```sql
CREATE TABLE lists (
    id        INTEGER PRIMARY KEY,
    task      TEXT,
    date      TEXT,
    completed INTEGER
);
```
