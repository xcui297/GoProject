package main

import (
    "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
	"strconv"
)

var (
	indexWindow *widgets.QMainWindow
    helpWindow *widgets.QMainWindow
    helpMainWidget *widgets.QWidget
	mixColorWindow *widgets.QMainWindow
	mixColorOutputWindow *widgets.QMainWindow
	mixColorOutputMainWidget *widgets.QWidget
	mixColorOutputLayout *widgets.QVBoxLayout
	offColorWindow *widgets.QMainWindow
	colorRangeWindow *widgets.QMainWindow
	repositoryLoaded bool
	repositoryFileName string
	repository []Pigment
	func1Exists bool
	func1InStock bool
	func1Composition []Component
)

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)
	// Create index window. The following buttons connect to other windows
	CreateIndexWindow()
	repositoryLoaded = false
	// Create help window
    CreateHelpWindow()
	// Create the mix color window and the potential output window
	CreateMixColorWindow()
	CreateOffColorWindow()
	CreateColorRangeWindow()
	// Show the index window
	indexWindow.Show()
	// Execute app
	app.Exec()
}

func openFile() {
	repositoryFileName = widgets.QFileDialog_GetOpenFileName(nil, "Open Repository file", core.QDir_HomePath(), "", "", 0)
	repositoryLoaded, repository = ReadFile(repositoryFileName)
	if repositoryLoaded {
		widgets.QMessageBox_Information(nil, "OK", "Congratulations! Repository data successfully loaded!",
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else {
		widgets.QMessageBox_Information(nil, "OK", "An error has occured when loading repository file!",
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	}
}

func CreateIndexWindow() {
	// Create index window
    indexWindow = widgets.NewQMainWindow(nil,0)
    indexWindow.SetWindowTitle("Automobile paint calculator")
	indexWindow.SetMinimumSize2(500, 500)
    indexWindow.SetMaximumSize2(500, 500)
	// Create layout of index window
	indexLayout := widgets.NewQVBoxLayout()
	// Create main widget of index window and set the layout
	indexMainWidget := widgets.NewQWidget(nil, 0)
	indexMainWidget.SetLayout(indexLayout)
    indexWindow.SetWindowTitle("Automobile paint calculator")
    // Create introduction textbox on index window
    introTextContent := "Hello! This is a simple calculator for automotive paints!"+
    " Choose from the major functions to start your painting journey! Click \"Help\""+
    " for more information. Also, if you have not yet uploaded your repository, you" +
    " can either upload it now or do so after choosing the function you need!"
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    // Add intro text box to layout of index window
    indexLayout.AddWidget(introText, 0, 0)
    // Create layout for control panel
    controlLayout := widgets.NewQHBoxLayout()
	controlLayout.SetContentsMargins(0, 0, 0, 0)
    // Create control panel with help button and load repository function
    helpButton := widgets.NewQPushButton2("Help", nil)
	helpButton.ConnectClicked(func(_ bool) { openHelp() })
	loadButton := widgets.NewQPushButton2("Load repository", nil)
    loadButton.ConnectClicked(func(_ bool) { openFile() })
	// Add help and load buttons to control panel layout
    controlLayout.AddWidget(helpButton, 0, 0)
    controlLayout.AddWidget(loadButton, 0, 0)
    // Add control layout to index window layout
    indexLayout.AddLayout(controlLayout, 0)
    // Create major function box
    majorFunctionsBox := widgets.NewQGroupBox2("Major functions", nil)
    functionsLayout := widgets.NewQVBoxLayout()
    majorFunctionsBox.SetLayout(functionsLayout)
    // Create buttons for the 3 main functions, and add them to the layout
    function1Button := widgets.NewQPushButton2("1. Mix target color", majorFunctionsBox)
	function1Button.ConnectClicked(func(_ bool) { openFunction(mixColorWindow) } )

    function2Button := widgets.NewQPushButton2("2. Off-color hits", majorFunctionsBox)
	function2Button.ConnectClicked(func(_ bool) { openFunction(offColorWindow) })
	
    function3Button := widgets.NewQPushButton2("3. Color range", majorFunctionsBox)
	function3Button.ConnectClicked(func(_ bool) { openFunction(colorRangeWindow) })
	
    functionsLayout.AddWidget(function1Button, 0, 0)
    functionsLayout.AddWidget(function2Button, 0, 0)
    functionsLayout.AddWidget(function3Button, 0, 0)
    // Add major function box to index window layout
    indexLayout.AddWidget(majorFunctionsBox, 0, 0)
    // Create toolkit box
    toolkitBox := widgets.NewQGroupBox2("Toolkit", nil)
    toolkitLayout := widgets.NewQVBoxLayout()
    toolkitBox.SetLayout(toolkitLayout)
    // Create color visualization button and add it to toolkit box layout
    colorVisualizeButton := widgets.NewQPushButton2("Visualize RGB color", toolkitBox)
    toolkitLayout.AddWidget(colorVisualizeButton, 0, 0)
    // Add toolkit box to layout of the index window
    indexLayout.AddWidget(toolkitBox, 0, 0)
	//Set index main widget as the central widget of the index window
	indexWindow.SetCentralWidget(indexMainWidget)
}

func CreateHelpWindow() {
    // Create help window
    helpWindow = widgets.NewQMainWindow(nil, 0)
    helpWindow.SetWindowTitle("Help")
	helpWindow.SetMinimumSize2(500, 500)
    helpWindow.SetMaximumSize2(500, 500)
    // Create layout for help window
    helpLayout := widgets.NewQVBoxLayout()
    helpMainWidget = widgets.NewQWidget(nil, 0)
	helpMainWidget.SetLayout(helpLayout)
    // Create question links of the following questions
    // The first question is on format of the input data
    formatQuestionContent := "1. What format should the repository data be?"
    formatQuestionLayout, formatQuestionButton := CreateQuestionLayout(formatQuestionContent)
    formatAnswerContent := "blah blah"
    formatAnswerWindow := CreateAnswerWindow(formatQuestionContent, formatAnswerContent)
    formatQuestionButton.ConnectClicked(func(_ bool) { openAnswer(formatAnswerWindow) })
    // The 2nd question is on the first major function of the program
    mixColorQuestionContent := "2. What does the Mix-target-color function do?"
    mixColorQuestionLayout, mixColorQuestionButton := CreateQuestionLayout(mixColorQuestionContent)
	mixColorAnswerContent := "This function allows formulators to develop new basecoats with increased efficiency"+
	" and precision. Given a target color represented in RGB format, this program will recommend a pigment package"+
	" from the pigment grinds available to the chemist. The program will output not only the proportions of" +
	" the pigments to be added to the formulation but also the price added per gallon increase and any special" +
	" instructions required when using the recommended pigment package. The calculator will recommend the best"+
	" pigment package based on the criteria of minimizing cost, simplifying manufacturing, and reducing the propensity"+
	" for variability between batches."
	mixColorAnswerWindow := CreateAnswerWindow(mixColorQuestionContent, mixColorAnswerContent)
	mixColorQuestionButton.ConnectClicked(func(_ bool) { openAnswer(mixColorAnswerWindow) })
    // The 3rd question is on the second major function of the program
    offColorQuestionContent := "3. What are off-color hits?"
    offColorQuestionLayout, offColorQuestionButton := CreateQuestionLayout(offColorQuestionContent)
	offColorAnswerContent := "For every batch made, a sample is taken and sprayed to compare against the standard"+
	" set by the customer. Whenever a batch is off color, chemists are required to make \"hits\" to the batch to "+
	"bring it back within the designated color space. Given a target color, the current color, and the repository"+
	" of pigments, this calculator would recommend which pigments to add to the batch to bring it back to spec."
	offColorAnswerWindow := CreateAnswerWindow(offColorQuestionContent, offColorAnswerContent)
	offColorQuestionButton.ConnectClicked(func(_ bool) { openAnswer(offColorAnswerWindow) })
    // The 4th question is on the third major function of the program
    colorRangeQuestionContent := "4. What does the Color-range function do?"
    colorRangeQuestionLayout, colorRangeQuestionButton := CreateQuestionLayout(colorRangeQuestionContent)
	colorRangeAnswerContent := "The function takes the current repository of pigments listed within it and return"+
	" the entire range of color able to be achieved. This will allow the chemists to foresee gaps in our capabilities"+
	" so that they can either add or replace the pigments we have on stock to either reach new color spaces or reach"+
	" current color spaces for less money."
	colorRangeAnswerWindow := CreateAnswerWindow(colorRangeQuestionContent, colorRangeAnswerContent)
	colorRangeQuestionButton.ConnectClicked(func(_ bool) { openAnswer(colorRangeAnswerWindow) })
    // More questions
    moreQuestionContent := "I have other questions?"
    moreQuestionLayout, moreQuestionButton := CreateQuestionLayout(moreQuestionContent)
	moreAnswerContent := `Contact us!
	Echo: ziyic@andrew.cmu.edu
	Xiaoyue: xcui297@gmail.com
	Xinyu: xinyuh1@andrew.cmu.edu`
	moreAnswerWindow := CreateAnswerWindow(moreQuestionContent, moreAnswerContent)
	moreQuestionButton.ConnectClicked(func(_ bool) { openAnswer(moreAnswerWindow) })
    // Add the questions to layout of the help window
    helpLayout.AddLayout(formatQuestionLayout, 0)
    helpLayout.AddLayout(mixColorQuestionLayout, 0)
    helpLayout.AddLayout(offColorQuestionLayout, 0)
    helpLayout.AddLayout(colorRangeQuestionLayout, 0)
    helpLayout.AddLayout(moreQuestionLayout, 0)
    // Create button directing back to the index window
    backButton := widgets.NewQPushButton2("Back", nil)
    backButton.ConnectClicked(func(_ bool) { goBackFromHelp() })
    helpLayout.AddWidget(backButton, 0, 0)
    helpWindow.SetCentralWidget(helpMainWidget)
}

func CreateMixColorWindow() {
	mixColorWindow = widgets.NewQMainWindow(nil, 0)
	mixColorWindow.SetWindowTitle("Mix target color")
	mixColorWindow.SetMinimumSize2(500,500)
	mixColorWindow.SetMaximumSize2(500,500)
	mainWidget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout()
	mainWidget.SetLayout(layout)
	mixColorWindow.SetCentralWidget(mainWidget)
	// Create instruction textbox
	instructionTextContent := "Select your target color in RGB format: only integers from 0 to 255 are accepted."+
	" Please also input quantity required. This should be a float64 number between 0 and 100."
    instructionText := widgets.NewQLabel(nil, 0)
    instructionText.SetWordWrap(true)
    instructionText.SetText(instructionTextContent)
	// Create target color selection box
	targetColor := widgets.NewQGroupBox2("Target color", nil)
	targetColorLayout := widgets.NewQVBoxLayout()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("R:", nil, 0)
	redSpinBox := widgets.NewQSpinBox(nil)
	redSpinBox.SetMaximum(255)
	redSpinBox.SetMinimum(0)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("G:", nil, 0)
	greenSpinBox := widgets.NewQSpinBox(nil)
	greenSpinBox.SetMaximum(255)
	greenSpinBox.SetMinimum(0)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("B:", nil, 0)
	blueSpinBox := widgets.NewQSpinBox(nil)
	blueSpinBox.SetMaximum(255)
	blueSpinBox.SetMinimum(0)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	// Construct box for quantity input
	quantityGroup := widgets.NewQGroupBox(nil)
	quantityLabel := widgets.NewQLabel2("Required quantity:", nil, 0)
	quantitySpinBox := widgets.NewQDoubleSpinBox(nil)
	quantitySpinBox.SetMaximum(100.0)
	quantitySpinBox.SetMinimum(0.0)
	quantitySpinBox.SetSingleStep(1.0)
	quantityLayout := widgets.NewQGridLayout2()
	quantityLayout.AddWidget(quantityLabel, 0, 0, 0)
	quantityLayout.AddWidget(quantitySpinBox, 0, 1, 0)
	quantityGroup.SetLayout(quantityLayout)
	// Add RGB and quantity input box to target color layout
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	targetColorLayout.AddWidget(quantityGroup, 0, 0)
	// Create load repository button
	loadButton := widgets.NewQPushButton2("Load repository", nil)
	loadButton.ConnectClicked(func(_ bool) { openFile() })
	// Construct visualization button
	visualizeButton := widgets.NewQPushButton2("Visualize target color", nil)
	// Create back button (go back to index window)
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToIndexFromFunction(mixColorWindow) } )
	// Create mix button
	mixButton := widgets.NewQPushButton2("Mix target color!", nil)
	mixButton.ConnectClicked(func(_ bool) {mixColor(redSpinBox, greenSpinBox, blueSpinBox, quantitySpinBox)} )
	layout.AddWidget(instructionText, 0, 0)
	layout.AddWidget(targetColor, 0, 0)
	layout.AddWidget(visualizeButton, 0, 0)
	layout.AddWidget(loadButton, 0, 0)
	layout.AddWidget(mixButton, 0, 0)
	layout.AddWidget(backButton, 0, 0)
	
}


func CreateOffColorWindow() {
	offColorWindow = widgets.NewQMainWindow(nil, 0)
	offColorWindow.SetWindowTitle("Off-color hits")
	offColorWindow.SetMinimumSize2(500,500)
	offColorWindow.SetMaximumSize2(500,500)
	mainWidget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout()
	mainWidget.SetLayout(layout)
	offColorWindow.SetCentralWidget(mainWidget)
	// Create instruction textbox
	instructionTextContent := "Select your target color and off-color batch color in RGB format:"+
	" only integers from 0 to 255 are accepted."
    instructionText := widgets.NewQLabel(nil, 0)
    instructionText.SetWordWrap(true)
    instructionText.SetText(instructionTextContent)
	// Create target color selection box
	targetColor := widgets.NewQGroupBox2("Target color", nil)
	targetColorLayout := widgets.NewQHBoxLayout()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for target red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("R:", nil, 0)
	redSpinBox := widgets.NewQSpinBox(nil)
	redSpinBox.SetMaximum(255)
	redSpinBox.SetMinimum(0)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for target green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("G:", nil, 0)
	greenSpinBox := widgets.NewQSpinBox(nil)
	greenSpinBox.SetMaximum(255)
	greenSpinBox.SetMinimum(0)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for target blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("B:", nil, 0)
	blueSpinBox := widgets.NewQSpinBox(nil)
	blueSpinBox.SetMaximum(255)
	blueSpinBox.SetMinimum(0)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	// Add RGB input box to target color layout
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	// Create off-color batch color selection box
	offColorBatch := widgets.NewQGroupBox2("Off-color batch", nil)
	offColorBatchLayout := widgets.NewQHBoxLayout()
	offColorBatch.SetLayout(offColorBatchLayout)
	// Construct box for batch red input
	batchRedGroup := widgets.NewQGroupBox(nil)
	batchRedLabel := widgets.NewQLabel2("R:", nil, 0)
	batchRedSpinBox := widgets.NewQSpinBox(nil)
	batchRedSpinBox.SetMaximum(255)
	batchRedSpinBox.SetMinimum(0)
	batchRedLayout := widgets.NewQGridLayout2()
	batchRedLayout.AddWidget(batchRedLabel, 0, 0, 0)
	batchRedLayout.AddWidget(batchRedSpinBox, 0, 1, 0)
	batchRedGroup.SetLayout(batchRedLayout)
	// Construct box for batch green input
	batchGreenGroup := widgets.NewQGroupBox(nil)
	batchGreenLabel := widgets.NewQLabel2("G:", nil, 0)
	batchGreenSpinBox := widgets.NewQSpinBox(nil)
	batchGreenSpinBox.SetMaximum(255)
	batchGreenSpinBox.SetMinimum(0)
	batchGreenLayout := widgets.NewQGridLayout2()
	batchGreenLayout.AddWidget(batchGreenLabel, 0, 0, 0)
	batchGreenLayout.AddWidget(batchGreenSpinBox, 0, 1, 0)
	batchGreenGroup.SetLayout(batchGreenLayout)
	// Construct box for batch blue input
	batchBlueGroup := widgets.NewQGroupBox(nil)
	batchBlueLabel := widgets.NewQLabel2("B:", nil, 0)
	batchBlueSpinBox := widgets.NewQSpinBox(nil)
	batchBlueSpinBox.SetMaximum(255)
	batchBlueSpinBox.SetMinimum(0)
	batchBlueLayout := widgets.NewQGridLayout2()
	batchBlueLayout.AddWidget(batchBlueLabel, 0, 0, 0)
	batchBlueLayout.AddWidget(batchBlueSpinBox, 0, 1, 0)
	batchBlueGroup.SetLayout(batchBlueLayout)
	// Add RGB input box to off-color batch layout
	offColorBatchLayout.AddWidget(batchRedGroup, 0, 0)
	offColorBatchLayout.AddWidget(batchGreenGroup, 0, 0)
	offColorBatchLayout.AddWidget(batchBlueGroup, 0, 0)
	// Create load repository button
	loadButton := widgets.NewQPushButton2("Load repository", nil)
	loadButton.ConnectClicked(func(_ bool) { openFile() })
	// Create make hits button
	makeHitsButton := widgets.NewQPushButton2("Make hits!", nil)
	makeHitsButton.ConnectClicked(func(_ bool) {makeHits()} )
	// Create back button
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToIndexFromFunction(offColorWindow) })
	layout.AddWidget(instructionText, 0, 0)
	layout.AddWidget(targetColor, 0, 0)
	layout.AddWidget(offColorBatch, 0, 0)
	layout.AddWidget(loadButton, 0, 0)
	layout.AddWidget(makeHitsButton, 0, 0)
	layout.AddWidget(backButton, 0, 0)
}

