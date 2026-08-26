# API Design

## 基本方針

- Base URL: `/api/v1`
- Content-Type: `application/json`
- APIのレスポンスはJSON形式
- エラー時はHTTPステータスコードとエラーメッセージを返す

# Customer API

## セッション開始

QRコードからアクセスしたテーブルの利用セッションを開始する。

### Endpoint

POST `/api/v1/tables/{table_id}/sessions`

### Path Parameters

| Name | Type | Required | Description |
|---|---|---|---|
| table_id | SMALLINT | Yes | テーブルID |

### Request

なし

### Response

#### 201 Created

```json
{
  "session_id": 123,
  "table_id": 10
}
```

### Error

#### 404 Not Found

```json
{
  "message": "table not found"
}
```

## メニュー一覧取得

店舗で現在注文可能なメニュー一覧を取得する。

### Endpoint

GET ```/api/v1/stores/{store_id}/menus```

### Path Parameters

| Name | Type | Required | Description |
|---|---|---|---|
| store_id | SMALLINT | Yes | 店舗ID |

### Request

なし

### Response

#### 200 OK

```json
{
  "menus": [
    {
      "id": 1,
      "name": "醤油ラーメン",
      "price": 900,
      "is_available": true
    },
    {
      "id": 2,
      "name": "餃子",
      "price": 500,
      "is_available": true
    }
  ]
}
```

## カート取得

現在のセッションのカート内容を取得する。

### Endpoint

GET ```/api/v1/sessions/{session_id}/cart```

### Request

なし

### Response

#### 200 OK

```json
{
  "items": [
    {
      "menu_id": 1,
      "name": "醤油ラーメン",
      "unit_price": 900,
      "quantity": 2
    }
  ],
  "total_price": 1800
}
```

## カート追加

メニューをカートに追加する。

### Endpoint

POST ```/api/v1/sessions/{session_id}/cart/items```

### Request

```json
{
  "menu_id": 1,
  "quantity": 2
}
```

### Response

#### 201 Created

```json
{
  "menu_id": 1,
  "quantity": 2
}
```

### Error

#### 400 Bad Request

```json
{
  "message": "invalid quantity"
}
```

#### 409 Conflict

```json
{
  "message": "menu is not available"
}
```

## カートの商品削除

### Endpoint

DELETE ```/api/v1/sessions/{session_id}/cart/items/{menu_id}```

### Request

なし

### Response

#### 204 No Content

レスポンスボディなし

## 注文確定

カートの商品を注文として確定する。

### Endpoint

POST ```/api/v1/sessions/{session_id}/orders```

### Request

なし(※注文内容はサーバー側で現在のカートから取得する)

### Response

#### 201 Created

```json
{
  "order_id": 1001,
  "items": [
    {
      "menu_id": 1,
      "name": "醤油ラーメン",
      "unit_price": 900,
      "quantity": 2,
      "status": 0
    }
  ],
  "total_price": 1800
}
```

## 注文履歴取得

### Endpoint

GET ```/api/v1/sessions/{session_id}/orders```

### Request

なし

### Response

#### 200 OK

```json
{
  "orders": [
    {
      "order_id": 1001,
      "ordered_at": "2026-08-26T12:00:00+09:00",
      "items": [
        {
          "menu_id": 1,
          "name": "醤油ラーメン",
          "quantity": 2,
          "unit_price": 900,
          "status": 1
        }
      ]
    }
  ]
}
```

# Staff API

## 注文一覧取得

店舗に入った注文一覧を取得する。

### Endpoint

GET ```/api/v1/stores/{store_id}/orders```

### Request

なし

### Response

#### 200 OK

```json
{
  "orders": [
    {
      "order_id": 1001,
      "table_id": 10,
      "items": [
        {
          "menu_id": 1,
          "name": "醤油ラーメン",
          "quantity": 2,
          "status": 0
        }
      ]
    }
  ]
}
```

## 注文ステータス変更

注文商品のステータスを変更する。

### Endpoint

PATCH ```/api/v1/order-items/{order_item_id}/status```

### Request

```json
{
  "status": 1
}
```

### Status

| Value | Description |
| --- | --- |
| 0 | 未提供 |
| 1 | 提供済 |

### Response

#### 200 OK

```json
{
  "order_item_id": 2001,
  "status": 1
}
```

## メニュー販売状態変更

### Endpoint

PATCH ```/api/v1/menus/{menu_id}/availability```

### Request

```json
{
  "is_available": false
}
```

### Response

#### 200 OK

```json
{
  "menu_id": 1,
  "is_available": false
}
```




