package dao

import "gorm.io/gen"

type UsersInterface interface {
	// sql(select * from @@table where userid=@userid limit 1)
	FindOneUserByID(userid string) (gen.T, error)
	// sql(select * from @@table where email=@email limit 1)
	FindOneUserByEmail(email string) (gen.T, error)
}
