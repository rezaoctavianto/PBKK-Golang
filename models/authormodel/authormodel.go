package authormodel

import (
	"Authors/config"
	"Authors/entities"
)

func GetAll() []entities.Author {
	rows, err := config.DB.Query(`SELECT * FROM authors`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var authors []entities.Author

	for rows.Next() {
		var author entities.Author
		if err := rows.Scan(&author.Id, &author.Name, &author.DoB, &author.Updated_At); err != nil {
			panic(err)
		}

		authors = append(authors, author)
	}
	return authors
}
