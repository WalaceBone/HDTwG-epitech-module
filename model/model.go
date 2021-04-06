package model

type Location struct {
	UUID    string `json:"uuid" db:"uuid"`
	Address string `json:"network" db:"network"`
}

type Translation struct {
	UUID       string `json:"uuid" db:"uuid"`
	Continent  string `json:"continent" db:"continent"`
	Country    string `json:"country" db:"country"`
	Region     string `json:"region" db:"region"`
	Department string `json:"department" db:"department"`
	City       string `json:"city" db:"city"`
}

type TranslationES struct {
	UUID       string `json:"uuid" db:"uuid"`
	Continent  string `json:"continent" db:"continent"`
	Country    string `json:"country" db:"country"`
	Region     string `json:"region" db:"region"`
	Department string `json:"department" db:"department"`
	City       string `json:"city" db:"city"`
}

type TranslationEN struct {
	UUID       string `json:"uuid" db:"uuid"`
	Continent  string `json:"continent" db:"continent"`
	Country    string `json:"country" db:"country"`
	Region     string `json:"region" db:"region"`
	Department string `json:"department" db:"department"`
	City       string `json:"city" db:"city"`
}
