package repository

import (
	"fmt"
	"log/slog"

	"github.com/michelemendel/genvaeg/entity"
)

func (r *Repo) CreateUser(user entity.User) error {
	_, err := r.DB.Exec(`
    INSERT INTO users(
      uuid, 
      name, 
      hashedpassword
    ) VALUES(?, ?, ?)`,
		user.UUID,
		user.Name,
		user.HashedPassword,
	)
	if err != nil {
		slog.Error(err.Error(), "name", user.Name)
		return fmt.Errorf("error creating user, error: %w", err)
	}
	slog.Info("CreateUser", "name", user.Name)
	return nil
}

func (r *Repo) CreateURLPair(url entity.URLPair, user entity.User) error {
	_, err := r.DB.Exec(`
    INSERT INTO urls(
      shorturlpath, 
      fullurl, 
      fk_user_uuid
    ) VALUES(?, ?, ?)`,
		url.ShortURL,
		url.FullURL,
		user.UUID,
	)
	if err != nil {
		slog.Error(err.Error(), "shortURL", url.ShortURL)
		return fmt.Errorf("error creating URL, error: %w", err)
	}
	slog.Info("CreateURLPair", "full", url.FullURL, "short", url.ShortURL, "user", user.Name)
	return nil
}
