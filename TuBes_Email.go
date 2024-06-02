package main

import "fmt"

const NMAX int = 100

type temp struct {
	addtemp, passtemp string
	acc               int
}

type register struct {
	addemail, passemail string
	mail                [NMAX]pesan
	acc, jpesan         int
}

type pesan struct {
	pmasuk, apesan string
	npesan         int
}

type tabtemp [NMAX]temp

type tabreg [NMAX]register

func main() {
	var email tabreg
	var emailtemp tabtemp
	var siapa, pilih string
	var n, nakun int
	menuutama()
	fmt.Print("Masuk Sebagai: ")
	fmt.Scan(&siapa)
	for siapa != "Exit" {
		if siapa == "Admin" {
			menuadmin()
			loginadmin(&emailtemp, &email, &n, &nakun)
		} else if siapa == "Pengguna" {
			menureglog()
			fmt.Scan(&pilih)
			for pilih != "Exit" {
				if pilih == "Register" {
					menuregisterpengguna()
					registerpengguna(&emailtemp, &n)
				} else if pilih == "Login" {
					menuloginpengguna()
					loginpengguna(&email, nakun)
				}
				menureglog()
				fmt.Scan(&pilih)
			}
		}
		menuutama()
		fmt.Print("Masuk Sebagai: ")
		fmt.Scan(&siapa)
	}
}

func menuutama() {
	fmt.Println("-------------------------")
	fmt.Println("          EMAIL          ")
	fmt.Println("-------------------------")
	fmt.Println("          Admin          ")
	fmt.Println("         Pengguna        ")
	fmt.Println("           Exit          ")
	fmt.Println("-------------------------")
}

func menuadmin() {
	fmt.Println("-------------------------")
	fmt.Println("          LOGIN          ")
	fmt.Println("-------------------------")
	fmt.Print("Masukkan Username dan Password: ")
}

func loginadmin(Mtemp *tabtemp, M *tabreg, n, nakun *int) {
	var user, pass string
	fmt.Scan(&user, &pass)
	if user == "admin" && pass == "admin" {
		menuacc(&*Mtemp, &*M, &*n, &*nakun)
	}
}

func menuacc(Mtemp *tabtemp, M *tabreg, n, nakun *int) {
	fmt.Println("Login Berhasil!")
	fmt.Println("-------------------------")
	fmt.Println("          ADMIN          ")
	fmt.Println("-------------------------")
	fmt.Println("Daftar Email yang Harus Diacc:")
	var i int
	var email string
	for i = 0; i < *n+1; i++ {
		if Mtemp[i].acc == -1 {
			fmt.Println(Mtemp[i].addtemp)
		}
	}
	fmt.Println("Acc Email atau Logout: ")
	fmt.Scan(&email)
	for email != "Logout" {
		accemail(&*Mtemp, email, &*M, &*n, &*nakun)
		fmt.Scan(&email)
	}
}

func accemail(Mtemp *tabtemp, email string, M *tabreg, n, nakun *int) {
	var i, idx int
	idx = -1
	for idx == -1 && i < *n+1 {
		if Mtemp[i].addtemp == email {
			idx = i
			Mtemp[i].acc = 1
			M[*nakun].acc = 1
			M[*nakun].addemail = Mtemp[i].addtemp
			M[*nakun].passemail = Mtemp[i].passtemp
			*nakun++
			hapusacc(idx, &*Mtemp, &*n)
		}
		i++
	}
}

func hapusacc(idx int, Mtemp *tabtemp, n *int) {
	var i int
	for i = idx; i < *n; i++ {
		Mtemp[i] = Mtemp[i+1]
	}
	*n--
}

func menureglog() {
	fmt.Println("-------------------------")
	fmt.Println("         Register        ")
	fmt.Println("          Login          ")
	fmt.Println("           Exit          ")
	fmt.Println("-------------------------")
	fmt.Print("Pilih Salah Satu: ")
}

func menuregisterpengguna() {
	fmt.Println("-------------------------")
	fmt.Println("         REGISTER        ")
	fmt.Println("-------------------------")
	fmt.Print("Masukkan Username dan Password: ")
}

func registerpengguna(Mtemp *tabtemp, n *int) {
	fmt.Scan(&Mtemp[*n].addtemp, &Mtemp[*n].passtemp)
	Mtemp[*n].acc = -1
	*n++
	fmt.Println("Mohon tunggu Admin acc emailmu")
}

func menuloginpengguna() {
	fmt.Println("-------------------------")
	fmt.Println("          LOGIN          ")
	fmt.Println("-------------------------")
	fmt.Print("Masukkan Username dan Password: ")
}

