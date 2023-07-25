package main

import (
	"arquivo"
	"estatisticas"
)

func main() {
	// Instanciando variavel
	table := arquivo.LerArquivo()

	NascimentosPorMunicipio := estatisticas.MunicipioPorCodigo(table)

	arquivo.EscreverDat(NascimentosPorMunicipio)
}
