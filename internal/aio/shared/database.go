package shared

import (
	"aresbot/internal/aio/constants"
	"aresbot/pkg/logger"
	"database/sql"
)

func GetValue(db *sql.DB, key string) (string, bool) {
	row, err := db.Query("SELECT * FROM settings")
	if err != nil {
		logger.ErrorLogger.Println(err)
		logger.WarningLogger.Println(constants.NewGeneralError("can't read database"))
		return "", false
	}
	var value string
	var key_ string

	for row.Next() {
		_ = row.Scan(&key_, &value)
		if key_ == key {
			return value, true
		}
	}
	logger.WarningLogger.Println(constants.NewGeneralError("can't read database"))
	return "", false
}

func SetValue(db *sql.DB, key string, value string) bool {
	c, err := db.Prepare("update settings set value=? WHERE key_=?")
	if err != nil {
		logger.WarningLogger.Println(constants.NewGeneralError("can't set database"))
		logger.ErrorLogger.Println(err)
		return false
	}

	_, err = c.Exec(value, key)
	if err != nil {
		logger.WarningLogger.Println(constants.NewGeneralError("can't set database"))
		logger.ErrorLogger.Println(err)
		return false
	}
	return true
}
