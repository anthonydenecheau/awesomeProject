package db

import (
	"fmt"
	"github.com/go-pg/pg"
	"log"
	"strconv"
	"time"
)

type Config struct {
	host     string `required:"true"`
	port     int    `required:"true"`
	user     string `required:"true"`
	password string `required:"true"`
	name     string `required:"true"`
}
type dbLogger struct{}

var config Config

func panicIf(err error) {
	if err != nil {
		log.Println("Connection postgreSql {}", err)
		panic(err)
	}
}
func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}
func Connect() *pg.DB {

	db := pg.Connect(&pg.Options{
		User:                  config.user,
		Password:              config.password,
		Database:              config.name,
		Addr:                  config.host + ":" + strconv.Itoa(config.port),
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	panicIf(err)
	fmt.Println(n)

	db.AddQueryHook(dbLogger{})

	return db
}

func init() {
	config = Config{user: "ws_dev", password: "ws_dev", name: "ws_dev", host: "10.3.2.5", port: 5432}
}
