package query

import "database/sql"

func CreateTablesIfNotCreated(db *sql.DB) error {
	err := createThreadsIfNotCreated(db)
	if err != nil {
		return err
	}

	err = createUsersIfNotCreated(db)
	if err != nil {
		return err
	}

	err = createMessagesIfNotCreated(db)
	if err != nil {
		return err
	}

	return nil
}

func createThreadsIfNotCreated(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS threads(
			name TEXT NOT NULL PRIMARY KEY
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func createUsersIfNotCreated(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			name TEXT NOT NULL PRIMARY KEY,
			password TEXT NOT NULL
		) 
	`)
	if err != nil {
		return err
	}

	return nil
}

func createMessagesIfNotCreated(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS messages(
			id SERIAL PRIMARY KEY,
			author_name TEXT NOT NULL,
			thread_name TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			reference_id INTEGER,
			text TEXT,
			image_name TEXT,
			FOREIGN KEY (thread_name) REFERENCES threads(name) ON DELETE CASCADE,
			FOREIGN KEY (author_name) REFERENCES users(name) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_thread_name on messages(thread_name);
	`)
	if err != nil {
		return err
	}

	return nil
}
