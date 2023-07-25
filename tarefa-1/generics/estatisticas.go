package generics

import (
	"fmt"
	"math"
	"strconv"
)

func Verifica_Total_Ano(table [][]string) bool {

	for i := 1; i < len(table[0])-1; i++ {
		sum := 0
		for j := 1; j < len(table)-1; j++ {
			//Iterando linha a linha da tabela para realizar soma dos valores de cada ano
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])
				sum += value
			}
		}

		total, _ := strconv.Atoi(table[len(table)-1][i])
		if sum != total {
			fmt.Println("O total do ano", table[0][i], "não condiz com o esperado!")
			return false
		}
	}

	fmt.Println("Todos os resultados condizem com o esperado")
	return true
}

func Verifica_Total_Municipio(line []string) bool {
	sum := 0
	//Percorrendo linha
	for i := 1; i < len(line)-1; i++ {
		if len(line[i]) > 0 {
			//Somatório de todos os nascimentos
			value, _ := strconv.Atoi(line[i])
			sum += value
		}
	}
	//Total é igual o último valor da linha
	total, _ := strconv.Atoi(line[len(line)-1])

	if sum != total {
		fmt.Println("O total do município", line[0], "não condiz com o esperado!")
		return false
	} else {
		return true
	}

}

// Função responsável por encontrar o maior número de nascimento em cada ano
func MaiorNumeroDeNascimentosPorAno(table [][]string) []int {
	maioresNumNascimentos := []int{}

	for i := 1; i < len(table[0])-1; i++ {
		maiorNumNascimento := 0
		for j := 1; j < len(table)-1; j++ {
			//Iterando linha a linha da tabela para conferir qual o maior número de nascimento em cada ano
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])

				if value > maiorNumNascimento {
					maiorNumNascimento = value
				}
			}
		}

		maioresNumNascimentos = append(maioresNumNascimentos, maiorNumNascimento)
	}

	return maioresNumNascimentos
}

// Função responsável por encontrar o menor número de nascimentos em cada ano
func MenorNumeroDeNascimentosPorAno(table [][]string) []int {
	menoresNumNascimentos := []int{}

	for i := 1; i < len(table[0])-1; i++ {
		var menorNumNascimento int
		for j := 1; j < len(table)-1; j++ {
			//Iterando linha a linha da tabela para conferir qual o maior número de nascimento em cada ano
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])

				if menorNumNascimento == 0 {
					menorNumNascimento = value
					continue
				}

				if value < menorNumNascimento {
					menorNumNascimento = value
				}
			}
		}
		menoresNumNascimentos = append(menoresNumNascimentos, menorNumNascimento)
	}

	return menoresNumNascimentos
}

// Função responsável por calcular a média de nascimentos em cada ano
func MediaPorAno(table [][]string) []float32 {
	mediasNascimentos := []float32{}

	for i := 1; i < len(table[0])-1; i++ {
		sum := 0
		for j := 1; j < len(table)-1; j++ {
			//Iterando linha a linha da tabela para realizar o somatório
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])
				sum += value
			}
		}
		//Variável para receber a média
		var media float32
		media = float32(sum) / float32(len(table)-2)
		mediasNascimentos = append(mediasNascimentos, media)
	}
	return mediasNascimentos
}

// Função responsável por somar o total de nascimentos em cada ano
func TotalDeNascimentosPorAno(table [][]string) []int {
	totaisNascimentos := []int{}

	for i := 1; i < len(table[0])-1; i++ {
		sum := 0
		for j := 1; j < len(table)-1; j++ {
			//Iterando linha a linha da tabela para realizar o somatório
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])
				sum += value
			}
		}
		totaisNascimentos = append(totaisNascimentos, sum)
	}
	return totaisNascimentos
}

func DesvioPadraoPorAno(table [][]string) []float32 {
	desviosPadrao := []float32{}

	for i := 1; i < len(table[0])-1; i++ {
		// Calculando a média
		sum := 0
		for j := 1; j < len(table)-1; j++ {
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])
				sum += value
			}
		}
		media := float32(sum) / float32(len(table)-2)

		//Calculando as diferenças e elevaando ao quadrado
		somatorioDiferencasAoQuadrado := 0.0
		for j := 1; j < len(table)-1; j++ {
			if len(table[j][i]) > 0 {
				value, _ := strconv.Atoi(table[j][i])
				diferenca := float32(value) - media
				somatorioDiferencasAoQuadrado += float64(diferenca * diferenca)
			}
		}

		// Calculando o desvio padrão
		desvioPadrao := float32(math.Sqrt(somatorioDiferencasAoQuadrado / float64(len(table)-2)))
		desviosPadrao = append(desviosPadrao, desvioPadrao)
	}

	return desviosPadrao
}
