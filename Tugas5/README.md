#### Kelompok 4
##### Anggota : 
##### Faishol Muzaky - 1301160369

<br></br>
## Web Server Design

Finite State Machine Web Server :

<p align="center">
  <img width="700" height="100" src="fsm.png">
</p>

cara kerja :

<p align="justify">
Berdasarkan dengan FSM diatas maka, cara kerja dari web server yang kami rancang secara umum adalah sebagai berikut :
</p>
<p align="justify">
Pada browser meminta data web page kepada server, maka instruksi permintaan data oleh browser tersebut akan dikemas di dalam TCP yang merupakan protokol transport  dan dikirim ke alamat yang dalam hal ini merupakan protokol berikutnya yaitu HTTP dan atau HTTPS.
</p>
<p align="justify">
Data yang diminta dari browser ke web server disebut dengan HTTP request kemudian akan dicarikan oleh web server di dalam data server. Jika ditemukan, data tersebut akan dikemas oleh web server dalam TCP dan dikirim kembali ke browser untuk ditampilkan.</p>
<p align="justify">
Data yang dikirim dari server ke browser disebut dengan HTTP response. Jika data yang diminta oleh browser tersebut tidak ditemukan oleh web server, maka web server akan menolak permintaan tersebut dan browser akan menampilkan notifikasi Page Not Found atau Error 404.
</p>


## Web Server Implementation

contoh implemetasi :
<p align="center">
  <img width="100%" height="100%" src="implementasii.PNG">
</p>
<br></br>

Untuk Menjalankannya :
```{r, engine='bash', count_lines}
go run *.go
