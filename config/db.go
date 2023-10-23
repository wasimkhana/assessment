package config

import (
	"fmt"
	"github.com/wasimkhandot01/assessment/models"
	"github.com/wasimkhandot01/assessment/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// tried to use factory approach
type DBClient struct {
	db *gorm.DB
}

// connect to postgres-db inside container
func NewDBFactory() (*DBClient, error) {
	dsn := "host=172.17.0.2 user=devuser dbname=devdb password=changeme sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return &DBClient{db}, nil
}

func (f *DBClient) Ping() (string, error) {
	var dbName string
	err := f.db.Raw("SELECT current_database()").Scan(&dbName).Error
	return dbName, err
}

func (f *DBClient) TablesCreationTransaction(queries interface{}) error {
	err := f.DropAllTables()
	if err != nil {
		return err
	}

	tx := f.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	switch queryList := queries.(type) {
	case []string:
		for _, query := range queryList {
			if err := tx.Exec(query).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

	case map[string]string:
		for _, query := range queryList {
			if err := tx.Exec(query).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	default:
		return fmt.Errorf("Unsupported query parameter type")
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (f *DBClient) DropAllTables() error {

	tx := f.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, statement := range utils.DropTableStatements {
		if err := tx.Exec(statement).Error; err != nil {
			f.db.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (f *DBClient) PrepareDB() error {

	//var statementarray []string
	//statementarray = append(statementarray, utils.CreateDepartmentTableSQL,
	//	utils.InsertDepartmentsSQL, utils.CreateEmployeeTableSQL, utils.InsertEmployeesSQL)

	tx := f.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	count := 0
	for _, statement := range utils.Statmentarray {
		count += 1
		if err := tx.Exec(statement).Error; err != nil {
			f.db.Rollback()
			return err
		}

	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (f *DBClient) TopThreeHighEarners(query string) ([]models.HighEarner, error) {
	var highEarners []models.HighEarner
	tx := f.db.Begin()
	rows := tx.Raw(query).Scan(&highEarners)
	if rows.Error != nil {
		return nil, rows.Error
	}

	// to remove tables for creating each time new
	tx.Exec("DROP TABLE IF EXISTS Employee")
	tx.Exec("DROP TABLE IF EXISTS Department")
	tx.Commit()

	return highEarners, nil
}
