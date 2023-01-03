CREATE TABLE users
(
    id integer primary key,
    login varchar(100) not null unique,
    first_name varchar(100) not null,
    surname varchar(100) not null,
    password_hash varchar(255) not null
);

CREATE TABLE categories
(
    id integer primary key,
    category_name varchar(255) not null unique
);

CREATE TABLE products
(
    id integer primary key,
    title varchar(255) not null,
    price float not null,
    description varchar(255) not null,
    category varchar(255) references categories (category_name),
    image varchar(255) not null,
    rating varchar(255) not null
);