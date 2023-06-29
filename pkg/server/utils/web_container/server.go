package web_container

type WebContainer interface {
	RunWebContainer(addr string) error
}
