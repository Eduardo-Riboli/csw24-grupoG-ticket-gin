package configs

import (
	"fmt"
	"log"
	"time"

	"github.com/grupoG/csw24-grupoG-ticket-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Função para configurar e migrar o banco de dados
func SetupDatabase() *gorm.DB {
	host := "postgresdb"
	user := "postgres"
	password := "postgres"
	dbname := "mydb"
	port := "5432"

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

	// Inserir dados iniciais
	createInitialData(db)

	return db
}
