package hashtags

import ()

type hashtagUsecase struct {
	hashtagRepository Repository
}

func NewHashtagUsecase(hr Repository) Usecase {
	return &hashtagUsecase{
		hashtagRepository: hr,
	}
}

func (hr *hashtagUsecase) GetAll() []Domain {
	return hr.GetAll()
}
func (hr *hashtagUsecase) GetByID(id_hashtag string) Domain {
	return hr.GetByID(id_hashtag)
}
