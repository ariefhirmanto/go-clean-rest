package user

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type Storage interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
	FindAll() ([]User, error)
}

type storage struct {
	ProductDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		ProductDB: db,
	}
}

func (s *storage) Save(data InsertProductRequest) (ProductResponse, error) {
	var resp ProductResponse

	var id int64
	if err := s.ProductDB.QueryRowContext(ctx, addProductQuery,
		data.Name,
		data.Description,
		data.Price,
		data.Rating,
		data.ImageURL,
		pq.Array(&data.AdditionalImageURL),
	).Scan(&id); err != nil {
		log.Println("[ProductModule][AddProduct][Storage] problem querying to db, err: ", err.Error())
		return resp, err
	}

	resp = ProductResponse{
		ID: id,
	}
	return resp, nil
}
