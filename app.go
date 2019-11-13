package main

import "github.com/gorilla/mux"

// App struct
type App struct {
	Router *mux.Router
}

// Initialize function
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}

// Run function
func (a *App) Run(addr string) {}
