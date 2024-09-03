package http

type HTTPServer interface {
	Run(add string) error
}
