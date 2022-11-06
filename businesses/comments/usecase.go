package comments

import ()

type commentUsecase struct {
	commentRepository Repository
}

func NewCommentUsecase(comr Repository) Usecase {
	return &commentUsecase{
		commentRepository: comr,
	}
}

func (comr *commentUsecase) GetAll(idPost string) []Domain {
	return comr.commentRepository.GetAll(idPost)
}

func (comr *commentUsecase) GetByID(idPost string,id string) Domain {
	return comr.commentRepository.GetByID(idPost, id)
}

func (comr *commentUsecase) Create(idUser string, idPost string, commentDomain *Domain) Domain {
	return comr.commentRepository.Create(idUser, idPost, commentDomain)
}

func (comr *commentUsecase) Update(idUser string, idPost string, idComment string, commentDomain *Domain) Domain {
	return comr.commentRepository.Update(idUser, idPost, idComment, commentDomain)
}

func (comr *commentUsecase) Delete(idPost, id string) bool {
	return comr.commentRepository.Delete(idPost, id)
}
