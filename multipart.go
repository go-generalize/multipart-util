package multipartutil

import (
	"io"
	"net/textproto"
)

// MultipartPayload is a payload of multipart/form-data for api_gen
type MultipartPayload struct {
	Filename   string
	Data       io.Reader
	MIMEHeader textproto.MIMEHeader
}
