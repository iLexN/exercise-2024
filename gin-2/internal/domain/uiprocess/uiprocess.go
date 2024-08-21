package uiprocess

import (
	"payment-portal/internal/domain/gateway"
	"time"
)

type UipathProcesses struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UipathProcessId    int
	UiPathKey          string
	UiPathProcessKey   string
	OrganizationUnitId int

	GatewaysID int
	Gateway    gateway.Gateway `gorm:"foreignKey:GatewaysID"`
}
