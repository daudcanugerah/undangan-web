package cmd

import (
	"basic-service/pkg/otel"

	"github.com/spf13/cobra"
)

var SampleService = cobra.Command{
	Use:   "basic",
	Short: "Sample service command",
	RunE: func(cmd *cobra.Command, args []string) error {
		return otel.CobraFuncEWithLogger(cmd.Context(), "basic", logLevelInput, func(otl otel.Otel) error {
			otl.Log.Infof(cmd.Context(), "Sample service is running with log level: %s", logLevelInput)
			otl.Log.Warnf(cmd.Context(), "This is a sample service command. It does not perform any real operations.")
			otl.Log.Error(cmd.Context(), "This is an error message for demonstration purposes.")
			otl.Log.Debug(cmd.Context(), "Debugging information can be added here.")

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(&SampleService)
}
