package service

import "sesi_8/model"

type BookService interface {
	GetAllBook() ([]model.Book, error)
	GetBookById(id int64) (model.Book, error)
	AddBook(in model.Book) (model.Book, error)
	UpdateBook(id int64, in model.Book) (model.Book, error)
	DeleteBook(id int64) error
}

func (s *Service) GetAllBook() ([]model.Book, error) {
	return s.repo.GetAllBook()
}

func (s *Service) GetBookById(id int64) (model.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *Service) AddBook(in model.Book) (model.Book, error) {
	return s.repo.AddBook(in)
}

func (s *Service) UpdateBook(id int64, in model.Book) (model.Book, error) {
	isAvail, err := s.repo.CheckBookById(id)
	if !isAvail {
		return model.Book{}, err
	}
	return s.repo.UpdateBook(id, in)
}
func (s *Service) DeleteBook(id int64) error {
	isAvail, err := s.repo.CheckBookById(id)
	if !isAvail {
		return err
	}
	return s.repo.DeleteBook(id)
}
