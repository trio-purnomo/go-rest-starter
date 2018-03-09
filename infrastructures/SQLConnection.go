package infrastructures

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ISQLConnection is
type ISQLConnection interface {
	GetPlayerWriteDb() *sql.DB
	GetPlayerReadDb() *sql.DB
	CloseConnection()
}

// SQLConnection define sql connection.
type SQLConnection struct{}

var (
	dbPlayerRead, dbPlayerWrite *sql.DB
	err                         error
)

func createDbConnection(descriptor string, maxIdle, maxOpen int, dbName string) *sql.DB {
	conn, err := sql.Open("mysql", descriptor)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "connection for mysql",
			"event":  "mysql_error_connection",
		}).Error(err)
		os.Exit(0)
	}

	conn.SetMaxIdleConns(maxIdle)
	conn.SetMaxOpenConns(maxOpen)
	return conn
}

//GetPlayerWriteDb used for connect to write database
func (s *SQLConnection) GetPlayerWriteDb() *sql.DB {
	if dbPlayerWrite == nil {
		dbPlayerWrite = createDbConnection(
			viper.GetString("database.player.write"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
			"PlayerWriteDB")
	}
	if dbPlayerWrite.Ping() != nil {
		dbPlayerWrite = createDbConnection(
			viper.GetString("database.player.write"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
			"PlayerWriteDB")
	}
	return dbPlayerWrite
}

//GetPlayerReadDb used for connect to read database
func (s *SQLConnection) GetPlayerReadDb() *sql.DB {
	if dbPlayerRead == nil {
		dbPlayerRead = createDbConnection(
			viper.GetString("database.player.read"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
			"PlayerReadDB")
	}
	if dbPlayerRead.Ping() != nil {
		dbPlayerRead = createDbConnection(
			viper.GetString("database.player.read"),
			viper.GetInt("database.player.max_idle"),
			viper.GetInt("database.player.max_cons"),
			"PlayerReadDB")
	}

	return dbPlayerRead
}

// CloseConnection used for close database connection
func (s *SQLConnection) CloseConnection() {

	if dbPlayerRead != nil {
		err = dbPlayerRead.Close()
	}

	if dbPlayerWrite != nil {
		err = dbPlayerWrite.Close()
	}
	if err != nil {
		log.Errorf("db Close Connection Error: %s", err)
	}
}
