Teknologi Yang Di Gunakan :

- Go Gin Gonic
- GORM
- Mysql
- godotenv

Set up local :
- clone repo
- ubah nama file .env.sample jadi .env
- ubah file .env sesuai dengan nama database, username, password, host dan port masing masing (jangan ubah MYSQL_ROOT_PASSWORD karna digunakan untuk kebutuhan service mysql di docker)
- buat database sesuai dengan nama database yang ada di .env
- jalankan di terminal go run main.go

Step deploy vps:
- buat Dockerfile yang nantinya digunakan untuk build aplikasi go
- buat docker-compose.yml, buat container untuk aplikasi go dan buat container untuk database mysql
- buat folder .github/workflows lalu isi dengan file deploy.yml, didalam deploy.yml
- project sudah harus ada di repository remote github, jika belum buat dan push terlebih dahulu
- lalu masuk ke setting di repo github masuk ke Secrets and variables -> Actions, lalu klik New repository secret.
- masukan SSH_HOST yang berisi alamat host server, SSH_USER, berisi user server, SSH_PRIVATE_KEY, berisi private key yang cocok dengan public key pada server. 
- login ke server via SSH
- jalankan sudo apt update di server
- install docker dan docker-compose, sudo apt install docker.io docker-compose -y (jika sudah ada tidak usah)
- lalu masuk ke directory server yang sama seperti yang ada di deploy.yml bagian script
- clone reposotory dengan git clone <url_repo> . (jangan lupa titik diakhir, agar tidak terbuat folder baru lagi)
- cek remote, git remote -v, jika belum ada lakukan remote ke github tersebut
- buat file .env, dengan cara touch .env lalu edit di nano dan masukan code sama seperti .env.sample
- lalu tes ci/cd dengan buat perubahan di code lalu push di main, jika berhasil server akan berjalan dan ci/cd akan hijau


URL Documentation API Inventaris :
https://documenter.getpostman.com/view/22397647/2sB34mhxzp
