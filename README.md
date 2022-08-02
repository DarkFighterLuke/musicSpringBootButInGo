# musicSpringBootButInGo
A Spring Boot example project rewritten in Go.

## Requirements
The project uses Go modules, but ODPI-C client has to be installed in order to get the Oracle DB GORM driver to work. See [ODPI-C installation](https://oracle.github.io/odpi/doc/installation.html).

## How to run
```
git clone https://github.com/DarkFighterLuke/musicSpringBootButInGo.git
cd musicSpringBootButInGo
go mod tidy
go run .
```

## API collection
Sample requests and responses are available [here](https://www.getpostman.com/collections/29845a8e5f6ba00c4c38) on Postman.
