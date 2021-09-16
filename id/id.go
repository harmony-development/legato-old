package id

type Generator interface {
	NextID() (uint64, error)
}
