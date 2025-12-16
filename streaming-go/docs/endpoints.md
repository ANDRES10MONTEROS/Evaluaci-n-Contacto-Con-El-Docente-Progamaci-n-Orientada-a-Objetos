
Documentación de Endpoints

Servicios Web utilizados.


Base URL:
http://localhost:8080/api

1. Usuarios
-Crear Usuario
POST /users
Body:
{
  "id": "u1",
  "name": "Andy",
  "email": "andy@example.com"
}

-Listar Usuarios
GET /users

2. Contenido
-Crear Contenido
POST /content
Body:
{
  "id": "c1",
  "title": "Avengers",
  "type": "movie",
  "duration": 120
}

-Listar Contenido
GET /content

3. Reproducción
-Reproducir Contenido
POST /play
Body:
{
  "user_id": "u1",
  "content_id": "c1"
}

4. Historial
-Historial por usuario
GET /history/:id

5. Actualizar título
PUT /content/:id/title
Body:
{
  "title": "Nuevo título"
}

6. Eliminar contenido
DELETE /content/:id
