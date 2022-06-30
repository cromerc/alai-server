package database

import (
	"backend/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() *gorm.DB {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST_IP")
	dbName := os.Getenv("MYSQL_DATABASE")

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 256,
	}), &gorm.Config{
		CreateBatchSize: 1000,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	return gdb
}

func AutoMigrate(gdb *gorm.DB) {
	var populate = make(map[string]bool)
	if (!gdb.Migrator().HasTable(&models.OS{})) {
		populate["os"] = true
	}
	if (!gdb.Migrator().HasTable(&models.Level{})) {
		populate["level"] = true
	}
	if (!gdb.Migrator().HasTable(&models.User{})) {
		populate["user"] = true
	}
	gdb.AutoMigrate(&models.GodotVersion{})
	gdb.AutoMigrate(&models.Game{})
	gdb.AutoMigrate(&models.Level{})
	gdb.AutoMigrate(&models.OS{})
	gdb.AutoMigrate(&models.Player{})
	gdb.AutoMigrate(&models.ObjectState{})
	gdb.AutoMigrate(&models.ObjectName{})
	gdb.AutoMigrate(&models.Object{})
	gdb.AutoMigrate(&models.Frame{})
	gdb.AutoMigrate(&models.User{})
	Populate(gdb, populate)
}

func DropAll(gdb *gorm.DB) {
	err := gdb.Migrator().DropTable(
		&models.GodotVersion{},
		&models.Object{},
		&models.Frame{},
		&models.Game{},
		&models.Level{},
		&models.OS{},
		&models.Player{},
		&models.ObjectState{},
		&models.ObjectName{},
		&models.User{})
	if err != nil {
		panic(err.Error())
	}
}

func Populate(gdb *gorm.DB, populate map[string]bool) {
	if _, ok := populate["os"]; ok {
		os_list := []models.OS{
			{Name: "Android"},
			{Name: "iOS"},
			{Name: "HTML5"},
			{Name: "OSX"},
			{Name: "Server"},
			{Name: "Windows"},
			{Name: "UWP"},
			{Name: "X11"},
		}

		gdb.Create(&os_list)
	}

	if _, ok := populate["level"]; ok {
		level_list := []models.Level{
			{Name: "Prototype"},
			{Name: "PrototypeR"},
			{Name: "Level1"},
			{Name: "Level2"},
			{Name: "Level3"},
			{Name: "Level4"},
			{Name: "Level5"},
			{Name: "Level6"},
		}

		gdb.Create(&level_list)
	}

	if _, ok := populate["user"]; ok {
		user_list := []models.User{
			{
				Name:     "Chris Cromer",
				Username: "cromer",
				Email:    "chris@cromer.cl",
			},
			{
				Name:     "Martín Araneda",
				Username: "martin",
				Email:    "martix.araneda@gmail.com",
			},
		}
		user_list[0].HashPassword(os.Getenv("ADMIN_PASSWORD"))
		user_list[1].HashPassword(os.Getenv("ADMIN_PASSWORD"))

		gdb.Create(&user_list)
	}
}

func Close(gdb *gorm.DB) {
	sqlDB, err := gdb.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
