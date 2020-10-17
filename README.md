# Proyecto 1 - Servicio Web

Estuidiante: Jose Pablo Fontana Vindas
Base de Datos escogida: [TV shows on Netflix, Prime Video, Hulu and Disney+ ](https://www.kaggle.com/ruchi798/tv-shows-on-netflix-prime-video-hulu-and-disney)

### Relacion de Entidades:
Debido a que la Base de Datos escogida no era una opcion viable para normalizar sus entidades, se llego a un acuerdo con el profesor para hacer solamente las funciones  para obetener todos los shows de un año especifico, obetener todos los shows de una calificacion especifica, obetener todos los shows de una edad especifica. Ademas del CRUD de los shows

Por lo tanto, las "entidades" serian las siguentes:
* shows
* years
* ages
* IMDb(calificacion)

### Comandos de prueba

* Creadon un show
```
curl -X POST -H "Content-Type: application/json" -d '{"idShow": "6000","Title": "test","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' http://localhost:8080/show
curl -X POST -H "Content-Type: application/json" -d '{"idShow": "6000","Title": "test","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' https://go-first-example.herokuapp.com/show
```

* Obteniendo un show
```
curl http://localhost:8080/show/1
curl https://go-first-example.herokuapp.com/show/1
```

* Actualizando un show
```
curl -d '{"idShow": "6000","Title": "test2","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' -X PUT http://localhost:8080/show/update
curl -d '{"idShow": "6000","Title": "test2","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' -X PUT https://go-first-example.herokuapp.com/show/update
```

* Eliminando un show
```
curl -X DELETE http://localhost:8080/show/6000
curl -X DELETE https://go-first-example.herokuapp.com/show/6000
```

* Obteniendo todos los shows de un año especifico
```
curl http://localhost:8080/years/2005
curl https://go-first-example.herokuapp.com/years/2005
```

* Obteniendo todos los shows de una calificacion especifica
```
curl http://localhost:8080/rates/9
curl https://go-first-example.herokuapp.com/rates/9
```

* Obteniendo todos los shows de un edad especifica
```
curl http://localhost:8080/ages/7+
curl https://go-first-example.herokuapp.com/ages/7+
```
