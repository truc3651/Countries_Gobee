# lưu ý 6 packages quan trọng

- db : kết nối với database
- models : khai báo các trường của bảng, và hàm tương tác với data
- controllers : điều chỉnh các view và data, khi router gọi đến
- routers : khai báo các route
- views : giao diện tương tác trược tiếp với user
- statics : các files và documents (vd: ảnh, css, js, csv)

# Tạo table (id là trường tự tăng)

create table countries (
id SERIAL primary key,
country_code text,
country_name text
)

# chỉnh sửa file db/postgres.go

- postgres://<username>:<password>@localhost/<dbname>?sslmode=disable

# lưu ý quan trọng khi đặt TableName, fields cho struct và cho table

# vì đây là ORM framework, nên Id là trường dữ liệu bắt buộc

- Ở table
- TableName : countries
- id - country_code - country_name

- Ở struct
- TableName : Countries
- Id - CountryCode - CountryName

# trong trường hợp chạy bị lỗi version postgres

- run : go get github.com/lib/pq

# (optional) đối với views\*.tpl

- có thể chỉnh language mode về html, để dễ tạo html code
- ngay dưới status bar phía bên phải

# chạy project

- bee run
- http://localhost:8080/
