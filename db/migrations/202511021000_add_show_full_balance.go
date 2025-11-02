package migrations

import (
	_ "embed"
	"text/template"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const showFullBalanceMigration = `ALTER TABLE apps ADD COLUMN show_full_balance BOOLEAN DEFAULT FALSE;`

var showFullBalanceMigrationTmpl = template.Must(template.New("showFullBalanceMigration").Parse(showFullBalanceMigration))

var _202511021000_add_show_full_balance = &gormigrate.Migration{
	ID: "202511021000_add_show_full_balance",
	Migrate: func(tx *gorm.DB) error {

		err := exec(tx, showFullBalanceMigrationTmpl)
		if err != nil {
			return err
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}