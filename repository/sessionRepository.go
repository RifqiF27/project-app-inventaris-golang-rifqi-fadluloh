package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type SessionRepository interface {
	CreateSession(userID int) (uuid.UUID, error)
	GetUserIDBySessionID(sessionID string) (int, error)
	DeleteSession(sessionID string) error
}

type sessionRepository struct {
    db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
    return &sessionRepository{db: db}
}

func (r *sessionRepository) CreateSession(userID int) (uuid.UUID, error) {
    sessionID := uuid.New()
    _, err := r.db.Exec(`INSERT INTO "Session" (session_id, user_id, login_time) VALUES ($1, $2, $3)`, sessionID, userID, time.Now())
    if err != nil {
        return uuid.UUID{}, err
    }
    return sessionID, nil
}

func (r *sessionRepository) GetUserIDBySessionID(sessionID string) (int, error) {
    var userID int
    err := r.db.QueryRow(`SELECT user_id FROM "Session" WHERE session_id=$1`, sessionID).Scan(&userID)
    if err != nil {
        return 0, err
    }
    return userID, nil
}

func (r *sessionRepository) DeleteSession(sessionID string) error {
    _, err := r.db.Exec(`DELETE FROM "Session" WHERE session_id=$1`, sessionID)
    return err
}