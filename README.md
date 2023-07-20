# Hotelifoz Backend

## Description
Hotelifoz adalah aplikasi berbasis Android yang dapat membantu pengguna dalam mencari dan memesan hotel. Dalam hal pencarian hotel, Hotelifoz menyediakan fitur pencarian nama hotel dan filter berdasarkan lokasi terdekat, terpopuler, dan low budget. Pengguna juga dapat menyimpan hotel favorit mereka ke dalam daftar bookmark, sehingga dapat mengaksesnya dengan mudah di kemudian hari. Fitur reservasi hotel juga tersedia di Hotelifoz, sehingga pengguna dapat memesan kamar hotel secara online dengan cepat dan mudah. Secara keseluruhan, Hotelifoz adalah aplikasi yang sangat berguna bagi pengguna yang membutuhkan akomodasi hotel dengan cepat dan mudah.

## Fitur Utama

-   Routing HTTP: Aplikasi ini menyediakan routing HTTP yang kuat dan fleksibel untuk menangani permintaan dari aplikasi klien. Pengguna dapat menentukan rute dan penanganan permintaan sesuai kebutuhan.
-   Koneksi Database: Aplikasi ini dapat terhubung ke database, seperti MySQL Hal ini memungkinkan pengguna untuk menyimpan dan mengambil data dari database sesuai kebutuhan.
-   Keamanan: Aplikasi ini memperhatikan aspek keamanan dengan menggunakan fitur-fitur seperti autentikasi dan otorisasi. Pengguna dapat mengimplementasikan lapisan keamanan sesuai dengan kebutuhan aplikasi.
-   Middleware: Aplikasi ini mendukung penggunaan middleware untuk memanipulasi dan memproses permintaan sebelum dan setelah rute dijalankan. Middleware dapat digunakan untuk melakukan validasi data, logging, dan lain sebagainya.
-   Pengaturan Konfigurasi: Aplikasi ini memungkinkan pengguna untuk mengatur konfigurasi melalui file atau variabel lingkungan (environment variables). Ini mempermudah pengguna untuk menyesuaikan aplikasi sesuai dengan lingkungan yang digunakan.

## Requirement & Stack

-   Golang >=1.18
-   MySQL 8.0.33

## Local Installation

1. Clone this project
    ```
    git clone https://github.com/berrylradianh/go-hotelifoz.git
    ```

2. Copy `.env.example` to `.env`
    ```
    cp .env.example .env
    ```
3. Configure environment variables for database connection
    ```
    DB_CONNECTION=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=hotelifoz_db
    DB_USERNAME=root
    DB_PASSWORD=
    ```

4.  Run the application
    ```
    go run main.go
    ```