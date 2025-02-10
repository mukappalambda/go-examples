package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Item struct {
	ID        int
	Name      string
	Price     int
	CreatedAt time.Time
	Duration  time.Duration
}

var _ json.Marshaler = (*Item)(nil)

func (i *Item) MarshalJSON() ([]byte, error) {
	type Tmp struct {
		ID        int
		Name      string
		Price     int
		CreatedAt string
		Duration  string
	}
	tmp := &Tmp{
		ID:        i.ID,
		Name:      i.Name,
		Price:     i.Price,
		CreatedAt: i.CreatedAt.Format(time.RFC3339),
		Duration:  i.Duration.String(),
	}
	return json.Marshal(tmp)
}

var _ json.Unmarshaler = (*Item)(nil)

func (i *Item) UnmarshalJSON(data []byte) error {
	type Tmp struct {
		ID        int
		Name      string
		Price     int
		CreatedAt string
		Duration  string
	}
	var tmp Tmp
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	createdAt, err := time.Parse(time.RFC3339, tmp.CreatedAt)
	if err != nil {
		return fmt.Errorf("error parsing CreatedAt: %s", err.Error())
	}
	duration, err := time.ParseDuration(tmp.Duration)
	if err != nil {
		return fmt.Errorf("error parsing Duration: %s", err.Error())
	}
	item := Item{
		ID:        tmp.ID,
		Name:      tmp.Name,
		Price:     tmp.Price,
		CreatedAt: createdAt,
		Duration:  duration,
	}
	*i = item
	return nil
}

func main() {
	firstItem := &Item{
		ID:        1,
		Name:      "food",
		Price:     10,
		CreatedAt: time.Now(),
		Duration:  24 * time.Hour,
	}
	data, err := json.Marshal(firstItem)
	if err != nil {
		log.Fatalf("error marshaling item: %s\n", err)
	}
	fmt.Println(string(data))

	var secondItem Item

	err = json.Unmarshal(data, &secondItem)
	if err != nil {
		log.Fatalf("error unmarshaling item: %s\n", err)
	}
	fmt.Printf("%+v\n", secondItem)
}
