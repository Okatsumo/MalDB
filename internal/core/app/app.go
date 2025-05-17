package app

import (
	"MalDB/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type App struct {
	cfg  config.Config
	ctx  context.Context
	rdb  *redis.Client
	pgdb *pgx.Conn
}

func (app *App) Close() {
	if app.pgdb != nil {
		_ = app.pgdb.Close(app.ctx)
	}
	if app.rdb != nil {
		_ = app.rdb.Close()
	}
}

// New init app
func New(ctx context.Context, cfg *config.Config) *App {
	return &App{cfg: *cfg, ctx: ctx}
}

// CreatePgDB create postgres connection
func (app *App) CreatePgDB() (*pgx.Conn, error) {
	urlDB := fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s",
		app.cfg.DbHost,
		app.cfg.DbPort,
		app.cfg.DB,
		app.cfg.DbUser,
		app.cfg.DbPassword)

	conn, err := pgx.Connect(app.ctx, urlDB)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(app.ctx)
	if err != nil {
		return nil, err
	}

	app.pgdb = conn

	return app.pgdb, nil
}

// CreateRedis create redis connection
func (app *App) CreateRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", app.cfg.RedisHost, app.cfg.RedisPort),
		Password: app.cfg.RedisPassword,
		DB:       0,
	})

	_, err := rdb.Ping(app.ctx).Result()
	if err != nil {
		return nil, err
	}

	app.rdb = rdb
	return rdb, nil
}

// Config return app configuration
func (app *App) Config() config.Config {
	return app.cfg
}

// Redis return connection
func (app *App) Redis() *redis.Client {
	return app.rdb
}

// PgDB return connection
func (app *App) PgDB() *pgx.Conn {
	return app.pgdb
}
