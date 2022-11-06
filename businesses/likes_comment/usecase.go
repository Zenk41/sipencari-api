package likescomment

import ()

type likeCommentUsecase struct {
	likeCommentRepository Repository
}

func NewLikeCommentUsecase(lcr Repository) Usecase {
	return &likeCommentUsecase{
		likeCommentRepository: lcr,
	}
}

func (lmu *likeCommentUsecase) GetAll() []Domain {
	return lmu.likeCommentRepository.GetAll()
}

func (lmu *likeCommentUsecase) GetByID(idLike string, idMissing string, idComment int) Domain {
	return lmu.likeCommentRepository.GetByID(idLike, idMissing, idComment)
}

func (lmu *likeCommentUsecase) Like(idUser string, idMissing string,  idComment int, likeDomain *Domain) Domain {
	return lmu.likeCommentRepository.Like(idUser, idMissing, idComment, likeDomain)
}

func (lmu *likeCommentUsecase) Unlike(idUser string, idMissing string,  idComment int) bool{
	return lmu.likeCommentRepository.Unlike(idUser, idMissing, idComment)
}
