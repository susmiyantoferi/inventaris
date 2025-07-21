#!/bin/sh

echo "Menunggu MySQL siap di host $DB_HOST:$DB_PORT..."

# Loop sampai koneksi berhasil
until nc -z $DB_HOST $DB_PORT; do
  echo "MySQL belum siap - tunggu 1 detik"
  sleep 1
done

echo "MySQL siap - menjalankan aplikasi"
exec "$@"
