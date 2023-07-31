package api

import (
	"Ecommerce/middleware/jwt"
	"Ecommerce/user_service/auth"
	"Ecommerce/user_service/models"
	"Ecommerce/user_service/store"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserStore struct {
	UsersStore store.UserStore
}

func NewUserStoreApi(storeDependency store.Dependency) *UserStore {
	return &UserStore{
		UsersStore: storeDependency.UsersStore,
	}
}

type UsersService interface {
	SignUpUser(models.SignUpUserRequest) (string, error)
	LoginUser(models.SignInUserRequest) (string, error)
	UpdateUser(models.UpdateUserRequest) error
	GetUserById(models.GetUserByIDRequest) (*models.Users, error)
	DeleteUser(models.DeleteUserByIDRequest) error
}

func (userstore *UserStore) SignUpUser(user models.SignUpUserRequest) (string, error) {
	var jwtString string
	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		return "", errJsonValidation
	}
	//Hash the password
	hashPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return "", errInternalServerError
	}

	dbUser := models.Users{
		User_ID:  uuid.New().String(),
		Email:    user.Email,
		Password: hashPassword,
	}

	if err := userstore.UsersStore.Create(dbUser); err != nil {
		return "", errInternalServerError
	}

	if user.IsAdmin {
		jwtString, err = jwt.NewToken(dbUser.User_ID, jwt.AdminScope)
	} else {
		jwtString, err = jwt.NewToken(dbUser.User_ID, jwt.UserScope)
	}
	if err != nil {
		return "", errInternalServerError
	}

	return jwtString, nil
}

func (userstore *UserStore) LoginUser(user models.SignInUserRequest) (string, error) {
	//Fetch user from DB
	var jwtString string
	dbUser, err := userstore.UsersStore.GetOne(models.Users{
		Email: user.Email,
	})

	if err != nil {
		return "", errUserNotFound
	}

	//Check the password salted pass == actual pass
	if err = auth.CheckPassword(user.Password, dbUser.Password); err != nil {
		return "", errAuthenticationFailed
	}

	//generate jwt for user
	if dbUser.IsAdmin {
		jwtString, err = jwt.NewToken(dbUser.User_ID, jwt.UserScope)
	} else {
		jwtString, err = jwt.NewToken(dbUser.User_ID, jwt.UserScope)
	}

	if err != nil {
		return "", errInternalServerError
	}

	return jwtString, nil
}

func (userstore *UserStore) UpdateUser(user models.UpdateUserRequest) error {

	//TODO Check user_id or email from JWT
	dbUser, err := userstore.UsersStore.GetOne(models.Users{
		User_ID: user.UserID,
	})

	if err != nil {
		return errUserNotFound
	}

	//TODO  handel empty pass err
	if user.Password != "" {
		if user.Password, err = auth.HashPassword(user.Password); err != nil {
			//TODO Err
			return err
		}
	}

	updatedDbUser := models.Users{
		User_ID:  dbUser.User_ID,
		Email:    user.Email,
		Password: user.Password,
	}

	if err = userstore.UsersStore.Update(updatedDbUser); err != nil {
		return errInternalServerError
	}
	return nil
}

func (userstore *UserStore) GetUserById(id models.GetUserByIDRequest) (*models.Users, error) {
	dbUser, err := userstore.UsersStore.GetOne(models.Users{
		User_ID: id.User_ID,
	})
	if err != nil {
		return nil, errUserNotFound
	}
	return dbUser, nil
}

func (userstore *UserStore) DeleteUser(id models.DeleteUserByIDRequest) error {
	dbUser := models.Users{
		User_ID: id.User_ID,
	}
	if err := userstore.UsersStore.Delete(dbUser); err != nil {
		return errInternalServerError
	}
	return nil
}
