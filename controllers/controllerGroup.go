package controllers

import "github.com/jessemillar/rytsar/accessors"

// ControllerGroup holds all config information for the controllers
type ControllerGroup struct {
	Accessors *accessors.AccessorGroup
}