func CreateColorRangeWindow() {
	colorRangeWindow = widgets.NewQMainWindow(nil, 0)
	colorRangeWindow.SetWindowTitle("Color range")
	colorRangeWindow.SetMinimumSize2(500,500)
	colorRangeWindow.SetMaximumSize2(500,500)
	mainWidget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout()
	mainWidget.SetLayout(layout)
	colorRangeWindow.SetCentralWidget(mainWidget)
	// Create instruction textbox
	instructionTextContent := "The function takes the current repository of pigments listed within it and return"+
	" the entire range of color able to be achieved."
    instructionText := widgets.NewQLabel(nil, 0)
    instructionText.SetWordWrap(true)
    instructionText.SetText(instructionTextContent)
	// Create load repository button
	loadButton := widgets.NewQPushButton2("Load repository", nil)
	loadButton.ConnectClicked(func(_ bool) { openFile() })
	// Create report color range button
	rangeButton := widgets.NewQPushButton2("Color range", nil)
	rangeButton.ConnectClicked(func(_ bool) { colorRange() })
	// Create back button (go back to index window)
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToIndexFromFunction(colorRangeWindow) } )
	layout.AddWidget(instructionText, 0, 0)
	layout.AddWidget(loadButton, 0, 0)
	layout.AddWidget(rangeButton, 0, 0)
	layout.AddWidget(backButton, 0, 0)
}

