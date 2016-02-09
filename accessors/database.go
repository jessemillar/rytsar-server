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
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		fmt.Printf("%T %v\n", entry["latitude"], entry["latitude"])

		latitude, err := strconv.ParseFloat(entry["latitude"].(string), 64)
		if err != nil {
			log.Panic(err)
		}

		longitude, err := strconv.ParseFloat(entry["longitude"].(string), 64)
		if err != nil {
			log.Panic(err)
		}

		if withinRadius(latitude, longitude, userLatitude, userLongitude) { // Only return enemies that are close to the player
			tableData = append(tableData, entry)
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
	radius := float64(1000000)

	p := geo.NewPoint(lat1, lon1)
	p2 := geo.NewPoint(lat2, lon2)

	dist := p.GreatCircleDistance(p2) // Find the great circle distance between points

	if dist < radius { // Return whether we're inside the radius or not
		return true
	} else {
		return false
	}
}
