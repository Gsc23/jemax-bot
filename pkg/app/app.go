package app

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gsc23/jemax-bot/internal/database"
	"github.com/gsc23/jemax-bot/pkg/config"
	"github.com/gsc23/jemax-bot/pkg/logs"
	"gorm.io/gorm"
)

const (
	keyAppContext = "jemax-chatbot"
)

type App struct {
	Config        *config.Config
	Database      *database.Database
	Logger        logs.Logger
	setupFinished chan struct{}
}

func NewApp(ctx context.Context) *App {
	a := &App{
		setupFinished: make(chan struct{}, 1),
		Config:        config.Load(),
	}

	logs.Setup(
		a.Config.Logger.Level,
		a.Config.Logger.Color,
		a.Config.Logger.Trace,
	)
	a.Logger = logs.DefaultLoggerWithName("app")

	go func() {
		defer func() {
			a.Logger.Debug(ctx, "finishing initialization")
			a.setupFinished <- struct{}{}
		}()

		dbConfig := database.Config{
			Host:     a.Config.Database.Host,
			Port:     a.Config.Database.Port,
			Database: a.Config.Database.Database,
			Username: a.Config.Database.User,
			Password: a.Config.Database.Pass,
		}

		a.Logger.Info(ctx, "connecting to database")
		database, err := database.NewDatabase(dbConfig)
		if err != nil {
			a.Logger.Error(ctx, "could not connect to the database: %v\n", err)
			return
		}

		a.Logger.Info(ctx, "banco de dados iniciado")
		a.Database = database

	}()
	return a
}

func (a *App) WaitReady(ctx context.Context) error {
	select {
	case <-a.setupFinished:
	case <-ctx.Done():
	}
	if a.Database == nil {
		return errors.New("database not initialized")
	}
	return nil
}

func (a *App) Start(ctx context.Context) {
	if err := a.WaitReady(ctx); err != nil {
		a.Logger.Error(ctx, "Could not start app: %v\n", err)
	} else {
		a.Logger.Info(ctx, "app started")
	}
}

func (a *App) InsertIn(c *gin.Context) {
	c.Set(keyAppContext, a)
}

func DB(c *gin.Context) *gorm.DB {
	return c.MustGet(keyAppContext).(*App).Database.DB
}

func Database(c *gin.Context) *database.Database {
	return c.MustGet(keyAppContext).(*App).Database
}

func Logger(c *gin.Context) logs.Logger {
	return c.MustGet(keyAppContext).(*App).Logger
}

func Config(c *gin.Context) *config.Config {
	return c.MustGet(keyAppContext).(*App).Config
}

func VerifyToken(c *gin.Context) string {
	return c.MustGet(keyAppContext).(*App).Config.Whatsapp.VerifyToken
}

func AccessToken(c *gin.Context) string {
	return c.MustGet(keyAppContext).(*App).Config.Whatsapp.AccessToken
}

func PhoneID(c *gin.Context) string {
	return c.MustGet(keyAppContext).(*App).Config.Whatsapp.PhoneID
}
