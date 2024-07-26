package main

type Handler interface {
	Handle() error
}

type S struct {
	//
}

// interface guard
var _ Handler = (*S)(nil)

func (s S) Handle() error {
	return nil
}

func main() {
	_ = doWork(S{})
}

type T struct {
	Payload string
}

func doWork(h Handler) T {
	err := h.Handle()
	if err != nil {
		return T{Payload: "failed"}
	}
	return T{Payload: "done"}
}
