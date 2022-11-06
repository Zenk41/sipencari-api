package locationsmissing

import ()

type locationMissingUsecase struct {
	locationMissingRepository Repository
}

func NewLocationMissingUsecase(lcmr Repository) Usecase {
	return &locationMissingUsecase{
		locationMissingRepository: lcmr,
	}
}

func (lcmu *locationMissingUsecase) GetAll() []Domain {
	return lcmu.locationMissingRepository.GetAll()
}

func (lcmu *locationMissingUsecase) GetByID(idMissing, idLocation string) Domain {
	return lcmu.locationMissingRepository.GetByID(idMissing, idLocation)
}

func (lcmu *locationMissingUsecase) Create(idMissing string, LocationMissingDomain *Domain) Domain {
	return lcmu.locationMissingRepository.Create(idMissing, LocationMissingDomain)
}

func (lcmu *locationMissingUsecase) Update(idMissing, idLocation string, LocationMissingDomain *Domain) Domain {
	return lcmu.locationMissingRepository.Update(idMissing, idLocation, LocationMissingDomain)
}

func (lcmu *locationMissingUsecase) Delete(idMissing, idLocation string) bool {
	return lcmu.locationMissingRepository.Delete(idMissing, idLocation)
}