func CreateQuestionLayout(questionContent string) (*widgets.QHBoxLayout, *widgets.QPushButton) {
    // This function is to create a box containing a question and a question mark button
    // that directs to the answer when clicked, for the help window
    questionLayout := widgets.NewQHBoxLayout()
    questionLayout.SetContentsMargins(0, 0, 0, 0)
    question := widgets.NewQLabel(nil, 0)
    question.SetWordWrap(true)
    question.SetText(questionContent)
    questionLayout.AddWidget(question, 0, 0)
    questionButton := widgets.NewQPushButton(nil)
    questionButton.SetIcon(helpMainWidget.Style().StandardIcon(widgets.QStyle__SP_MessageBoxQuestion, nil, nil))
    questionButton.SetMaximumSize2(20,20)
    questionLayout.AddWidget(questionButton, 0, 0)
    return questionLayout, questionButton
}

func CreateAnswerWindow(questionContent string, answerContent string) *widgets.QMainWindow {
    answerWindow := widgets.NewQMainWindow(nil, 0)
    answerWindow.SetWindowTitle(questionContent)
	answerWindow.SetMinimumSize2(500, 500)
    answerWindow.SetMaximumSize2(500, 500)
    layout := widgets.NewQVBoxLayout()
    mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)
    answer := widgets.NewQLabel(nil, 0)
    answer.SetWordWrap(true)
    answer.SetText(answerContent)
    layout.AddWidget(answer, 0, 0)
    backButton := widgets.NewQPushButton2("Back", nil)
    backButton.ConnectClicked(func(_ bool) { goBackToHelp(answerWindow) } )
    layout.AddWidget(backButton, 0, 0)
    answerWindow.SetCentralWidget(mainWidget)
    return answerWindow
}

