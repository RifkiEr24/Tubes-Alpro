package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type soal struct {
	pertanyaan string
	jawaban    string
	options    [4]string
}

type quiz struct {
	nama  string
	soals []soal
}

type tugas struct {
	pertanyaan string
	jawaban    string
}

type forum struct {
	pertanyaan string
	jawaban    []string
}

type mataKuliah struct {
	nama    string
	quizzes []quiz
	tugases []tugas
	forums  []forum
}

type user struct {
	nama     string
	username string
	password string
	role     string
	nilai    []mataKuliah
}

type users []user

var arrUser users
var arrMataKuliah []mataKuliah

func main() {
	arrUser = append(arrUser, user{
		nama:     "Rizky",
		username: "rizky",
		password: "rizky",
		role:     "guru",
		nilai:    nil,
	})

	var currentUser user

	login(&arrUser, &currentUser)
}

func login(users *users, currentUser *user) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Login")

	for {
		fmt.Println("Masukkan username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Println("Masukkan password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		for i := 0; i < len(*users); i++ {
			if (*users)[i].username == username && (*users)[i].password == password {
				fmt.Println("Login berhasil")
				*currentUser = (*users)[i]
				menu(currentUser)
				return
			}
		}
		fmt.Println("Login gagal, silahkan coba lagi")
	}
}

func menu(currentUser *user) {
	fmt.Printf("Selamat datang, %s\n", currentUser.nama)
	fmt.Println("Anda adalah seorang " + currentUser.role)

	if currentUser.role == "guru" {
		menuGuru(currentUser)
	} else {
		menuSiswa(currentUser)
	}
}

