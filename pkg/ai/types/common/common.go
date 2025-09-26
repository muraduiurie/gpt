package common

type Responser interface {
	Unmarshal(b []byte) error
}

type Requester interface {
	Marshal() ([]byte, error)
}

type Response struct {
	TextResponse Responser
}

type Request struct {
	TextRequest Requester
}