func openHelp() {
    // Show the help window hand hide the index window
    helpWindow.Show()
    indexWindow.Hide()
}

func goBackFromHelp() {
    // Go back to the index window from the help window
    helpWindow.Hide()
    indexWindow.Show()
}

func openAnswer(answerWindow *widgets.QMainWindow) {
    // Open an answer window and hide the help window
    answerWindow.Show()
    helpWindow.Hide()
}

func goBackToHelp(answerWindow *widgets.QMainWindow) {
    // Go back to the help window from the current answer window
    answerWindow.Hide()
    helpWindow.Show()
}

func openFunction(functionWindow *widgets.QMainWindow) {
	functionWindow.Show()
	indexWindow.Hide()
}

func backToIndexFromFunction(functionWindow *widgets.QMainWindow) {
	functionWindow.Hide()
	indexWindow.Show()
}

func backToMixFromResults() {
	mixColorWindow.Show()
	mixColorOutputWindow.Hide()
}

func mixColor(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox, quantitySpinBox *widgets.QDoubleSpinBox) {
	if !repositoryLoaded {
		noFileErrorMessage := "Sorry! You have not uploaded a valid repository file."
		widgets.QMessageBox_Information(nil, "OK", noFileErrorMessage,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else{
		// Get RGB and required quantity value from the boxes
		targetR := redSpinBox.Value()
		targetG := greenSpinBox.Value()
		targetB := blueSpinBox.Value()
		quantityRequired := quantitySpinBox.Value()
		func1Exists = false
		func1InStock = false
		func1Composition = []Component{}
		repositoryLoaded, repository = ReadFile(repositoryFileName)
		func1Exists, func1InStock, func1Composition = FindComponents(targetR, targetG, targetB, quantityRequired, repository)
		ShowMixColorDialog(func1Exists, func1InStock, func1Composition)
		
		targetRstring := strconv.FormatBool(func1Exists)
		widgets.QMessageBox_Information(nil, "OK", targetRstring,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		//CreateMixColorOutputWindow(func1Exists, func1InStock, func1Composition)
		//mixColorOutputWindow.Show()
		//mixColorWindow.Hide()
	}
}

func CreateMixColorOutputWindow(exists, inStock bool, composition []Component) {
	mixColorOutputWindow = widgets.NewQMainWindow(nil, 0)
	mixColorOutputWindow.SetWindowTitle("Mix target color results")
	mixColorOutputWindow.SetMinimumSize2(500,500)
	//mixColorOutputWindow.SetMaximumSize2(500,500)
	mixColorOutputMainWidget = widgets.NewQWidget(nil, 0)
	mixColorOutputLayout = widgets.NewQVBoxLayout()
	mixColorOutputMainWidget.SetLayout(mixColorOutputLayout)
	mixColorOutputWindow.SetCentralWidget(mixColorOutputMainWidget)
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToMixFromResults() } )
	if !exists {
		resultTextContent := "We are very sorry. With the given repository, your target color cannot be obtained." +
		" Perhaps you could try another repository or some other target color."
		resultText := widgets.NewQLabel(nil, 0)
		resultText.SetWordWrap(true)
		resultText.SetText(resultTextContent)
		mixColorOutputLayout.AddWidget(resultText, 0, 0)
	} else if !inStock {
		resultTextContent := "Congratulations! Your target color can be mixed using the following pigments. But "+
		"note that one or more pigments are out of stock."
		resultText := widgets.NewQLabel(nil, 0)
		resultText.SetWordWrap(true)
		resultText.SetText(resultTextContent)
		componentLayout := widgets.NewQGridLayout2()
		componentNum := len(composition)
		pigmentLabel := widgets.NewQLabel2("Pigment", nil, 0)
		percentageLabel := widgets.NewQLabel2("Percentage", nil, 0)
		inStockLabel := widgets.NewQLabel2("In stock", nil, 0)
		componentLayout.AddWidget(pigmentLabel, 0, 0, 0)
		componentLayout.AddWidget(percentageLabel, 0, 1, 0)
		componentLayout.AddWidget(inStockLabel, 0, 2, 0)
		for i:=0; i<componentNum; i++ {
			component := composition[i]
			pigment := component.pigment
			pigmentName := pigment.name
			pigmentPercentage := strconv.FormatFloat(component.percentage, 'f', -1, 64)
			var pigmentInStock string
			if component.inStock {
				pigmentInStock = "yes"
			} else {
				pigmentInStock = "no"
			}
			pigmentNameLabel := widgets.NewQLabel2(pigmentName, nil, 0)
			pigmentPercentageLabel := widgets.NewQLabel2(pigmentPercentage, nil, 0)
			pigmentInStockLabel := widgets.NewQLabel2(pigmentInStock, nil, 0)
			componentLayout.AddWidget(pigmentNameLabel, i+1, 0, 0)
			componentLayout.AddWidget(pigmentPercentageLabel, i+1, 1, 0)
			componentLayout.AddWidget(pigmentInStockLabel, i+1, 2, 0)
		}
		mixColorOutputLayout.AddWidget(resultText, 0, 0)
		mixColorOutputLayout.AddLayout(componentLayout, 0)
	} else {
		resultTextContent := "Congratulations! Your target color can be mixed using the following pigments (with lowest price)."
		resultText := widgets.NewQLabel(nil, 0)
		resultText.SetWordWrap(true)
		resultText.SetText(resultTextContent)
		componentLayout := widgets.NewQGridLayout2()
		componentNum := len(composition)
		pigmentLabel := widgets.NewQLabel2("Pigment", nil, 0)
		percentageLabel := widgets.NewQLabel2("Percentage", nil, 0)
		componentLayout.AddWidget(pigmentLabel, 0, 0, 0)
		componentLayout.AddWidget(percentageLabel, 0, 1, 0)
		for i:=0; i<componentNum; i++ {
			component := composition[i]
			pigment := component.pigment
			pigmentName := pigment.name
			pigmentPercentage := strconv.FormatFloat(component.percentage, 'f', -1, 64)
			pigmentNameLabel := widgets.NewQLabel2(pigmentName, nil, 0)
			pigmentPercentageLabel := widgets.NewQLabel2(pigmentPercentage, nil, 0)
			componentLayout.AddWidget(pigmentNameLabel, i+1, 0, 0)
			componentLayout.AddWidget(pigmentPercentageLabel, i+1, 1, 0)
		}
		mixColorOutputLayout.AddWidget(resultText, 0, 0)
		mixColorOutputLayout.AddLayout(componentLayout, 0)
	}
	mixColorOutputLayout.AddWidget(backButton, 0, 0)
}

func ShowMixColorDialog(exists, inStock bool, composition []Component) {
	if !exists {
		message :=" We are very sorry. With the given repository, your target color cannot be obtained." +
		" Perhaps you could try another repository or some other target color."
		widgets.QMessageBox_Information(nil, "OK", message,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else {
		CreateMixColorOutputWindow(exists, inStock, composition)
		mixColorOutputWindow.Show()
		mixColorWindow.Hide()
	}
}

func makeHits() {
	if !repositoryLoaded {
		noFileErrorMessage := "Sorry! You have not uploaded a valid repository file."
		widgets.QMessageBox_Information(nil, "OK", noFileErrorMessage,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	}
}

func colorRange() {
	
}
