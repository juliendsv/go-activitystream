package activitystream

import (
	"encoding/json"
)

type Activitystream struct {
	Stream []Activity
}

func (a Activitystream) ToJson() ([]byte, error) {
	activtyAsjson, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return activtyAsjson, nil
}

func (a Activitystream) FromJson(activities string) error {
	if err := json.Unmarshal([]byte(activities), a); err != nil {
		return err
	}
	return nil
}

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

func (a Activity) ToJsonSteam() ([]byte, error) {
	astream := &Activitystream{
		Stream: make([]Activity, 0),
	}
	astream.Stream = append(astream.Stream, a)
	return astream.ToJson()
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
