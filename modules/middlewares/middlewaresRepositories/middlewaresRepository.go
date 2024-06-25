package middlewaresRepositories

import (
	"fmt"
	"log"

	_pkgMiddlewares "github.com/MarkTBSS/075_Role_Based_Authorization/modules/middlewares"
	"github.com/jmoiron/sqlx"
)

type IMiddlewaresRepository interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*_pkgMiddlewares.Role, error)
}

type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) IMiddlewaresRepository {
	return &middlewaresRepository{
		db: db,
	}
}

func (r *middlewaresRepository) FindAccessToken(userId, accessToken string) bool {
	query := `
	SELECT
		(CASE WHEN COUNT(*) = 1 THEN TRUE ELSE FALSE END)
	FROM "oauth"
	WHERE "user_id" = $1
	AND "access_token" = $2;`
	var exists bool
	err := r.db.Get(&exists, query, userId, accessToken)
	if err != nil {
		log.Printf("Database error: %v", err)
		return false
	}
	return exists
}

func (r *middlewaresRepository) FindRole() ([]*_pkgMiddlewares.Role, error) {
	query := `
	SELECT
		"id",
		"title"
	FROM "roles"
	ORDER BY "id" DESC;`

	roles := make([]*_pkgMiddlewares.Role, 0)
	if err := r.db.Select(&roles, query); err != nil {
		return nil, fmt.Errorf("roles are empty")
	}
	return roles, nil
}
