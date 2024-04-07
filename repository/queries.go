package repository

import (
	"fmt"
	"log/slog"

	"github.com/michelemendel/genvaeg/entity"
)

func (r *Repo) GetUserByUUID(uuid string) (entity.User, error) {
	var user entity.User
	err := r.DB.QueryRow(`
    SELECT 
      uuid, 
      name, 
      hashedpassword 
    FROM users 
    WHERE uuid = ?`,
		uuid,
	).Scan(&user.UUID, &user.Name, &user.HashedPassword)
	if err != nil {
		slog.Error(err.Error(), "uuid", uuid)
		return user, fmt.Errorf("error getting user by uuid, error: %w", err)
	}
	return user, nil
}

func (r *Repo) GetUserByName(name string) (entity.User, error) {
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

func (r *Repo) GetAllUsers() []entity.User {
	var users []entity.User
	rows, err := r.DB.Query(`
    SELECT 
      uuid, 
      name 
    FROM users`)
	if err != nil {
		slog.Error(err.Error())
		return users
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.UUID, &user.Name)
		if err != nil {
			slog.Error(err.Error())
			return users
		}
		users = append(users, user)
	}
	slog.Info("GetAllUsers")
	return users

}

func (r *Repo) GetFullURLByShortURL(shortURLPath string) (string, error) {
	var url entity.URLPair
	err := r.DB.QueryRow(`
  SELECT 
  fullurl, 
  shorturlpath 
  FROM urls 
  WHERE shorturlpath = ?`,
		shortURLPath,
	).Scan(&url.FullURL, &url.ShortURL)
	if err != nil {
		slog.Error(err.Error(), "shortURLPath", shortURLPath)
		return "", fmt.Errorf("error getting fullURL by shortURLPath, error: %w", err)
	}
	slog.Info("GetURLPairByFullURL", "fullURL", shortURLPath)
	return url.FullURL, nil
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
