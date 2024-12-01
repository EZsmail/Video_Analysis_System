package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type PostgreSQL struct {
	DB          *sql.DB
	TableStatus string
}

func ConnectPostgreSQL(host string, port int, user, password, dbname, tableStatus string, logger *zap.Logger) (*PostgreSQL, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL", zap.Error(err))
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error("PostgreSQL connection test failed", zap.Error(err))
		return nil, err
	}

	logger.Info("Connected to PostgreSQL", zap.String("host", host))
	return &PostgreSQL{
		DB:          db,
		TableStatus: tableStatus,
	}, nil
}

// Insert status
func (pg *PostgreSQL) InsertStatus(processingID, status string) error {
	query := fmt.Sprintf("INSERT INTO %s (processing_id, status) VALUES ($1, $2)", pg.TableStatus)
	_, err := pg.DB.Exec(query, processingID, status)
	return err
}

// Update status
func (pg *PostgreSQL) UpdateStatus(processingID, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status=$2 WHERE processing_id=$1", pg.TableStatus)
	_, err := pg.DB.Exec(query, processingID, status)
	return err
}

// Get Status
func (pg *PostgreSQL) GetStatus(processingID string) (string, error) {
	query := fmt.Sprintf("SELECT status FROM %s WHERE processing_id=$1", pg.TableStatus)
	var status string
	err := pg.DB.QueryRow(query, processingID).Scan(&status)
	return status, err
}
