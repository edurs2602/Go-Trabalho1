package estatisticas

import (
	"arquivo"
	"fmt"
	"strconv"
)

func MunicipioPorCodigo(table [][]string) []int {
	codigos := make(map[int][]int)
	var nascimentos []int
	var munCod []int
	municipiosQntd := 0
	munCodLista := arquivo.AbrirArquivoDat()

	// Transforma os elementos da lista em inteiro e os adiciona em uma lista de inteitos
	for i := 1; i < len(munCodLista); i++ {
		codigo, err := strconv.Atoi(munCodLista[i])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		munCod = append(munCod, codigo)
		// Verifica a quantidade de códigos, consequentemente de municipios
		municipiosQntd++
	}

	// Iterar sobre as linhas da tabela
	for i := 1; i < len(table)-1; i++ {
		codigo, _ := strconv.Atoi(table[i][0])

		// Verificação para saber se o código está na lista
		for _, munCod := range munCod {
			if codigo == munCod {
				// Adicionando os nascimentos ao map
				for j := 1; j < len(table[i]); j++ {
					nascimento, _ := strconv.Atoi(table[i][j])
					codigos[codigo] = append(codigos[codigo], nascimento)
				}
				break
			}
		}
	}

	for _, v := range codigos {
		nascimentos = append(nascimentos, v...)
	}

	return nascimentos
}
