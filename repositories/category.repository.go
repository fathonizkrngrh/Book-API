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

type CategoryRepo interface {
	GetAll(ctx context.Context) ([]data.Category, error)
	GetById(ctx context.Context, id int) (data.Category, error)
	IsExist(ctx context.Context, id int) (bool, error)
	IsDuplicate(ctx context.Context, name string) (bool, error)
	Insert(ctx context.Context, category forms.InsertCategory) (error)
	UpdateByID(ctx context.Context, category forms.UpdateCategory, id int) ( error)
	DeleteByID(ctx context.Context, id int) ( error)
}

type categoryRepo struct {
}

func NewCategoryRepo() *categoryRepo {
	return &categoryRepo{}
}

const (
	categoryTable          = "category"
	layoutDateTime = "2006-01-02 15:04:05"
)

func (repo *categoryRepo) GetAll(ctx context.Context ) ([]data.Category, error) {
	var categories []data.Category
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", categoryTable)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category data.Category
		var createdAt, updatedAt string
		if err = rows.Scan(&category.ID,
			&category.Name,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repo *categoryRepo) GetById(ctx context.Context, id int ) (data.Category, error) {
	
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v WHERE id = %v", categoryTable, id)
	row := db.QueryRowContext(ctx, queryText)

	var category data.Category
	var createdAt, updatedAt string
	if err = row.Scan(&category.ID,
		&category.Name,
		&createdAt,
		&updatedAt); err != nil {
		return data.Category{}, err
	}

	category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
	if err != nil {
		log.Fatal(err)
	}

	return category, nil
}

func (repo *categoryRepo) IsExist(ctx context.Context, id int) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE id = %v", categoryTable, id)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (repo *categoryRepo) IsDuplicate(ctx context.Context, name string) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE name = '%v'", categoryTable, name)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (repo *categoryRepo) Insert(ctx context.Context, category forms.InsertCategory) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	
	queryText := fmt.Sprintf("INSERT IGNORE INTO %v (name, created_at, updated_at) values('%v', NOW(), NOW())", categoryTable,
	  	category.Name)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
	  return err
	}

	return nil
}

func (repo *categoryRepo) UpdateByID(ctx context.Context, category forms.UpdateCategory, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET name = '%v', updated_at = NOW() WHERE id = %v",
		categoryTable,
		category.Name,
		id)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func (repo *categoryRepo) DeleteByID(ctx context.Context, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id = %v", categoryTable, id)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}