package main

import (
	"fmt"
)

func main() {
	var (
		nums []int
	)
	fmt.Println(nums, len(nums), cap(nums)) //caps é para indicar a capacidade

	nums = make([]int, 5) //tipos primitivos são sempre inicializados
	fmt.Println(nums, len(nums), cap(nums))

	paises := []string{"Brasil", "França"}
	fmt.Println(paises, len(paises), cap(paises))

	paises = append(paises, "Uruguai")
	fmt.Println(paises, len(paises), cap(paises))

	cidades := make([]string, 5)
	cidades[0] = "Tokio"
	cidades[1] = "Londes"
	cidades[2] = "Brasilia"
	cidades[3] = "Nova York"
	fmt.Println(cidades, len(cidades), cap(cidades))

	cidades = append(cidades, "Singapura")
	fmt.Println(cidades, len(cidades), cap(cidades))

	cidades[4] = "Gothan"
	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {
		fmt.Println("no indice", indice, "temos a cidade", cidade)
	}

	cidades[1] = "Xing ling"
	capitaisAsiaticas := cidades[0:2] // na segunda coluna ele não conta a partir do 0, e sim a partir do 1
	fmt.Println(capitaisAsiaticas, len(capitaisAsiaticas), cap(capitaisAsiaticas))
	// ou como quero pegar do inicio eu posso ignorar o primeiro termo
	capitaisAsiaticas = cidades[:2]
	fmt.Println(capitaisAsiaticas, len(capitaisAsiaticas), cap(capitaisAsiaticas))
	//para o contrário....
	capitaisAsiaticas = cidades[3:]
	fmt.Println(capitaisAsiaticas, len(capitaisAsiaticas), cap(capitaisAsiaticas))

	// para remover um item do slice...
	indiceDoItem := 3
	sliceTemp := cidades[:indiceDoItem]
	println("Vou remover o item: ", cidades[indiceDoItem])
	fmt.Println(sliceTemp, len(sliceTemp), cap(sliceTemp))
	sliceTemp = append(sliceTemp, cidades[indiceDoItem+1:]...)
	copy(cidades, sliceTemp)
	fmt.Println(cidades, len(cidades), cap(cidades))
}
