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
	WarningLogger   *log.Logger
	InfoLogger      *log.Logger
	DebuggingLogger *log.Logger
	ErrorLogger     *log.Logger

	// Configuration
	config     Config
	configPath string

	// Debugging
	debugging bool
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	DebuggingLogger = log.New(os.Stdout, "DEBUGGING: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
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

	if debugging {
		for _, m := range config.Manips {
			DebuggingLogger.Printf("type: %s; column: %s; type: %s; ids: %v; modifiers: %d\n", m.Type, m.Column, m.Type, m.UniqueIDs, len(m.Pairs))
			for _, p := range m.Pairs {
				DebuggingLogger.Printf("\tkey: %s; value: %v\n", p.Key, p.Value)
			}
		}
	}
}

func main() {
	flag.StringVar(&databaseDsn, "dsn", "root:ascent@tcp(localhost:3306)/emucoach_v15_vip_world", "The database DSN string")
	flag.StringVar(&configPath, "config", "./manipulations.toml", "The configuration file to use")
	flag.BoolVar(&debugging, "debugging", false, "The configuration file to use")
	flag.Parse()

	// Instantiate Database connection
	database(databaseDsn)
	InfoLogger.Printf("connected to database: %s\n", databaseDsn)

	configuration(configPath)
	InfoLogger.Printf("successfully loaded configuration: %s\n", configPath)

	switch config.Repack {
	case "catav15":
		parseCataV15Repack(config.Manips)
	default:
		ErrorLogger.Fatalln("unknown repack value provided")
	}
}
