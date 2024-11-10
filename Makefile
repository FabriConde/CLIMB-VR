# Variables
BINARY_NAME=CLIMB-VR
SOURCE_DIR=./internal
CHECK_FILE=internal/sector.go

# Tareas
.PHONY: build test clean run install check

//Construir el proyecto
build:
    @echo "Compilando el proyecto..."
    go build -o $(BINARY_NAME) $(SOURCE_DIR)

//Ejecutar pruebas
test:
    @echo "Ejecutando pruebas..."
    go test ./...

//Limpiar proyecto
clean:
    @echo "Limpiando archivos binarios..."
    rm -f $(BINARY_NAME)

//Ejecutar la aplicación
run: build
    @echo "Ejecutando la aplicación..."
    ./$(BINARY_NAME)

//Instalar dependencias
install:
    @echo "Instalando dependencias..."
    go mod tidy
    go mod download

//Comprobar la sintaxis de la entidad
check:
    @echo "Comprobando si el archivo $(CHECK_FILE) existe..."
    @if [ -f "$(CHECK_FILE)" ]; then \
        echo "El archivo $(CHECK_FILE) existe."; \
    else \
        echo "El archivo $(CHECK_FILE) no existe."; \
        exit 1; \
    fi