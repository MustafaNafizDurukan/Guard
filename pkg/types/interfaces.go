package types

type IInjector interface {
	Init(uint32)
	Descriptor() *Descriptor
	Inject(string) error
}

type IHunter interface {
	Init()
	Descriptor() *Descriptor
	Hunt() error
}
