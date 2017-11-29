package main

import (
	"os"
  "encoding/csv"
  "fmt"
  "bufio"
  "io"
  "strings"
  "strconv"
)

type Pigment struct {
  name string
  ID string
  color Color
  opacity int
  percentSolids int
  price int
  stock int
}

type Color struct{
  r,g,b int
}

//MixColor takes a slice of pigments and their weights 
//and returns the mixed new color, as well as the price per gal
func MixColor(p []Pigment, weight []float64) Color{
  r := 0.0
  g := 0.0
  b := 0.0
  totalWeight := 0.0
  totalPrice := 0.0
  for i := range p{
    r += float64(p[i].color.r)*weight[i]
    g += float64(p[i].color.g)*weight[i]
    b += float64(p[i].color.r)*weight[i]
    totalWeight += weight[i]
    totalPrice += p[i].price*weight[i]
  }
  r /= totalWeight
  g /= totalWeight
  b /= totalWeight
  totalPrice /= totalWeight
  return Color{int(r),int(g),int(b)}, totalPrice
}

// ReadFile read a given csv file and create a Pigment struct for each pigment
func ReadFile(filename string) (bool,[]Pigment) {
  // noError reports whether the function works well
  var noError bool
  // a slice to hold all pigment structs
  pigments := make([]Pigment,0)
  // Open the csv file
  in, err := os.Open(filename)
	if err != nil {
		return noError, pigments
	}

  // Read the csv file line by line
	lines := csv.NewReader(bufio.NewReader(in))
	for {
    record,err1 := lines.Read()
    if err1 == io.EOF {
      break
    }
    if err1 != nil {
			return noError, pigments
		}
    // Omit the heading line
    if record[0] == "Pigment Name (string)" { ///need modification according to heading name
      continue
    } else {
      // Remove the white space inside the name
      name := SpaceFieldsJoin(record[0])
      record[0] = name
      // Build a struct to hold each pigment
      var p Pigment
      p.name = record[0]
      p.ID = record[1]
      L,err2 := strconv.ParseFloat(record[2],64)
      p.L = L
      if err2 != nil {
    		return noError, pigments
    	}
      a, err3 := strconv.ParseFloat(record[3],64)
      p.a = a
      if err3 != nil {
    		return noError, pigments
    	}
      b, err4 := strconv.ParseFloat(record[4],64)
      p.b = b
      if err4 != nil {
    		return noError, pigments
    	}
      opacity, err5 := strconv.Atoi(record[5])
      p.opacity = opacity
      if err5 != nil {
    		return noError, pigments
    	}
      percentSolids, err6 := strconv.Atoi(record[6])
      p.percentSolids =percentSolids
      if err6 != nil {
    		return noError, pigments
    	}
      price, err7 := strconv.Atoi(record[7])
      p.price = price
      if err7 != nil {
    		return noError, pigments
    	}
			l := len(record)
			stock, err8 := strconv.Atoi(record[l-1])
      p.stock = stock
      if err8 != nil {
    		return noError, pigments
    	}
      pigments = append(pigments,p)
    }
  }
  noError = true
  return noError, pigments
}

// SpaceFieldsJoin strips the space inside a string.
func SpaceFieldsJoin(str string) string {
    return strings.Join(strings.Fields(str), "")
}

func main() {
  filename := os.Args[1]
  noError, pigments := ReadFile(filename)
  fmt.Println(noError)
  fmt.Println(pigments)
}
