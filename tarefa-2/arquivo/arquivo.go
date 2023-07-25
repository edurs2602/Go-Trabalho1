package arquivo

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func AbrirArquivo() *os.File {
	// Abrindo o arquivo CSV
	file, err := os.Open("../../Nascimentos_RN.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo ", err)
		log.Fatal(err)
		return nil
	}

	return file
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

	}

	print("")
	return table
}

func AbrirArquivoDat() []string {

	// Abrindo o arquivo dat
	data, err := os.ReadFile("../main/alvos.dat")
	if err != nil {
		panic(err)
	}

	str := string(data)
	lines := strings.Split(str, "\n")

	return lines
}

func EscreverCSV(NascimentosCod []int) *os.File {
	// Criando arquivo CSV
	anos := []int{1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020}

	filecsv, err := os.Create("nascimentos-alvos.csv")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo", err)
		return nil
	}
	defer func(filecsv *os.File) {
		err := filecsv.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o arquivo", err)
		}
	}(filecsv)

	writer := csv.NewWriter(filecsv)

	// Escrevendo o cabeçalho do arquivo
	header := []string{"Ano"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Erro ao escrever o cabeçalho do arquivo", err)
		return nil
	}

	// Escrevendo dados para o arquivo
	for i := 0; i < len(anos); i++ {
		ano := strconv.Itoa(anos[i])
		var nascimentos []string
		for _, n := range NascimentosCod {
			nascimentos = append(nascimentos, strconv.Itoa(n))
		}

		record := []string{ano}
		for _, n := range nascimentos {
			record = append(record, n)
		}
		err = writer.Write(record)
		if err != nil {
			fmt.Println("Erro ao escrever a linha no arquivo", err)
			return nil
		}
	}

	writer.Flush()

	return filecsv
}

func EscreverDat(alvos []int) *os.File {
	// Criando arquivo dat a partir do CSV
	csvFile := EscreverCSV(alvos)

	csvContent, err := ioutil.ReadFile(csvFile.Name())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	datFile, err := os.Create("nascimentos-alvos.dat")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(datFile *os.File) {
		err := datFile.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(datFile)

	_, err = datFile.Write(csvContent)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Arquivo nascimentos-alvos.dat criado com sucesso!")
	return nil
}
