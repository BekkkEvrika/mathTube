package models

import "mathTube/base"

func AutoMigrate() error {
	if err := base.Migrate(&Person{}); err != nil {
		return err
	}
	if err := base.Migrate(&Material{}); err != nil {
		return err
	}
	return nil
}
