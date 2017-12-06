package mssql

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	// mssql driver
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	conn   *sql.DB
	config *Config
)

// Config mssql配置
type Config struct {
	ConnectionTimeout  int    `json:"ConnectionTimeout"`
	MaxOpenConnections int    `json:"MaxOpenConnections"`
	Usr                string `json:"Usr"`
	Pass               string `json:"Pass"`
	ProcPass           string `json:"ProcPass"`
	EncryptEnabled     bool   `json:"EncryptEnabled"`
}

// Startup 启动
func Startup(db string, host string) bool {
	query := url.Values{}
	query.Add("connection timeout", fmt.Sprintf("%d", config.ConnectionTimeout))
	query.Add("database", db)

	var connectionString string
	if config.EncryptEnabled {
		u := &url.URL{
			Scheme: "sqlserver",
			User:   url.UserPassword(config.Usr, config.Pass),
			Host:   host,
			// Path:  instance, // if connecting to an instance instead of a port
			RawQuery: query.Encode(),
		}

		connectionString = u.String()
	} else {
		connectionString = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=%d&log=63&encrypt=disable", config.Usr, config.Pass, host, db, config.ConnectionTimeout)
	}

	dbConn, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalln(db, "Cannot connect: ", err.Error())
		return false
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalln(db, "Cannot ping: ", err.Error())
		return false
	}

	dbConn.SetMaxOpenConns(config.MaxOpenConnections)
	log.Printf("DB %s connected.\n", db)
	return true
}

func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}
