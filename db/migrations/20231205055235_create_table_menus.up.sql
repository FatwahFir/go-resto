CREATE TABLE menus (
    id CHAR(36) NOT NULL,
    category_id char(36) NOT NULL,
    name varchar(100) NOT NULL,
    price int(11) NOT NULL,
    price_modal int(11) NOT NULL,
    description text NOT NULL,
    image varchar(100) NOT NULL,
    is_available tinyint(1) NOT NULL,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    FOREIGN KEY(category_id) REFERENCES categories(id),
    PRIMARY KEY(id) 
) engine = innodb;