= Activitystream

ActivityStreams Official Site (activitystrea.ms/) JSON Activity Streams 1.0 (activitystrea.ms/specs/json/1.0/)

== Installation

go get github.com/juliendsv/go-activitysteam

== Resources

ActivityStreams Official Site (http://activitystrea.ms/)
JSON Activity Streams 1.0 (http://activitystrea.ms/specs/json/1.0/)

== Examples

=== Single Activity

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


== TODO

* validator
* activity collection
