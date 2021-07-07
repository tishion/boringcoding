package datastorage

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// local packages
	"mmc.com/landingtask/be/internal/common"
)

// Represents the data store object
type DataStorage struct {
	db *sql.DB
}

// Creates a DataStore object
func NewStorage(connString string) *DataStorage {
	d, err := sql.Open("mysql", connString)
	if nil != err {
		panic(err)
	}

	// config the connection pool
	d.SetMaxOpenConns(64)
	d.SetMaxIdleConns(64)

	return &DataStorage{
		db: d,
	}
}

// Closes the DataStore
func (s *DataStorage) Close() {
	s.db.Close()
}

// Gets the user password from the data store
func (s *DataStorage) GetUserPsw(username string) (string, error) {
	var password string
	err := s.db.QueryRow("select password from test.user where username=? limit 1", username).Scan(&password)
	switch {
	case err == sql.ErrNoRows:
		return "", errors.New("UsernameNotFound")
	case err != nil:
		return "", err
	}

	return password, nil
}

// Gets the user profile from the data store
func (s *DataStorage) GetUserProfile(username string) (string, string, error) {
	var nick string
	var avatar string
	err := s.db.QueryRow("select nickname, avatar from test.user where username=? limit 1", username).Scan(&nick, &avatar)
	switch {
	case err == sql.ErrNoRows:
		return "", "", errors.New("UsernameNotFound")
	case err != nil:
		return "", "", err
	}

	return nick, avatar, nil
}

// Updates the user profile to the data store
func (s *DataStorage) UpdateProfile(username string, nick string, avatar string) (string, string, error) {
	ret, err := s.db.Exec("update test.user set nickname=?, avatar=? where username=?", nick, avatar, username)
	if err != nil {
		common.LogE("Failed to update: %s", err.Error())
		return "", "", err
	}

	cnt, err := ret.RowsAffected()
	if err != nil {
		common.LogE("Failed to update: %s", err.Error())
		return "", "", err
	}
	common.LogD("Update %d rows: ", cnt)

	return "", "", nil
}

func (s *DataStorage) CreateUser(username string, password string, nick string, avatar string) {
	sqlStr := fmt.Sprintf("insert into user(username, password, nickname, avatar) values ('%s', '%s', '%s', '%s')", username, password, nick, avatar)

	ret, err := s.db.Exec(sqlStr)
	if err != nil {
		common.LogE("Failed to insert new record, err:%v", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		common.LogE("Failed to get last insert id, err:%v", err)
		return
	}

	common.LogD("insert success, the id is %d.\n", theID)
}

// The global singleton instance
var dataStorageInstance *DataStorage

// Initialize the singleton instance
func InitDataStore(connString string) {
	dataStorageInstance = NewStorage(connString)
}

// Get the global DataStore instance
func GetInstance() *DataStorage {
	return dataStorageInstance
}
