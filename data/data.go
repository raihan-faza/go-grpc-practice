package data

type User struct {
	UserID   int64
	UserName string
	UserAge  int64
}

var DummyData = []User{{1, "user1", 12}, {2, "user2", 22}}
var MappedDummyData = make(map[int]User)
