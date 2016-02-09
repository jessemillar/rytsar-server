package accessors

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// Returns an array of all loot locations and values to plot on the map in iOS
func (ag *AccessorGroup) DumpDatabase(userLatitude float64, userLongitude float64, radius float64) (string, error) {
	rows, err := ag.DB.Query("SELECT * FROM enemies")
	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Panic(err)
	}

	count := len(columns)
	tableData := make([]map[string]string, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		entry := make(map[string]string)

		for i, col := range columns {
			val := values[i]
			if val != nil {
				entry[col] = fmt.Sprintf("%s", string(val.([]byte))) // Save the data as a string
			}
		}

		if len(entry["latitude"]) > 0 && len(entry["latitude"]) > 0 { // Make sure we don't have bad data
			latitude, err := strconv.ParseFloat(entry["latitude"], 64)
			if err == nil {
				longitude, err := strconv.ParseFloat(entry["longitude"], 64)
				if err == nil {
					if withinRadius(latitude, longitude, userLatitude, userLongitude, radius) { // Only return enemies that are close to the player
						tableData = append(tableData, entry)
					}
				} else {
					log.Panic(err)
				}
			} else {
				log.Panic(err)
			}
		}
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		log.Panic(err)
	}

	return string(jsonData), nil
}

func (ag *AccessorGroup) CountNearbyEnemies(userLatitude float64, userLongitude float64, radius float64) (int, error) {
	enemyCount := 0

	rows, err := ag.DB.Query("SELECT * FROM enemies")
	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Panic(err)
	}

	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		entry := make(map[string]string)

		for i, col := range columns {
			val := values[i]
			if val != nil {
				entry[col] = fmt.Sprintf("%s", string(val.([]byte))) // Save the data as a string
			}
		}

		if len(entry["latitude"]) > 0 && len(entry["latitude"]) > 0 { // Make sure we don't have bad data
			latitude, err := strconv.ParseFloat(entry["latitude"], 64)
			if err == nil {
				longitude, err := strconv.ParseFloat(entry["longitude"], 64)
				if err == nil {
					if withinRadius(latitude, longitude, userLatitude, userLongitude, radius) { // Only return enemies that are close to the player
						enemyCount++
					}
				} else {
					log.Panic(err)
				}
			} else {
				log.Panic(err)
			}
		}
	}

	return String(enemyCount), nil
}

func withinRadius(lat1 float64, lon1 float64, lat2 float64, lon2 float64, radius float64) bool {
	p := geo.NewPoint(lat1, lon1)
	p2 := geo.NewPoint(lat2, lon2)

	dist := p.GreatCircleDistance(p2) // Find the great circle distance between points

	if dist < radius { // Return whether we're inside the radius or not
		return true
	} else {
		return false
	}
}
