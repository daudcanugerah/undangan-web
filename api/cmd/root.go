package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"basic-service/config"
	"basic-service/pkg/otel"
	"basic-service/system"

	"github.com/spf13/cobra"
)

var (
	ctx               context.Context = context.Background()
	mainOtel          otel.Otel
	shudtdownFuncList = []func(ctx context.Context) error{}
	systemConfig      *config.Config
)

var (
	defaultHttpClient = &http.Client{Timeout: time.Second * 60}
	proxyClient       = &http.Client{Timeout: time.Second * 60}
)

var (
	cfgFileInput  string
	logLevelInput string
)
var rootCmd = &cobra.Command{}

func Initialize(ctx context.Context) error {
	var err error

	if systemConfig, err = config.InitConfig(cfgFileInput); err != nil {
		return err
	}

	// setup otel
	otelConfig := otel.SetupOption{
		EnableMetric:   systemConfig.Otel.EnableMetric,
		EnableTrace:    systemConfig.Otel.EnableTrace,
		EnableLog:      systemConfig.Otel.EnableLog,
		ServiceName:    system.APP_NAME,
		ServiceVersion: system.APP_VERSION,
	}

	shutdownFunc, err := otel.SetupOTelSDK(ctx, otelConfig)
	if err != nil {
		return err
	}

	mainOtel = otel.NewOtel("main", logLevelInput)
	shudtdownFuncList = append(shudtdownFuncList, shutdownFunc)

	// setup timezone
	if err := config.SetUpTimezone(systemConfig.App.Timezone); err != nil {
		return err
	}

	return nil
}

func Execute() int {
	defer func() {
		for _, v := range shudtdownFuncList {
			fmt.Println("shutting down function", v)
			v(ctx)
		}
	}()

	cobra.OnInitialize(func() {
		if err := Initialize(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "%+v", err)
			os.Exit(1)
		}
	})

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		if !mainOtel.IsLogAvailable() {
			fmt.Fprintf(os.Stderr, "error when running command %+v", err)
			return 1
		}

		mainOtel.Log.Error(ctx, fmt.Sprintf("app exit error: %s", err), "stacktrace", fmt.Errorf("%+v", err))
		return 1
	}

	return 0
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFileInput, "config", "", "", "config file (default is $HOME/.config.toml")
	rootCmd.PersistentFlags().StringVarP(&logLevelInput, "log-level", "", "info", "log level, available level debug,info,warn,error")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
