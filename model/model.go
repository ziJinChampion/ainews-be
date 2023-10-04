package model

import (
	"fmt"
	"log"

	"github.com/southwind/ainews/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
	DeletedAt  int `json:"deleted_at"`
}

var DB *gorm.DB

func InitDB(conf lib.ServerConfig) {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		conf.Host,
		conf.User,
		conf.Password,
		conf.DbName,
	)
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ai_",
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatal(2, err)
	}

	conn, _ := DB.DB()
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(1000)

	fmt.Println("database init on port ", conf.Host)
}
