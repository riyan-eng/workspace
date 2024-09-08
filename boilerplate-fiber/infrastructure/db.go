package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"server/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var SqlDB *sql.DB
var SqlxDB *sqlx.DB

func ConnectSqlDB() {
	dsn := fmt.Sprintf(`
		host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v
	`, env.NewEnv().POSTGRE_HOST, env.NewEnv().POSTGRE_USERNAME, env.NewEnv().POSTGRE_PASSWORD, env.NewEnv().POSTGRE_DATABASE, env.NewEnv().POSTGRE_PORT, env.NewEnv().POSTGRE_TIMEZONE)
	var err error
	SqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("sql database: can't connect to database - ", err.Error())
		os.Exit(1)
	}
	SqlDB.SetMaxIdleConns(env.NewEnv().POSTGRE_CONN_MAX_IDLE)
	SqlDB.SetMaxOpenConns(env.NewEnv().POSTGRE_CONN_MAX_OPEN)
	SqlDB.SetConnMaxLifetime(time.Minute * env.NewEnv().POSTGRE_CONN_MAX_LIFETIME)
	if err := SqlDB.Ping(); err != nil {
		fmt.Printf("sql database: can't ping to database - %v \n", err)
		os.Exit(1)
	}

	fmt.Println("sql database: connection opened to database")
}

func ConnectSqlxDB() {
	dsn := fmt.Sprintf(`
		host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v
	`, env.NewEnv().POSTGRE_HOST, env.NewEnv().POSTGRE_USERNAME, env.NewEnv().POSTGRE_PASSWORD, env.NewEnv().POSTGRE_DATABASE, env.NewEnv().POSTGRE_PORT, env.NewEnv().POSTGRE_TIMEZONE)
	var err error
	SqlxDB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("sqlx database: can't connect to database")
		os.Exit(1)
	}
	SqlxDB.SetMaxIdleConns(env.NewEnv().POSTGRE_CONN_MAX_IDLE)
	SqlxDB.SetMaxOpenConns(env.NewEnv().POSTGRE_CONN_MAX_OPEN)
	SqlxDB.SetConnMaxLifetime(time.Minute * env.NewEnv().POSTGRE_CONN_MAX_LIFETIME)
	if err := SqlxDB.Ping(); err != nil {
		fmt.Printf("sqlx database: can't ping to database - %v \n", err)
		os.Exit(1)
	}

	fmt.Println("sqlx database: connection opened to database")
}
