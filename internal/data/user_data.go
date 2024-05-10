package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nartvt/smart-service/ent"
	"github.com/nartvt/smart-service/ent/user"
	"github.com/nartvt/smart-service/internal/biz"
	"github.com/nartvt/smart-service/internal/constants"
	"github.com/nartvt/smart-service/internal/module"
)

type userRepo struct {
	db  *Data
	log *log.Helper
}

func NewUserRepo(db *Data) biz.IUserRepo {
	return &userRepo{
		db:  db,
		log: log.NewHelper(log.With(log.DefaultLogger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user module.User) error {
	_, err := r.db.client.User.Create().
		SetPhoneNumber(user.PhoneNumber).
		SetUsername(user.UserName).
		SetEmail(user.Email).
		SetHashPassword(user.HashPassword).
		SetStatus(string(constants.ACTIVE)).
		SetVerified(false).
		Save(ctx)
	return err
}
func (r *userRepo) GetUserByUserName(ctx context.Context, userName string) (module.User, error) {
	user, err := r.db.client.User.Query().Where(user.UsernameEQ(userName)).Only(ctx)
	if err != nil {
		return module.User{}, err
	}
	return mapUserToModuleUser(user), nil
}

func (r *userRepo) GetUserByUserNameOrEmail(ctx context.Context, userName, email string) (module.User, error) {
	user, err := r.db.client.User.Query().Where(user.UsernameEQ(userName), user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return module.User{}, err
	}
	return mapUserToModuleUser(user), nil
}

func mapUserToModuleUser(user *ent.User) module.User {
	return module.User{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhoneNumber:  user.PhoneNumber,
		Street:       user.Street,
		CityId:       user.CityID,
		DistrictId:   user.DistrictID,
		CountryCode:  user.CountryCode,
		HashPassword: user.HashPassword,
		UserName:     user.Username,
		Avatar:       user.Avatar,
		Address:      user.Address,
		Gender:       user.Gender,
		BirthDay:     user.BirthDay,
		Status:       user.Status,
		Verified:     user.Verified,
	}
}
