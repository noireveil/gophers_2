package repository

import (
	"errors"
	"gophers_2/internal/entity"
)

type InventoryRepository interface {
	Save(barang entity.Barang) error
	FindAll() []entity.Barang
	FindByID(id int) (*entity.Barang, error)
	Update(barang entity.Barang) error
}

type inventoryRepository struct {
	data   []entity.Barang
	autoID int
}

func NewInventoryRepository() InventoryRepository {
	return &inventoryRepository{
		data:   make([]entity.Barang, 0),
		autoID: 1,
	}
}

func (r *inventoryRepository) Save(barang entity.Barang) error {
	barang.ID = r.autoID
	r.autoID++
	r.data = append(r.data, barang)
	return nil
}

func (r *inventoryRepository) FindAll() []entity.Barang {
	return r.data
}

func (r *inventoryRepository) FindByID(id int) (*entity.Barang, error) {
	for _, b := range r.data {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, errors.New("barang tidak ditemukan")
}

func (r *inventoryRepository) Update(barang entity.Barang) error {
	for i, b := range r.data {
		if b.ID == barang.ID {
			r.data[i] = barang
			return nil
		}
	}
	return errors.New("barang tidak ditemukan")
}
