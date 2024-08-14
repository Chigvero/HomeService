package postgres

import (
	"Avito/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FlatPostgres struct {
	db *sqlx.DB
}

func NewFlatPostgres(db *sqlx.DB) *FlatPostgres {
	return &FlatPostgres{
		db: db,
	}
}

func (r *FlatPostgres) Create(flat model.Flat) (model.Flat, error) {
	flatQuery := fmt.Sprintf("INSERT INTO %s(house_id,price,rooms) VALUES($1,$2,$3) RETURNING id,status", flatTable)
	tx, err := r.db.Begin()
	if err != nil {
		return model.Flat{}, err
	}
	err = tx.QueryRow(flatQuery, flat.HouseId, flat.Price, flat.Rooms).Scan(&flat.Id, &flat.Status)
	if err != nil {
		tx.Rollback()
		return model.Flat{}, err
	}
	houseQuery := fmt.Sprintf("UPDATE %s SET update_at=CURRENT_TIMESTAMP WHERE id = $1", houseTable)
	_, err = tx.Exec(houseQuery, flat.HouseId)
	if err != nil {
		tx.Rollback()
		return model.Flat{}, err
	}
	err = tx.Commit()
	if err != nil {
		return model.Flat{}, err
	}
	return flat, nil
}
func (r *FlatPostgres) Update(id int, status string, user_id uuid.UUID) (model.Flat, error) {
	query := fmt.Sprintf("UPDATE %s SET status=$1, moderator_id=$2 WHERE id=$3 ", flatTable)
	_, err := r.db.Exec(query, status, user_id, id)
	if err != nil {
		return model.Flat{}, err
	}
	return r.GetById(id)
}

func (r *FlatPostgres) GetById(id int) (model.Flat, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", flatTable)
	var f model.Flat
	err := r.db.QueryRow(query, id).Scan(&f.Id, &f.HouseId, &f.Price, &f.Rooms, &f.Status, &f.ModeratorId)
	if err != nil {
		return model.Flat{}, err
	}
	return f, nil
}
