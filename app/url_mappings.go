package app

import (
	"github.com/radhakrushna/hpe_employees_api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

}
