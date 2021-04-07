package network

import (
	"HDTwG/internal/store"
	Stores "HDTwG/internal/store"
	"HDTwG/model"
	"archive/zip"
	"context"
	"io/ioutil"
	"strings"
)

type PutCmd func(ctx context.Context) error

func Put(stores ...Stores.Store) PutCmd {
	return func(ctx context.Context) error {
		var translations store.Translations
		var locations []model.Location
		translations, locations, err := unzip("ressources/IP-locations.zip")
		for _, store := range stores {
			err = store.Put(ctx, translations, locations)
			if err != nil {
				//TODO error models
				/*if err != model.ErrNotFound {
					logrus.Error(err)
				}*/
			}
		}
		return err
	}
}

func readAll(file *zip.File) ([]byte, error) {
	// Open File and check error
	fc, err := file.Open()
	if err != nil {
		return nil, err
	}

	//Close file at the end
	defer func() {
		err := fc.Close()
		if err != nil {
			panic(err)
		}
	}()

	//Read all file and check error
	content, err := ioutil.ReadAll(fc)
	if err != nil {
		return nil, err
	}

	return content, err
}

// Fill database model with unziped file
func createObject(file *zip.File) []model.Translation {
	var translation []model.Translation
	rawBytes, _ := readAll(file)
	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines[1:] {
		values := strings.Split(line, ";")
		if len(values) > 5 {
			translation = append(translation,
				model.Translation{
					UUID:       values[0],
					Continent:  values[1],
					Country:    values[2],
					Region:     values[3],
					Department: values[4],
					City:       values[5],
				},
			)
		}
	}
	return translation
}

func unzip(src string) (store.Translations, []model.Location, error) {
	var translations store.Translations
	var locations []model.Location

	//Unzip ressource folder
	r, err := zip.OpenReader(src)
	if err != nil {
		return store.Translations{}, []model.Location{}, err
	}

	//Close folder at the end
	defer r.Close()

	//Loop through each file
	for _, file := range r.File {
		if file.Name == "IP-locations/IP-locations.csv" {
			rawBytes, _ := readAll(file)
			lines := strings.Split(string(rawBytes), "\n")
			for _, line := range lines[1:] {
				values := strings.Split(line, ",")
				if len(values) > 1 {
					locations = append(locations,
						model.Location{
							UUID:    values[1],
							Address: values[0],
						},
					)
				}
			}
		} else if strings.Contains(file.Name, "IP-locations/Locations-FR") {
			translations.TranslationFR = createObject(file)
		} else if strings.Contains(file.Name, "IP-locations/Locations-EN") {
			translations.TranslationEN = createObject(file)
		} else if strings.Contains(file.Name, "IP-locations/Locations-ES") {
			translations.TranslationES = createObject(file)
		}
	}
	return translations, locations, err
}
