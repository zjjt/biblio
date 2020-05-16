package user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	us "github.com/zjjt/biblio/user_service/proto/user"
)

//BeforeCreate is a postgresql gorm related function enabling us to
//add uuid for example
func (model *us.User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
