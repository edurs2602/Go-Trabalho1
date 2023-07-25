package reader

import (
	"encoding/csv"
	"fmt"
	"generics"
	"io"
	"os"
	"strconv"
)

func AbrirArquivo() *os.File {
	// Abrindo o arquivo CSV
	file, err := os.Open("../../Nascimentos_RN.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo ", err)
		return nil
	}

	return file
}

func FecharArquivo(file *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o arquivo ", err)
			return
		}
	}(file)
}

func LerArquivo() [][]string {
	//Lendo o arquivo CSV

	reader := csv.NewReader(AbrirArquivo()) //Leitor de CSV

	table := [][]string{}

	//Ler as linhas do arquivo
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Erro ao ler arquivo ", err)
			return nil
		}
		table = append(table, line)

		// Verifica se é o cabeçalho ou a última linha
		if line[0] == "Município" || line[0] == "Total" {
			continue
		}

		//Verificando se o total de nacimentos em cada munípio realmente corresponde ao somatório do número de nascidos
		//TODO Qual o comportamento esperado quando não for?
		generics.Verifica_Total_Municipio(line)
	}
	generics.Verifica_Total_Ano(table)
	print("")
	return table
}

func EscreverDat(totais []int) *os.File {
	anos := []int{1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020}

	file, err := os.Create("totais.dat")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erro ao fechar arquivo", err)
		}
	}(file)

	for i := 0; i < len(anos); i++ {
		_, err := fmt.Fprintf(file, "%d %d\n", anos[i], totais[i])
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	fmt.Println("Arquivo totais.dat criado com sucesso!")
	return file
}

func GerarEstatisticasCSV(maiorNascimentos []int, menorNascimentos []int, totalNascimentos []int, mediasNascimentos []float32, desviosPadrao []float32) {
	anos := []int{1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020}
	file, err := os.Create("estatisticas.csv")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o arquivo")
		}
	}(file)

	writer := csv.NewWriter(file)

	// Escrevendo o cabeçalho do arquivo
	header := []string{"Ano", "Maior Nascimentos", "Menor Nascimentos", "Total Nascimentos", "Media Nascimentos", "Desvio Padrao Nascimentos"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Erro ao escrever o cabeçalho do arquivo", err)
		return
	}

	// Escrevendo as estatísticas de cada ano
	for i := 0; i < len(anos); i++ {
		ano := strconv.Itoa(anos[i])
		maior := strconv.Itoa(maiorNascimentos[i])
		menor := strconv.Itoa(menorNascimentos[i])
		total := strconv.Itoa(totalNascimentos[i])
		media := strconv.FormatFloat(float64(mediasNascimentos[i]), 'f', 2, 32)
		desvio := strconv.FormatFloat(float64(desviosPadrao[i]), 'f', 2, 32)

		record := []string{ano, maior, menor, total, media, desvio}
		err = writer.Write(record)
		if err != nil {
			fmt.Println("Erro ao escrever a linha no arquivo", err)
			return
		}
	}

	writer.Flush()

	fmt.Println("Arquivo estatisticas.csv gerado com sucesso!")
}
