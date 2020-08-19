package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"rabie.com/portfolio/apimanager"
)

func main() {
	apimanager.Wg.Add(1)
	go apimanager.Allapi()
	apimanager.Wg.Wait()

}
