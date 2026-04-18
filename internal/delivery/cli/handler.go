package cli

import (
	"bufio"
	"fmt"
	"gophers_2/internal/usecase"
	"os"
	"strconv"
	"strings"
)

type Handler struct {
	usecase usecase.InventoryUsecase
}

func NewHandler(u usecase.InventoryUsecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== TOKO KELONTONG INVENTORY MANAGER ===")
		fmt.Println("1. Tambah Barang ke Gudang")
		fmt.Println("2. Lihat Semua Stok Barang")
		fmt.Println("3. Beli Barang")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih Menu (1-4): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			h.tambahBarang(reader)
		case "2":
			h.lihatBarang()
		case "3":
			h.beliBarang(reader)
		case "4":
			fmt.Println("[SISTEM]: Program selesai. Selamat beristirahat, Juragan!")
			return
		default:
			fmt.Println("[SISTEM]: Pilihan tidak valid.")
		}
	}
}

func (h *Handler) tambahBarang(reader *bufio.Reader) {
	fmt.Print("Masukkan Nama Barang: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan Harga: ")
	hargaStr, _ := reader.ReadString('\n')
	harga, _ := strconv.Atoi(strings.TrimSpace(hargaStr))

	fmt.Print("Masukkan Stok Awal: ")
	stokStr, _ := reader.ReadString('\n')
	stok, _ := strconv.Atoi(strings.TrimSpace(stokStr))

	h.usecase.TambahBarang(nama, harga, stok)
	fmt.Println("[SISTEM]: Barang berhasil ditambahkan ke gudang!")
}

func (h *Handler) lihatBarang() {
	barang := h.usecase.LihatSemua()
	fmt.Println("\n=== DAFTAR STOK GUDANG ===")
	for _, b := range barang {
		fmt.Printf("ID: %d | Nama: %s | Harga: Rp %d | Stok: %d pcs\n", b.ID, b.Nama, b.Harga, b.Stok)
	}
	fmt.Printf("Total Jenis Barang: %d\n", len(barang))
}

func (h *Handler) beliBarang(reader *bufio.Reader) {
	fmt.Println("\n=== MENU PEMBELIAN ===")
	barangList := h.usecase.LihatSemua()
	for _, b := range barangList {
		fmt.Printf("ID: %d | Nama: %s | Stok: %d | Harga: Rp %d\n", b.ID, b.Nama, b.Stok, b.Harga)
	}

	fmt.Print("\nPilih ID Barang yang mau dibeli: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Jumlah yang mau dibeli: ")
	jumlahStr, _ := reader.ReadString('\n')
	jumlah, _ := strconv.Atoi(strings.TrimSpace(jumlahStr))

	b, err := h.usecase.GetBarangByID(id)
	if err != nil {
		fmt.Println("[SISTEM]:", err.Error())
		return
	}

	total := b.Harga * jumlah
	fmt.Printf("[SISTEM]: Total Harga: Rp %d\n", total)

	fmt.Print("Masukkan Uang Anda: ")
	uangStr, _ := reader.ReadString('\n')
	uang, _ := strconv.Atoi(strings.TrimSpace(uangStr))

	kembalian, err := h.usecase.BeliBarang(id, jumlah, uang)
	if err != nil {
		fmt.Printf("[SISTEM]: Transaksi Gagal - %s\n", err.Error())
		return
	}

	bUpdated, _ := h.usecase.GetBarangByID(id)

	fmt.Println("[SISTEM]: Transaksi Berhasil!")
	fmt.Printf("[SISTEM]: Kembalian Anda: Rp %d\n", kembalian)
	fmt.Printf("[SISTEM]: Stok %s sekarang: %d pcs\n", bUpdated.Nama, bUpdated.Stok)
}
