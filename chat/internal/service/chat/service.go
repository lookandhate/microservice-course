package chat

type Service struct {
	repo interface{}
}

func NewService(repo interface{}) *Service {
	return &Service{repo: repo}
}
