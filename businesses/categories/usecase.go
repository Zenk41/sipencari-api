package categories

import ()

type categoryUsecase struct {
	categoryRepository Repository
}

func NewCategoryUsecase(cr Repository) Usecase {
	return &categoryUsecase{
		categoryRepository: cr,
	}
}

func (cr *categoryUsecase) GetAll() []Domain {
	return cr.categoryRepository.GetAll()
}

func (cr *categoryUsecase) GetByID(category_id string) Domain {
	return cr.categoryRepository.GetByID(category_id)
}

func (cr *categoryUsecase) Create(categoryDomain *Domain) Domain {
	return cr.categoryRepository.Create(categoryDomain)
}

func (cr *categoryUsecase) Update(category_id string, categoryDomain *Domain) Domain {
	return cr.categoryRepository.Update(category_id, categoryDomain)
}

func (cr *categoryUsecase) Delete(category_id string) bool {
	return cr.categoryRepository.Delete(category_id)
}