package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *sql.DB

func NewDB() (*bun.DB, error) {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("error time setting: %w", err)
	}
	c := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Addr:      os.Getenv("MYSQL_DRIVER"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	engine, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// ここでクローズするNewDBの呼び出し完了時にクローズするのでSQL実行時に接続エラーになる
	//	defer engine.Close()

	if err := engine.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	db := bun.NewDB(engine, mysqldialect.New())
	// ここでクローズするNewDBの呼び出し完了時にクローズするのでSQL実行時に接続エラーになる
	//	defer db.Close()

	// クエリーフックを追加すると標準出力に実行したSQLが出力される
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	return db, nil
}
