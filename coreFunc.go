package main

import (
	"os"
  "encoding/csv"
  //"fmt"
  "bufio"
  "io"
  "strings"
  "strconv"

	"github.com/draffensperger/golp"
)

type Pigment struct {
	name string
  ID string
  L float64
  a float64
  b float64
  opacity int
  percentSolids int
  price int
	R int
	G int
	B int
	stock int
}

type Paint struct {
	name string
	R int
	G int
	B int
	amount float64
}

type Component struct {
	pigment *Pigment
	percentage float64
	amount float64
	inStock bool
}

/*==============================================================================
 * 1. Load repository
 *
 *============================================================================*/
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
			R, err8 := strconv.Atoi(record[8])
			p.R = R
			if err8 != nil {
    		return noError, pigments
    	}
			G, err9 := strconv.Atoi(record[9])
			p.G = G
			if err9 != nil {
    		return noError, pigments
    	}
			B, err10 := strconv.Atoi(record[10])
			p.B = B
			if err10 != nil {
    		return noError, pigments
    	}
			l := len(record)
			stock, err11 := strconv.Atoi(record[l-1])
      p.stock = stock
      if err11 != nil {
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

/*==============================================================================
 * 2. Mix to get the target color
 *
 *============================================================================*/
// FindComponents uses linear programming to find pigments that can mix together in certain proportion
// to give the targret Color.
// It returns two boolean value indicating whether can find components from
// the given repository and whether all components are in stock, respectively.
// It also returns a slice of those found components.
func FindComponents(R, G, B int, targetAmount float64, repository []Pigment) (bool, bool, []Component) {
	var targetColor Paint
	targetColor.R = R
	targetColor.G = G
	targetColor.B = B
	targetColor.amount = targetAmount
	vars := SolveLPFunc1(targetColor, repository)
	return InterpretFunc1LPResults(vars, targetAmount, repository)
}

// SolveLP gives the component proportions at a minimized price when it is
// constrained by the targeted R,G,B values and the same total unit amount.
func SolveLPFunc1(targetColor Paint, repository []Pigment) []float64 {
	// Make a new linear program structure.
	numPigments := len(repository)
	lp := golp.NewLP(0, numPigments)

	// Make slices to hold the coefficients for R, G, B, N and slices.
	// N indicates the amount of each pigment.
	rowR := make([]float64,0)
	rowG := make([]float64,0)
	rowB := make([]float64,0)
	rowN := make([]float64,0)
	rowPrice:= make([]float64, 0)

	// Add coefficents to the above slices.
	for _,pigment := range repository {
		r := float64(pigment.R)
		g := float64(pigment.G)
		b := float64(pigment.B)
		p := -float64(pigment.price)
		rowR = append(rowR, r)
		rowG = append(rowR, g)
		rowB = append(rowR, b)
		rowN = append(rowN, 1.0)
		rowPrice = append(rowPrice, p)
	}

	// Add constraints to the linear programming.
	targetR := float64(targetColor.R)
	targetG := float64(targetColor.G)
	targetB := float64(targetColor.B)
	// This constraint requires pigments to give the target value of R.
	lp.AddConstraint(rowR, golp.EQ, targetR)
	// This constraint requires pigments to give the target value of G.
	lp.AddConstraint(rowG, golp.EQ, targetG)
	// This constraint requires pigments to give the target value of B.
	lp.AddConstraint(rowB, golp.EQ, targetB)
	// This constraint requires the total amount for each pigment combination
	// to be the same, which is 1.0.
	lp.AddConstraint(rowN, golp.EQ, 1.0)

	// Minimize the objective function, which is the total price for each pigment
	// combination. We convert the minimization to a maximination by setting the
	// coefficents of price to negative.
	lp.SetObjFn(rowPrice)
	vars := SolveLP(*lp)

	return vars
}

func SolveLP(lp golp.LP) []float64 {
	// Maximizethe objective function, restricted by constraints of lp
	lp.SetMaximize()

	// Solve the linear programming.
	lp.Solve()
	vars := lp.Variables()

	return vars
}

// InterpretLPResults evaluates the results from linear programming and gives:
// 1. a boolean value, which indicates the feasability to find components for the
// target value.
// 2. a boolean value tells whether all components are in stock to form the
// required amount of the target color.
// 3. a slice of Comonent, each of which gives the pigment, the proportion, the
// required amount, the stock status.
func InterpretFunc1LPResults(vars []float64, targetAmount float64, repository []Pigment) (bool, bool, []Component) {
	var exists, inStock bool
	composition:= make([]Component, 0)

	for idx := range vars {
		// Check negative values. Since we do not want negative values, simply return
		// exists = false, inStock = false indicating pigments in repository can
		// not form the target color.
		// It also returns an empty slice of Component.
		if vars[idx] < 0.0 {
			composition = []Component{}
			return exists, inStock, composition

		// Add each component with positve percentage into composition.
		// Calculate their required amount according to the percentage and amount of
		// the target color.
		} else if vars[idx] > 0.0 {
			var component Component
			component.pigment = &(repository[idx])
			component.percentage = vars[idx]
			component.amount = targetAmount * vars[idx]
			stock := float64(repository[idx].stock)
			// Compare the amount required with the stock. If in stock, return true.
			if stock < component.amount {
				component.inStock = false
			} else {
				component.inStock = true
			}
			composition = append(composition, component)
		// "0" indicates this pigment is not used.
		} else {
			continue
		}
	}

	// Check whether the composition is empty.
	if len(composition) > 0 {
		exists = true
	}

	// Check whether all the components are in stock.
	for _,component := range composition {
		if component.inStock == false {
			inStock = false
			return exists, inStock, composition
		}
	}

	inStock = true
	return exists,inStock, composition
}

/*==============================================================================
 *
 *
 *============================================================================*/
//MixColor takes a slice of pigments and their weights
//and returns the mixed new color, as well as the price per gal
func MixColor(p []Pigment, weight []float64) (Paint, float64) {
	var targetColor Paint
  r := 0.0
  g := 0.0
  b := 0.0
  totalWeight := 0.0
  totalPrice := 0.0
  for i := range p{
    r += float64(p[i].R)*weight[i]
    g += float64(p[i].G)*weight[i]
    b += float64(p[i].B)*weight[i]
    totalWeight += weight[i]
    totalPrice += float64(p[i].price)*weight[i]
  }
  r /= totalWeight
  g /= totalWeight
  b /= totalWeight
  totalPrice /= totalWeight

	targetColor.R = int(r)
  targetColor.G = int(g)
  targetColor.B = int(b)
  targetColor.amount = totalWeight
  return targetColor, totalPrice
}

/*==============================================================================
 * main for test
 *
 *============================================================================*/
/*func main() {
  filename := os.Args[1]
  noError, repository := ReadFile(filename)
	if noError == true {
		fmt.Println("Load repository: Succeeded.")
	} else {
		fmt.Println("Load repository: Failed.")
	}

	R := 88
	G := 92
	B := 103
	targetAmount := 30.0
	for i := 0; i<10; i++ {
		exists, inStock, composition := FindComponents(R, G, B, targetAmount, repository)
		if exists == true {
			fmt.Fprintf(os.Stdout, "Components found for targetColor from the given repository.\n")
		} else {
			fmt.Fprintf(os.Stdout,"CANNOT find components for from the given repository.\n")
		}

		if inStock == true {
			fmt.Fprintf(os.Stdout,"All Components to form %v of target color are in stock.\nPlease check details below:\n",targetAmount)
		} else {
			fmt.Fprintf(os.Stdout,"NOT all Components for %v of target color are in stock.\nPlease check details below:\n",targetAmount)
		}
	
		for _, component := range composition {
			pigment := *(component.pigment)
			status := "in stock"
			if component.inStock == false {
				status = "out of stock"
			}
			fmt.Fprintf(os.Stdout,"Name = %v, ID = %v, proportion = %v, required amount = %v, %v \n",
			pigment.name, pigment.ID, component.percentage, component.amount, status)
		}
	}
}*/
