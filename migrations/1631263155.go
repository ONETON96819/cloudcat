package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/scriptscat/cloudcat/internal/domain/user/entity"
	"github.com/scriptscat/cloudcat/pkg/utils"
	"gorm.io/gorm"
)

func T1631263155() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1631263155",
		Migrate: func(db *gorm.DB) error {
			return utils.Errs(
				db.AutoMigrate(&entity.User{}),
				db.AutoMigrate(&entity.WechatOauthUser{}),
				db.AutoMigrate(&entity.BbsOauthUser{}),
			)
		},
		Rollback: func(db *gorm.DB) error {
			return utils.Errs(
				db.Migrator().DropTable(&entity.User{}),
				db.Migrator().DropTable(&entity.WechatOauthUser{}),
				db.Migrator().DropTable(&entity.BbsOauthUser{}),
			)
		},
	}
}