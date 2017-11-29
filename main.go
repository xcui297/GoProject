package main

import (
  //"github.com/therecipe/qt/widgets"
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
  L float64
  a float64
  b float64
  opacity int
  percentSolids int
  price int
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
/*
  // Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Automobile Paint Calculator")
	window.SetMinimumSize2(200, 200)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create a line edit and add it to the layout
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 0, 0)

	// Create a button and add it to the layout
	button := widgets.NewQPushButton2("2. click me", nil)
	layout.AddWidget(button, 0, 0)

	// Connect event for button
	button.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(),
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)

	// Show the window
	window.Show()

	// Execute app
	app.Exec()
*/
}
