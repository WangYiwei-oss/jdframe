package jdft

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc interface{}
}

type RoutesInfo []RouteInfo
