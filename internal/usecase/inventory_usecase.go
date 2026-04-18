package usecase

import (
	"errors"
	"gophers_2/internal/entity"
	"gophers_2/internal/repository"
)

type InventoryUsecase interface {
	TambahBarang(nama string, harga, stok int) error
	LihatSemua() []entity.Barang
	GetBarangByID(id int) (*entity.Barang, error)
	BeliBarang(id, jumlah, uang int) (int, error)
}

type inventoryUsecase struct {
	repo repository.InventoryRepository
}

func NewInventoryUsecase(repo repository.InventoryRepository) InventoryUsecase {
	return &inventoryUsecase{repo: repo}
}

func (u *inventoryUsecase) TambahBarang(nama string, harga, stok int) error {
	barang := entity.Barang{
		Nama:  nama,
		Harga: harga,
		Stok:  stok,
	}
	return u.repo.Save(barang)
}

func (u *inventoryUsecase) LihatSemua() []entity.Barang {
	return u.repo.FindAll()
}

func (u *inventoryUsecase) GetBarangByID(id int) (*entity.Barang, error) {
	return u.repo.FindByID(id)
}

func (u *inventoryUsecase) BeliBarang(id, jumlah, uang int) (int, error) {
	barang, err := u.repo.FindByID(id)
	if err != nil {
		return 0, err
	}

	if barang.Stok < jumlah {
		return 0, errors.New("stok tidak mencukupi")
	}

	total := barang.Harga * jumlah
	if uang < total {
		return 0, errors.New("uang tidak mencukupi")
	}

	barang.Stok -= jumlah
	u.repo.Update(*barang)

	kembalian := uang - total
	return kembalian, nil
}
