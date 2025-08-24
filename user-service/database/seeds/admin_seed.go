package seeds

import (
	"log"
	"user-service/utils/conv" 
	"user-service/internal/core/domain/model"
	"gorm.io/gorm"
)


func SeedAdmin(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	modelRole := model.Role{}
	err = db.Where("name = ?", "Super Admin").First(&modelRole).Error
	if err != nil {
		log.Fatalf("%s: %v, err.Error(), err)")
	}

	admin := model.User{
		Name:     "super admin",
		Email:    "superadmin@email.com",
		Password: bytes,
		Address:  "Jalan Panjang",
		IsVerified: true,
		Roles: []model.Role{modelRole},
	}
	if err := db.FirstOrCreate(&admin, model.User{Email: admin.Email}).Error; err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	} else {
		log.Printf("Admin user %s created", admin.Name)
	}

}