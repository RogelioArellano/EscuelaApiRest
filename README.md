# EscuelaApiRest

## Esta Api REST simula la gestión escolar, maneja operaciones CRUD de estudiantes en una base de datos
## Al ejecutar este programa, se buscará la base de datos y la tabla necesaria, de no encontrarlas, se crearán 

##Endpoints
GET: localhost:8080/estudiantes

POST: localhost:8080/estudiantes 
BodyExample: {"nombre": "Juan Pérez","direccion": "Av. Siempre Viva 123","email":"juanperez@example.com","telefono": "1234567890","altaLocal": false,"altaSep": false}

PUT: localhost:8080/estudiantes/id-reemplazar por id a modificar 
BodyExample: {"nombre": "Juan Pérez Arellano","direccion": "Av. Siempre Viva 123","email": "juan.perez@example.com","telefono": "1234567890","altaLocal": false,"altaSep": false}