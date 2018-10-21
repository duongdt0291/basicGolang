package test

type WorkerUnexportedField struct {
	name string
	age  int
}

type Worker struct {
	Name     string
	Position string
	Address  Address
}

type Address struct {
	Street string
	City   string
}

type Worker2 struct {
	Name, Position string
	FamilyMember
}

type FamilyMember struct {
	Wife, Daughter string
}
