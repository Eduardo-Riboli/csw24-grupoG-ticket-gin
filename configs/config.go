package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Função para configurar e migrar o banco de dados
func SetupDatabase() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// String de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Conectando ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	// Migrar o schema
	err = db.AutoMigrate(
		&models.Sample{},
		&models.Tenant{},
		&models.User{},
		&models.Event{},
		&models.Ticket{},
		&models.Transaction{},
		&models.NotificationPreferences{},
	)
	if err != nil {
		log.Fatal("Falha ao migrar o schema:", err)
	}

	return db
}
