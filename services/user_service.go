package services

import (
	"fmt"
	"github.com/abnereel/lottery/comm"
	"github.com/abnereel/lottery/dao"
	"github.com/abnereel/lottery/datasource"
	"github.com/abnereel/lottery/models"
	"github.com/gomodule/redigo/redis"
	"log"
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
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}

func (s *userService) GetAll(page, size int) []models.LtUser {
	return s.dao.GetAll(page, size)
}

func (s *userService) CountAll() int {
	return s.dao.CountAll()
}

func (s *userService) Get(id int) *models.LtUser {
	data := s.getByCache(id)
	if data == nil || data.Id <= 0 {
		data = s.dao.Get(id)
		if data == nil || data.Id <= 0 {
			data = &models.LtUser{Id: id}
		}
		s.setByCache(data)
	}
	return data
}

func (s *userService) Update(data *models.LtUser, columns []string) error {
	s.updateByCache(data, columns)
	return s.dao.Update(data, columns)
}

func (s *userService) Create(data *models.LtUser) error {
	return s.dao.Create(data)
}

// 从缓存中得到信息
func (s *userService) getByCache(id int) *models.LtUser {
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		log.Println("user_service.getByCache HGETALL key=", key, ", error=", err)
		return nil
	}
	dataId := comm.GetInt64FromStringMap(dataMap, "Id", 0)
	if dataId <= 0 {
		return nil
	}
	data := &models.LtUser{
		Id:         int(dataId),
		Username:   comm.GetStringFromStringMap(dataMap, "Username", ""),
		Blacktime:  int(comm.GetInt64FromStringMap(dataMap, "Blacktime", 0)),
		Realname:   comm.GetStringFromStringMap(dataMap, "Realname", ""),
		Mobile:     comm.GetStringFromStringMap(dataMap, "Mobile", ""),
		Address:    comm.GetStringFromStringMap(dataMap, "Address", ""),
		SysCreated: int(comm.GetInt64FromStringMap(dataMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(dataMap, "SysUpdated", 0)),
		SysIp:      comm.GetStringFromStringMap(dataMap, "SysIp", ""),
	}
	return data
}

// 将信息更新到缓存
func (s *userService) setByCache(data *models.LtUser) {
	if data == nil || data.Id <= 0 {
		return
	}
	id := data.Id
	key := fmt.Sprintf("info_user_%d", id)
	rds := datasource.InstanceCache()
	// 数据更新到redis缓存
	params := redis.Args{key}
	params = append(params, "Id", id)
	if data.Username != "" {
		params = append(params, "Username", data.Username)
		params = append(params, "Blacktime", data.Blacktime)
		params = append(params, "Realname", data.Realname)
		params = append(params, "Mobile", data.Mobile)
		params = append(params, "Address", data.Address)
		params = append(params, "SysCreated", data.SysCreated)
		params = append(params, "SysUpdated", data.SysUpdated)
		params = append(params, "SysIp", data.SysIp)
	}
	_, err := rds.Do("HMSET", params...)
	if err != nil {
		log.Println("user_service.setByCache HMSET params=", params, ", error=", err)
	}
}

// 数据更新了，直接清空缓存数据
func (s *userService) updateByCache(data *models.LtUser, column []string) {
	if data == nil || data.Id <= 0 {
		return
	}
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_user_%d", data.Id)
	rds := datasource.InstanceCache()
	// 删除redis中的缓存
	_, _ = rds.Do("DEL", key)
}
