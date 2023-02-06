package codetype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type Image struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (Image) TableName() string {
	return "images"
}

// Scan nhận vào json value, implement sql.Scanner interface
func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img

	return nil
}

// Value return json value, implement driver.Valuer interface
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type Images []Image

// Scan nhận vào json value, implement sql.Scanner interface
func (i *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img

	return nil
}

// Value return json value, implement driver.Valuer interface
func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}
