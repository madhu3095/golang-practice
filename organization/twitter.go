package organization

import (
	"awesomeProject/organization"
)

func main() {
	p :=organization.NewPerson( "madhusmita", "panigrahi")
	println(p.ID())
	println(p.FullName())
}
