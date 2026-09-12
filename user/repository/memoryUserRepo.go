package repository

import (
	"time"

	"github.com/BradleyZhang/ecommerce/user/domain"
)

type MemUserRepo struct {
}

func NewMemRepo() *MemUserRepo {
	return &MemUserRepo{}
}

var (
	userMap   = make(map[int]domain.User)
	maxUserId = 0
)

func (r *MemUserRepo) Create(user domain.User) (id int, err error) {
	user.Id = maxUserId + 1
	maxUserId++
	user.CreateAt = int64(time.Now().Unix())
	userMap[user.Id] = user
	return user.Id, nil
}
func (r *MemUserRepo) GetAll() (users []*domain.User, total int64, err error) {
	return nil, 0, nil
}
func (r *MemUserRepo) GetById(id int) (*domain.User, error) {
	return nil, nil
}
func (r *MemUserRepo) Exist(username string) (exist bool, err error) {
	for _, user := range userMap {
		if user.UserName == username {
			return true, nil
		}
	}
	return false, nil
}
func (r *MemUserRepo) Update(user domain.User) error {
	return nil
}
func (r *MemUserRepo) DeleteById(id int) error {
	return nil
}

// GetAll() (users []*User, total int64, err error)
// GetById(id int) (*User, error)
// Update(user User) error
// DeleteById(id int) error
