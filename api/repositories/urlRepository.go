package repositories

import (
	"database/sql"
	"errors"
	"lilurl/api/entities"
)

type UrlRepository struct {
	DB *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{DB: db}
}

func (r *UrlRepository) Create(url *entities.Url) error {
	query := `INSERT INTO urls (shortId, url) VALUES (?, ?)`
	_, err := r.DB.Exec(query, url.ShortId, url.Url)
	return err
}

func (r *UrlRepository) FindByShortId(shortId string) (*entities.Url, error) {
	query := `SELECT * FROM urls WHERE shortId = ? LIMIT 1`
	row := r.DB.QueryRow(query, shortId)

	var url entities.Url
	err := row.Scan(&url.Id, &url.ShortId, &url.Url)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}
