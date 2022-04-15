package user

import (
	"errors"

	"github.com/Chubacabrazz/picus-storeApp/storage/helper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func GetUserByID(uid int) (User, error) {
	var r *UserRepository
	var u User

	if err := r.db.First(&u, uid).Error; err != nil {
		return u, errors.New("user not found")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(Email string, password string) (string, error) {
	var r *UserRepository
	var err error

	u := User{}

	err = r.db.Model(User{}).Where("Email = ?", Email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := helper.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (r *UserRepository) Migration() {
	r.db.AutoMigrate(&User{})
}

func (u *User) SaveUser() (*User, error) {
	var r *UserRepository
	err := r.db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

/* func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

} */
