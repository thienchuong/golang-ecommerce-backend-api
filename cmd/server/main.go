package main

import (
	"github.com/thienchuong/golang-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8080")
}
