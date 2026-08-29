-- +goose Up

CREATE TABLE stores (
    id TINYINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE tables (
    id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
    store_id TINYINT UNSIGNED NOT NULL,
    table_number TINYINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_tables_store_id_table_number (store_id, table_number),
    CONSTRAINT fk_tables_store_id
        FOREIGN KEY (store_id)
        REFERENCES stores (id)
);

CREATE TABLE table_sessions (
    id BIGINT NOT NULL AUTO_INCREMENT,
    table_id SMALLINT UNSIGNED NOT NULL,
    guest_no INT NOT NULL,
    status TINYINT(1) NOT NULL,
    started_at DATETIME NOT NULL,
    closed_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_table_sessions_table_id
        FOREIGN KEY (table_id)
        REFERENCES tables (id)
);

CREATE TABLE menus (
    id MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT,
    store_id TINYINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    price SMALLINT UNSIGNED NOT NULL,
    description TEXT NULL,
    image_url VARCHAR(255) NULL,
    is_available TINYINT(1) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_menus_store_id
        FOREIGN KEY (store_id)
        REFERENCES stores (id)
);

CREATE TABLE orders (
    id BIGINT NOT NULL AUTO_INCREMENT,
    table_session_id BIGINT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_orders_table_session_id
        FOREIGN KEY (table_session_id)
        REFERENCES table_sessions (id)
);

CREATE TABLE order_items (
    id BIGINT NOT NULL AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    menu_id MEDIUMINT UNSIGNED NOT NULL,
    quantity TINYINT UNSIGNED NOT NULL,
    unit_price SMALLINT UNSIGNED NOT NULL,
    status TINYINT(1) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_order_items_order_id
        FOREIGN KEY (order_id)
        REFERENCES orders (id),
    CONSTRAINT fk_order_items_menu_id
        FOREIGN KEY (menu_id)
        REFERENCES menus (id)
);

CREATE TABLE table_sales (
    id BIGINT NOT NULL AUTO_INCREMENT,
    table_session_id BIGINT NOT NULL,
    total_amount INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_table_sales_table_session_id (table_session_id),
    CONSTRAINT fk_table_sales_table_session_id
        FOREIGN KEY (table_session_id)
        REFERENCES table_sessions (id)
);

CREATE TABLE admin_users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    store_id TINYINT UNSIGNED NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY uq_admin_users_email (email),
    CONSTRAINT fk_admin_users_store_id
        FOREIGN KEY (store_id)
        REFERENCES stores (id)
);


-- +goose Down

DROP TABLE IF EXISTS admin_users;
DROP TABLE IF EXISTS table_sales;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS table_sessions;
DROP TABLE IF EXISTS tables;
DROP TABLE IF EXISTS stores;
