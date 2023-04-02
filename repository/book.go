package repository

import (
	"sesi_8/model"
)

// clean architectures -> handler->service->repo

// interface employee
type BookRepo interface {
	AddBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookByID(id int64) (res model.Book, err error)
	CheckBookById(id int64) (bool, error)
	UpdateBook(id int64, in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) AddBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookByID(id int64) (res model.Book, err error) {
	err = r.gorm.First(&res, id).Error
	if err != nil {

		return res, err
	}

	return res, nil
}

func (r Repo) CheckBookById(id int64) (isAvail bool, err error) {
	err = r.gorm.First(&model.Book{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r Repo) UpdateBook(id int64, in model.Book) (res model.Book, err error) {

	err = r.gorm.Model(&res).Where("id = ?", id).Updates(model.Book{
		Title:  in.Title,
		Author: in.Author,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {

	err = r.gorm.Debug().Model(&model.Book{}).Delete(&model.Book{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
