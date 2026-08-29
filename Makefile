GOOSE_DRIVER=mysql
GOOSE_DBSTRING=root:password@tcp(127.0.0.1:3307)/qr_order_db?parseTime=true
GOOSE_MIGRATION_DIR=migrations

export GOOSE_DRIVER
export GOOSE_DBSTRING
export GOOSE_MIGRATION_DIR

.PHONY: migrate-up migrate-down migrate-status

migrate-up:
	goose up

migrate-down:
	goose down

migrate-status:
	goose status