package likesmissing

import ()

type likeMissingUsecase struct {
	likeMissingRepository Repository
}

func NewLikeMissingUsecase(lmr Repository) Usecase {
	return &likeMissingUsecase{
		likeMissingRepository: lmr,
	}
}

func (lmu *likeMissingUsecase) GetAll(idMissing string) []Domain {
	return lmu.likeMissingRepository.GetAll(idMissing)
}

func (lmu *likeMissingUsecase) GetByID(idLike string, idMissing string) Domain {
	return lmu.likeMissingRepository.GetByID(idLike, idMissing)
}

func (lmu *likeMissingUsecase) Like(idUser string, idMissing string, likeDomain *Domain) Domain {
	return lmu.likeMissingRepository.Like(idUser, idMissing, likeDomain)
}

func (lmu *likeMissingUsecase) Unlike(idUser string, idMissing string) bool{
	return lmu.likeMissingRepository.Unlike(idUser, idMissing)
}
