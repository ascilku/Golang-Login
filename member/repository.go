package member

import "gorm.io/gorm"

type Respository interface {
	SaveRespository(member Member) (Member, error)
	FindByEmail(nama string) (Member, error)
}

type respository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *respository {
	return &respository{db}
}

func (r *respository) SaveRespository(member Member) (Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *respository) FindByEmail(nama string) (Member, error) {
	var member Member
	err := r.db.Where("nama = ?", nama).Find(&member).Error
	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}
