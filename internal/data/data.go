package data

import (
	"context"
	"github.com/SuKaiFei/go-wxxcx/internal/biz"
	"github.com/SuKaiFei/go-wxxcx/internal/conf"
	"github.com/go-redis/redis/v8"
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
	NewChatGPTRepo,
	NewCommunityRepo,
	NewSecurityRepo,
	NewWechatRepo,
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // no password set
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
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
		new(biz.ChatGPT),
		new(biz.ChatGPTQuota),
		new(biz.CommunityUser),
		new(biz.CommunityArticle),
		new(biz.CommunityComment),
		new(biz.CommunityLike),
		new(biz.CommunityFeedback),
	)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		rdb.Close()
		sqlDB.Close()
		log.NewHelper(logger).Info("closed the data resources")
	}

	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}
