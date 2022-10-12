1) Install Golang (https://go.dev/doc/install)
2) Install PostgreSQL (https://www.postgresql.org/download/)
3) Install goose migration tool (https://github.com/pressly/goose)
4) Create Postgresql database (you can create more than one)
5) Run migrations `goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" status` using your db credentials and db name. if you have several databases you should run migrations for each of them
6) Add `127.0.0.1	katinka.local` to your hosts file )
7) Run project
   1) build executable `go build`
   2) run executable
8) go to http://katinka.local 
9) Wilkommen!