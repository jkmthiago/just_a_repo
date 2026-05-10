package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Bora primeiro conectar essa bagaça no banco
	// Primeiro nos declaramos a DSN (Data Source Name, ou em protugues Nome da Fonte de Dados)
	// Basicamente a config que tu precisa passar para conectar ao bd
	dsn := `host=just_a_postgres_container
	        port=5432
			user=just_a_postgres_container
			password=759134680
			dbname=just_a_postgres_container
			sslmode=disable`

	// DSN criada agora pora conectar nessa bagaça
	// gorm.open inicializa a conexção por meio do uso de um driver de conexão e as configurações de conexão e outras opções de uso
	// postgres.open e justamente a passagem desse driver para que a integração seja feita com o banco por meio da DSN passada
	// &gorm.Config é exatamente algumas configs de uso, no caso aqui estou passando um prefixo de esquema (schema) por meio de
	// TablePrefix e indicando não singulariedade de nome de tabela por meio de SingularTable, o que significa que uma struct
	// Product{} ficaria no bd products com essa configuração em falso e manteria singulariedade caso verdadeiro
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "just_a_schema.",
			SingularTable: false,
		},
	})
	if err != nil { // aquela checada se deu merda que o go adora
		log.Fatal(err)
	}

	// Aqui eu crio, caso não exista, o esquema (schema) que a aplicação vai utilizar
	if err := db.Exec(`CREATE SCHEMA IF NOT EXISTS "just_a_schema"`).Error; err != nil { // aquela checada parte 2 - o retorno dos que não foram
		log.Fatal(fmt.Errorf("failed to create schema %s: %w", "just_a_schema", err))
	}

	// Agora que a conexão foi feita e o esquema criado, bora criar a tabela
	// AutoMigrate é um método do gorm que tem a função de criar ou atualizar a estrutura da tabela no banco de 
	// dados com base na definição da struct passada como argumento
	db.AutoMigrate(&Product{})
}
