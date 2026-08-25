# Database Design

## 概要

本システムでは、注文・メニュー・テーブル利用・売上などの業務データをMySQLで管理する。

RAGで使用するベクトルデータはPostgreSQL + pgvectorで管理する。

---

## ER図

erDiagram
    stores ||--o{ tables : has
    stores ||--o{ menus : has
    stores ||--o{ admin_users : has
    
    tables ||--o{ table_sessions : has
    
    table_sessions ||--o{ orders : has
    table_sessions ||--o| table_sales : has
    
    orders ||--o{ order_items : has
    
    menus ||--o{ order_items : referenced_by

---

## テーブル定義

### stores

店舗情報を管理する。

| カラム | 型 | NULL | 説明 |
| --- | --- | --- | --- |
| id | BIGINT | NO | 店舗ID |
| name | VARCHAR(255) | NO | 店舗名 |
| created_at | DATETIME | NO | 作成日時 |
| updated_at | DATETIME | NO | 更新日時 |

### Foreign Key

なし

---

- id
- name
- created_at
- updated_at

tables
- id
- store_id
- table_number
- created_at
- updated_at

table_sessions
- id
- table_id
- guest_no
- status
- started_at
- closed_at
- created_at
- updated_at

menus
- id
- store_id
- name
- price
- description
- image_url
- is_available
- created_at
- updated_at

orders
- id
- table_session_id
- created_at
- updated_at

order_items
- id
- order_id
- menu_id
- quantity
- unit_price
- status
- created_at
- updated_at

table_sales
- id
- table_session_id
- total_amount
- created_at

admin_users
- id
- store_id
- email
- password_hash
- created_at
- updated_at