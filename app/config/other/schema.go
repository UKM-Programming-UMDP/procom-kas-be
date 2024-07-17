package other

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDatabaseSchema(db *gorm.DB) {
	schemaNames := []string{"development", "production"}
	for _, schemaName := range schemaNames {
		if err := CreateSchema(db, schemaName); err != nil {
			fmt.Printf("Error creating schema %s: %v\n", schemaName, err)
		}
	}
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
