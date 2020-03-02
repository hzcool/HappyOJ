package adminApp

import "time"

type Admin struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"_"`
	Name string `gorm:"size:20;UNIQUE_INDEX" form:"name"  binding:"required" validate:"lte=20,required" json:"name"`
	Password string `gorm:"size:16" form:"password"  binding:"required" validate:"gte=6,lte=16,required,printascii" json:"password"`
	Access string `gorm:"size:12" json:"access"`   //'super_admin和admin'
	Avatar string `json:"avatar"`                  //头像路径
}


func models_init()  {
	db.AutoMigrate(&Admin{})
	if super_admin := GetAdminByName("super_admin");super_admin==nil{
		db.Create(&Admin{
			Name:      "super_admin",
			Password:  "123456",
			Access:    "super_admin",
			Avatar:    "",
		})
	}
}

