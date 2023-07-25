package main

import (
	"fmt"
	"generics"
	"reader"
)

func main() {
	table := reader.LerArquivo()
	//fmt.Println(table)

	maioresNascimentosPorAno := generics.MaiorNumeroDeNascimentosPorAno(table)
	//fmt.Println(maioresNascimentosPorAno)

	menoresNascimentosPorAno := generics.MenorNumeroDeNascimentosPorAno(table)
	//fmt.Println(menoresNascimentosPorAno)

	mediasNascimentosPorAno := generics.MediaPorAno(table)
	//fmt.Println(mediasNascimentosPorAno)

	desviosPadraoPorAno := generics.DesvioPadraoPorAno(table)

	totaisNascimentosPorAno := generics.TotalDeNascimentosPorAno(table)
	//fmt.Println(totaisNascimentosPorAno)

	file := reader.EscreverDat(totaisNascimentosPorAno)
	fmt.Println(file)

	reader.GerarEstatisticasCSV(maioresNascimentosPorAno, menoresNascimentosPorAno, totaisNascimentosPorAno, mediasNascimentosPorAno, desviosPadraoPorAno)
}
