BINARY_NAME=climb-vr
SOURCE_DIR=./internal
CHECK_FILES=internal/sector.go internal/via.go internal/ubicacion.go

.PHONY: build test clean run install check

build:
    @echo "Compilando el proyecto..."
    go build -o $(BINARY_NAME) $(SOURCE_DIR)

test:
    @echo "Ejecutando pruebas..."
    go test ./...

clean:
    @echo "Limpiando archivos binarios..."
    rm -f $(BINARY_NAME)
    
install:
    @echo "Instalando dependencias..."
    go mod download

check:
    @echo "Comprobando la sintaxis de los archivos..."
    gofmt -e $(CHECK_FILES)