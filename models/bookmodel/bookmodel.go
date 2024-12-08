package bookmodel

import (
	"Authors/config"
	"Authors/entities"
)

func GetAll() []entities.Book {
	rows, err := config.DB.Query(`
		SELECT 
			books.id,
			books.title,
			authors.name as author_name,
			books.genre
		FROM books
		JOIN authors ON books.author_id = authors.id
		`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var books []entities.Book

	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author.Name,
			&book.Genre,
		); err != nil {
			panic(err)
		}

		books = append(books, book)
	}
	return books
}

func Create(book entities.Book) bool {
	result, err := config.DB.Exec(`
	INSERT INTO books (title, authors, genre)
	VALUE (?, ?, ?)`,
		book.Title, book.Author, book.Genre,
	)
	if err != nil {
		panic(err)
	}
	LastInsertedId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return LastInsertedId > 0
}

func Detail(id int) entities.Book {
	row := config.DB.QueryRow(`SELECT id, title, authors, genre FROM books WHERE id = ?`, id)

	var book entities.Book
	if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Genre); err != nil {
		panic(err.Error())
	}
	return book
}

func Update(id int, book entities.Book) bool {
	query, err := config.DB.Exec(`UPDATE books SET title = ?, authors = ?, genre = ? WHERE id = ?`, book.Title, book.Author, book.Genre, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM books WHERE id = ?`, id)
	return err
}
