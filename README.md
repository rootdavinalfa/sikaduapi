## Unofficial API untuk Sistem Informasi Akademik Universitas Banten Jaya 

Live URL hosted on Heroku [API live](https://sikadu-unbaja.herokuapp.com/)


## Untuk Apa API ini?


API ini dapat digunakan untuk berbagai client diberbagai platform mulai mobile sampai desktop.Dengan adanya API ini,saya berharap developer kampus dapat membuatkan versi mobile dsb



## **Perhatian untuk developer/enthusiast**


 1. API ini tidak bisa melakukan pembuatan KRS,FRS dan sebagainya.Hanya dapat membaca data yang sudah ada di akun anda
 2. API ini tidak menyimpan **credential** yang anda masukkan.
 3. API ini memakai token yang hanya berlaku 1 jam sejak diterbitkan
 4. Token dapat digunakan diendpoint mana saja selama tidak **expired**



## Cara kerja :


Dikarenakan website sikadu.unbaja.ac.id memakai FW PHP codeigniter maka session yang diterbitkan adalah 'ci_session' .ci_session ini dibungkus kedalam Webtoken yang hanya berlaku 1 jam.Jika sudah lewat 1 jam,anda harus mengissue kembali token yang baru.



## Endpoint :



 1. **/** (Home) hanya berisi tentang API ini,siapa yang membuat dan sebagainya
 2. **/login/mahasiswa** dengan POST FORM (param yang diperlukan adalah user dan password) untuk mengissue token baru. [**HARAP DIINGAT,SETIAP REQUEST YANG ANDA LAKUKAN DI ENDPOINT MANAPUN KECUALI '/' DIHARUSKAN MEMASUKKAN TOKEN**]
 3. **/mahasiswa/info/{token}** untuk melihat informasi mahasiswa yang sedang login.Ganti {token} dengan token yang anda dapatkan dari /login **contoh https://sikadu-unbaja.herokuapp.com/mahasiswa/info/kqwjas.ash9qwe.asdas**
 4. **/mahasiswa/schedule/{year}/{quart}/{token}** untuk melihat jadwal anda.Year anda isi dengan tahun akademik anda,quart anda isi dengan 1/2.Angka 1 untuk semester ganjil dan 2 untuk genap.Untuk token sama seperti diatas.

## Persyaratan di Server sendiri

>  - Golang 1.12.7
>  - Not behind proxy server
>  - Allow compiled binary to access internet

## Cara compile

Cukup dengan cara clone project ini ke komputer lokal anda,lalu cd ke directory project ini,kemudian jalankan 

    go build main.go -o unbajaapi

## Contribution


Saya menerima masukan maupun perbaikan dari kode yang dibuat.Anda bisa membuka 'issue' maupun membuat pull request
