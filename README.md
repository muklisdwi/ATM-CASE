# ATM-CASE
Study Case Atm

## ATM Case - Tahap 1
Buatlah aplikasi console untuk simulasi mesin ATM	
Aplikasi ATM menyimpan data nasabah, saldo, dan transaksi. Untuk tahap pertama, data-data tersebut disimpan dalam memori.

### Informasi Login
Pada tahap 1 ini, data akun tersimpan dalam memory dan dibuat saat program dijalankan. Berikut daftar info login yang bisa digunakan :
#### Akun1
`No.Rekening : 111122223333
Pin : 123456
Saldo Awal : 500000`
#### Akun2
`No.Rekening : 222233334444
Pin : 123456
Saldo Awal : 1000000`
#### Akun3
`No.Rekening : 123412341234
Pin : 123456
Saldo Awal : 1500000`

### Menjalankan Program
Untuk menjalankan program, pastikan telah memiliki instalasi program golang yang sudah terpasang. Lalu jalankan perintah berikut di root folder aplikasi melalui terminal/cmd.
```
$ go build

$ ./atmcase
```
### Contoh Login
```
>>>>>> Selamat Datang ! <<<<<<

1. Login
0. Keluar

Pilih Menu [1]/[0] : 1

Silahakan Login !

Masukan Nomor Rekening :
123412341234

Masukan PIN :
123456
```

Gunakan salah satu dari daftar login yang sudah disediakan diatas, pada bagian awal readme.md

1. Login => tekan 1 pada pilih menu, maka akan keluar tampilan login seperti diatas.
0. Keluar => tekan 0 lalu enter, maka aplikasi akan berhenti berjalan.

### Contoh tampilan menu setelah login
```
Selamat Datang, Akun3
1. Tarik Uang
2. Setor Uang
3. Transfer
4. Lihat Riwayat Transaksi
0. Keluar

Pilih Menu [1]/[2]/[3]/[4]/[0] : 4
```

Jenis jenis menu yang tersedia
1. Tarik Uang => melakukan tarik uang, proses mengurangi saldo
2. Setor Uang => melakukan setor uang, kelipatan 50000 dan proses menambah saldo
3. Transfer => proses transfer, mengurangi saldo pada pengirim dan menambah saldo pada penerima
4. Lihat Riwayat Transaksi => menampilkan daftar riwayat transaksi yang pernah dilakukan
0. Keluar => keluar dari menu dan kembali ke tampilan login

*) riwayat yang tersimpan adalah setelah aplikasi dijalankan, setelah aplikasi ditutup riwayat di reset ulang.

#### Contoh tampilan tarik uang
```
Tarik Uang :
1. 50000
2. 100000
3. 200000
4. 300000
0. Keluar

Pilih Menu [1]/[2]/[3]/[4]/[0] : 1

Anda akan melakukan tarik uang 50000

[1] Ya / [0] Tidak : 1

>>>>> Penarikan diproses <<<<
```

#### Contoh tampilan setor uang
```
Setor Uang :
Setor hanya menerima kelipatan 50000

Masukan jumlah nominal : 50000

Anda akan melakukan setor uang 50000

[1] Ya / [0] Tidak : 1

>>>>>> Setoran diproses <<<<<
```

#### Contoh tampilan transfer
```
Transfer :
Masukan nomor rekening tujuan :
111122223333

Masukan nominal transfer :
50000

Anda akan melakukan transfer uang 50000
ke rekening 111122223333 a/n Akun1

[1] Ya / [0] Tidak : 1

>>>>> Transfer Berhasil ! <<<<
```

#### Contoh tampilan riwayat transaksi
```
Daftar Riwayat Transaksi :
2022-04-27 09:00:14 | Setor | 1500000 | 1500000
2022-04-27 09:00:47 | Kirim | 500000 | 1000000
2022-04-27 09:01:17 | Kirim | 500000 | 500000
2022-04-27 09:01:35 | Tarik | 300000 | 200000
2022-04-27 09:01:49 | Setor | 50000 | 250000
2022-04-27 09:02:16 | Terima | 50000 | 300000
2022-04-27 09:02:16 | Kirim | 50000 | 250000
2022-04-27 09:06:56 | Tarik | 200000 | 50000
2022-04-27 09:07:04 | Tarik | 50000 | 0
2022-04-27 09:08:54 | Terima | 1000000 | 1000000
2022-04-27 10:05:13 | Tarik | 50000 | 950000
2022-04-27 10:07:26 | Setor | 50000 | 1000000
2022-04-27 10:09:13 | Kirim | 50000 | 950000

(Tekan Enter)
```
Keterangan :
Tanggal | Jenis Transaksi | Nominal Transaksi | Saldo Akhir