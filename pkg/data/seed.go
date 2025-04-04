package data

import "my-app/internal/models"

var Countries = []models.Country{
	// North America (7)
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
		Name:      "Mexico",
		Code:      "MX",
		Capital:   "Mexico City",
		Continent: "North America",
	},
	{
		Name:      "Costa Rica",
		Code:      "CR",
		Capital:   "San José",
		Continent: "North America",
	},
	{
		Name:      "Panama",
		Code:      "PA",
		Capital:   "Panama City",
		Continent: "North America",
	},
	{
		Name:      "Jamaica",
		Code:      "JM",
		Capital:   "Kingston",
		Continent: "North America",
	},
	{
		Name:      "Cuba",
		Code:      "CU",
		Capital:   "Havana",
		Continent: "North America",
	},

	// South America (6)
	{
		Name:      "Brazil",
		Code:      "BR",
		Capital:   "Brasília",
		Continent: "South America",
	},
	{
		Name:      "Argentina",
		Code:      "AR",
		Capital:   "Buenos Aires",
		Continent: "South America",
	},
	{
		Name:      "Colombia",
		Code:      "CO",
		Capital:   "Bogotá",
		Continent: "South America",
	},
	{
		Name:      "Peru",
		Code:      "PE",
		Capital:   "Lima",
		Continent: "South America",
	},
	{
		Name:      "Chile",
		Code:      "CL",
		Capital:   "Santiago",
		Continent: "South America",
	},
	{
		Name:      "Ecuador",
		Code:      "EC",
		Capital:   "Quito",
		Continent: "South America",
	},

	// Europe (12)
	{
		Name:      "France",
		Code:      "FR",
		Capital:   "Paris",
		Continent: "Europe",
	},
	{
		Name:      "Germany",
		Code:      "DE",
		Capital:   "Berlin",
		Continent: "Europe",
	},
	{
		Name:      "Italy",
		Code:      "IT",
		Capital:   "Rome",
		Continent: "Europe",
	},
	{
		Name:      "Spain",
		Code:      "ES",
		Capital:   "Madrid",
		Continent: "Europe",
	},
	{
		Name:      "United Kingdom",
		Code:      "GB",
		Capital:   "London",
		Continent: "Europe",
	},
	{
		Name:      "Greece",
		Code:      "GR",
		Capital:   "Athens",
		Continent: "Europe",
	},
	{
		Name:      "Sweden",
		Code:      "SE",
		Capital:   "Stockholm",
		Continent: "Europe",
	},
	{
		Name:      "Norway",
		Code:      "NO",
		Capital:   "Oslo",
		Continent: "Europe",
	},
	{
		Name:      "Netherlands",
		Code:      "NL",
		Capital:   "Amsterdam",
		Continent: "Europe",
	},
	{
		Name:      "Belgium",
		Code:      "BE",
		Capital:   "Brussels",
		Continent: "Europe",
	},
	{
		Name:      "Portugal",
		Code:      "PT",
		Capital:   "Lisbon",
		Continent: "Europe",
	},
	{
		Name:      "Switzerland",
		Code:      "CH",
		Capital:   "Bern",
		Continent: "Europe",
	},

	// Asia (12)
	{
		Name:      "Japan",
		Code:      "JP",
		Capital:   "Tokyo",
		Continent: "Asia",
	},
	{
		Name:      "China",
		Code:      "CN",
		Capital:   "Beijing",
		Continent: "Asia",
	},
	{
		Name:      "India",
		Code:      "IN",
		Capital:   "New Delhi",
		Continent: "Asia",
	},
	{
		Name:      "South Korea",
		Code:      "KR",
		Capital:   "Seoul",
		Continent: "Asia",
	},
	{
		Name:      "Thailand",
		Code:      "TH",
		Capital:   "Bangkok",
		Continent: "Asia",
	},
	{
		Name:      "Vietnam",
		Code:      "VN",
		Capital:   "Hanoi",
		Continent: "Asia",
	},
	{
		Name:      "Indonesia",
		Code:      "ID",
		Capital:   "Jakarta",
		Continent: "Asia",
	},
	{
		Name:      "Malaysia",
		Code:      "MY",
		Capital:   "Kuala Lumpur",
		Continent: "Asia",
	},
	{
		Name:      "Singapore",
		Code:      "SG",
		Capital:   "Singapore",
		Continent: "Asia",
	},
	{
		Name:      "Philippines",
		Code:      "PH",
		Capital:   "Manila",
		Continent: "Asia",
	},
	{
		Name:      "Turkey",
		Code:      "TR",
		Capital:   "Ankara",
		Continent: "Asia",
	},
	{
		Name:      "Saudi Arabia",
		Code:      "SA",
		Capital:   "Riyadh",
		Continent: "Asia",
	},

	// Africa (8)
	{
		Name:      "Egypt",
		Code:      "EG",
		Capital:   "Cairo",
		Continent: "Africa",
	},
	{
		Name:      "South Africa",
		Code:      "ZA",
		Capital:   "Pretoria",
		Continent: "Africa",
	},
	{
		Name:      "Nigeria",
		Code:      "NG",
		Capital:   "Abuja",
		Continent: "Africa",
	},
	{
		Name:      "Kenya",
		Code:      "KE",
		Capital:   "Nairobi",
		Continent: "Africa",
	},
	{
		Name:      "Morocco",
		Code:      "MA",
		Capital:   "Rabat",
		Continent: "Africa",
	},
	{
		Name:      "Ghana",
		Code:      "GH",
		Capital:   "Accra",
		Continent: "Africa",
	},
	{
		Name:      "Ethiopia",
		Code:      "ET",
		Capital:   "Addis Ababa",
		Continent: "Africa",
	},
	{
		Name:      "Tanzania",
		Code:      "TZ",
		Capital:   "Dodoma",
		Continent: "Africa",
	},

	// Oceania (5)
	{
		Name:      "Australia",
		Code:      "AU",
		Capital:   "Canberra",
		Continent: "Oceania",
	},
	{
		Name:      "New Zealand",
		Code:      "NZ",
		Capital:   "Wellington",
		Continent: "Oceania",
	},
	{
		Name:      "Fiji",
		Code:      "FJ",
		Capital:   "Suva",
		Continent: "Oceania",
	},
	{
		Name:      "Papua New Guinea",
		Code:      "PG",
		Capital:   "Port Moresby",
		Continent: "Oceania",
	},
	{
		Name:      "Samoa",
		Code:      "WS",
		Capital:   "Apia",
		Continent: "Oceania",
	},
}
