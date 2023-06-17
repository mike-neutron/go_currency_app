<b>Hello, it is a rates app!</b><br>
It uses Golang, Fiber, Gorm.

<b>Functional:</b><br>
- Get rates for date
- Get rate from one to another currency


<b>Swagger:</b><br>
`/swagger/index.html`


<b>Commands:</b><br>
- `docker compose up`<br>
for run server<br>
- `docker exec -it rates_app go run commands/getRates.go`<br>cron to parse rates from CBR<br>