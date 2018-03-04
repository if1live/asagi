package main

type Response struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Blank   bool   `json:"blank"`
	Message string `json:"message"`
}

func NewEchoResponse(msg string, blank bool) EchoResponse {
	return EchoResponse{
		Message: msg,
		Blank:   blank,
	}
}

func main() {
	mainWithNet()
	//mainWithEvent()
}
