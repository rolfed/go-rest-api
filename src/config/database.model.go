package config

import "github.com/jinzhu/gorm"

// GormClient is a wrapper for database
type GormClient struct {
	crDB *gorm.DB
}
