package postgres

import (
	"fmt"

	"github.com/mrumyantsev/pastebin/internal/domain/user"
	"github.com/mrumyantsev/pastebin/internal/pkg/core"
	"github.com/mrumyantsev/pastebin/internal/pkg/database"
)

type UserPostgresRepository struct {
	database *database.Database
}

func NewUserPostgresRepository(db *database.Database) *UserPostgresRepository {
	return &UserPostgresRepository{database: db}
}

func (r *UserPostgresRepository) CreateUser(usr user.User) error {
	query := fmt.Sprintf("INSERT INTO %s (%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6)",
		user.TabUsers,
		user.ColUsername,
		user.ColPasswordHash,
		user.ColFirstName,
		user.ColLastName,
		user.ColEmail,
		user.ColCreatedAt,
	)

	_, err := r.database.Exec(
		query,
		usr.Username,
		usr.PasswordHash,
		usr.FirstName,
		usr.LastName,
		usr.Email,
		usr.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgresRepository) GetUsers(pg core.Pagination) ([]user.User, error) {
	query := fmt.Sprintf(`SELECT %s
FROM %s
WHERE %s > (CASE WHEN %d > 0 THEN
				(SELECT MAX(%s)
				FROM (SELECT %s
					FROM %s
					LIMIT %d))
			ELSE
				-1
			END)
LIMIT %d`,
		user.AllCols,
		user.TabUsers,
		user.ColId,
		pg.Page,
		user.ColId,
		user.ColId,
		user.TabUsers,
		pg.Page,
		pg.Limit,
	)

	var users []user.User

	r.database.Select(&users, query)

	return users, nil
}

func (r *UserPostgresRepository) GetUser(id int) (user.User, error) {
	var usr user.User

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1",
		user.AllCols,
		user.TabUsers,
		user.ColId,
	)

	if err := r.database.Get(&usr, query, id); err != nil {
		return usr, err
	}

	return usr, nil
}

func (r *UserPostgresRepository) UpdateUser(id int, usr user.User) error {
	query := fmt.Sprintf(`UPDATE %s
SET
%s = CASE WHEN $1 = '' THEN %s ELSE $1 END,
%s = CASE WHEN $2 = '' THEN %s ELSE $2 END,
%s = CASE WHEN $3 = '' THEN %s ELSE $3 END,
%s = CASE WHEN $4 = '' THEN %s ELSE $4 END,
%s = CASE WHEN $5 = '' THEN %s ELSE $5 END
WHERE %s = $6`,
		user.TabUsers,
		user.ColUsername, user.ColUsername,
		user.ColPasswordHash, user.ColPasswordHash,
		user.ColFirstName, user.ColFirstName,
		user.ColLastName, user.ColLastName,
		user.ColEmail, user.ColEmail,
		user.ColId,
	)

	_, err := r.database.Exec(
		query,
		usr.Username,
		usr.PasswordHash,
		usr.FirstName,
		usr.LastName,
		usr.Email,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgresRepository) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1",
		user.TabUsers,
		user.ColId,
	)

	if _, err := r.database.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (r *UserPostgresRepository) IsUserExists(username string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1",
		user.TabUsers,
		user.ColUsername,
	)

	var cond bool

	if err := r.database.Get(&cond, query, username); err != nil {
		return false, err
	}

	return cond, nil
}

func (r *UserPostgresRepository) IsEmailExists(email string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1",
		user.TabUsers,
		user.ColEmail,
	)

	var cond bool

	if err := r.database.Get(&cond, query, email); err != nil {
		return false, err
	}

	return cond, nil
}
