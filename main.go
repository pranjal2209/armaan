package main

import (
	"fmt"

	"github.com/pranjal2209/armaan/config"
	"github.com/pranjal2209/armaan/services"
)

const (
	Host     = "ec2-54-157-78-113.compute-1.amazonaws.com"
	Database = "d5ndt2glmk4mvu"
	User     = "ozbtcczyruamxy"
	Port     = 5432
	Password = ""
)

func main() {
	fmt.Println("OM")
	config.LoadConfig()
	services.HandlerFunc()
	return
}
