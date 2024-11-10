package configs

import (
	"fmt"
	"log"
	"os"
	"time"

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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)

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

	createInitialData(db)
	
	return db
}

// Função para criar registros iniciais
func createInitialData(db *gorm.DB) {
	// Criando Samples
	sample1 := models.Sample{ID: 1, Name: "Sample A", Email: "Description A"}
	sample2 := models.Sample{ID: 2, Name: "Sample B", Email: "Description B"}
	db.Create(&sample1)
	db.Create(&sample2)

	// Criando Tenants
	tenant1 := models.Tenant{Name: "Tenant A", ContactInfo: "contact@tenanta.com", SpecificConfigurations: "Config A"}
	tenant2 := models.Tenant{Name: "Tenant B", ContactInfo: "contact@tenantb.com", SpecificConfigurations: "Config B"}
	db.Create(&tenant1)
	db.Create(&tenant2)

	// Criando Users
	user1 := models.User{Name: "Alice", Email: "alice@example.com", TenantID: tenant1.ID}
	user2 := models.User{Name: "Bob", Email: "bob@example.com", TenantID: tenant2.ID}
	db.Create(&user1)
	db.Create(&user2)

	// Criando Events
	event1 := models.Event{TenantID: tenant1.ID, Name: "Conference", Type: "Business", Location: "New York", Date: time.Now().AddDate(0, 1, 0)}
	event2 := models.Event{TenantID: tenant2.ID, Name: "Music Festival", Type: "Entertainment", Location: "Los Angeles", Date: time.Now().AddDate(0, 2, 0)}
	db.Create(&event1)
	db.Create(&event2)

	// Criando Tickets
	ticket1 := models.Ticket{EventID: event1.ID, TenantID: tenant1.ID, OriginalPrice: 100.00, SellerID: user1.ID, VerificationCode: "ABC123", Status: "Available"}
	ticket2 := models.Ticket{EventID: event2.ID, TenantID: tenant2.ID, OriginalPrice: 150.00, SellerID: user2.ID, VerificationCode: "XYZ789", Status: "Available"}
	db.Create(&ticket1)
	db.Create(&ticket2)

	// Criando Transactions
	transaction1 := models.Transaction{TenantID: tenant1.ID, BuyerID: user2.ID, TicketID: ticket1.ID, SalePrice: 95.00, TransactionDate: time.Now(), TransactionStatus: "Completed"}
	db.Create(&transaction1)

	// Criando NotificationPreferences
	preferences1 := models.NotificationPreferences{UserID: user1.ID, ReceiveEmails: true}
	preferences2 := models.NotificationPreferences{UserID: user2.ID, ReceiveEmails: false}
	db.Create(&preferences1)
	db.Create(&preferences2)

	fmt.Println("Dados iniciais criados com sucesso!")
}