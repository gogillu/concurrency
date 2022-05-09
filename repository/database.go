package repository

import (
	"concurrency/config"
	"concurrency/item"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:generate mockgen -source=database.go -destination=mocks/database.go -package=repository
type Database interface {
	ReadItems() ([]item.Item, error)
}

type Repository struct {
	db *gorm.DB
}

func Init(cfg config.Config) (*gorm.DB, func(), error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, nil, errors.Wrap(err, "error connecting to db ")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error getting generic db interface ")
	}

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			fmt.Println("error closing db conn : ", err)
		}
	}

	return db, cleanup, nil
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) ReadItems() ([]item.Item, error) {
	var items []item.Item
	r.db.Find(&items).Debug()
	return items, nil
}
