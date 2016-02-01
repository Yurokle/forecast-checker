package conditions_persistence

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"forecast-checker/conditions"
	"forecast-checker/config"
)

func openDB() *sql.DB {
	pg_db_uri := fmt.Sprintf(
		"dbname=%s user=%s password=%s sslmode=disable",
		config.GetConfigKey("db_name"),
		config.GetConfigKey("db_user"),
		config.GetConfigKey("db_pass"),
	)

	db, _ := sql.Open("postgres", pg_db_uri)
	if err := db.Ping(); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return nil
	}
	return db
}

var db *sql.DB

func insertConditions(conditions *conditions.Conditions) {
	current_observation := conditions.Current_observation
	result, err := db.Exec(
		"INSERT INTO conditions (zip, station_id, observation_epoch, temp_f, temp_c, relative_humidity, pressure_mb, pressure_in, pressure_trend, dewpoint_f, dewpoint_c, feelslike_f, feelslike_c, visibility_mi, visibility_km, icon, icon_url, json, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)",
		conditions.Zip,
		current_observation.Station_id,
		current_observation.Observation_epoch,
		current_observation.Temp_f,
		current_observation.Temp_c,
		current_observation.Relative_humidity,
		current_observation.Pressure_mb,
		current_observation.Pressure_in,
		current_observation.Pressure_trend,
		current_observation.Dewpoint_f,
		current_observation.Dewpoint_c,
		current_observation.Feelslike_f,
		current_observation.Feelslike_c,
		current_observation.Visibility_mi,
		current_observation.Visibility_km,
		current_observation.Icon,
		current_observation.Icon_url,
		conditions.Json,
		conditions.Created_at,
	)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
}

func StoreConditions(conditions *conditions.Conditions) error {
	if db == nil {
		db = openDB()
	}
	insertConditions(conditions)
	return nil
}
