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
	if db, err := sql.Open("sqlite3", "/tmp/test.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		log.Fatal("Database initialisation")
	}
}

func Create() int {
	const sql = "CREATE TABLE IF NOT EXISTS Tracks" +
		"(Id TEXT PRIMARY KEY, Audio TEXT)"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Clear() int {
	const sql = "DELETE FROM Tracks"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Read(id string) (Track, int64) {
	const sql = "SELECT * FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
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
	const sql = "SELECT Id FROM Tracks"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		rows, err := stmt.Query()
		if err == nil {
			var ids []string
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
	const sql = "UPDATE Tracks SET Audio = ? WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(track.Audio, track.Id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Insert(track Track) int64 {
	const sql = "INSERT INTO Tracks (Id, Audio) VALUES (?, ?)"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(track.Id, track.Audio); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Delete(id string) int64 {
	const sql = "DELETE FROM Tracks WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}