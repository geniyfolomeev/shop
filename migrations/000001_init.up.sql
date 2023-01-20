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
    image bytea not null,
    rating varchar(255) not null
);

INSERT INTO users (login, first_name, surname, password_hash)
values ('admin', 'admin', 'admin', '666d3265666a323066656a326e77667c222fb2927d828af22f592134e8932480637c0d');

INSERT INTO categories (category_name)
values ('Украшения'), ('Мужская одежда'), ('Женская одежда'), ('Канцтовары'), ('Продукты'), ('Одежда для детей'), ('Товары для животных'), ('Электроника');

INSERT INTO products (title, price, description, category, image, rating)
values ('картинка', 145.3, 'описание', 'Мужская одежда', decode('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', 'base64'), 'Жабий рейтинг');