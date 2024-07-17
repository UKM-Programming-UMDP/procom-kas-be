package other

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDatabaseSchema(db *gorm.DB) {
	CreateSchema(db, "development")
	CreateSchema(db, "production")
}

func CreateSchema(db *gorm.DB, schemaName string) error {
	var existingSchemaName string
	checkQuery := "SELECT schema_name FROM information_schema.schemata WHERE schema_name = ?"
	if err := db.Raw(checkQuery, schemaName).Scan(&existingSchemaName).Error; err != nil {
		return err
	}

	if existingSchemaName == "" {
		fmt.Printf("===== Initialize %s Schema =====\n", schemaName)
		createSchemaQuery := "CREATE SCHEMA " + schemaName
		if err := db.Exec(createSchemaQuery).Error; err != nil {
			return err
		}
		fmt.Printf("Schema %s created\n", schemaName)
	}

	return nil
}
