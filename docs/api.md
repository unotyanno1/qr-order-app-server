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

404 Not Found

```json
{
  "message": "table not found"
}
```
