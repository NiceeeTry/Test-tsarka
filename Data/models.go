package Data

import (
	"database/sql"

	"Alikhan.Aitbayev/Data/redisDB"
	sqlitedb "Alikhan.Aitbayev/Data/sqliteDB"
	"github.com/go-redis/redis"
)

type Models struct {
	Users interface {
		Insert(user *sqlitedb.User) (int, error)
		Get(id int) (*sqlitedb.User, error)
		Update(id int, user *sqlitedb.User) error
		Delete(id int) error
	}
	Counter redisDB.CounterModel
}

func NewModels(db *sql.DB, client *redis.Client) Models {
	return Models{
		Users:   sqlitedb.UserModel{DB: db},
		Counter: redisDB.CounterModel{RedisDB: client},
	}
}
