package user

import (
	"log"

	"github.com/alaa2amz/g1/service"
)

func init() {
	log.Println(Path + "init")
	if service.R == nil {
		log.Fatal("main router not initialized")
	}

	R = service.R
	Register(R)

	if service.DB == nil {
		log.Fatal("main database not initialized")
	}

	DB = service.DB

//	for _, dropColumn := range DroppedColumns {
//		DB.Migrator().DropColumn(Proto(), dropColumn)
//	}
	DB.AutoMigrate(Proto())
	service.Paths[Path] = Proto()
	service.Index = append(service.Index,Path)
}

func Init() {
	//TODO:log only
	DB.AutoMigrate(Proto())
	DB = service.DB
}
