package main

import (
	"context"
	"log"
	"os"

	gooauthinterface "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-interface"
	gooauthservice "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-service"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	godotenv.Load()
	var oauthin gooauthinterface.GoOauthInterface
	database := NewDatabase("host=localhost user=postgres password=root dbname=go-user-auth port=5432 TimeZone=America/Bogota")
	oauthin = gooauthservice.NewGoOauthService(database, "jd_secret")

	jwt, err := oauthin.LoginUser(context.Background(), "judas3991@gmail.com", "1234567")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*jwt)
}

func NewDatabase(pg_url string) *gorm.DB {
	config := &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "go-oauth_"}}

	db, err := gorm.Open(postgres.Open(pg_url), config)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return db
}
