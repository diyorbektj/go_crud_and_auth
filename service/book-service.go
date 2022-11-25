package service

import "test/dto"

type BookService interface {
	InsertBook(book dto.BookCreateDTO)
}
