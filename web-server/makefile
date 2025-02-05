include .env

test-add-stock:
	 curl -X POST http://localhost:${APPLICATION_PORT}/inventory -H "Content-Type: application/json" -d '{"price": 100, "title": "shoes", "company": "okta", "description": "run fast", "stock_count": 30}'

test-get-stock:
	curl -X GET http://localhost:${APPLICATION_PORT}/inventory/$(ID) -H "Content-Type: application/json"