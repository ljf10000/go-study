package librelation

import (
	"errors"
	"fmt"
)

type MemberID string

type Member struct {
	ID        MemberID `json:id`
	FatherID  MemberID `json:father`
	OtherInfo string
	Children  []*Member `json:children`
}

type Relation map[MemberID]*Member

func (me Relation) Get(id MemberID) *Member {
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
	}

	if "" != member.FatherID {
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
