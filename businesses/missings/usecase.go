package missings

import ()

type missingUsecase struct {
	missingRepository Repository
}

func NewMissingUsecase(mr Repository) Usecase {
	return &missingUsecase{
		missingRepository: mr,
	}
}

func (mr *missingUsecase) GetAll() []Domain {
	return mr.missingRepository.GetAll()
}

func (mr *missingUsecase) GetByID(id string) Domain {
	return mr.missingRepository.GetByID(id)
}

func (mr *missingUsecase) Create(idUser string, missingDomain *Domain) Domain {
	return mr.missingRepository.Create(idUser, missingDomain)
}

func (mr *missingUsecase) Update(idUser string, id string, missingDomain *Domain) Domain {
	return mr.missingRepository.Update(idUser, id, missingDomain)
}

func (mr *missingUsecase) Delete(id string) bool {
	return mr.missingRepository.Delete(id)
}