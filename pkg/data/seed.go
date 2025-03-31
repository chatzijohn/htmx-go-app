package data

import "my-app/models"

var Countries = []models.Country{
	{
		Name:      "United States",
		Code:      "US",
		Capital:   "Washington, D.C.",
		Continent: "North America",
	},
	{
		Name:      "Canada",
		Code:      "CA",
		Capital:   "Ottawa",
		Continent: "North America",
	},
	{
		Name:      "France",
		Code:      "FR",
		Capital:   "Paris",
		Continent: "Europe",
	},
	{
		Name:      "Japan",
		Code:      "JP",
		Capital:   "Tokyo",
		Continent: "Asia",
	},
	{
		Name:      "Greece",
		Code:      "GR",
		Capital:   "Athens",
		Continent: "Europe",
	},
}
