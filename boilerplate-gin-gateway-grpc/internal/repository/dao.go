package repository

import (
	"database/sql"

	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DAO interface {
	NewExampleRepository() ExampleRepository
	NewAuthRepository() AuthRepository
	NewPerangkatRepository() PerangkatRepository
	NewObjectRepository() ObjectRepository
}

type dao struct {
	SqlDB      *sql.DB
	SqlxDB     *sqlx.DB
	Cache      *redis.Client
	Permission *casbin.Enforcer
}

func NewDAO(sqlDB *sql.DB, sqlxDB *sqlx.DB, cache *redis.Client, permission *casbin.Enforcer) DAO {
	return &dao{
		SqlDB:      sqlDB,
		SqlxDB:     sqlxDB,
		Cache:      cache,
		Permission: permission,
	}
}

func (m *dao) NewExampleRepository() ExampleRepository {
	return &exampleRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
	}
}

func (m *dao) NewAuthRepository() AuthRepository {
	return &authRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
		cache:  m.Cache,
	}
}

func (m *dao) NewPerangkatRepository() PerangkatRepository {
	return &perangkatRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
	}
}

func (m *dao) NewObjectRepository() ObjectRepository {
	return &objectRepository{
		sqlDB:  m.SqlDB,
		sqlxDB: m.SqlxDB,
	}
}
