package users

import (
	"sipencari-api/app/middlewares"

)

type userUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *userUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}

func (uu *userUsecase) GetByID(id string) Domain {
	return uu.userRepository.GetByID(id)
}
func (uu *userUsecase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *userUsecase) CreateAdmin(userDomain *Domain) Domain {
	return uu.userRepository.CreateAdmin(userDomain)
}

func (uu *userUsecase) Update(id string, userDomain *Domain) Domain {
	return uu.userRepository.Update(id, userDomain)
}

func (uu *userUsecase) Delete(id string) bool {
	return uu.userRepository.Delete(id)
}

func (uu *userUsecase) Login(userDomain *Domain) string {
	user := uu.userRepository.Login(userDomain)
	if user.ID == "" {
		return ""
	}
	token := uu.jwtAuth.GenerateToken(user.ID, user.IsAdmin)
	return token
}
