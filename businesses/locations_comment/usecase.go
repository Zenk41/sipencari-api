package locationscomment

import ()

type locationCommentUsecase struct {
	locationCommentRepository Repository
}

func NewLocationCommentUsecase(lccmr Repository) Usecase {
	return &locationCommentUsecase{
		locationCommentRepository: lccmr,
	}
}

func (lccmu *locationCommentUsecase) GetAll() []Domain {
	return lccmu.locationCommentRepository.GetAll()
}

func (lccmu *locationCommentUsecase) Create(idMissing string, LocationMissingDomain *Domain) Domain {
	return lccmu.locationCommentRepository.Create(idMissing, LocationMissingDomain)
}

func (lccmu *locationCommentUsecase) GetByID(idComment, idLocation string) Domain {
	return lccmu.locationCommentRepository.GetByID(idComment, idLocation)
}

func (lccmu *locationCommentUsecase) Update(idComment, idLocation string, LocationCommentDomain *Domain) Domain {
	return lccmu.locationCommentRepository.Update(idComment, idLocation, LocationCommentDomain)
}

func (lccmu *locationCommentUsecase) Delete(idComment, idLocation string)  bool {
	return lccmu.locationCommentRepository.Delete(idComment, idLocation)
}