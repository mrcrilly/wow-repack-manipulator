package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// Database connection
	databaseDsn string
	db          *gorm.DB

	// Loggers
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger

	// Configuration
	config     Config
	configPath string
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func database(dsn string) {
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		ErrorLogger.Fatalf("unable to connect to database via given dsn: %s: %s\n", dsn, err)
	}
}

func configuration(path string) {
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		ErrorLogger.Fatalf("unable to load %s: %s\n", path, err)
	}
}

func main() {
	flag.StringVar(&databaseDsn, "dsn", "root:ascent@tcp(localhost:3306)/emucoach_v15_vip_world", "The database DSN string")
	flag.StringVar(&configPath, "config", "./manipulations.toml", "The configuration file to use")

	// Instantiate Database connection
	database(databaseDsn)
	InfoLogger.Printf("connected to database: %s\n", databaseDsn)

	configuration(configPath)
	InfoLogger.Printf("successfully loaded configuration: %s\n", configPath)

	switch config.Repack {
	case "catav15":
		parseCataV15Repack(config.Manips)
	}
}
