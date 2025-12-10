
Sistema de Gestión de Streaming  
Evaluación Contacto con el Docente 

**Asignatura:** Programación Orientada a Objetos    
**Carrera:** Ingenieria en Sistemas de la Información    
**Estudiante:** Andrés Monteros  
**Lenguaje:** Go (Golang)  
**Fecha:** 21/07/2025

**1. Introducción**

Este proyecto corresponde al desarrollo de un sistema de streaming.  
El software implementa:

- Gestión de usuarios  
- Gestión de contenido multimedia  
- Reproducción y registro de historial  
- Servicios Web   
- Manejo de errores  
- Encapsulación, POO y funciones  
- Repositorios en memoria  
- Serialización en JSON  

**2. Arquitectura del Proyecto**

El proyecto sigue una arquitectura que separa cada responsabilidad:

/cmd → punto de entrada (main.go) 

/internal 

    /handlers → controladores HTTP (API REST) 
    /services → lógica de negocio 
    /models → modelos y estructura de datos 
    /repos → repositorios en memoria 

/errors → errores personalizados 

/docs → documentación (endpoints, visión del futuro, etc.) 

go.mod, go.sum → dependencias 

**3. Ejecución del Proyecto**

**Requisitos**
- Go 1.20+
- Git
- Postman o Thunder Client

**Instalar dependencias:**
go mod tidy

**Ejecutar:**
go run ./cmd/app

**4. Servicios Web**
1. Crear usuario
2. Listar usuarios
3. Crear contenido
4. Listar contenido
5. Reproducir contenido
6. Obtener historial de un usuario
7. Actualizar título de contenido
8. Eliminar contenido


**5. Tecnologías utilizadas**
- Go
- Git Gonic
- JSON
- Programación Orientada a Objetos
- Funciones y encapsulación
- Repositorios en memoria
- Manejo de errores personalizados

**6. Visualización del Futuro**

La evolución tecnológica permitirá plataformas de streaming con: IA para recomendaciones hiperpersonalizadas

- Aqui el contenido generado por IA en tiempo real se lo vera con realidad aumentada integrada.
- El Streaming adaptativo inteligente
- Integración con dispositivos hápticos y neuronales
- Escenarios personalizados para cada usuario


**7. Conclusiones**
El proyecto permitió desarrollar lógica de negocio modular, arquitectura por capas, programación funcional y orientada a objetos, servicios web reales, control de versiones con Git y GitHub, y documentación técnica, fortaleciendo las habilidades necesarias para proyectos más complejos y profesionales.


**9. Licencia**

Este proyecto se distribuye bajo la licencia MIT.

