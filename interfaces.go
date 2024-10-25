package onvif

type Function interface {
	Request() any
	Response() any
}
