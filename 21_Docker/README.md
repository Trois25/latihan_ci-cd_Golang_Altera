* Docker 
** Container merupakan aplikasi tersebut sendiri yang berisi aplikasi dan apa yang dibutuhkan aplikasi tersebut. 

** Docker infrastructure
* CLient, berisi comment untuk menjalankan doker (docker build, docker pull, docker run)
* Docker host, merupakan mesin yang digunakan untuk menjalankan 1 atau lebih container dan teridiri dari docker images, daemon, Networks, Storage, and Containers. 
* registry,Registri Docker lebih seperti lokasi di mana semua Image Docker disimpan. 

** yang membedakan vm dan container adalah container tidak dapat menduplicate OS.

Dilakukan dengan cara :
* docker file (berisi bagaimana cara menjalankan aplikasi dan dalam aplikasi merupakan code yang kita buat)
* setelah membuat docker file maka dibuatlah docker image (gambaran dari aplikasi dan tidak dapat dijalankan)
* container registry ct: docker hub (merupakan kumpulan dari image)