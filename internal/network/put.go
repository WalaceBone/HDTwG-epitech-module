package network

import (
	Stores "HDTwG/internal/store"
	"archive/zip"
	"context"
	"fmt"
	"io/ioutil"
	"strings"
)

type PutCmd func(ctx context.Context) error

func Put(stores ...Stores.Store) PutCmd {
	return func(ctx context.Context) error {
		err := unzip("ressources/IP-locations.zip")
		for _, store := range stores {
			err = store.Put(ctx)
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

func unzip(src string) error {
	//Unzip ressource folder
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	//Close folder at the end
	defer r.Close()

	//Loop through each file
	for _, file := range r.File {
		if file.Name == "IP-locations/IP-locations.csv" {
			rawBytes, _ := readAll(file)
			lines := strings.Split(string(rawBytes[13:]), ",")
			for i, line := range lines {
				if i < 20 {
					fmt.Println(line)
				} else {
					break
				}
			}
		} else if strings.Contains(file.Name, "IP-locations/Locations-") {
			rawBytes, _ := readAll(file)
			lines := strings.Split(string(rawBytes[47:]), ";")
			for n, line := range lines {
				if n < 20 {
					fmt.Println(line)
				} else {
					break
				}
			}
		}
	}
	return err
}
