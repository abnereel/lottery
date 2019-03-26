package services

import (
	"github.com/abnereel/lottery/dao"
	"github.com/abnereel/lottery/datasource"
	"github.com/abnereel/lottery/models"
)

type BlackipService interface {
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Search(ip string) []models.LtBlackip
	Get(id int) *models.LtBlackip
	Update(user *models.LtBlackip, columns []string) error
	Create(user *models.LtBlackip) error
	GetByIp(ip string) *models.LtBlackip
}

type blackipService struct {
	dao *dao.BlackipDao
}

func NewBlackService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDao(datasource.InstanceDbMaster()),
	}
}

func (s *blackipService) GetAll(page, size int) []models.LtBlackip {
	return s.dao.GetAll(page, size)
}

func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *blackipService) Search(ip string) []models.LtBlackip {
	return s.dao.Search(ip)
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}

func (s *blackipService) Update(user *models.LtBlackip, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *blackipService) Create(user *models.LtBlackip) error {
	return s.dao.Create(user)
}

func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	return s.dao.GetByIp(ip)
}