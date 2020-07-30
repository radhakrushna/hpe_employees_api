package app

import (
	"github.com/federicoleon/bookstore_users-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

}
