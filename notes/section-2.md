## Agenda
- Mengenal Go Modules
- Membuat Module
- Rilis Module
- Menambah Dependency
- Upgrade Module
- Upgrade Dependency

## Sebelum Ada Go Modules
- Saat kita membuat aplikasi, biasanya kita akan menggunakan library atau dependency dari project lain.
- Sebelum ada Go Modules, management untuk dependency sangat sulit dilakukan di Go-Lang.
- Biasanya kita akan meng-copy semua source code library atau dependency lain ke project kita, hal ini membuat project kita menjadi bengkak karena penuh dengan library orang lain.
- Atau biasanya library orang lain akan kita save di GOPATH directory, namun hal ini akan sulit jika ternyata beberapa aplikasi butuh library yang sama dengan versi yang berbeda

## Go-Modules
- Go-Modules mulai dikenalkan di Go-lang 1.11 and 1.12
- Dengan Go-Modules, kita bisa membuat project dengan mudah dan dependency management yang sangat mudah

## Membuat Module
- Untuk membuat module baru, kita bisa menggunakan perintah : `go mod init nama-module`
- Go-Lang akan secara otomatis membuat file go.mod yang berisikan informasi nama module dan juga versi Go-Lang yang kita gunakan

## Rilis Module
- Go-Lang terintegrasi baik dengan Git
- Untuk merilis module, kita hanya perlu membuat Tag di Git

## Menambah Dependency
```
go get nama-module
```

## Upgrade Module
- Untuk melakukan upgrade module, kita hanya cukup membuat tag baru di Git

## Upgrade Dependency
- Untuk upgrade dependency ke versi terbaru, kita bisa mengubah isi go.mod, lalu mengubah tag nya menjadi tag terbaru
- Untuk mendownload versi terbaru, gunakan perintah : `go get`

## Major Upgrade
- Major upgrade biasanya terjadi dikarenakan ada perubahan pada isi kode Program kita, sehingga membuatnya tidak backward compatible
- Sebaiknya hal ini sebisa mungkin dihindari
- Namun jika tidak bisa dihindari, strategy terbaik adalah merubah nama module