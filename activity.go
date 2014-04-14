package activitystream

import (
	"encoding/json"
)

type Activity struct {
	Id        string
	Verb      string
	Title     string  `json:",omitempty"`
	Generator *Object `json:",omitempty"`
	Actor     *Object
	Object    *Object
	Target    *Object
	Provider  *Object `json:",omitempty"`
	Published int64
	Updated   int64  `json:",omitempty"`
	Url       string `json:",omitempty"`
}

type Object struct {
	Id          string  `json:",omitempty"`
	Title       string  `json:",omitempty"`
	Url         string  `json:",omitempty"`
	DisplayName string  `json:",omitempty"`
	ObjectType  string  `json:",omitempty"`
	Actor       *Object `json:",omitempty"`
	Object      *Object `json:",omitempty"`
}

func (a Activity) AsJson() ([]byte, error) {
	json, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return json, nil
}
