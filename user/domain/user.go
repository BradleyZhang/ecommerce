package domain

type User struct {
	Id           int    `json:"id"`
	UserName     string `json:"username" validate:"max=20"`
	Password     string `json:"password" validate:"min=8,max=20"`
	DisplayName  string `json:"display_name" validate:"max=20"`
	Setting      string `json:"setting"`
	CreateAt     int64  `json:"create_at"`
	LastLogintAt string `json:"last_logint_at"`
	Role         int    `json:"role"`
}
type ShippingAddress struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Region        string `json:"region"`
	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	Street        string `json:"street"`
	AddressLine1  string `json:"address_line_1"`
	AddressLine2  string `json:"address_line_2"`
	RecipientName string `json:"recipient_name"`
	PhoneNumber   string `json:"phone_number"`
	PostalCode    string `json:"postal_code"`
}
type UserRepository interface {
	Create(user User) (id int, err error)
	GetAll() (users []*User, total int64, err error)
	GetById(id int) (*User, error)
	Exist(username string) (exist bool, err error)
	Update(user User) error
	DeleteById(id int) error
}
type AddressesRepository interface {
	Create(address ShippingAddress) error
	GetAllByUserId(userId int) (addresses []*ShippingAddress, total int64, err error)
	Update(address ShippingAddress) error
	DeleteById(id int) error
	DeleteAllByUserId(userId int) error
}
type Validater interface {
	Struct(s interface{}) error
}

const (
	//Role
	UserRole = 1
)
