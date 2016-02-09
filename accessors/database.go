package accessors

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/kellydunn/golang-geo"
)

// Returns an array of all loot locations and values to plot on the map in iOS
func (ag *AccessorGroup) DumpDatabase(userLatitude float64, userLongitude float64) (string, error) {
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
				entry[col] = fmt.Sprintf("%s", string(val.([]byte)))
			}
		}

		fmt.Printf("%T %v\n", entry["latitude"], entry["latitude"])

		if len(entry["latitude"]) > 0 && len(entry["latitude"]) > 0 {
			latitude, err := strconv.ParseFloat(entry["latitude"], 64)
			if err == nil {
				longitude, err := strconv.ParseFloat(entry["longitude"], 64)
				if err == nil {
					if withinRadius(latitude, longitude, userLatitude, userLongitude) { // Only return enemies that are close to the player
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

	fmt.Println(string(jsonData))
	return string(jsonData), nil
}

func withinRadius(lat1 float64, lon1 float64, lat2 float64, lon2 float64) bool {
	radius := float64(1000)

	p := geo.NewPoint(lat1, lon1)
	p2 := geo.NewPoint(lat2, lon2)

	dist := p.GreatCircleDistance(p2) // Find the great circle distance between points

	if dist < radius { // Return whether we're inside the radius or not
		return true
	} else {
		return false
	}
}
