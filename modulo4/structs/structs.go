package main

import "fmt"

/* Structs:
- Forma de criar sua própria estrutura de dados
- Personalizar de acordo com a sua necessidade
- Podemos usar vários tipos diferentes
*/

// type <nome da nossa estrutura> struct { <campos> }
type Pessoa struct {
	Nome string
	Sobrenome string
	Idade int
}

func main(){
	fmt.Println(Pessoa{Nome: "Mariana", Sobrenome: "Orsi", Idade: 29})

	pessoa1 := Pessoa{Nome: "Bob", Idade: 3}
	fmt.Println(pessoa1.Nome)
	fmt.Println(pessoa1.Idade)

	pessoa2 := Pessoa{Nome: "Eliana", Idade: 50}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, pessoa1, pessoa2)
	fmt.Println(pessoas)



	// --- structs + map ---

	// alunos := map[string][]Pessoa{}
	// alunos["programação"] = pessoas
	// fmt.Println(alunos)

	// var alunos = map[string][]Pessoa{
	// 	"Programação": {{Nome: "Mariana", Idade: 29}, {Nome: "Salem", Idade: 2}},
	// 	"Engenharia":  {{Nome: "Mariana2", Idade: 29}, {Nome: "Salem2", Idade: 2}},
	// }
	// fmt.Println(alunos)

	// struct herdando campos de outra struct
	// prof := Profissao{p2, "dev"}
	// fmt.Println(prof)
	// fmt.Println(prof.Pessoa.Nome)
	// fmt.Println(prof.Pessoa.Idade)
	// fmt.Println(prof.Tipo)
}
