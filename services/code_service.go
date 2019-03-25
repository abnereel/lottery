package services

import (
	"github.com/abnereel/lottery/dao"
	"github.com/abnereel/lottery/models"
)

type CodeService interface {
	GetAll(page, size int) []models.LtCode
	CountAll() int64
	CountByGift(giftId int) int64
	Search(giftId int) []models.LtCode
	Get(id int) *models.LtCode
	Delete(id int) error
	Update(user *models.LtCode, columns []string) error
	Create(user *models.LtCode) error
	NextUsingCode(giftId, codeId int) *models.LtCode
	UpdateByCode(data *models.LtCode, columns []string) error
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(nil),
	}
}

func (s *codeService) GetAll(page, size int) []models.LtCode {
	return s.dao.GetAll(page, size)
}

func (s *codeService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *codeService) CountByGift(giftId int) int64 {
	return s.dao.CountByGift(giftId)
}

func (s *codeService) Search(giftId int) []models.LtCode {
	return s.dao.Search(giftId)
}

func (s *codeService) Get(id int) *models.LtCode {
	return s.dao.Get(id)
}

func (s *codeService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *codeService) Update(user *models.LtCode, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *codeService) Create(user *models.LtCode) error {
	return s.dao.Create(user)
}

func (s *codeService) NextUsingCode(giftId, codeId int) *models.LtCode {
	return s.dao.NextUsingCode(giftId, codeId)
}

func (s *codeService) UpdateByCode(data *models.LtCode, columns []string) error {
	return s.dao.UpdateByCode(data, columns)
}