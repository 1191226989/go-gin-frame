package article

type Service interface {
	Create(articleData *CreateArticleData) (uint, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
