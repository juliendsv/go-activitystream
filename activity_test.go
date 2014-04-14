package activitystream

import (
	"strconv"
	"testing"
	"time"

	uuid "code.google.com/p/go-uuid/uuid"
)

func TestSimpleActivity(t *testing.T) {

	id := uuid.New()
	now := time.Now().Unix()

	activity := Activity{
		Id:   id,
		Verb: "post",
		Actor: &Object{
			Url:         "http://twitter.com/juliendsv",
			DisplayName: "juliendsv",
			ObjectType:  "person",
		},
		Object: &Object{
			Url: "https://twitter.com/juliendsv/status/454965935745814528",
		},
		Target: &Object{
			Url:        "https://twitter.com/juliendsv/",
			ObjectType: "tweet",
			Id:         "https://twitter.com/juliendsv/status/454965935745814528",
		},
		Published: now,
	}

	json, err := activity.AsJson()
	if err != nil {
		t.Fatal("Error marshalling activity: %v, error: %v", activity, err)
	}

	expected := "{\"Id\":\"" + id + "\",\"Verb\":\"post\",\"Actor\":{\"Url\":\"http://twitter.com/juliendsv\",\"DisplayName\":\"juliendsv\",\"ObjectType\":\"person\"},\"Object\":{\"Url\":\"https://twitter.com/juliendsv/status/454965935745814528\"},\"Target\":{\"Id\":\"https://twitter.com/juliendsv/status/454965935745814528\",\"Url\":\"https://twitter.com/juliendsv/\",\"ObjectType\":\"tweet\"},\"Published\":" + strconv.Itoa(int(now)) + "}"

	if string(json) != expected {
		t.Fatal("we were expecting: %s, got: %s", expected, string(json))
	}
}
