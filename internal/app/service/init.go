package service

import (
	"jupiter-test/internal/app/model"
	"jupiter-test/internal/app/service/user"
	"jupiter-test/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)

//Init instantiate the service
func Init() {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}
