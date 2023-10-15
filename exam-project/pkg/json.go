package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"playground/newProject/models"
)

func Read(fileName string, obj any) (any, error) {
	var (
		objects any
	)

	switch obj.(type) {
	case []models.Branch:
		objects = obj
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &objects)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return objects, nil
}

func Write(fileName string, obj any) error {

	body, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return err
}
