package model

import (
	"fmt"
	"log"
	"time"

	"github.com/southwind/ainews/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(1000)

	fmt.Println("database init on host ", conf.Host)
}
