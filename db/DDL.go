package db

const UserDDL = `Create table if not exists users(
	id 			integer primary key autoincrement,
	name 		text 	not null,
	surname 	text 	not null,
	phone	 	text 	not null,
	age 		integer not null,
	email		text	not null,
	gender 		text 	not null,
	Role		text	not null,
	login 		text 	not null unique,
	password	text 	not null,
	remove 		boolean not null default false 
);`
