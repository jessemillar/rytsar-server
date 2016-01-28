package controllers

import "github.com/jessemillar/rytsar-server/accessors"

// ControllerGroup holds all config information for the controllers
type ControllerGroup struct {
	Accessors *accessors.AccessorGroup
}
