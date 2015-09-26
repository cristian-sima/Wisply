package controllers

type RemoteRepository interface {
	Validate() bool
	Collect() bool
}
