snmp-simulator:
	docker run --rm -d -p 161:161/udp --name snmp_sim tandrup/snmpsim

build-go:
	docker compose -f ./Go/compose.yml build

run-go:
	docker compose -f ./Go/compose.yml run --rm app

test-go:
	docker compose -f ./Go/compose.yml run --rm test

build-php:
	docker compose -f ./Php/compose.yml build

run-php:
	docker compose -f ./Php/compose.yml run --rm app

test-php:
	docker compose -f ./Php/compose.yml run --rm test

build-php-nokia:
	docker compose -f ./PhpNokia/compose.yml build

run-php-nokia:
	docker compose -f ./PhpNokia/compose.yml run --rm app

build-go-nokia:
	docker compose -f ./GoNokia/compose.yml build

run-go-nokia:
	docker compose -f ./GoNokia/compose.yml run --rm app

run-all-three-times:
	echo "Running Go application..."
	make run-go
	make run-go
	make run-go

	echo "----------------------------------------"
	echo ""
	echo "Running PHP application..."
	make run-php
	make run-php
	make run-php
	