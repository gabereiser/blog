package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(test bool) {
	if test {
		// Initialize a test database connection
		ConnectTest()
	} else {
		// Initialize a production database connection
		Connect()
	}
}

func DB() *gorm.DB {
	return db
}

func Where[T any](model T, query string, args ...any) ([]T, error) {
	var results []T
	if err := db.Where(query, args...).Find(&results).Error; err != nil {
		return results, err
	}
	return results, nil
}

func Find[T any](model T, id uint) (T, error) {
	var result T
	if err := db.First(&result, id).Error; err != nil {
		return result, err
	}
	return result, nil
}

func First[T any](model T) (T, error) {
	var result T
	if err := db.First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func FindBy[T any](model T, field string, value any) (T, error) {
	var result T
	if err := db.Where(field+" = ?", value).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func Create[T any](model T) (T, error) {
	if err := db.Create(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func Update[T any](model T) (T, error) {
	if err := db.Save(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func Delete[T any](model T) error {
	if err := db.Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func Close() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if sqlDB != nil {
		return sqlDB.Close()
	}
	return nil
}
