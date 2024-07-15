package user

import (
	"gorm.io/gorm"
	"payment-portal/internal/model"
	"payment-portal/internal/paginator"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) GetByEmailOrName(input string) (*model.User, error) {

	var u model.User

	result := r.Db.Where("email = ?", input).
		Or("name = ?", input).
		First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (r *Repository) GetByEmail(email string) (*model.User, error) {
	var u model.User

	result := r.Db.Where("email = ?", email).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (r *Repository) GetById(id uint) (*model.User, error) {
	var u model.User

	result := r.Db.Where("id = ?", id).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (r *Repository) GetPaginatorWithFilter(p *paginator.Paginator, search string) *ListResult {
	var users []model.User
	var totalItems int64

	query := r.Db.Model(&model.User{})

	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.
		Order("created_at DESC").
		Count(&totalItems).
		Limit(p.GetPageSize()).
		Offset(p.GetOffset()).
		Find(&users)

	paginatorResult := p.ToPaginatorResult(totalItems, len(users))

	return &ListResult{
		Users:     users,
		Paginator: *paginatorResult,
	}
}
