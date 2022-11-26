package server

import (
	"bytes"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/encoding"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"mime/multipart"
	"strings"
)

// Name is the name registered for the json codec.
const Name = "form-data"

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with json.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return []byte{}, nil
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	topLine := strings.Split(string(data), "\r\n")
	boundary := topLine[0][2:]
	writer := multipart.NewReader(bytes.NewReader(data), boundary)

	tmp := make(map[string]interface{})
	for {
		part, err := writer.NextPart()
		if err != nil {
			break
		}
		partByte, err := io.ReadAll(part)
		if err != nil {
			return err
		}
		if part.FormName() == "file" {
			tmp[part.FormName()] = partByte
			tmp["file_name"] = part.FileName()
		} else {
			tmp[part.FormName()] = string(partByte)
		}
	}
	marshal, err := json.Marshal(tmp)
	if err != nil {
		return err
	}

	return json.Unmarshal(marshal, v)
}

func (codec) Name() string {
	return Name
}
