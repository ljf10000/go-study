package librelation

import (
	"errors"
	"fmt"
)

type Member struct {
	ID        string `json:id`
	FatherID  string `json:father`
	OtherInfo string
	Children  []*Member `json:children`
}

type Relation map[string]*Member

func (me Relation) Get(id string) *Member {
	if "" == id {
		return nil
	} else if v, ok := me[id]; ok {
		return v
	} else {
		return nil
	}
}

func (me Relation) Insert(member *Member) error {
	id := member.ID

	if "" == id {
		return errors.New("empty id")
	} else if nil != me.Get(id) {
		return errors.New(fmt.Sprintf("%s exist", id))
	} else {
		me[id] = member

		return nil
	}
}

func (me Relation) MultiInsert(members []*Member) error {
	for _, member := range members {
		if err := me.Insert(member); nil != err {
			return err
		}
	}

	return nil
}

func (me Relation) Build() {
	for _, member := range me {
		father := me.Get(member.FatherID)
		if nil != father {
			father.Children = append(father.Children, member)
		}
	}
}
