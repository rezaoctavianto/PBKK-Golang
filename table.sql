
DELETE TABLE authors;
DROP TABLE IF EXISTS authors;
CREATE TABLE authors (
  id INT AUTO_INCREMENT NOT NULL,
  name VARCHAR(128) NOT NULL,
  date_of_birth DATE NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

INSERT INTO authors (name, date_of_birth) VALUES
('Harper Lee', '1926-04-28'),
('George Orwell', '1903-06-25'),
('F. Scott Fitzgerald', '1896-09-24'),
('Jane Austen', '1775-12-16'),
('J.D. Salinger', '1919-01-01');

DROP TABLE IF EXISTS books;
CREATE TABLE books (
  id INT AUTO_INCREMENT NOT NULL,
  title VARCHAR(128) NOT NULL,
  author_id INT NOT NULL,
  genre VARCHAR(128),
  added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);

INSERT INTO books (title, author_id, genre) VALUES
('To Kill a Mockingbird', 1, 'Fiction'),
('1984', 2, 'Dystopian'),
('The Great Gatsby', 3, 'Classic'),
('Pride and Prejudice', 4, 'Romance'),
('The Catcher in the Rye', 5, 'Literary Fiction');

