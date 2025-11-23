package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type ID xid.ID

func NewID() ID {
	return ID(xid.New())
}

func IDFromString(s string) (ID, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return ID{}, err
	}
	return ID(id), err
}

func IDFromBytes(bs []byte) (ID, error) {
	id, err := xid.FromBytes(bs)
	if err != nil {
		return ID{}, err
	}
	return ID(id), err
}

func (id ID) String() string {
	return xid.ID(id).String()
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (id *ID) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		newID, err := IDFromString(v)
		if err != nil {
			return errors.WithStack(err)
		}
		*id = newID
		return nil

	case []byte:
		newID, err := IDFromBytes(v)
		if err != nil {
			return errors.WithStack(err)
		}
		*id = newID
		return nil
	}
	return fmt.Errorf("%T is not a ID", v)

}

// MarshalGQL implements the graphql.Marshaler interface
func (id ID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(id.String()))
}
