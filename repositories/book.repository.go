package repositories

import (
	"context"
	"fmt"
	"log"
	"quiz3/config"
	"quiz3/data"
	"quiz3/forms"
	"time"
)

type BookRepo interface {
	GetAll(ctx context.Context) ([]data.Book, error)
	GetById(ctx context.Context, id int) (data.Book, error)
	IsExist(ctx context.Context, id int) (bool, error)
	IsDuplicate(ctx context.Context, name string) (bool, error)
	Insert(ctx context.Context, book forms.InsertBook) (error)
	UpdateByID(ctx context.Context, book forms.UpdateBook, id int) ( error)
	DeleteByID(ctx context.Context, id int) ( error)
}

type bookRepo struct {
}

func NewBookRepo() *bookRepo {
	return &bookRepo{}
}

const (
	bookTable          = "book"
)

func (repo *bookRepo) GetAll(ctx context.Context ) ([]data.Book, error) {
	var books []data.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", bookTable)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book data.Book
		var createdAt, updatedAt string
		if err = rows.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	return books, nil
}

func (repo *bookRepo) GetById(ctx context.Context, id int ) (data.Book, error) {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v WHERE id = %v", bookTable, id)
	row := db.QueryRowContext(ctx, queryText)

	var book data.Book
	var createdAt, updatedAt string
	if err = row.Scan(&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
		&createdAt,
		&updatedAt); err != nil {
		return data.Book{}, err
	}

	book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return book, nil
}

func (repo *bookRepo) IsExist(ctx context.Context, id int) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE id = %v", bookTable, id)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (repo *bookRepo) IsDuplicate(ctx context.Context, title string) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE title = '%v'", bookTable, title)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (repo *bookRepo) Insert(ctx context.Context, book forms.InsertBook) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	
	queryText := fmt.Sprintf("INSERT IGNORE INTO '%v' (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) VALUES ('%v', '%v', '%v', %v, '%v', %v, '%v', NOW(), NOW(), %v)",
		bookTable,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		)
	_, err = db.ExecContext(ctx, queryText,
		)

	if err != nil {
	  return err
	}

	return nil
}

func (repo *bookRepo) UpdateByID(ctx context.Context, book forms.UpdateBook, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE '%v' SET  title = '%v', description = '%v', image_url = '%v', release_year = %v, price '%v', total_page = %v, thickness = '%v', updated_at = NOW(), category_id = %v WHERE id = %v",
		bookTable,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		id)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func (repo *bookRepo) DeleteByID(ctx context.Context, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id = %v", bookTable, id)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}