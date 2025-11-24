package utils

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ToPrettyString(msg proto.Message) (string, error) {
	m := protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}

	out, err := m.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