func loginpengguna(M *tabreg, nakun int) {
	var user, pass, pilihbh string
	var i int = 0
	var idxakun int = -1
	var idxakuntujuan int = -1
	var nomorpesan int = -1
	var pilih int
	fmt.Scan(&user, &pass)
	for i < nakun && idxakun == -1 {
		if M[i].addemail == user && M[i].passemail == pass && M[i].acc == 1 {
			idxakun = i
			fmt.Println("Anda Berhasil Login!")
			for pilih != 4 {
				menupengguna(*M, idxakun)
				fmt.Print("Pilih (1/2/3/4)? ")
				fmt.Scan(&pilih)
				if pilih == 1 {
					kirimpesan(&idxakuntujuan, &nakun, &*M, idxakun)
				} else if pilih == 2 {
					lbhpesan(&*M, idxakun)
					fmt.Print("Balas/Hapus Pesan: ")
					fmt.Scan(&pilihbh)
					if pilihbh == "Balas" {
						balaspesan(idxakuntujuan, idxakun, nakun, nomorpesan, &*M)
					} else if pilihbh == "Hapus" {
						hapuspesan(nomorpesan, idxakun, &*M)
					}
				} else if pilih == 3 {
					cetakpesan(&*M, idxakun)
				}
			}
		}
		i++
	}
	if idxakun == -1 {
		fmt.Println("Mohon Coba Lagi")
	}
}

func menupengguna(M tabreg, idxakun int) {
	fmt.Println("-------------------------")
	fmt.Println("         PENGGUNA        ")
	fmt.Println("-------------------------")
	fmt.Println("Akun :", M[idxakun].addemail)
	fmt.Println("1. Kirim Pesan")
	fmt.Println("2. Lihat/Balas/Hapus Pesan")
	fmt.Println("3. Cetak Pesan")
	fmt.Println("4. Logout")
	fmt.Println("-------------------------")
}

func kirimpesan(idxakuntujuan, nakun *int, M *tabreg, idxakun int) {
	fmt.Print("Kirim pesan untuk: ")
	cariuser(&*idxakuntujuan, &*M, *nakun)
	if *idxakuntujuan != -1 {
		fmt.Print("Tulis Pesan: ")
		fmt.Scan(&M[*idxakuntujuan].mail[M[*idxakuntujuan].jpesan].pmasuk)
		M[*idxakuntujuan].mail[M[*idxakuntujuan].jpesan].apesan = "dari " + M[idxakun].addemail
		M[*idxakuntujuan].jpesan++
		fmt.Println("Pesan Terkirim")
	}
}

func cariuser(idxakuntujuan *int, M *tabreg, nakun int) {
	var user string
	var i int = 0
	fmt.Scan(&user)
	for i < nakun && *idxakuntujuan == -1 {
		if M[i].addemail == user {
			*idxakuntujuan = i
		}
		i++
	}
}

func lbhpesan(M *tabreg, idxakun int) {
	var i int
	fmt.Println("List Pesan: ")
	M[idxakun].mail[i].npesan = 0
	for i = 0; i < M[idxakun].jpesan; i++ {
		M[idxakun].mail[i].npesan = i + 1
		fmt.Println(M[idxakun].mail[i].npesan, M[idxakun].mail[i].pmasuk, M[idxakun].mail[i].apesan)
	}
}

func balaspesan(idxakuntujuan, idxakun, nakun, nomorpesan int, M *tabreg) {
	fmt.Print("Balas pesan dari : ")
	cariuser(&idxakuntujuan, &*M, nakun)
	if idxakuntujuan != -1 {
		fmt.Print("Tulis Pesan: ")
		fmt.Scan(&M[idxakuntujuan].mail[M[idxakuntujuan].jpesan].pmasuk)
		M[idxakuntujuan].mail[M[idxakuntujuan].jpesan].pmasuk = M[idxakuntujuan].mail[M[idxakuntujuan].jpesan].pmasuk + " (Balasan)"
		M[idxakuntujuan].mail[M[idxakuntujuan].jpesan].apesan = "dari " + M[idxakun].addemail
		M[idxakuntujuan].jpesan++
	}
}

func caripesan(nomorpesan *int, idxakun int, M *tabreg) {
	var nomor int
	var i int = 0
	fmt.Scan(&nomor)
	for *nomorpesan == -1 && i < M[idxakun].jpesan {
		if nomor == M[idxakun].mail[i].npesan && nomor < M[idxakun].jpesan+1 {
			*nomorpesan = i
		}
		i++
	}
}

func hapuspesan(nomorpesan, idxakun int, M *tabreg) {
	var i int
	fmt.Print("Hapus Pesan nomor: ")
	caripesan(&nomorpesan, idxakun, &*M)
	M[idxakun].jpesan--
	for i = nomorpesan; i < M[idxakun].jpesan+1; i++ {
		M[idxakun].mail[i] = M[idxakun].mail[i+1]
	}
}

func cetakpesan(M *tabreg, idxakun int) {
	var i int
	fmt.Println("Daftar Email yang masuk: ")
	for i = 0; i < M[idxakun].jpesan; i++ {
		M[idxakun].mail[i].npesan = i + 1
		fmt.Printf("%d. %s dari %s\n", i+1, M[idxakun].mail[i].pmasuk, M[idxakun].mail[i].apesan)
	}
}
