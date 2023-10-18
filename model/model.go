package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/southwind/ainews/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	Id        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

var client *gorm.DB

func InitDB(conf lib.ServerConfig) {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5438 sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host,
		conf.User,
		conf.Password,
		conf.DbName,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	client, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(2, err)
	}

	conn, _ := client.DB()
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(1000)

	fmt.Println("database init on host ", conf.Host)
}

func MigrateDb(conf lib.ServerConfig) {

	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5438 sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host,
		conf.User,
		conf.Password,
		conf.DbName,
	)
	client, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ai_",
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Fatal(2, err)
	}

	conn, _ := client.DB()

	driver, err := migratePg.WithInstance(conn, &migratePg.Config{})
	if err != nil {
		log.Fatal("database migration instance created failed", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/db",
		"postgres", driver)
	if err != nil {
		log.Fatal("database migration failed", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("database migration up failed", err)
	} else {
		log.Println("database migration up success")
	}

}
