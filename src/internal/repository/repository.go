package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"tcc-test/internal/entity"
)

type userRepository struct {
	userTable *sql.DB
}

func NewUserRepository(
	userTable *sql.DB,
) *userRepository {
	return &userRepository{
		userTable: userTable,
	}
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (err error) {
	db := u.userTable
	sqlStatement := `INSERT INTO users (name, surname, age, birth_date, address) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.QueryRowContext(ctx, sqlStatement, user.Name, user.Surname, user.Age, user.BirthDate, user.Address).Scan(&user.ID)
	if err != nil {
		log.Println("insert db error: ", err.Error())
		err = fmt.Errorf("insert db error: %s", err.Error())
		return
	}
	return
}

func (u *userRepository) FindMany(ctx context.Context) (users []entity.User, err error) {
	db := u.userTable
	sqlStatement := `SELECT * FROM users ORDER BY id`
	res, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		log.Println("query db error: ", err.Error())
		err = fmt.Errorf("query db error: %s", err.Error())
		return
	}

	defer res.Close()

	for res.Next() {
		var u entity.User
		if err = res.Scan(&u.ID, &u.Name, &u.Surname, &u.Address, &u.BirthDate, &u.Age); err != nil {
			log.Println("internal error: ", err.Error())
			err = fmt.Errorf("internal error: %s", err.Error())
			return
		}

		users = append(users, u)
	}

	return users, nil
}

func (u *userRepository) FindOne(ctx context.Context, id string) (user entity.User, err error) {
	db := u.userTable
	sqlStatement := `SELECT * FROM users WHERE id = $1`
	row := db.QueryRowContext(ctx, sqlStatement, id)
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.BirthDate, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no user in postgresql")
			err = fmt.Errorf("no user in postgresql")
		} else {
			log.Println("internal error")
			err = fmt.Errorf("internal error")
		}
		return
	}

	return
}
