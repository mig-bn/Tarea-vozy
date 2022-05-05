# Tarea-vozy

Desarrollar una APIRest implementando la librería estándar de Go (golang) que permita hacer un
CRUD de Usuarios y sus respectivos Posts, como motor de base de datos se debe implementar
mongoDB utilizando la librería oficial de mongo.

## Base de datos Mongo
	usr      = "miguel"
	pwd      = "123456"
	host     = "localhost"
	port     = 27017
	database = "test"
  
## Correr APIRest

```
go run main.go
```
luego desde una herramienta para el testing de APIRest puede ser Postman. Con a esta herramienta testearemos API REST:

## insertar:
para realizar el insertado de datos se manda los Keys:Value desde el query params	
```
http://localhost:8080/insert
```
ejemplo:
[![ejemplo-Insert.png](https://i.postimg.cc/j2ST1kXx/ejemplo-Insert.png)](https://postimg.cc/ph43hC44)

## Leer:
para realizar el leer de datos se manda:
```
http://localhost:8080/read
```
ejemplo:
[![ejemplo-Read.png](https://i.postimg.cc/SQf7713X/ejemplo-Read.png)](https://postimg.cc/grrZGsYd)

## Modificar:
para realizar el modificar se necesita copiar un ID que arroja el Leer y pasarlo con datos se manda los Keys:Value desde el query params	
```
http://localhost:8080/update
```
ejemplo:
[![ejemplo-Update.png](https://i.postimg.cc/wM414NDr/ejemplo-Update.png)](https://postimg.cc/kR8MBB1N)

## Eliminar:
para realizar el modificar se necesita copiar un ID que arroja el Leer y se manda los Keys:Value desde el query params
```
http://localhost:8080/delete
```
ejemplo:
[![ejemplo-Delete.png](https://i.postimg.cc/cH91dS6g/ejemplo-Delete.png)](https://postimg.cc/bDnfR4xp)
