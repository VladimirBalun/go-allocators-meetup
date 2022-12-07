package allocation

type Data struct {
	iValue int
	sValue string
	bValue bool
}

//go:noinline
func NewDataByValue() Data {
	return Data{iValue: 100, sValue: "100", bValue: true}
}

//go:noinline
func NewDataByPointer() *Data {
	return &Data{iValue: 100, sValue: "100", bValue: true}
}
