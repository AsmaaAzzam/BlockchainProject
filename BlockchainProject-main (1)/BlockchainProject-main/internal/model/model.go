package model

// Model is an interface that all models should implement
type Model interface {
	GetID() string // Method to get the ID of the model
}
