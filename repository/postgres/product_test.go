package postgres

import (
	"context"
	"publisher/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO products (name, price, rating, image_url) VALUES ($1, $2, $3, $4) RETURNING id`)).
		WithArgs("baju", 10000, 1, "http://google.com").
		WillReturnRows(sqlmock.NewRows([]string{"col1"}).AddRow(1))

	repo := NewProductRepository(db)

	// now we execute our method
	_, err = repo.Store(context.Background(), models.Product{
		Name:     "baju",
		Price:    10000,
		Rating:   1,
		ImageURL: "http://google.com",
	})

	assert.NoError(t, err)
}

func Test_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`DELETE FROM products WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"col1"}).AddRow(1))

	repo := NewProductRepository(db)

	// now we execute our method
	err = repo.Delete(context.Background(), 1)

	assert.NoError(t, err)
}

func Test_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE products SET (name=$1, price=$2, image_url=$3, rating=$4) WHERE id = $5`)).
		WithArgs("baju", 10000, "http://google.com", 1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"col1"}).AddRow(1))

	repo := NewProductRepository(db)

	// now we execute our method
	err = repo.Update(context.Background(), models.Product{
		ID:       1,
		Name:     "baju",
		Price:    10000,
		Rating:   1,
		ImageURL: "http://google.com",
	})

	assert.NoError(t, err)
}

func Test_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// row := []interface{}{1, "baju", 10000, 1, "google.com"}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, price, rating, image_url FROM products WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "rating", "image_url"}).AddRow(1, "", 1, 1, ""))

	repo := NewProductRepository(db)

	// now we execute our method
	result, err := repo.Get(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
