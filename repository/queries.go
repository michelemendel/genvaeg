package repository

import (
	"fmt"
	"log/slog"

	"github.com/michelemendel/genvaeg/entity"
)

func (r *Repo) GetUser(name string) (entity.User, error) {
	var user entity.User
	err := r.DB.QueryRow(`
    SELECT 
      uuid, 
      name, 
      hashedpassword 
    FROM users 
    WHERE name = ?`,
		name,
	).Scan(&user.UUID, &user.Name, &user.HashedPassword)
	if err != nil {
		slog.Error(err.Error(), "name", name)
		return user, fmt.Errorf("error getting user, error: %w", err)
	}
	slog.Info("GetUser", "name", user.Name)
	return user, nil
}

func (r *Repo) GetURLPairByFullURL(fullURL string) (entity.URLPair, error) {
	var url entity.URLPair
	err := r.DB.QueryRow(`
  SELECT 
  fullurl, 
  shorturlpath 
  FROM urls 
  WHERE fullurl = ?`,
		fullURL,
	).Scan(&url.FullURL, &url.ShortURL)
	if err != nil {
		slog.Error(err.Error(), "fullURL", fullURL)
		return url, fmt.Errorf("error getting URL pair by full URL, error: %w", err)
	}
	slog.Info("GetURLPairByFullURL", "fullURL", fullURL)
	return url, nil
}

func (r *Repo) DoesShortURLExists(shortURL string) bool {
	var count int
	err := r.DB.QueryRow(`
  SELECT 
  COUNT(*) 
  FROM urls 
  WHERE shorturlpath = ?`,
		shortURL,
	).Scan(&count)
	if err != nil {
		slog.Error(err.Error(), "shortURL", shortURL)
		return false
	}
	slog.Info("DoesShortURLExists", "shortURL", shortURL)
	return count > 0
}

func (r *Repo) GetAllURLPairs() ([]entity.URLPair, error) {
	var urls []entity.URLPair
	rows, err := r.DB.Query(`
  SELECT 
  fullurl, 
  shorturlpath
  FROM urls`)
	if err != nil {
		slog.Error(err.Error())
		return urls, fmt.Errorf("error getting all URL pairs, error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var url entity.URLPair
		err := rows.Scan(&url.FullURL, &url.ShortURL)
		if err != nil {
			slog.Error(err.Error())
			return urls, fmt.Errorf("error getting all URL pairs, error: %w", err)
		}
		urls = append(urls, url)
	}
	slog.Info("GetAllURLPairs")
	return urls, nil
}