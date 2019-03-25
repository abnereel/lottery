package services

import (
	"github.com/abnereel/lottery/dao"
	"github.com/abnereel/lottery/models"
)

type UserdayService interface {
	GetAll(page, size int) []models.LtUserday
	CountAll() int64
	Search(uid, day int) []models.LtUserday
	Count(uid, day int) int
	Get(id int) *models.LtUserday
	Update(user *models.LtUserday, columns []string) error
	Create(user *models.LtUserday) error
	GetUserToday(uid int) *models.LtUserday
}

type userdayService struct {
	dao *dao.UserdayDao
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao: dao.NewUserdayDao(nil),
	}
}

func (s *userdayService) GetAll(page, size int) []models.LtUserday {
	return s.dao.GetAll(page, size)
}

func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *userdayService) Search(uid, day int) []models.LtUserday {
	return s.dao.Search(uid, day)
}

func (s *userdayService) Count(uid, day int) int {
	return s.dao.Count(uid, day)
}

func (s *userdayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}

func (s *userdayService) Update(user *models.LtUserday, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *userdayService) Create(user *models.LtUserday) error {
	return s.dao.Create(user)
}

func (s *userdayService) GetUserToday(uid int) *models.LtUserday {
	return nil
}