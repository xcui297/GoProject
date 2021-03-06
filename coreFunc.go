package main
// #include <stdlib.h>
import "C"
import (
	"os"
  "encoding/csv"
  "fmt"
  "bufio"
  "io"
  "strings"
  "strconv"
	"github.com/draffensperger/golp"
	"unsafe"
)
// The pigments are the current colors we have in the repository that can be
// used for mixing to get a desired new color.
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

// The paint is a desired color that we want to achieve.
type Paint struct {
	name string
	R int
	G int
	B int
	amount float64
	price int
	componentFound bool
}

// The components are colors required for getting a desired new color.
type Component struct {
	pigment *Pigment
	percentage float64
	amount float64
	inStock bool
	supplement float64
}
/*==============================================================================
 * 1. Load repository
 * This module reads the given .csv files and creates a slice of Pigments.
 *============================================================================*/
// ReadFile reads a given .csv file and creates a Pigment struct for each pigment.
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
    if record[0] == "Pigment Name (string)" {
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
 * This module tries to find the components reguired for mixing of a given target
 * color, with a minimized price.
 *============================================================================*/
// FindComponents uses linear programming to find pigments that can mix together
// in certain proportion to give the targret Color.
// It returns two boolean values indicating whether can find components in
// the given repository and whether all components are in stock, respectively.
// It also returns a slice of those found components and the final price of the
// targeted color.
func FindComponents(R, G, B int, targetAmount float64, repository []Pigment) (bool, bool, []Component, int) {
	var targetColor Paint
	targetColor.R = R
	targetColor.G = G
	targetColor.B = B
	targetColor.amount = targetAmount
	vars, finalUnitPrice := SolveLPFunc1(targetColor, repository)
	return InterpretFunc1LPResults(vars, finalUnitPrice, targetAmount, repository)
}

// SolveLPFunc1 gives the component proportions at a minimized price when it is
// constrained by the targeted R,G,B values and the same total unit amount.
// It utilized a 'golp' package that coordinate C language linear programming
// into Go programming.
func SolveLPFunc1(targetColor Paint, repository []Pigment) ([]float64, int) {
	// Make a new liear program structure.
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

	// It calls SolveLP to solve the linear programming.
	vars, finalUnitPrice := SolveLP(*lp)
	C.free(unsafe.Pointer(lp))
	return vars, finalUnitPrice
}

// SolveLP tries to solve the given linear programming, and return the variables'
// value at the optimal condition.
// It also returns the value of the objective function at its minimum or maximum.
func SolveLP(lp golp.LP) ([]float64, int) {
	// Maximize the objective function, restricted by constraints of lp
	lp.SetMaximize()

	// Solve the linear programming.
	lp.Solve()
	vars := lp.Variables()
	finalUnitPrice := -int(lp.Objective())
	return vars, finalUnitPrice
}


// InterpretLPResults evaluates the results from linear programming and gives:
// 1. a boolean value, which indicates the feasability to find components for the
// target value.
// 2. a boolean value tells whether all components are in stock to form the
// required amount of the target color.
// 3. a slice of Comonent, each of which gives the pigment, the proportion, the
// required amount, the stock status.
func InterpretFunc1LPResults(vars []float64, finalUnitPrice int, targetAmount float64, repository []Pigment) (bool, bool, []Component, int) {
	var exists, inStock bool
	composition:= make([]Component, 0)

	for idx := range vars {
		// Check negative values. Since we do not want negative values, simply return
		// exists = false, inStock = false indicating pigments in repository can
		// not form the target color.
		// It also returns an empty slice of Component.
		if vars[idx] < 0.0 {
			composition = []Component{}
			finalUnitPrice = 0
			return exists, inStock, composition, finalUnitPrice

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
				component.supplement = component.amount - stock
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
	} else {
		composition = []Component{}
		finalUnitPrice = 0
		return exists, inStock, composition, finalUnitPrice
	}

	// Check whether all the components are in stock.
	for _,component := range composition {
		if component.inStock == false {
			inStock = false
			return exists, inStock, composition, finalUnitPrice
		}
	}

	inStock = true
	return exists,inStock, composition,finalUnitPrice
}
/*==============================================================================
 * 3. Off-color hit
 * This module improves a current color to a desired color by adding pigments
 * and it also ensures the use of the current color to avoid waste.
 *============================================================================*/
 // OffColorHit calculates the pigments required to add into the a color given，
 // in order to achieve another desired color.
 func OffColorHit(R0, G0, B0, R, G, B, currentPrice int, repository []Pigment) (bool,float64, []Component, int) {
 	var currentColor,targetColor Paint
 	currentColor.R = R0
 	currentColor.G = G0
 	currentColor.B = B0
 	currentColor.price = currentPrice

 	targetColor.R = R
 	targetColor.G = G
 	targetColor.B = B
 	vars, finalUnitPrice := SolveLPFunc2(currentColor, targetColor, repository)
 	return InterpretFunc2LPResults(vars, finalUnitPrice, repository)
 }

// SolveLPFunc2 is the function that sets up the linear programming structure
// for solving OffColorHit.
 func SolveLPFunc2(currentColor, targetColor Paint, repository []Pigment) ([]float64, int) {
 	// Make a new liear program structure.
 	numPigments := len(repository)
 	lp := golp.NewLP(0, numPigments+1)

 	// Make slices to hold the coefficients and the objective function.
 	// N indicates the amount of each pigment.
 	// rowCurrentColor requires currentColor must be used.
 	rowR := make([]float64,0)
 	rowG := make([]float64,0)
 	rowB := make([]float64,0)
 	rowN := make([]float64,0)
 	rowCurrentColor:= make([]float64,0)
 	rowPrice:= make([]float64, 0)
 	// Add coefficients for currentColor to the constraints
 	rowR = append(rowR, float64(currentColor.R))
 	rowG = append(rowG, float64(currentColor.G))
 	rowB = append(rowB, float64(currentColor.B))
 	rowN = append(rowN, 1.0)
 	rowCurrentColor = append(rowCurrentColor, 1.0)
 	rowPrice = append(rowPrice,-float64(currentColor.price))
 	// Add coefficents for pigment repository to the above slices.
 	for _,pigment := range repository {
 		rowR = append(rowR, float64(pigment.R))
 		rowG = append(rowR, float64(pigment.G))
 		rowB = append(rowR, float64(pigment.B))
 		rowN = append(rowN, 1.0)
 		rowCurrentColor = append(rowCurrentColor, 0.0)
		rowPrice = append(rowPrice, -float64(pigment.price))
 	}

 	// Add constraints to the linear programming.
 	// This constraint requires pigments to give the target value of R.
 	lp.AddConstraint(rowR, golp.EQ, float64(targetColor.R))
 	// This constraint requires pigments to give the target value of G.
 	lp.AddConstraint(rowG, golp.EQ, float64(targetColor.G))
 	// This constraint requires pigments to give the target value of B.
 	lp.AddConstraint(rowB, golp.EQ, float64(targetColor.B))
 	// This constraint requires the total amount for each pigment combination
 	// to be the same, which is 1.0.
 	lp.AddConstraint(rowN, golp.EQ, 1.0)
 	// currentColor must be used.
 	lp.AddConstraint(rowCurrentColor,golp.GE, 0.001)
 	lp.SetObjFn(rowPrice)

 	return SolveLP(*lp)
 }

 // InterpretLPFunc2 Results evaluates the results from linear programming and gives:
 // 1. a boolean value, which indicates the feasability to find components for the
 // target value.
 // 2. a percentage of the currentcolor.
 // 3. a slice of Comonent, each of which gives the pigment, the proportion, the
 // required amount, the stock status.
 func InterpretFunc2LPResults(vars []float64, finalUnitPrice int, repository []Pigment) (bool,float64, []Component, int) {
 	var exists bool
	var currentColorProportion float64
 	composition:= make([]Component, 0)

 	for idx := range vars {
 		// Check negative values. Since we do not want negative values, simply return
 		// exists = false
 		// It also returns an empty slice of Component.
 		if vars[idx] < 0.0 {
 			composition = []Component{}
			finalUnitPrice = 0
 			return exists, 0.0, composition, finalUnitPrice
 		// Add each component with positve percentage into composition.
 		} else if vars[idx] > 0.0 {
 			if idx == 0 {
 			currentColorProportion = vars[idx]
 			} else {
 				var component Component
 				component.pigment = &(repository[idx])
 				component.percentage = vars[idx]
 				composition = append(composition, component)
 			}
 		// "0" indicates this pigment is not used.
 		} else {
 			continue
 		}
 	}
 	// Check whether the composition is empty.
 	if len(composition) > 1 {
 		exists = true
 	} else {
		composition = []Component{}
		finalUnitPrice = 0
		return exists, 0.0, composition, finalUnitPrice
	}
 	return exists, currentColorProportion, composition, finalUnitPrice
 }
/*==============================================================================
* 4. Color range from mixing the pigments in the repository
* This module can give out all the colors can get from mixing some of the
* pigments in the repository.
* (Since the package used to slove linear programming utilized C language,
* the parallel programming of Go is not able to use this package. We still
* retain the parallel version of this function, since it is too slow to
* implement in serial, which is meaningless. )
*=============================================================================*/
// ColorRange gives out the color range that pigments in a given repository can
// give.
func ColorRange(repository []Pigment) [][]int {
	// We narrow down the range by finding the max and min for R, G and B values.
	maxR,maxG,maxB,minR,minG,minB := 130, 130, 130, 130, 125,120
	AllCandidateColors := make([]Paint,0)
	var candidate Paint
	for i := minR; i <= maxR; i ++ {
		for j := minG; j <= maxG; j ++ {
			for k := minB; k <= maxB; k ++ {
				candidate.R = i
				candidate.G = j
				candidate.B = k
				AllCandidateColors = append(AllCandidateColors, candidate)
			}
		}
	}
	fmt.Println(len(AllCandidateColors))
	c := make(chan bool)
	colorRange := make([][]int,0)

	// We call the CheckComponent for each candidate parallelly but the golp package
	// used call functions from C language, which cannot be parallelled in Go.
	// *** Bug still remains.
	for _,candidate := range AllCandidateColors {
		go candidate.CheckComponent(c, repository)
	}

	for _,candidate := range AllCandidateColors {
    <-c
		if candidate.componentFound == true {
			colorRange = append(colorRange, []int{candidate.R, candidate.G, candidate.B})
		}
  }
	return colorRange
}

// CheckComponent calls the FindCompnoent function.
// It simply return a bool value to tell whether pigments in repository can mix
// to get this candidate color.
func (candidate *Paint) CheckComponent(c chan bool, repository []Pigment) {
	r := candidate.R
	g := candidate.G
	b := candidate.B
	exist, _, _, _ :=FindComponents(r,g,b, 0.0, repository)
	if exist == true {
		candidate.componentFound = true
	}
	c <- true
}

//MaxAndMinChannle returns the maximum R,G,B value and the minimum R,G,B value
//of a slice of pigments
func RGBValueRange(repository []Pigment) (int,int,int,int,int,int){
  maxR := 0
  maxG := 0
  maxB := 0
  minR := 255
  minG := 255
  minB := 255
  for _,pigment:=range repository {
    maxR = MaxInt(pigment.R,maxR)
    maxG = MaxInt(pigment.G,maxG)
    maxB = MaxInt(pigment.B,maxB)
    minR = MinInt(pigment.R,minR)
    minG = MinInt(pigment.G,minG)
    minB = MinInt(pigment.B,minB)
  }
  return maxR,maxG,maxB,minR,minG,minB
}

//MaxInt takes two integers and returns the maximum
func MaxInt(a,b int)int{
  if a > b {
    return a
  } else {
    return b
  }
}

//MaxInt takes two integers and returns the minimum
func MinInt(a,b int)int{
  if a > b {
    return b
  } else {
    return a
  }
}
/*==============================================================================
* 5. Mix Color
* This function allows user to explore the color. Simply by inputing pigments
* and proportion, they can check what color they can get.
*=============================================================================*/
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
 * Since this project will hava a main file that calls functions in this script,
 * the main function in this scripts is commented out.
 * However, the test serves as an important part of the build-up for this
 * project, and it gives out results in the shell, which is easy to read and
 * understand.
 *============================================================================*/
 /*
func main() {
  filename := os.Args[1]
  noError, repository := ReadFile(filename)
	if noError == true {
		fmt.Println("Load repository: Succeeded.")
	} else {
		fmt.Println("Load repository: Failed.")
	}

	//------------------test for func 1-------------------------------------------
	R := 10
	G := 128
	B := 128
	targetAmount := 15.0
	start1 := time.Now()
	exists, inStock, composition, finalUnitPrice := FindComponents(R, G, B, targetAmount, repository)
	elapsed1 := time.Since(start1)
	log.Printf("func1 took %v", elapsed1)
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
			fmt.Fprintf(os.Stdout,"Name = %v, ID = %v, proportion = %v, required amount = %v, %v, supplement = %v \n",
			pigment.name, pigment.ID, component.percentage, component.amount, status, component.supplement)
		} else {
		fmt.Fprintf(os.Stdout,"Name = %v, ID = %v, proportion = %v, required amount = %v, %v \n",
		pigment.name, pigment.ID, component.percentage, component.amount, status)
		}
	}
	fmt.Fprintf(os.Stdout, "The lowest unit price for this target color is %v.", finalUnitPrice)
	fmt.Println()
	//------------------test for func 2-------------------------------------------
	R0 := 128
	G0 := 120
	B0 := 110
	currentP := 20
	R = 128
	G = 110
	B = 120
	exists2, currentProp, composition2, finalUnitPrice := OffColorHit(R0, G0, B0, R, G, B, currentP, repository)
	fmt.Println(exists2, currentProp)
	for _, component := range composition2 {
		pigment := *(component.pigment)
		fmt.Fprintf(os.Stdout,"Name = %v, ID = %v, proportion = %v\n",
		pigment.name, pigment.ID, component.percentage)
	}
	fmt.Println("final unit price is %v", finalUnitPrice)
	fmt.Println()

	//------------------test for func 3---------------------------------------------
	a,b,c,d,e,f := RGBValueRange(repository)
	fmt.Println(a,b,c,d,e,f)
	colorRange := ColorRange(repository)
	fmt.Println(colorRange)
}
*/
