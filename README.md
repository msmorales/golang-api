# Golang API

Esta es una API escrita en Golang que permite realizar operaciones CRUD en las entidades `User`, `Video` y `Challenge`, así como llenar estas tablas con datos generados.

## Requisitos

- Go 1.16 o superior
- Una base de datos MySQL

## Configuración

1. Clona el repositorio:

   ```sh
   git clone https://github.com/msmorales/golang-api.git
   cd golang-api

2. Configura la conexión a la base de datos en el archivo config/db.go o en .env:

    const (
    Host     = "localhost"
    Port     = 5432
    User     = "tu-usuario"
    Password = "tu-contraseña"
    DBName   = "tu-base-de-datos"
)

3. go run main.go


//////////////////////////////////////////////////////////////////////////////////////////

Endpoints

Users

GET /users: Obtiene todos los usuarios.
GET /users/:id: Obtiene un usuario por ID.
POST /users: Crea un nuevo usuario.
PUT /users/:id: Actualiza un usuario existente.
DELETE /users/:id: Elimina un usuario por ID.

Videos

GET /videos: Obtiene todos los videos.
GET /videos/:id: Obtiene un video por ID.
POST /videos: Crea un nuevo video.
PUT /videos/:id: Actualiza un video existente.
DELETE /videos/:id: Elimina un video por ID.

Challenges

GET /challenges: Obtiene todos los desafíos.
GET /challenges/:id: Obtiene un desafío por ID.
POST /challenges: Crea un nuevo desafío.
PUT /challenges/:id: Actualiza un desafío existente.
DELETE /challenges/:id: Elimina un desafío por ID.

Fill Data

POST /fill: Llena la tabla especificada con datos generados.

Cuerpo de la solicitud:

{
  "entity": "users | videos | challenges",
  "count": 10
}

Ejemplos de uso
Usando cURL

Llenar la tabla users con 20 registros

    curl -X POST "http://localhost:8080/fill" -H "Content-Type: application/json" -d '{"entity": "users", "count": 20}'

Llenar la tabla videos con 15 registros

    curl -X POST "http://localhost:8080/fill" -H "Content-Type: application/json" -d '{"entity": "videos", "count": 15}'

Llenar la tabla challenges con 10 registros

    curl -X POST "http://localhost:8080/fill" -H "Content-Type: application/json" -d '{"entity": "challenges", "count": 10}'

Usando Postman

1. Abrir Postman: Abre la aplicación Postman en tu computadora.
2. Crear una nueva solicitud: Haz clic en "New" y selecciona "Request".
3. Configurar la solicitud:
    - Método: Selecciona POST.
    - URL: Ingresa la URL de la API, por ejemplo, http://localhost:8080/fill.
    - Headers: Añade un header Content-Type con el valor application/json.
    - Body: Selecciona raw y JSON y añade el JSON con la entidad y la cantidad de registros.
4. Enviar la solicitud: Haz clic en "Send".

Ejemplo de configuración en Postman para llenar la tabla users con 20 registros
    - Método: POST
    - URL: http://localhost:8080/fill
    - Headers: Content-Type: application/json
    - Body:

        {
            "entity": "users",
            "count": 20
        }

Contribuciones
Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request.

Licencia
Este proyecto está licenciado bajo la Licencia MIT.

