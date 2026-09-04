INSERT INTO stores (id, name)
VALUES
    (1, 'テスト店舗')
ON DUPLICATE KEY UPDATE
    name = VALUES(name);

INSERT INTO tables (id, store_id, table_number)
VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3)
ON DUPLICATE KEY UPDATE
    store_id = VALUES(store_id),
    table_number = VALUES(table_number);
