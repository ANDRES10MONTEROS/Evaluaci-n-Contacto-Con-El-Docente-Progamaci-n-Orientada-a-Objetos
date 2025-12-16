Estructura del Proyecto

Proyecto desarrollado en Go con arquitectura modular por capas.

Carpeta principal: streaming-go

/cmd
 /app
   main.go - punto de entrada

/internal
 /handlers - controladores HTTP (handlers.go)
 /services - lógica del negocio (streaming_service.go)
 /repos - repositorios en memoria (memory_repo.go)
 /models - entidades del dominio (content.go y user.go)
 /errors - errores personalizados (custom_error.go)

/docs
 endpoints.md - documentación API
 vision_futuro.md - visualización del futuro
 estructura_proyecto.md
 manual_instalacion.md

README.md - documentación general
go.mod / go.sum → dependencias


Explicación por Capas

Handlers
Conectan la API con el servicio.  
`CreateUser`, `ListContent`.

Services
Lógica de negocio:
- Registrar usuario
- Crear contenido
- Reproducir
- Historial

Repos
Simulan base de datos con almacenamiento en memoria.

Models
Define entidades:
- User
- Content

Errors
Define errores personalizados:
- ErrNotFound
- ErrAlreadyExists

Esta estructura permite escalabilidad y un mantenimiento profesional.