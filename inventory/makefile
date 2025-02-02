include .env

db:
	docker run --name inventory_db \
	  -e POSTGRES_USER=${DB_USERNAME} \
	  -e POSTGRES_PASSWORD=${DB_PASSWORD} \
	  -e POSTGRES_DB=${DB_DATABASE} \
	  -p 5432:${DB_PORT} \
	  -d postgres

test-create-item:
	 grpcurl -d '{"inventory_item": {"price": 100, "title": "shoes", "company": "okta", "description": "run fast", "stock_count": 30}}' -plaintext localhost:3030 Inventory/CreateItem

# ex) make test-get-item ID=2
test-get-item:
	grpcurl -d '{"id": $(ID)}' -plaintext localhost:3030 Inventory/GetItem

.PHONY: db test-create-item test-get-item