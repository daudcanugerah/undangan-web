package cmd

import (
	"log"
	"net/http"

	"basic-service/interface/rest"
	"basic-service/interface/sql"
	"basic-service/usecase"

	"github.com/spf13/cobra"
)

var appCmd = cobra.Command{
	Use: "app",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := sql.NewSQLite(systemConfig.DBSqlite.DBFile)
		if err != nil {
			return err
		}
		guestManager := sql.NewGuestManager(db)
		publicTemplate := sql.NewPublicTemplateRepository(db)
		userTemplate := sql.NewUserTemplateRepository(db)
		userManager := sql.NewUserRepository(db)

		auth := usecase.NewAuth(userManager, "secret")
		publicTemplateUseCase := usecase.NewPublicTemplateUseCase(publicTemplate)
		userTemplateCase := usecase.NewUserTemplate(userTemplate)
		guestUsecase := usecase.NewGuestUsecase(guestManager)

		r := rest.SetupRouter(auth, publicTemplateUseCase, userTemplateCase, guestUsecase)

		log.Println("Server starting on :8085")
		if err := http.ListenAndServe(":8085", r); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(&appCmd)
}
