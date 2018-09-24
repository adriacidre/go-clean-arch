package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/adriacidre/go-clean-arch/middleware"
	httpDeliver "github.com/adriacidre/go-clean-arch/payment/delivery/http"
	repo "github.com/adriacidre/go-clean-arch/payment/repository"
	ucase "github.com/adriacidre/go-clean-arch/payment/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbConn := getDBConnection()
	defer dbConn.Close()

	e := echo.New()
	e.Debug = true
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	ar := repo.NewMysqlPayment(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := ucase.NewPayment(ar, timeoutContext)
	httpDeliver.NewPaymentHTTPHandler(e, au)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}

func getDBConnection() *sql.DB {
	dbHost := viper.GetString(`ºdatabase.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Europe/Paris")
	val.Add("allowNativePasswords", "true")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return dbConn
}
