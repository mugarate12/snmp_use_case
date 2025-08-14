snmp-simulator:
	docker run --rm -d -p 161:161/udp --name snmp_sim tandrup/snmpsim

build-go:
	docker compose -f ./Go/compose.yml build

run-go:
	docker compose -f ./Go/compose.yml run --rm app

build-php:
	docker compose -f ./Php/compose.yml build

run-php:
	docker compose -f ./Php/compose.yml run --rm app
