**(c)Copyright 2019 , Davin Alfarizky Putra Basudewa all right reserved in touch with dvnlabs.ml**

Email : dbasudewa@gmail.com / moshi2_davin@dvnlabs.ml

## Unofficial API untuk Sistem Informasi Akademik Universitas Banten Jaya   
Live URL hosted on Heroku [API live](https://sikadu-unbaja.herokuapp.com/)  
  
  
## Untuk Apa API ini?  
  
  
API ini dapat digunakan untuk berbagai client diberbagai platform mulai mobile sampai desktop.Dengan adanya API ini,saya berharap developer kampus dapat membuatkan versi mobile dsb  
  
  
  
## **Perhatian untuk developer/enthusiast**  
  
  
 1. API ini tidak bisa melakukan pembuatan KRS,FRS dan sebagainya.Hanya dapat membaca data yang sudah ada di akun anda  
 2. API ini tidak menyimpan **credential** yang anda masukkan.  
 3. API ini memakai **token** yang **hanya berlaku 1 jam** sejak diterbitkan  
 4. Token dapat digunakan diendpoint mana saja selama tidak **expired**  
  
  
  
## Cara kerja :  
  
  
Dikarenakan website sikadu.unbaja.ac.id memakai FW PHP codeigniter maka session yang diterbitkan adalah 'ci_session' .ci_session ini dibungkus kedalam Webtoken yang hanya berlaku 1 jam.Jika sudah lewat 1 jam,anda harus mengissue kembali token yang baru.  
  
  
  
## Endpoint :  
  
  
  
 1. **/** (Home) hanya berisi tentang API ini,siapa yang membuat dan sebagainya  
 2. **/login/mahasiswa** dengan POST FORM (param yang diperlukan adalah user dan password) untuk mengissue token baru. [**HARAP DIINGAT,SETIAP REQUEST YANG ANDA LAKUKAN DI ENDPOINT MANAPUN KECUALI '/' DIHARUSKAN MEMASUKKAN TOKEN**]  
 3. **/mahasiswa/info/{token}** untuk melihat informasi mahasiswa yang sedang login.Ganti {token} dengan token yang anda dapatkan dari /login **contoh https://sikadu-unbaja.herokuapp.com/mahasiswa/info/kqwjas.ash9qwe.asdas**  
 4. **/mahasiswa/schedule/{year}/{quart}/{token}** untuk melihat jadwal anda.Year anda isi dengan tahun akademik anda,quart anda isi dengan 1/2.Angka 1 untuk semester ganjil dan 2 untuk genap.Untuk token sama seperti diatas.  
 5. **/mahasiswa/grade/summary/{token}** untuk melihat secara umu nilai anda selama di kampus  
 6. **/mahasiswa/grade/{year}/{quart}/{token}** untuk melihat detail nilai anda di semester terkait  
  
## Example Request  
  **NOTE:**

>   Jika tahun akademik saat ini adalah 2019/2020,maka tahun yang anda
> masukkan adalah tahun 2019,jika semester ganjil quart bernilai 1 jika
> genap bernilai 2

  
**POST REQUEST**

    <POST> https://sikadu-unbaja.herokuapp.com/login/mahasiswa Dengan urlencoded form user : USER_ANDA password : PASSWORD_ANDA

**Get Request** 
Mendapatkan info mahasiswa yang login

    <GET> https://https://sikadu-unbaja.herokuapp.com/mahasiswa/info/{TOKEN_YANG_DIDAPAT_DARI_LOGIN}

**Mendapatkan Jadwal pada tahun akademik 2018/2019 semester 2**

    <GET> https://https://sikadu-unbaja.herokuapp.com/mahasiswa/schedule/2018/2/{TOKEN_YANG_DIDAPAT_DARI_LOGIN}

**Mendapat nilai selama berkuliah**

    <GET> https://https://sikadu-unbaja.herokuapp.com/mahasiswa/grade/summary/{TOKEN_YANG_DIDAPAT_DARI_LOGIN}

**Mendapat detail mata kuliah nilai di tahun akademik 2018/2019 semester 1**

    <GET> https://https://sikadu-unbaja.herokuapp.com/mahasiswa/grade/2018/1/{TOKEN_YANG_DIDAPAT_DARI_LOGIN}

  
## Persyaratan di Server sendiri  
  
>  - Golang 1.12.7  
>  - Not behind proxy server  
>  - Allow compiled binary to access internet  
  
## Cara compile  
  
Cukup dengan cara clone project ini ke komputer lokal anda,lalu cd ke directory project ini,kemudian jalankan   
  

     go build main.go -o unbajaapi  

## Contribution  
  
  
Saya menerima masukan maupun perbaikan dari kode yang dibuat.Anda bisa membuka 'issue' maupun membuat pull request
