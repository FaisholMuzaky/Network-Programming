#### Kelompok 4
#### Tugas 7 Pemrograman Jaringan
##### Faishol Muzaky - 1301160369


### PhPloy Tool
<p align="justify">
PHPloy merupakan deployment tool Git FTP dan SFTP. PHPloy merupakan file PHP sederhana yang dapat dipanggil melewati bash. Tools ini bekerja dengan menyimpan file .revision yaitu hash commit yang telah dideploy pada server. Ketika PHPloy dijalankan, PHPloy akan men-download file tersebut dan membandingkannya dengan referensi dari commit di dalamnya dengan commit yang akan kita upload. Setelah itu PHPloy akan menyimpan .revision pada setiap submodule respository kita.
</p>

cara kerja :
<p align="justify">
Cara kerja phploy adalah melakukan proses upload/delete file yang telah ter commit pada git ke server sehingga tools ini mempermudah developer dalam deployment web project ke live server. Proses ini dilakukan menggunakan command pada terminal dengan perintah-perintah yang telah di definisikan.
</p>

### Stack Up Tool
<p align="justify">
Stack Up Tool atau SUP merupakan tool untuk menjalankan script pada beberapa host server. koneksi ssh di semua server dibutuhkan jika ingin menggunakan tool ini. Sup tidak menggunakan protocol FTP, melainkan sup menggunakan protocol SSH. Semua managemen dilakukan secara eksplisit pada file konfigurasi, yakni “Supfile”.
</p>

cara kerja :
<p align="justify">
cara kerja dari tools ini berbeda, perbedaannya terletak pada proses penggunaan perintah yang dilakukan secara parallel terhadap beberapa host secara bersamaan.
</p>
