package postgres

import (
	"Avito/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type HousePostgres struct {
	db *sqlx.DB
}

func NewHousePostgres(db *sqlx.DB) *HousePostgres {
	return &HousePostgres{
		db: db,
	}
}

func (r *HousePostgres) Create(house model.House) (model.House, error) {
	query := fmt.Sprintf("INSERT INTO %s(address,year,developer) VALUES($1,$2,$3) RETURNING id,address,year,developer,created_at,update_at", houseTable)
	err := r.db.QueryRow(query, house.Address, house.Year, house.Developer).Scan(&house.Id, &house.Address, &house.Year, &house.Developer, &house.CreatedAt, &house.UpdateAt)
	if err != nil {
		return model.House{}, err
	}
	return house, nil
}

func (r *HousePostgres) GetHouseModerFlatsList(houseId int) ([]model.Flat, error) {
	var flats []model.Flat
	query := fmt.Sprintf("SELECT * FROM %s where house_id=$1", flatTable)
	rows, err := r.db.Query(query, houseId)
	if err != nil {
		return flats, err
	}
	var zagl string
	for rows.Next() {
		var fl model.Flat
		err = rows.Scan(&fl.Id, &fl.HouseId, &fl.Price, &fl.Rooms, &fl.Status, &zagl)
		if err != nil {
			logrus.Print(err)
		}
		flats = append(flats, fl)
	}
	return flats, nil
}

func (r *HousePostgres) GetHouseClientFlatsList(houseId int) ([]model.Flat, error) {
	var flats []model.Flat
	query := fmt.Sprintf("SELECT * FROM %s where house_id=$1 AND status='approved'", flatTable)
	rows, err := r.db.Query(query, houseId)
	if err != nil {
		return flats, err
	}
	var zagl string
	for rows.Next() {
		var fl model.Flat
		err = rows.Scan(&fl.Id, &fl.HouseId, &fl.Price, &fl.Rooms, &fl.Status, &zagl)
		if err != nil {
			logrus.Print(err)
		}
		flats = append(flats, fl)
	}
	fmt.Println(flats)
	return flats, nil
}
