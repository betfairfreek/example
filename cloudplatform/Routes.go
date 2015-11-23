package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunk http.HandlerFunc
}

type Routes []Route
