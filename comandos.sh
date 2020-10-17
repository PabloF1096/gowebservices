TV shows
POST
curl -X POST -H "Content-Type: application/json" -d '{"idShow": "6000","Title": "test","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' http://localhost:8080/show
curl -X POST -H "Content-Type: application/json" -d '{"idShow": "6000","Title": "test","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' https://go-first-example.herokuapp.com/show
GET
curl http://localhost:8080/show/1
curl https://go-first-example.herokuapp.com/show/1
PUT
curl -d '{"idShow": "6000","Title": "test2","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' -X PUT http://localhost:8080/show/update
curl -d '{"idShow": "6000","Title": "test2","Year": "2020","Age": "18+","IMDb": "10","Rotten Tomatoes": "96%","Netflix": "1","Hulu": "0","Prime Video": "0","Disney+": "0","type": "0"}' -X PUT https://go-first-example.herokuapp.com/show/update
DELETE
curl -X DELETE http://localhost:8080/show/6000
curl -X DELETE https://go-first-example.herokuapp.com/show/6000

Get Years
curl http://localhost:8080/years/2005
Get Rates
curl http://localhost:8080/rates/9
Get Ages
curl http://localhost:8080/ages/7+