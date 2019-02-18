package cmd

import (
	"context"
	"fmt"
	breederDeliver "github.com/anthonydenecheau/gopocservice/breeder/delivery/http"
	breederRepo "github.com/anthonydenecheau/gopocservice/breeder/repository"
	breederUcase "github.com/anthonydenecheau/gopocservice/breeder/usecase"
	cfg "github.com/anthonydenecheau/gopocservice/config/env"
	"github.com/anthonydenecheau/gopocservice/config/middleware"
	healthDeliver "github.com/anthonydenecheau/gopocservice/health/delivery/http"
	"github.com/anthonydenecheau/gopocservice/health/delivery/renderings"
	"github.com/go-pg/pg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"time"
)

type dbLogger struct{}

var config cfg.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopocservice",
	Short: "Rest-Api Poc",
	Long:  `Cette Api est la migration du WS dog développé avec le stack Spirng Boot, Netflix.'`,

	Run: func(cmd *cobra.Command, args []string) {
		environment, _ := cmd.Flags().GetString("environment")
		if environment == "" {
			environment = "dev"
		}
		if viper.GetString("environment") != "" {
			environment = viper.GetString("environment")
		}
		fmt.Println(" > Start Environment :  " + environment)

		// Load configuration
		// [TODO]
		// CF. https://github.com/bxcodec/go-clean-arch
		fmt.Println(" > Load configuration :  " + environment)

		db := pg.Connect(&pg.Options{
			User:                  config.GetString(`database.user`),
			Password:              config.GetString(`database.pass`),
			Database:              config.GetString(`database.name`),
			Addr:                  config.GetString(`database.host`) + ":" + config.GetString(`database.port`),
			RetryStatementTimeout: true,
			MaxRetries:            4,
			MinRetryBackoff:       250 * time.Millisecond,
		})

		var n int
		_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
		if err != nil && config.GetBool("debug") {
			fmt.Println(err)
		}
		defer db.Close()
		db.AddQueryHook(dbLogger{})

		// Start API server
		fmt.Println(" > Start Api :  " + environment)
		e := echo.New()

		middL := middleware.InitMiddleware()
		e.Use(middL.CORS)
		e.Use(middL.RequestIDMiddleware)

		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}
			if he, ok := err.(*echo.HTTPError); ok {
				c.JSON(he.Code, renderings.Error{
					Status:  he.Code,
					Message: he.Error(),
				})
			}
		}

		br := breederRepo.NewPgBreederRepository(db)
		bu := breederUcase.NewBreederUsecase(br)
		breederDeliver.NewBreederHttpHandler(e, bu)

		healthDeliver.NewHealthHttpHandler(e)

		// Start server
		go func() {
			if err := e.Start(config.GetString("server.address")); err != nil {
				e.Logger.Info("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}

	},
}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}
func RootCommand() *cobra.Command {
	return rootCmd
}
func init() {

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
	//rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	rootCmd.Flags().StringP("environment", "n", viper.GetString("ENVIRONMENT"), "Set your environment")

	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}
func waitForShutdown(r *echo.Echo) {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Fatal(err)
	}
}
