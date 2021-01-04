package db

const AddNewStudent = `Insert into users(name, surname, phone, age, email, gender, Role, login, password)
values($1, $2, $3, $4, $5, $6, $7, $8, $9);`
