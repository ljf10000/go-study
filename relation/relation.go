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
	if v, ok := me[id]; ok {
		return v
	} else {
		return nil
	}
}

func (me Relation) Insert(member *Member) error {
	id := member.ID

	if nil != me.Get(id) {
		return errors.New(fmt.Sprintf("%s exist", id))
	} else if "" != member.FatherID {
		if father := me.Get(member.FatherID); nil != father {
			father.Children = append(father.Children, member)
		}
	}

	me[id] = member

	return nil
}

func (me Relation) MultiInsert(members []*Member) error {
	for _, member := range members {
		if err := me.Insert(member); nil != err {
			return err
		}
	}

	return nil
}
