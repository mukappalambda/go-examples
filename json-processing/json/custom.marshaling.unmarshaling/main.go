package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Active bool     `json:"is_active"`
	Emails []string `json:"emails"`
	Desc   string   `json:"desc"`
}

var _ json.Marshaler = (*Data)(nil)

func (d *Data) MarshalJSON() ([]byte, error) {
	if d.Desc == "" {
		d.Desc = "null description"
	}
	return json.Marshal(*d)
}

var _ json.Unmarshaler = (*Data)(nil)

func (d *Data) UnmarshalJSON(data []byte) error {
	type tmpType Data
	var out tmpType
	if err := json.Unmarshal(data, &out); err != nil {
		return err
	}
	if out.Desc == "null description" {
		out.Desc = ""
	}
	*d = Data(out)
	return nil
}

func main() {
	// marshaling
	firstData := &Data{
		Id:     1,
		Name:   "alpha",
		Active: true,
		Emails: []string{"alpha@email.com", "alpha.backup@email.com"},
	}
	buf, err := json.Marshal(firstData)
	if err != nil {
		log.Fatalf("error marshalling data: %s\n", err)
	}
	fmt.Println(string(buf))

	//unmarshaling
	var secondData Data
	err = json.Unmarshal(buf, &secondData)
	if err != nil {
		log.Fatalf("error unmarshalling byte slice: %s\n", err)
	}
	fmt.Printf("%+v\n", secondData)
}
