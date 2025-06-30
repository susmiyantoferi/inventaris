CREATE database inventaris;

CREATE table produk(
	id int primary key auto_increment,
	nama varchar(100),
	deskripsi varchar(225),
	harga decimal,
	kategori varchar(225),
	gambar varchar(255)
);

CREATE table inventaris(
	id int primary key auto_increment,
	id_produk int,
	jumlah int,
	lokasi varchar(100),
	foreign key (id_produk) references produk(id)
);

CREATE table pesanan(
	id int primary key auto_increment,
	id_produk int,
	jumlah int,
	tanggal_pesanan date,
	foreign key (id_produk) references produk(id)
	
);

INSERT into produk (nama, deskripsi, harga, kategori)
values ("indomi goreng", "mie instan", 5000, "makanan"),
 ("indomi kuah soto", "mie instan", 3000, "makanan"),
 ("indomi goreng aceh", "mie instan", 5000, "makanan"),
 ("indomi goreng kari", "mie instan", 4000, "makanan"),
 ("indomi goreng rendang", "mie instan", 5000, "makanan"),
 ("aqua mineral", "air mineral", 5000, "minuman"),
 ("pocari sweet", "air ion electrolit", 10000, "minuman"),
 ("moo milk", "susu murni", 10000, "minuman"),
 ("teh pucuk", "minuman teh manis", 5000, "minuman"),
("bear brand", "susu murni", 11000, "minuman");


INSERT into inventaris (id_produk, jumlah, lokasi)
values (1, 3, 'gudang A'),
(1, 3, 'gudang A'),
(2, 11, 'Gudang A'),
(5, 7, 'Gudang C'),
(8, 10, 'Gudang A'),
(10, 7, 'Gudang C');

INSERT into pesanan (id_produk, jumlah, tanggal_pesanan)
values (1, 2, '2025-06-01'),
(2, 9, '2025-06-01'),
(10, 5, '2025-06-01'),
(8, 3, '2025-06-01'),
(1, 1, '2025-06-01');

SELECT * FROM  produk p ;

SELECT i.id, nama, jumlah, lokasi 
from inventaris i 
join produk p on i.id_produk = p.id 
WHere nama = 'indomi goreng';

SELECT p.id, nama, jumlah, tanggal_pesanan 
from pesanan p 
join produk p2 on p.id_produk = p2.id ;

SELECT pr.nama,
  SUM(p.jumlah) as jumlah,
  SUM(jumlah*harga) as total_harga
FROM pesanan p
JOIN produk pr ON p.id_produk = pr.id 
GROUP BY pr.nama;

SELECT lokasi,
  SUM(jumlah) AS total_stok
FROM inventaris i 
GROUP BY lokasi;


SELECT 
  pr.nama,
  SUM(p.jumlah) AS total_pesanan,
  sum(jumlah*harga) as total_harga
FROM pesanan p
JOIN produk pr ON p.id_produk = pr.id 
WHERE pr.nama = 'indomi kuah soto'
GROUP BY pr.nama;
