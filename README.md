<b>Hello, it is a rates app!</b><br>
It uses Golang, Fiber, Gorm.

<b>Methods:</b><br>
`GET /api/date/:date`<br>
for get rates by date, example /api/date/1970-01-01

`GET /api/exchange/:from/:to/1000`<br>
for exchange, example /api/exchange/RUB/USD/1000

<b>Commands:</b><br>
`docker compose up`<br>
for run server<br><br>
`go run getRates.go`<br>from container from commands folder  for parse rates from CBR<br>