# Proyecto Go - Servidor HTTP "Hola"

Este proyecto es un microservicio simple en Go que expone un endpoint HTTP y mide el tiempo de respuesta de las solicitudes entrantes mediante un middleware.

---

## 📦 Estructura del Proyecto

.
├── main.go
├── go.mod
└── internal/
├── handler/
│ └── hola.go
└── middleware/
└── logger.go

yaml
Copiar
Editar

---

## 🚀 Cómo correr el proyecto

### 1. Clona el repositorio

```bash
git clone https://github.com/tuusuario/tu-repo.git
cd tu-repo
2. Inicializa los módulos de Go
bash
go mod tidy
3. Corre el servidor
bash
go run main.go
Verás en la consola:
nginx
Servidor escuchando en http://localhost:8080
🧪 Cómo probarlo
Abrí otra terminal y ejecutá:
bash
curl http://localhost:8080
La respuesta será:

¡Hola desde el servidor Go!
En la terminal donde corre el servidor, verás un log como:
csharp
[GET] localhost:8080/ - 72.441µs
