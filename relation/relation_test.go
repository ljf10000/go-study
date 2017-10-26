package librelation

import (
	"encoding/json"
	"fmt"
	"testing"
)

func newmember(id, father string) *Member {
	return &Member{
		ID:       id,
		FatherID: father,
	}
}

func newmembers(members []*Member, root int) []*Member {
	return nil
}

func Test1(t *testing.T) {
	members := []*Member{
		newmember("001", ""),
		newmember("011", "001"),
		newmember("111", "011"),
		newmember("211", "011"),
		newmember("311", "011"),

		newmember("021", "001"),
		newmember("121", "021"),
		newmember("221", "021"),
		newmember("321", "021"),

		newmember("031", "001"),
		newmember("131", "031"),
		newmember("231", "031"),
		newmember("331", "031"),

		newmember("002", ""),
		newmember("012", "002"),
		newmember("112", "012"),
		newmember("212", "012"),
		newmember("312", "012"),

		newmember("022", "002"),
		newmember("122", "022"),
		newmember("222", "022"),
		newmember("322", "022"),

		newmember("032", "002"),
		newmember("132", "032"),
		newmember("232", "032"),
		newmember("332", "032"),
	}

	r := &Relation{}

	r.MultiInsert(members)
	r.Build()

	b, err := json.MarshalIndent(r.Get("001"), "", "    ")
	if nil != err {
		t.Error("json error")
	}

	fmt.Println(string(b))
}
