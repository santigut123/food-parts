package foodParser

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"github.com/santigut123/food-parts/fpStructs"
)

type RDAParser struct {
	fileName string
	fileType string
}
func GetReferenceRDA() *fpStructs.RDA{
	RDAP := RDAParser{
		fileName: "ReferenceRDA",
		fileType: "../assets/rdi.txt",
	}
	return 	RDAP.ReadRDA()

}
func NewRDAParser(fileName string, fileType string) *RDAParser {
	// Eventually you will be able to change edit RDA properties in order
	// to account for macro and micro-nutirient differences in various diets

	nRDAP := RDAParser{
		fileName: fileName,
		fileType: fileType,
	}
	return &nRDAP
}
func (RDAP *RDAParser) parseRDAFile(rda *fpStructs.RDA) {
	rdaFile, err := os.Open(RDAP.fileName)
	if err != nil {
		log.Println("Did not the RDA File")
		panic(err)
	}
	// The RDA file is stored in CSV format
	rdaReader := csv.NewReader(rdaFile)
	var NType rune
	var typeString string
	var NName string
	var NMassTemp float64
	var NMass float32

	for {
		line, err := rdaReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("Problems reading RDA file %s", err)
		}
		typeString = line[0]
		switch typeString {
		case "Vitamin":
			NType = 'v'
		case "Mineral":
			NType = 'm'
		case "AminoAcid":
			NType = 'a'
		case "Macro":
			NType = 'M'
		default:
			NType = rune(typeString[0])
		}
		NName = line[1]
		NMassTemp,_= strconv.ParseFloat(line[2],32)
		NMass= float32(NMassTemp)
		rda.UpdateNutrient(NName, NType, NMass)
	}

}
func (rp *RDAParser) ReadRDA() *fpStructs.RDA{
	newRDA := fpStructs.RDA{
		RDANutri: fpStructs.Nutrients{
			Vitamins:   map[string]*fpStructs.Nutrient{},
			Minerals:   map[string]*fpStructs.Nutrient{},
			AminoAcids: map[string]*fpStructs.Nutrient{},
			Macros:     map[string]*fpStructs.Nutrient{},
		},
	}
	rp.parseRDAFile(&newRDA)
	return &newRDA
}
