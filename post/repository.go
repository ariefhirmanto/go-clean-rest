package post

import (
	"database/sql"
	"log"
	"project-go/models"
)

type Storage interface {
	AddTicket(input models.InputPostRequest) (resp models.PostResponse, err error)
	FindByID(ID int64) (resp models.PostResponse, err error)
	Update(ID int64) (resp models.PostResponse, err error)
}

type storage struct {
	TicketDB *sql.DB
}

func newStorage(db *sql.DB) *storage {
	return &storage{
		TicketDB: db,
	}
}

func (s *storage) AddTicket(input models.InputPostRequest) (resp models.PostResponse, err error) {
	var id int64
	err = s.TicketDB.QueryRow(addTicketQuery,
		input.Title,
		input.Content,
		input.ImageURL,
		input.Category,
	).Scan(&id)
	if err != nil {
		log.Println("[Ticket][AddTicket][Storage] Problem to querying to db, err: ", err.Error())
		return resp, err
	}

	resp = models.PostResponse{
		ID: id,
	}
	return resp, nil
}

func (s *storage) FindByID(ID int64) (resp models.PostResponse, err error) {
	err = s.TicketDB.QueryRow(getTicketQuery, ID).Scan(
		&resp.Title,
		&resp.Content,
		&resp.ImageURL,
		&resp.Category,
	)
	if err != nil {
		log.Println("[Ticket][FindByID][Storage] Problem to querying to db, err: ", err.Error())
		return resp, err
	}

	return resp, nil
}

func (s *storage) Update(ID int64) (resp models.PostResponse, err error) {

}
