package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/louistwiice/go/simplebase/models"
)

type UserSQL struct {
	db *sql.DB
}

func NewUserSQL(db *sql.DB) *UserSQL {
	return &UserSQL{
		db: db,
	}
}

// List all Users
func (r *UserSQL) List() ([]*models.User, error) {
	stmt, err := r.db.Prepare(`SELECT id, email, first_name, last_name, is_active, is_staff, is_superuser, created_at, updated_at FROM user`)

	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var users []*models.User
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.IsActive, &u.IsStaff, &u.IsSuperuser, &u.CreatedAt, &u.UpdatedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

// Create a user
func (r *UserSQL) Create(u *models.User, id ...string) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	stmt, err := r.db.Prepare(`
			INSERT user SET id=?, email=?, first_name=?, last_name=?, password=?, is_active=?, is_staff=?, is_superuser=?, created_at=?, updated_at=?
	`)
	if err != nil {
		return err
	}

	// if the id has been sent, we use it. Otherwise, we generate a new one
	if len(id) > 0 {
		u.ID = id[0]
	} else {
		u.ID = uuid.New().String()
	}

	_, err = stmt.Exec(u.ID, u.Email, u.FirstName, u.LastName, u.Password, u.IsActive, u.IsStaff, u.IsSuperuser, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

// Get a user by ID
func (r *UserSQL) Get(id string) (*models.User, error) {
	return getUser(id, r.db)
}

func getUser(id string, db *sql.DB) (*models.User, error) {
	stmt, err := db.Prepare(`SELECT id, email, first_name, last_name, password, is_active, is_staff, created_at, updated_at FROM user WHERE id=?`)

	if err != nil {
		return &models.User{}, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return &models.User{}, err
	}
	var u models.User
	var nb_rows = 0
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Password, &u.IsActive, &u.IsStaff, &u.CreatedAt, &u.UpdatedAt)
		nb_rows += 1
	}
	if err != nil {
		return &models.User{}, err
	}

	if nb_rows > 0 {
		return &u, nil
	} else {
		return nil, nil
	}

}

// Update a user profile
func (r *UserSQL) Update(u *models.User) error {
	u.UpdatedAt = time.Now()
	_, err := r.db.Exec(`UPDATE user SET email=?, first_name=?, last_name=?, updated_at=? WHERE id=?`,
		u.Email, u.FirstName, u.LastName, u.UpdatedAt, u.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update a user profile password
func (r *UserSQL) UpdatePassword(u *models.User) error {
	u.UpdatedAt = time.Now()

	_, err := r.db.Exec(`UPDATE user SET password=?, updated_at=? WHERE id=?`,
		u.Password, u.UpdatedAt, u.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user
func (r *UserSQL) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM user WHERE id=?`, id)
	if err != nil {
		return err
	}
	return nil
}
