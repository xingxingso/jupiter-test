package user

import (
	"jupiter-test/internal/app/model/db"
)

type Repository interface {
	Get(id int) (user db.User, err error)
}
