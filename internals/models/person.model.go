package models

type Person struct {
	Name, Address, Phone string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name: name,
		Address: address,
		Phone: phone,
	}
}

func (p *Person) Print() string {
	return "Name: " + p.Name + "\nAddress: " + p.Address + "\nPhone: " + p.Phone
}

func (p *Person) Greet() string {
	return "Hello " + p.Name + " 👋"
}

func (p *Person) SetNamePerson(name string) {
	p.Name = name
}