package persistence

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	lib "github.com/southwind/ainews/common/config"
	"github.com/southwind/ainews/domain/entity"
	"github.com/southwind/ainews/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DAO struct {
	User    repository.UserRepository
	Tag     repository.TagRepository
	Article repository.ArticleRepository
	db      *gorm.DB
}

func NewDAO(conf lib.ServerConfig) (*DAO, error) {
	DBURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5438 sslmode=disable TimeZone=Asia/Shanghai",
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
	client, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ai_",
			SingularTable: true,
			NoLowerCase:   false,
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	conn, _ := client.DB()
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(1000)
	return &DAO{
		User:    NewUserDAO(client),
		Tag:     NewTagDAO(client),
		Article: NewArticleDAO(client),
		db:      client,
	}, nil

}

func (d *DAO) Close() error {
	db, _ := d.db.DB()
	return db.Close()
}

func (d *DAO) Migrate() error {
	return d.db.AutoMigrate(&entity.User{}, &entity.Tag{}, &entity.Article{}, &entity.ArticleTag{})
}
