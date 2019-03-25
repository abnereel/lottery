package services

import (
	"github.com/abnereel/lottery/dao"
	"github.com/abnereel/lottery/models"
)

type UserService interface {
	GetAll(page, size int) []models.LtUser
	CountAll() int
	Get(id int) *models.LtUser
	Update(user *models.LtUser, columns []string) error
	Create(user *models.LtUser) error
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(nil),
	}
}

func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}

func (s *userService) CountAll() int {
	return s.dao.CountAll()
}

func (s *userService) Get(id int) *models.LtUser {
	return s.dao.Get(id)
}

func (s *userService) Update(user *models.LtUser, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *userService) Create(user *models.LtUser) error {
	return s.dao.Create(user)
}