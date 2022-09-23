package orm

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"github.com/zazhedho/gorental/src/database/orm/models"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	RunE:  dbMigrate,
}

var migUp, migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "0001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.History{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.User{}, &models.Vehicle{}, &models.History{})
			},
		},
	})

	if migUp {
		if err = m.Migrate(); err != nil {
			return err
		}
		log.Println("Migration database successfully")
		return nil
	}

	if migDown {
		if err = m.RollbackLast(); err != nil {
			return err
		}
		log.Println("Rollback database successfully")
		return nil
	}

	log.Println("init schema database done")
	return nil
}
