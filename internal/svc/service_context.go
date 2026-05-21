package svc

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xxx-newbee/product/internal/config"
	"github.com/xxx-newbee/product/internal/model"
	"github.com/xxx-newbee/storage"
	"github.com/xxx-newbee/storage/cache"
	"github.com/xxx-newbee/storage/locker"
	"github.com/xxx-newbee/storage/queue"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config            config.Config
	Cache             storage.AdapterCache
	Locker            storage.AdapterLocker
	MemoryQueue       storage.AdapterQueue
	RedisQueue        storage.AdapterQueue
	CategoryModel     model.CategoryModel
	ProductModel      model.ProductModel
	ProductSkuModel   model.ProductSkuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := InitDB(c)
	redisClient := InitRedis(c)
	redisQueue := InitRedisQueue(c)

	svcCtx := &ServiceContext{
		Config:          c,
		Cache:           cache.NewRedis(redisClient, nil),
		Locker:          locker.NewRedisLocker(redisClient),
		MemoryQueue:     queue.NewMemoryQueue(c.Queue.Memory.PoolSize),
		RedisQueue:      redisQueue,
		CategoryModel:   model.NewCategoryModel(db),
		ProductModel:    model.NewProductModel(db),
		ProductSkuModel: model.NewProductSkuModel(db),
	}

	return svcCtx
}

func InitDB(c config.Config) *gorm.DB {
	dsn := c.Database.DataSource
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic("failed to get database: " + err.Error())
	}
	sqlDb.SetMaxOpenConns(c.Database.MaxOpenConns)
	sqlDb.SetMaxIdleConns(c.Database.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(c.Database.ConnMaxLifetime) * time.Second)

	if sqlDb.PingContext(context.Background()) != nil {
		panic("failed to ping database: " + err.Error())
	}
	println("PostgreSQL connected")
	return db
}

func InitRedis(c config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Cache.Redis.Addr,
		Password: c.Cache.Redis.Password,
		DB:       c.Cache.Redis.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return client
}

func InitRedisQueue(c config.Config) storage.AdapterQueue {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Queue.Redis.Addr,
		Password: c.Queue.Redis.Password,
		DB:       c.Queue.Redis.DB,
	})
	return queue.NewRedisQueue(client, c.Queue.Redis.Prefix, c.Queue.Redis.MaxRetry)
}