func menuGuru(currentUser *user) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu Guru:")
		fmt.Println("1. Manajemen Konten")
		fmt.Println("2. Lihat Nilai Mahasiswa")
		fmt.Println("3. Tambah Pengguna")
		fmt.Println("4. Logout")

		fmt.Print("Pilih menu: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		switch choice {
		case 1:
			manajemenKonten(reader)
		case 2:
			lihatNilaiMahasiswa()
		case 3:
			tambahPengguna()
		case 4:
			fmt.Println("Logout berhasil")
			login(&arrUser, currentUser)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuSiswa(currentUser *user) {
	fmt.Println("Fitur menu siswa belum diimplementasikan.")

}

func manajemenKonten(reader *bufio.Reader) {
	for {
		fmt.Println("\nDaftar Mata Kuliah:")
		for i, mk := range arrMataKuliah {
			fmt.Printf("%d. %s\n", i+1, mk.nama)
		}
		fmt.Println("0. Tambah Mata Kuliah Baru")
		fmt.Print("Pilih nomor mata kuliah (atau 0 untuk tambah baru): ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		if choice == 0 {
			tambahMataKuliah(reader)
		} else if choice > 0 && choice <= len(arrMataKuliah) {
			manageMataKuliah(&arrMataKuliah[choice-1], reader)
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahMataKuliah(reader *bufio.Reader) {
	var mk mataKuliah
	fmt.Print("Masukkan nama mata kuliah: ")
	mk.nama, _ = reader.ReadString('\n')
	mk.nama = strings.TrimSpace(mk.nama)
	arrMataKuliah = append(arrMataKuliah, mk)
	fmt.Println("Mata kuliah berhasil ditambahkan.")
}

func manageMataKuliah(mk *mataKuliah, reader *bufio.Reader) {
	for {
		fmt.Printf("\nManajemen Konten untuk Mata Kuliah: %s\n", mk.nama)
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Edit Konten")
		fmt.Println("3. Hapus Konten")
		fmt.Println("4. Kembali")

		fmt.Print("Pilih menu: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		switch choice {
		case 1:
			tambahKonten(mk, reader)
		case 2:
			editKonten(mk, reader)
		case 3:
			hapusKonten(mk, reader)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahKonten(mk *mataKuliah, reader *bufio.Reader) {
	for {
		fmt.Println("\nPilih jenis konten yang akan ditambah:")
		fmt.Println("1. Tugas")
		fmt.Println("2. Quiz")
		fmt.Println("3. Forum")
		fmt.Println("4. Selesai")

		fmt.Print("Pilih menu: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		switch choice {
		case 1:
			var t tugas
			fmt.Print("Masukkan pertanyaan tugas: ")
			t.pertanyaan, _ = reader.ReadString('\n')
			t.pertanyaan = strings.TrimSpace(t.pertanyaan)
			mk.tugases = append(mk.tugases, t)
		case 2:
			var q quiz
			fmt.Print("Masukkan nama quiz: ")
			q.nama, _ = reader.ReadString('\n')
			q.nama = strings.TrimSpace(q.nama)
			for {
				var s soal
				fmt.Print("Masukkan pertanyaan quiz: ")
				s.pertanyaan, _ = reader.ReadString('\n')
				s.pertanyaan = strings.TrimSpace(s.pertanyaan)
				fmt.Print("Masukkan pilihan A: ")
				s.options[0], _ = reader.ReadString('\n')
				s.options[0] = strings.TrimSpace(s.options[0])
				fmt.Print("Masukkan pilihan B: ")
				s.options[1], _ = reader.ReadString('\n')
				s.options[1] = strings.TrimSpace(s.options[1])
				fmt.Print("Masukkan pilihan C: ")
				s.options[2], _ = reader.ReadString('\n')
				s.options[2] = strings.TrimSpace(s.options[2])
				fmt.Print("Masukkan pilihan D: ")
				s.options[3], _ = reader.ReadString('\n')
				s.options[3] = strings.TrimSpace(s.options[3])
				fmt.Print("Masukkan jawaban benar (A/B/C/D): ")
				s.jawaban, _ = reader.ReadString('\n')
				s.jawaban = strings.ToUpper(strings.TrimSpace(s.jawaban))
				q.soals = append(q.soals, s)

				fmt.Print("Tambah pertanyaan lagi? (y/n): ")
				tambahLagi, _ := reader.ReadString('\n')
				if strings.ToLower(strings.TrimSpace(tambahLagi)) != "y" {
					break
				}
			}
			mk.quizzes = append(mk.quizzes, q)
		case 3:
			var f forum
			fmt.Print("Masukkan pertanyaan forum: ")
			f.pertanyaan, _ = reader.ReadString('\n')
			f.pertanyaan = strings.TrimSpace(f.pertanyaan)
			mk.forums = append(mk.forums, f)
		case 4:
			fmt.Println("Konten berhasil ditambahkan.")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func editKonten(mk *mataKuliah, reader *bufio.Reader) {
	for {
		fmt.Println("\nPilih jenis konten yang akan diedit:")
		fmt.Println("1. Edit Tugas")
		fmt.Println("2. Edit Quiz")
		fmt.Println("3. Edit Forum")
		fmt.Println("4. Selesai")

		fmt.Print("Pilih menu: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid")
			continue
		}

		switch choice {
		case 1:
			editTugas(mk, reader)
		case 2:
			editQuiz(mk, reader)
		case 3:
			editForum(mk, reader)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func editTugas(mk *mataKuliah, reader *bufio.Reader) {
	if len(mk.tugases) == 0 {
		fmt.Println("Belum ada tugas yang dapat diedit.")
		return
	}

	fmt.Println("\nDaftar Tugas:")
	for i, t := range mk.tugases {
		fmt.Printf("%d. %s\n", i+1, t.pertanyaan)
	}

	fmt.Print("Pilih nomor tugas yang akan diedit: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(mk.tugases) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Printf("Masukkan pertanyaan tugas baru (kosongkan jika tidak ingin mengubah): ")
	newPertanyaan, _ := reader.ReadString('\n')
	newPertanyaan = strings.TrimSpace(newPertanyaan)

	if newPertanyaan != "" {
		mk.tugases[choice-1].pertanyaan = newPertanyaan
		fmt.Println("Pertanyaan tugas berhasil diubah.")
	} else {
		fmt.Println("Pertanyaan tugas tidak diubah.")
	}
}

func editQuiz(mk *mataKuliah, reader *bufio.Reader) {
	if len(mk.quizzes) == 0 {
		fmt.Println("Belum ada quiz yang dapat diedit.")
		return
	}

	fmt.Println("\nDaftar Quiz:")
	for i, q := range mk.quizzes {
		fmt.Printf("%d. %s\n", i+1, q.nama)
	}

	fmt.Print("Pilih nomor quiz yang akan diedit: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(mk.quizzes) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Printf("Masukkan nama quiz baru (kosongkan jika tidak ingin mengubah): ")
	newNama, _ := reader.ReadString('\n')
	newNama = strings.TrimSpace(newNama)

	if newNama != "" {
		mk.quizzes[choice-1].nama = newNama
		fmt.Println("Nama quiz berhasil diubah.")
	} else {
		fmt.Println("Nama quiz tidak diubah.")
	}

	for j := 0; j < len(mk.quizzes[choice-1].soals); j++ {
		var s soal
		fmt.Printf("Masukkan pertanyaan quiz %d: ", j+1)
		s.pertanyaan, _ = reader.ReadString('\n')
		s.pertanyaan = strings.TrimSpace(s.pertanyaan)

		fmt.Printf("Masukkan pilihan A: ")
		s.options[0], _ = reader.ReadString('\n')
		s.options[0] = strings.TrimSpace(s.options[0])

		fmt.Printf("Masukkan pilihan B: ")
		s.options[1], _ = reader.ReadString('\n')
		s.options[1] = strings.TrimSpace(s.options[1])

		fmt.Printf("Masukkan pilihan C: ")
		s.options[2], _ = reader.ReadString('\n')
		s.options[2] = strings.TrimSpace(s.options[2])

		fmt.Printf("Masukkan pilihan D: ")
		s.options[3], _ = reader.ReadString('\n')
		s.options[3] = strings.TrimSpace(s.options[3])

		fmt.Printf("Masukkan jawaban benar (A/B/C/D): ")
		jawaban, _ := reader.ReadString('\n')
		s.jawaban = strings.ToUpper(strings.TrimSpace(jawaban))

		mk.quizzes[choice-1].soals[j] = s
		fmt.Printf("Pertanyaan quiz %d berhasil diubah.\n", j+1)

		fmt.Println("Apakah Anda ingin mengubah pertanyaan quiz lainnya? (y/n)")
		tambahLagi, _ := reader.ReadString('\n')
		tambahLagi = strings.TrimSpace(tambahLagi)
		if tambahLagi != "y" {
			break
		}
	}
	fmt.Println("Quiz berhasil diubah.")
}

func editForum(mk *mataKuliah, reader *bufio.Reader) {
	if len(mk.forums) == 0 {
		fmt.Println("Belum ada forum yang dapat diedit.")
		return
	}

	fmt.Println("\nDaftar Forum:")
	for i, f := range mk.forums {
		fmt.Printf("%d. %s\n", i+1, f.pertanyaan)
	}

	fmt.Print("Pilih nomor forum yang akan diedit: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(mk.forums) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Printf("Masukkan pertanyaan forum baru (kosongkan jika tidak ingin mengubah): ")
	newPertanyaan, _ := reader.ReadString('\n')
	newPertanyaan = strings.TrimSpace(newPertanyaan)

	if newPertanyaan != "" {
		mk.forums[choice-1].pertanyaan = newPertanyaan
		fmt.Println("Pertanyaan forum berhasil diubah.")
	} else {
		fmt.Println("Pertanyaan forum tidak diubah.")
	}

	fmt.Println("Forum berhasil diubah.")
}

func hapusKonten(mk *mataKuliah, reader *bufio.Reader) {
	fmt.Println("\nPilih jenis konten yang akan dihapus:")
	fmt.Println("1. Tugas")
	fmt.Println("2. Quiz")
	fmt.Println("3. Forum")

	var choice int
	fmt.Print("Pilih jenis konten (1-3): ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > 3 {
		fmt.Println("Pilihan tidak valid")
		return
	}

	switch choice {
	case 1:
		hapusTugas(mk, reader)
	case 2:
		hapusQuiz(mk, reader)
	case 3:
		hapusForum(mk, reader)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func hapusTugas(mk *mataKuliah, reader *bufio.Reader) {
	if len(mk.tugases) == 0 {
		fmt.Println("Belum ada tugas yang dapat dihapus.")
		return
	}

	fmt.Println("\nDaftar Tugas:")
	for i, t := range mk.tugases {
		fmt.Printf("%d. %s\n", i+1, t.pertanyaan)
	}

	fmt.Print("Pilih nomor tugas yang akan dihapus: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(mk.tugases) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	mk.tugases = append(mk.tugases[:choice-1], mk.tugases[choice:]...)
	fmt.Println("Tugas berhasil dihapus.")
}

func hapusQuiz(mk *mataKuliah, reader *bufio.Reader) {
	if len(mk.quizzes) == 0 {
		fmt.Println("Belum ada quiz yang dapat dihapus.")
		return
	}

	fmt.Println("\nDaftar Quiz:")
	for i, q := range mk.quizzes {
		fmt.Printf("%d. %s\n", i+1, q.nama)
	}

	fmt.Print("Pilih nomor quiz yang akan dihapus: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(mk.quizzes) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	// Hapus quiz dari slice menggunakan teknik slice-trick
	mk.quizzes = append(mk.quizzes[:choice-1], mk.quizzes[choice:]...)
	fmt.Println("Quiz berhasil dihapus.")
}

func hapusForum(mk *mataKuliah, reader *bufio.Reader) {
	if len(mk.forums) == 0 {
		fmt.Println("Belum ada forum yang dapat dihapus.")
		return
	}

	fmt.Println("\nDaftar Forum:")
	for i, f := range mk.forums {
		fmt.Printf("%d. %s\n", i+1, f.pertanyaan)
	}

	fmt.Print("Pilih nomor forum yang akan dihapus: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(mk.forums) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	// Hapus forum dari slice menggunakan teknik slice-trick
	mk.forums = append(mk.forums[:choice-1], mk.forums[choice:]...)
	fmt.Println("Forum berhasil dihapus.")
}

func lihatNilaiMahasiswa() {
	fmt.Println("Fitur lihat nilai mahasiswa belum diimplementasikan.")
}

func tambahPengguna() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nTambah Pengguna Baru:")

	var newUser user

	// Masukkan informasi pengguna baru
	fmt.Print("Nama: ")
	nama, _ := reader.ReadString('\n')
	newUser.nama = strings.TrimSpace(nama)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	newUser.username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	newUser.password = strings.TrimSpace(password)

	fmt.Print("Role (guru/siswa): ")
	role, _ := reader.ReadString('\n')
	newUser.role = strings.TrimSpace(role)

	// Validasi role
	if newUser.role != "guru" && newUser.role != "siswa" {
		fmt.Println("Role tidak valid. Pengguna hanya bisa memiliki role 'guru' atau 'siswa'.")
		return
	}

	// Tambahkan pengguna baru ke dalam slice arrUser
	arrUser = append(arrUser, newUser)

	fmt.Printf("Pengguna dengan username '%s' berhasil ditambahkan.\n", newUser.username)
}
