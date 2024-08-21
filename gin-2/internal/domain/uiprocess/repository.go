package uiprocess

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) FindByProcessId(processId int) (UipathProcesses, error) {
	var uipathProcess UipathProcesses
	result := r.Db.Preload("Gateway").Where("uipath_process_id = ?", processId).First(&uipathProcess)
	if result.Error != nil {
		return uipathProcess, result.Error
	}
	return uipathProcess, nil
}
