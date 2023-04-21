package service

import (
	"errors"
	"sesi_8/config"
	"sesi_8/model"
	"sesi_8/repository"
	"sesi_8/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookByID(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{
				ID:     1,
				Title:  "risqi",
				Author: "rizqi",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:     1,
			Title:  "risqi",
			Author: "rizqi",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			repo := repository.NewRepo(config.GORM.DB)
			BookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(BookRepo)
			}

			service := Service{
				repo: repo,
			}

			res, err := service.GetBookById(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			Title:  "risqi",
			Author: "rizqi",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:     1,
				Title:  "risqi",
				Author: "rizqi",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:     1,
			Title:  "Tes",
			Author: "rizqi",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			Title:  "risqi",
			Author: "rizqi",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid title length",
		wantError: true,
		input: model.Book{
			Title:  "ade",
			Author: "rizqi",
		},
		expectedError: errors.New("invalid title length"),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			repo := repository.NewRepo(config.GORM.DB)
			BookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(BookRepo)
			}

			service := Service{
				repo: repo,
			}

			res, err := service.AddBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}
