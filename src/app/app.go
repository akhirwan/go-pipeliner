package app

import (
	"context"
	"go-pipeliner/src/app/task"
	"go-pipeliner/src/infrastructure/config"
	"go-pipeliner/src/infrastructure/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	ID      string
	Name    string
	Version string
	Port    string
	Logger  logger.LogInterface
	Config  config.Config
}

func NewApp() *App {
	configuration := config.New()
	return &App{
		ID:      configuration.Get("APP_ID"),
		Name:    configuration.Get("APP_NAME"),
		Version: configuration.Get("APP_VERSION"),
		Port:    configuration.Get("APP_PORT"),
		Logger:  logger.NewLoggerFactory(logger.Logrus, logger.InfoLevel, true),
		Config:  configuration,
	}
}

func (app *App) Run() *App {
	app.Logger.Info("%s %s started in port %s", app.ID, app.Name, app.Port)

	/* == API WEB SERVICE == */

	/* == DEV2PROD == */
	go func() {
		// app.RunDevToProdSalesCustomerTask()
	}()

	/* == PROD2DEV == */
	go func() {
		// app.RunProdToDevSalesDatamartDailyTask()
	}()

	return app
}

func (app *App) RunDevToProdSalesCustomerTask() {
	from, _ := time.Parse("2006-01-02", "2013-01-01")
	to := time.Now()

	task := task.NewDevToProdSalesCustomerPipelineTask(app.Config)
	task.Execute(from, to)
}

func (app *App) RunProdToDevSalesDatamartDailyTask() {
	from := time.Now()
	to := time.Now()

	task := task.NewProdToDevSalesDatamartDailyPipelineTask(app.Config)
	task.Execute(from, to)
}

// WithGracefulShutdown ...
func (app *App) WithGracefulShutdown() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go app.Logger.Info("Press ctrl+c to stop the service")

	sig := <-c // Blocking for graceful
	app.Logger.Info("\n⛔️  Got %s signal. Shutting down app...", sig)
	// Give some time to finish ongoing request
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	defer ctx.Done()

	// app.WebServer.Shutdown(ctx)
	app.Logger.Info("App has been completely shutdown. Asta Lavista!")

	os.Exit(0)
	return nil
}
