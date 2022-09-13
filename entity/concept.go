/*
法律概念实体
*/
package entity

type Concept struct {
	Name       string
	Definition string
}

func NewConcept() *Concept {
	return &Concept{
		Name:       "",
		Definition: "",
	}
}
