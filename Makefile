migrate:
	${HOME}/.goose/bin/goose -dir ./migrations/ postgres "user=username password=password dbname=db sslmode=disable" up
