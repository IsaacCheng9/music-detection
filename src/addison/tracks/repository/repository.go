package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

var repo Repository

func Init() {
	if db, err := sql.Open("sqlite3", "test.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		log.Fatal(err)
	}
}

func Clear() int {
	const query = "DROP TABLE IF EXISTS Tracks"
	if _, err := repo.DB.Exec(query); err == nil {
		return 0
	} else {
		return -1
	}
}

func Create() int {
	const query = "CREATE TABLE IF NOT EXISTS Tracks" +
		"(Id TEXT PRIMARY KEY, Audio TEXT)"
	if _, err := repo.DB.Exec(query); err == nil {
		return 0
	} else {
		return -1
	}
}

func Read(id string) (Track, int64) {
	const query = "SELECT * FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(query); err == nil {
		defer func(stmt *sql.Stmt) {
			if err := stmt.Close(); err != nil {
				log.Fatal(err)
			}
		}(stmt)
		var track Track
		row := stmt.QueryRow(id)
		if err := row.Scan(&track.Id, &track.Audio); err == nil {
			return track, 1
		} else {
			return Track{}, 0
		}
	}
	return Track{}, -1
}

func ListAllIds() ([]string, int64) {
	const query = "SELECT Id FROM Tracks"
	if stmt, err := repo.DB.Prepare(query); err == nil {
		defer func(stmt *sql.Stmt) {
			if err := stmt.Close(); err != nil {
				log.Fatal(err)
			}
		}(stmt)
		rows, err := stmt.Query()
		if err == nil {
			var ids = make([]string, 0)
			for rows.Next() {
				var id string
				if err := rows.Scan(&id); err == nil {
					ids = append(ids, id)
				} else {
					return []string{}, -1
				}
			}
			return ids, int64(len(ids))
		}
	}
	return []string{}, -1
}

func Update(track Track) int64 {
	const query = "UPDATE Tracks SET Audio = ? WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(query); err == nil {
		defer func(stmt *sql.Stmt) {
			if err := stmt.Close(); err != nil {
				log.Fatal(err)
			}
		}(stmt)
		if res, err := stmt.Exec(track.Audio, track.Id); err == nil {
			if numRowsAffected, err := res.RowsAffected(); err == nil {
				return numRowsAffected
			}
		}
	}
	return -1
}

func Insert(track Track) int64 {
	const query = "INSERT INTO Tracks (Id, Audio) VALUES (?, ?)"
	if stmt, err := repo.DB.Prepare(query); err == nil {
		defer func(stmt *sql.Stmt) {
			if err := stmt.Close(); err != nil {
				log.Fatal(err)
			}
		}(stmt)
		if res, err := stmt.Exec(track.Id, track.Audio); err == nil {
			if numRowsAffected, err := res.RowsAffected(); err == nil {
				return numRowsAffected
			}
		}
	}
	return -1
}

func Delete(id string) int64 {
	const query = "DELETE FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(query); err == nil {
		defer func(stmt *sql.Stmt) {
			if err := stmt.Close(); err != nil {
				log.Fatal(err)
			}
		}(stmt)
		if res, err := stmt.Exec(id); err == nil {
			if numRowsAffected, err := res.RowsAffected(); err == nil {
				return numRowsAffected
			}
		}
	}
	return -1
}
