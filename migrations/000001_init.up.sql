CREATE TABLE users
(
    id serial primary key,
    login varchar(100) not null unique,
    first_name varchar(100) not null,
    surname varchar(100) not null,
    password_hash varchar(255) not null
);

CREATE TABLE categories
(
    id serial primary key,
    category_name varchar(255) not null unique
);

CREATE TABLE products
(
    id serial primary key,
    title varchar(255) not null,
    price float not null,
    description varchar(255) not null,
    category varchar(255) references categories (category_name),
    image varchar(255) not null,
    rating varchar(255) not null
);

INSERT INTO users (login, first_name, surname, password_hash)
values ('admin', 'admin', 'admin', '666d3265666a323066656a326e77667c222fb2927d828af22f592134e8932480637c0d');

INSERT INTO categories (category_name)
values ('Украшения'), ('Мужская одежда'), ('Женская одежда'), ('Канцтовары'), ('Продукты'), ('Одежда для детей'), ('Товары для животных'), ('Электроника');