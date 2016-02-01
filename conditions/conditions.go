package conditions

type ObservationLocation struct {
	Full string "full"
	City string "city"
	State string "state"
	Country string "country"
	Country_iso3166 string "country_iso3166"
	Latitude string "latitude"
	Longitude string "longitude"
	Elevation string "elevation"
}

type CurrentObservation struct {
	Observation_location ObservationLocation "observation_location"
	Station_id string "station_id"
	Observation_epoch string "observation_epoch"
	Temp_f float32 "temp_f"
	Temp_c float32 "temp_c"
	Relative_humidity string "relative_humidity"
	Pressure_mb string "pressure_mb"
	Pressure_in string "pressure_in"
	Pressure_trend string "pressure_trend"
	Dewpoint_f int "dewpoint_f"
	Dewpoint_c int "dewpoint_c"
	Feelslike_f string "feelslike_f"
	Feelslike_c string "feelslike_c"
	Visibility_mi string "visibility_mi"
	Visibility_km string "visibility_km"
	Icon string "icon"
	Icon_url string "icon_url"
}

type Conditions struct {
	Current_observation CurrentObservation "current_observation"
	Json string
}