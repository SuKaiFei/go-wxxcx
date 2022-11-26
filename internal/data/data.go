package data

import (
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewBqbRepo,
	NewVoiceRepo,
	NewArticleRepo,
	NewNavigationRepo,
	NewMusicRepo,
)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(100)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1000)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(
		new(biz.BiaoQingBao),
		new(biz.BiaoQingBaoIndex),
		new(biz.Voice),
		new(biz.Article),
		new(biz.Navigation),
		new(biz.Music),
	)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		db: db,
	}, cleanup, nil
}
