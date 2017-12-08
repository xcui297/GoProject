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
	offColorWindow *widgets.QMainWindow
	offColorOutputWindow *widgets.QMainWindow
	colorRangeWindow *widgets.QMainWindow
	visualizeWindow *widgets.QMainWindow
	converterWindow *widgets.QMainWindow
	repositoryLoaded bool
	repositoryFileName string
	repository []Pigment
	func1Exists bool
	func1InStock bool
	func1Composition []Component
	rgbToLabWindow *widgets.QMainWindow
	labWindow *widgets.QMainWindow
	xyzWindow *widgets.QMainWindow
	cmyWindow *widgets.QMainWindow
	cmykWindow *widgets.QMainWindow
	hsbWindow *widgets.QMainWindow
)

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)
	// Create index window. The following buttons connect to other windows
	CreateIndexWindow()
	repositoryLoaded = false
	// Create help window
    CreateHelpWindow()
	// Create the mix color window 
	CreateMixColorWindow()
	// Create the off-color hits window 
	CreateOffColorWindow()
	// Create the color range window
	CreateColorRangeWindow()
	// Create color visualization window
	CreateVisualizeWindow()
	// Create color converter window
	CreateConverterWindow()
	// Show the index window
	indexWindow.Show()
	// Execute app
	app.Exec()
}

func openFile() {
	// Select repository data file from the system
	repositoryFileName = widgets.QFileDialog_GetOpenFileName(nil, "Open Repository file", core.QDir_HomePath(), "", "", 0)
	// If repository file is successfully converted to a slice of Pigments, repositoryLoaded is true.
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
	colorVisualizeButton.ConnectClicked(func(_ bool) { openFunction(visualizeWindow) } )
    toolkitLayout.AddWidget(colorVisualizeButton, 0, 0)
	// Create color converter button and add it to toolkit box layout
	converterButton := widgets.NewQPushButton2("Color converter", toolkitBox)
	converterButton.ConnectClicked(func(_ bool) {openFunction(converterWindow)} )
	toolkitLayout.AddWidget(converterButton, 0, 0)
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
    // Create button directing back to the index window and add it to the layout of help window
    backButton := widgets.NewQPushButton2("Back", nil)
    backButton.ConnectClicked(func(_ bool) { goBackFromHelp() })
    helpLayout.AddWidget(backButton, 0, 0)
	// Set help main widget as the central widget of help window
    helpWindow.SetCentralWidget(helpMainWidget)
}

func CreateMixColorWindow() {
	// Create mix color window
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
	visualizeButton.ConnectClicked(func(_ bool) { visualize(redSpinBox, greenSpinBox, blueSpinBox) })
	// Create back button (go back to index window)
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToIndexFromFunction(mixColorWindow) } )
	// Create mix button
	mixButton := widgets.NewQPushButton2("Mix target color!", nil)
	mixButton.ConnectClicked(func(_ bool) {mixColor(redSpinBox, greenSpinBox, blueSpinBox, quantitySpinBox)} )
	// Add instruction text, target color box, and the buttons to layout of the mix color window
	layout.AddWidget(instructionText, 0, 0)
	layout.AddWidget(targetColor, 0, 0)
	layout.AddWidget(visualizeButton, 0, 0)
	layout.AddWidget(loadButton, 0, 0)
	layout.AddWidget(mixButton, 0, 0)
	layout.AddWidget(backButton, 0, 0)
}

func CreateOffColorWindow() {
	// Create off color hits window
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
	offColorBatchLayout := widgets.NewQGridLayout2()
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
	// Construct off-color batch price input
	batchPriceGroup := widgets.NewQGroupBox(nil)
	batchPriceLabel := widgets.NewQLabel2("Price:", nil, 0)
	batchPriceSpinBox := widgets.NewQSpinBox(nil)
	batchPriceSpinBox.SetMaximum(50)
	batchPriceSpinBox.SetMinimum(20)
	batchPriceLayout := widgets.NewQGridLayout2()
	batchPriceLayout.AddWidget(batchPriceLabel, 0, 0, 0)
	batchPriceLayout.AddWidget(batchPriceSpinBox, 0, 1, 0)
	batchPriceGroup.SetLayout(batchPriceLayout)
	// Add RGB and price input box to off-color batch layout
	offColorBatchLayout.AddWidget(batchRedGroup, 0, 0, 0)
	offColorBatchLayout.AddWidget(batchGreenGroup, 0, 1, 0)
	offColorBatchLayout.AddWidget(batchBlueGroup, 1, 0, 0)
	offColorBatchLayout.AddWidget(batchPriceGroup, 1, 1, 0)
	// Create visualization buttons
	visualizeButton1 := widgets.NewQPushButton2("Visualize target color", nil)
	visualizeButton1.ConnectClicked(func(_ bool) { visualize(redSpinBox, greenSpinBox, blueSpinBox) })
	visualizeButton2 := widgets.NewQPushButton2("Visualize batch color", nil)
	visualizeButton2.ConnectClicked(func(_ bool) { visualize(batchRedSpinBox, batchGreenSpinBox, batchBlueSpinBox) })
	// Create load repository button
	loadButton := widgets.NewQPushButton2("Load repository", nil)
	loadButton.ConnectClicked(func(_ bool) { openFile() })
	// Create make hits button
	makeHitsButton := widgets.NewQPushButton2("Make hits!", nil)
	makeHitsButton.ConnectClicked(func(_ bool) {makeHits(redSpinBox, greenSpinBox, blueSpinBox, batchRedSpinBox, batchGreenSpinBox, batchBlueSpinBox, batchPriceSpinBox)} )
	// Create back button
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToIndexFromFunction(offColorWindow) })
	// Add instruction text, target color selection box, off-color batch selection window, and the buttons to
	// layout of the off-color hits window
	layout.AddWidget(instructionText, 0, 0)
	layout.AddWidget(targetColor, 0, 0)
	layout.AddWidget(offColorBatch, 0, 0)
	layout.AddWidget(visualizeButton1, 0, 0)
	layout.AddWidget(visualizeButton2, 0, 0)
	layout.AddWidget(loadButton, 0, 0)
	layout.AddWidget(makeHitsButton, 0, 0)
	layout.AddWidget(backButton, 0, 0)
}

func CreateColorRangeWindow() {
	// Create color range window
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
	// Add the instruction text and the buttons to layout of the color range window
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
	// This function is to create a window for the answer to a certain question in the help window
	// taking content of both the question and the answer as inputs.
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
	// Create back button leading back to help window and add it to the layout of the answer window
    backButton := widgets.NewQPushButton2("Back", nil)
    backButton.ConnectClicked(func(_ bool) { goBackToHelp(answerWindow) } )
    layout.AddWidget(backButton, 0, 0)
    answerWindow.SetCentralWidget(mainWidget)
    return answerWindow
}

func CreateVisualizeWindow() {
	// Create window for color visualization toolkit
	visualizeWindow = widgets.NewQMainWindow(nil,0)
    visualizeWindow.SetWindowTitle("Visualize RGB color")
	visualizeWindow.SetMinimumSize2(500, 500)
    visualizeWindow.SetMaximumSize2(500, 500)
	// Create layout of visualization window
	visualizeLayout := widgets.NewQVBoxLayout()
	// Create main widget of visualization window and set the layout
	visualizeMainWidget := widgets.NewQWidget(nil, 0)
	visualizeMainWidget.SetLayout(visualizeLayout)
	introTextContent := "Please select target RGB color."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    // Add intro text box to layout of index window
    visualizeLayout.AddWidget(introText, 0, 0)
	targetColorLayout := widgets.NewQVBoxLayout()
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
	// Add the RGB selection boxes to target color group box layout
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	// Add target color selection box to layout of the color visualization window
	visualizeLayout.AddLayout(targetColorLayout, 0)
	visualizeWindow.SetCentralWidget(visualizeMainWidget)
	// Create visualization button and back button, and add them to layout of the color visualization window
	visualizeButton := widgets.NewQPushButton2("Visualize", nil)
	visualizeButton.ConnectClicked(func(_ bool) {visualize(redSpinBox, greenSpinBox, blueSpinBox)} )
	visualizeLayout.AddWidget(visualizeButton, 0, 0)
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToIndexFromFunction(visualizeWindow)} )
	visualizeLayout.AddWidget(backButton, 0, 0)
}

func CreateConverterWindow() {
	// Create window for the color converter toolkit
	converterWindow = widgets.NewQMainWindow(nil,0)
    converterWindow.SetWindowTitle("Color converter")
	converterWindow.SetMinimumSize2(500, 500)
    converterWindow.SetMaximumSize2(500, 500)
	// Create layout of converter window
	converterLayout := widgets.NewQVBoxLayout()
	// Create main widget of converter window and set the layout
	converterMainWidget := widgets.NewQWidget(nil, 0)
	converterMainWidget.SetLayout(converterLayout)
	// Create introduction textbox on index window
    introTextContent := "Pick from the following color formats to convert from."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    // Add intro text box to layout of converter window
    converterLayout.AddWidget(introText, 0, 0)
	
	CreateRGBtoLabWindow()
	CreateLabWindow()
	CreateXYZWindow()
	CreateCMYWindow()
	CreateCMYKWindow()
	CreateHsbWindow()
	rgbButton := widgets.NewQPushButton2("Convert from RGB", nil)
	rgbButton.ConnectClicked(func(_ bool) {openConverterFromConverter(rgbToLabWindow)} )
	converterLayout.AddWidget(rgbButton, 0, 0)
	labButton := widgets.NewQPushButton2("Convert from L*a*b*", nil)
	labButton.ConnectClicked(func(_ bool) {openConverterFromConverter(labWindow)} )
	converterLayout.AddWidget(labButton, 0, 0)
	xyzButton := widgets.NewQPushButton2("Convert from XYZ", nil)
	xyzButton.ConnectClicked(func(_ bool) {openConverterFromConverter(xyzWindow)} )
	converterLayout.AddWidget(xyzButton, 0, 0)
	cmyButton := widgets.NewQPushButton2("Convert from CMY", nil)
	cmyButton.ConnectClicked(func(_ bool) {openConverterFromConverter(cmyWindow)} )
	converterLayout.AddWidget(cmyButton, 0, 0)
	cmykButton := widgets.NewQPushButton2("Convert from CMYK", nil)
	cmykButton.ConnectClicked(func(_ bool) {openConverterFromConverter(cmykWindow)} )
	hsbButton := widgets.NewQPushButton2("Convert from Hsb", nil)
	hsbButton.ConnectClicked(func(_ bool) {openConverterFromConverter(hsbWindow)} )
	converterLayout.AddWidget(hsbButton, 0, 0)
	// Create back button and add it to the layout
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToIndexFromFunction(converterWindow) } )
	converterLayout.AddWidget(backButton, 0, 0)
	//Set converter main widget as the central widget of the color converter window
	converterWindow.SetCentralWidget(converterMainWidget)
}

func CreateRGBtoLabWindow() {
	rgbToLabWindow = widgets.NewQMainWindow(nil,0)
    rgbToLabWindow.SetWindowTitle("RGB to other color spaces")
	rgbToLabWindow.SetMinimumSize2(500, 500)
    rgbToLabWindow.SetMaximumSize2(500, 500)
	rgbToLabLayout := widgets.NewQVBoxLayout()
	rgbToLabMainWidget := widgets.NewQWidget(nil, 0)
	rgbToLabMainWidget.SetLayout(rgbToLabLayout)
	rgbToLabWindow.SetCentralWidget(rgbToLabMainWidget)
    introTextContent := "RGB inputs should be integers from 0 to 255."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    rgbToLabLayout.AddWidget(introText, 0, 0)
	targetColor := widgets.NewQGroupBox2("RGB color", nil)
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
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	toLabButton := widgets.NewQPushButton2("Convert to L*a*b*", nil)
	toLabButton.ConnectClicked(func(_ bool) {rgbToLabDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toXYZButton := widgets.NewQPushButton2("Convert to XYZ", nil)
	toXYZButton.ConnectClicked(func(_ bool) {rgbToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYButton := widgets.NewQPushButton2("Convert to CMY", nil)
	toCMYButton.ConnectClicked(func(_ bool) {rgbToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYKButton := widgets.NewQPushButton2("Convert to CMYK", nil)
	toCMYKButton.ConnectClicked(func(_ bool) {rgbToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toHsbButton := widgets.NewQPushButton2("Convert to Hsb", nil)
	toHsbButton.ConnectClicked(func(_ bool) {rgbToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToConverterFrom(rgbToLabWindow)} )
	rgbToLabLayout.AddWidget(targetColor, 0, 0)
	rgbToLabLayout.AddWidget(toLabButton, 0, 0)
	rgbToLabLayout.AddWidget(toXYZButton, 0, 0)
	rgbToLabLayout.AddWidget(toCMYButton, 0, 0)
	rgbToLabLayout.AddWidget(toCMYKButton, 0, 0)
	rgbToLabLayout.AddWidget(toHsbButton, 0, 0)
	rgbToLabLayout.AddWidget(backButton, 0, 0)
}

func CreateLabWindow() {
	labWindow = widgets.NewQMainWindow(nil,0)
    labWindow.SetWindowTitle("L*a*b* to other color spaces")
	labWindow.SetMinimumSize2(500, 500)
    labWindow.SetMaximumSize2(500, 500)
	labLayout := widgets.NewQVBoxLayout()
	labMainWidget := widgets.NewQWidget(nil, 0)
	labMainWidget.SetLayout(labLayout)
	labWindow.SetCentralWidget(labMainWidget)
    introTextContent := "L should be an integer from 0 to 100. a and b are integers from -128 to 128."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    labLayout.AddWidget(introText, 0, 0)
	targetColor := widgets.NewQGroupBox2("L*a*b* color", nil)
	targetColorLayout := widgets.NewQVBoxLayout()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("L:", nil, 0)
	redSpinBox := widgets.NewQSpinBox(nil)
	redSpinBox.SetMaximum(100)
	redSpinBox.SetMinimum(0)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("a:", nil, 0)
	greenSpinBox := widgets.NewQSpinBox(nil)
	greenSpinBox.SetMaximum(128)
	greenSpinBox.SetMinimum(-128)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("b:", nil, 0)
	blueSpinBox := widgets.NewQSpinBox(nil)
	blueSpinBox.SetMaximum(128)
	blueSpinBox.SetMinimum(-128)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	toRGBButton := widgets.NewQPushButton2("Convert to RGB", nil)
	toRGBButton.ConnectClicked(func(_ bool) {labToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toXYZButton := widgets.NewQPushButton2("Convert to XYZ", nil)
	toXYZButton.ConnectClicked(func(_ bool) {labToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYButton := widgets.NewQPushButton2("Convert to CMY", nil)
	toCMYButton.ConnectClicked(func(_ bool) {labToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYKButton := widgets.NewQPushButton2("Convert to CMYK", nil)
	toCMYKButton.ConnectClicked(func(_ bool) {labToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toHsbButton := widgets.NewQPushButton2("Convert to Hsb", nil)
	toHsbButton.ConnectClicked(func(_ bool) {labToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToConverterFrom(labWindow)} )
	labLayout.AddWidget(targetColor, 0, 0)
	labLayout.AddWidget(toRGBButton, 0, 0)
	labLayout.AddWidget(toXYZButton, 0, 0)
	labLayout.AddWidget(toCMYButton, 0, 0)
	labLayout.AddWidget(toCMYKButton, 0, 0)
	labLayout.AddWidget(toHsbButton, 0, 0)
	labLayout.AddWidget(backButton, 0, 0)
}

func CreateXYZWindow() {
	xyzWindow = widgets.NewQMainWindow(nil,0)
    xyzWindow.SetWindowTitle("XYZ to other color spaces")
	xyzWindow.SetMinimumSize2(500, 500)
    xyzWindow.SetMaximumSize2(500, 500)
	xyzLayout := widgets.NewQVBoxLayout()
	xyzMainWidget := widgets.NewQWidget(nil, 0)
	xyzMainWidget.SetLayout(xyzLayout)
	xyzWindow.SetCentralWidget(xyzMainWidget)
    introTextContent := "XYZ inputs should be float64 numbers from 0 to 1."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    xyzLayout.AddWidget(introText, 0, 0)
	targetColor := widgets.NewQGroupBox2("XYZ color", nil)
	targetColorLayout := widgets.NewQVBoxLayout()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("X:", nil, 0)
	redSpinBox := widgets.NewQDoubleSpinBox(nil)
	redSpinBox.SetMaximum(1.0)
	redSpinBox.SetMinimum(0.0)
	redSpinBox.SetSingleStep(0.01)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("Y:", nil, 0)
	greenSpinBox := widgets.NewQDoubleSpinBox(nil)
	greenSpinBox.SetMaximum(1.0)
	greenSpinBox.SetMinimum(0.0)
	greenSpinBox.SetSingleStep(0.01)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("Z:", nil, 0)
	blueSpinBox := widgets.NewQDoubleSpinBox(nil)
	blueSpinBox.SetMaximum(1.0)
	blueSpinBox.SetMinimum(0.0)
	blueSpinBox.SetSingleStep(0.01)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	toRGBButton := widgets.NewQPushButton2("Convert to RGB", nil)
	toRGBButton.ConnectClicked(func(_ bool) {xyzToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toLabButton := widgets.NewQPushButton2("Convert to L*a*b*", nil)
	toLabButton.ConnectClicked(func(_ bool) {xyzToLabDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYButton := widgets.NewQPushButton2("Convert to CMY", nil)
	toCMYButton.ConnectClicked(func(_ bool) {xyzToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYKButton := widgets.NewQPushButton2("Convert to CMYK", nil)
	toCMYKButton.ConnectClicked(func(_ bool) {xyzToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toHsbButton := widgets.NewQPushButton2("Convert to Hsb", nil)
	toHsbButton.ConnectClicked(func(_ bool) {xyzToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToConverterFrom(xyzWindow)} )
	xyzLayout.AddWidget(targetColor, 0, 0)
	xyzLayout.AddWidget(toRGBButton, 0, 0)
	xyzLayout.AddWidget(toLabButton, 0, 0)
	xyzLayout.AddWidget(toCMYButton, 0, 0)
	xyzLayout.AddWidget(toCMYKButton, 0, 0)
	xyzLayout.AddWidget(toHsbButton, 0, 0)
	xyzLayout.AddWidget(backButton, 0, 0)
}

func CreateCMYWindow() {
	cmyWindow = widgets.NewQMainWindow(nil,0)
    cmyWindow.SetWindowTitle("CMY to other color spaces")
	cmyWindow.SetMinimumSize2(500, 500)
    cmyWindow.SetMaximumSize2(500, 500)
	cmyLayout := widgets.NewQVBoxLayout()
	cmyMainWidget := widgets.NewQWidget(nil, 0)
	cmyMainWidget.SetLayout(cmyLayout)
	cmyWindow.SetCentralWidget(cmyMainWidget)
    introTextContent := "CMY inputs should be float64 numbers from 0 to 1."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    cmyLayout.AddWidget(introText, 0, 0)
	targetColor := widgets.NewQGroupBox2("CMY color", nil)
	targetColorLayout := widgets.NewQVBoxLayout()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("C:", nil, 0)
	redSpinBox := widgets.NewQDoubleSpinBox(nil)
	redSpinBox.SetMaximum(1.0)
	redSpinBox.SetMinimum(0.0)
	redSpinBox.SetSingleStep(0.01)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("M:", nil, 0)
	greenSpinBox := widgets.NewQDoubleSpinBox(nil)
	greenSpinBox.SetMaximum(1.0)
	greenSpinBox.SetMinimum(0.0)
	greenSpinBox.SetSingleStep(0.01)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("Y:", nil, 0)
	blueSpinBox := widgets.NewQDoubleSpinBox(nil)
	blueSpinBox.SetMaximum(1.0)
	blueSpinBox.SetMinimum(0.0)
	blueSpinBox.SetSingleStep(0.01)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	toRGBButton := widgets.NewQPushButton2("Convert to RGB", nil)
	toRGBButton.ConnectClicked(func(_ bool) {cmyToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toLabButton := widgets.NewQPushButton2("Convert to L*a*b*", nil)
	toLabButton.ConnectClicked(func(_ bool) {cmyToLabDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toXYZButton := widgets.NewQPushButton2("Convert to XYZ", nil)
	toXYZButton.ConnectClicked(func(_ bool) {cmyToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYKButton := widgets.NewQPushButton2("Convert to CMYK", nil)
	toCMYKButton.ConnectClicked(func(_ bool) {cmyToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toHsbButton := widgets.NewQPushButton2("Convert to Hsb", nil)
	toHsbButton.ConnectClicked(func(_ bool) {cmyToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToConverterFrom(cmyWindow)} )
	cmyLayout.AddWidget(targetColor, 0, 0)
	cmyLayout.AddWidget(toRGBButton, 0, 0)
	cmyLayout.AddWidget(toLabButton, 0, 0)
	cmyLayout.AddWidget(toXYZButton, 0, 0)
	cmyLayout.AddWidget(toCMYKButton, 0, 0)
	cmyLayout.AddWidget(toHsbButton, 0, 0)
	cmyLayout.AddWidget(backButton, 0, 0)
}

func CreateCMYKWindow() {
	cmykWindow = widgets.NewQMainWindow(nil,0)
    cmykWindow.SetWindowTitle("CMYK to other color spaces")
	cmykWindow.SetMinimumSize2(500, 500)
    cmykWindow.SetMaximumSize2(500, 500)
	cmykLayout := widgets.NewQVBoxLayout()
	cmykMainWidget := widgets.NewQWidget(nil, 0)
	cmykMainWidget.SetLayout(cmykLayout)
	cmykWindow.SetCentralWidget(cmykMainWidget)
    introTextContent := "CMYK inputs should be float64 numbers from 0 to 1."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    cmykLayout.AddWidget(introText, 0, 0)
	targetColor := widgets.NewQGroupBox2("CMYK color", nil)
	targetColorLayout := widgets.NewQGridLayout2()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("C:", nil, 0)
	redSpinBox := widgets.NewQDoubleSpinBox(nil)
	redSpinBox.SetMaximum(1.0)
	redSpinBox.SetMinimum(0.0)
	redSpinBox.SetSingleStep(0.01)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("M:", nil, 0)
	greenSpinBox := widgets.NewQDoubleSpinBox(nil)
	greenSpinBox.SetMaximum(1.0)
	greenSpinBox.SetMinimum(0.0)
	greenSpinBox.SetSingleStep(0.01)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("Y:", nil, 0)
	blueSpinBox := widgets.NewQDoubleSpinBox(nil)
	blueSpinBox.SetMaximum(1.0)
	blueSpinBox.SetMinimum(0.0)
	blueSpinBox.SetSingleStep(0.01)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	kGroup := widgets.NewQGroupBox(nil)
	kLabel := widgets.NewQLabel2("K:", nil, 0)
	kSpinBox := widgets.NewQDoubleSpinBox(nil)
	kSpinBox.SetMaximum(1.0)
	kSpinBox.SetMinimum(0.0)
	kSpinBox.SetSingleStep(0.01)
	kLayout := widgets.NewQGridLayout2()
	kLayout.AddWidget(kLabel, 0, 0, 0)
	kLayout.AddWidget(kSpinBox, 0, 1, 0)
	kGroup.SetLayout(kLayout)
	targetColorLayout.AddWidget(redGroup, 0, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 1, 0)
	targetColorLayout.AddWidget(blueGroup, 1, 0, 0)
	targetColorLayout.AddWidget(kGroup, 1, 1, 0)
	toRGBButton := widgets.NewQPushButton2("Convert to RGB", nil)
	toRGBButton.ConnectClicked(func(_ bool) {cmykToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox)} )
	toLabButton := widgets.NewQPushButton2("Convert to L*a*b*", nil)
	toLabButton.ConnectClicked(func(_ bool) {cmykToLabDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox)} )
	toXYZButton := widgets.NewQPushButton2("Convert to XYZ", nil)
	toXYZButton.ConnectClicked(func(_ bool) {cmykToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox)} )
	toCMYButton := widgets.NewQPushButton2("Convert to CMY", nil)
	toCMYButton.ConnectClicked(func(_ bool) {cmykToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox)} )
	toHsbButton := widgets.NewQPushButton2("Convert to Hsb", nil)
	toHsbButton.ConnectClicked(func(_ bool) {cmykToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox)} )
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToConverterFrom(cmykWindow)} )
	cmykLayout.AddWidget(targetColor, 0, 0)
	cmykLayout.AddWidget(toRGBButton, 0, 0)
	cmykLayout.AddWidget(toLabButton, 0, 0)
	cmykLayout.AddWidget(toXYZButton, 0, 0)
	cmykLayout.AddWidget(toCMYButton, 0, 0)
	cmykLayout.AddWidget(toHsbButton, 0, 0)
	cmykLayout.AddWidget(backButton, 0, 0)
}

func CreateHsbWindow() {
	hsbWindow = widgets.NewQMainWindow(nil,0)
    hsbWindow.SetWindowTitle("Hsb to other color spaces")
	hsbWindow.SetMinimumSize2(500, 500)
    hsbWindow.SetMaximumSize2(500, 500)
	hsbLayout := widgets.NewQVBoxLayout()
	hsbMainWidget := widgets.NewQWidget(nil, 0)
	hsbMainWidget.SetLayout(hsbLayout)
	hsbWindow.SetCentralWidget(hsbMainWidget)
    introTextContent := "H should be an integer from 0 to 360. s and b are float64 numbers from 0 to 1."
    introText := widgets.NewQLabel(nil, 0)
    introText.SetWordWrap(true)
    introText.SetText(introTextContent)
    hsbLayout.AddWidget(introText, 0, 0)
	targetColor := widgets.NewQGroupBox2("Hsb color", nil)
	targetColorLayout := widgets.NewQVBoxLayout()
	targetColor.SetLayout(targetColorLayout)
	// Construct box for red input
	redGroup := widgets.NewQGroupBox(nil)
	redLabel := widgets.NewQLabel2("H:", nil, 0)
	redSpinBox := widgets.NewQSpinBox(nil)
	redSpinBox.SetMaximum(360)
	redSpinBox.SetMinimum(0)
	redLayout := widgets.NewQGridLayout2()
	redLayout.AddWidget(redLabel, 0, 0, 0)
	redLayout.AddWidget(redSpinBox, 0, 1, 0)
	redGroup.SetLayout(redLayout)
	// Construct box for green input
	greenGroup := widgets.NewQGroupBox(nil)
	greenLabel := widgets.NewQLabel2("s:", nil, 0)
	greenSpinBox := widgets.NewQDoubleSpinBox(nil)
	greenSpinBox.SetMaximum(1.0)
	greenSpinBox.SetMinimum(0.0)
	greenSpinBox.SetSingleStep(0.01)
	greenLayout := widgets.NewQGridLayout2()
	greenLayout.AddWidget(greenLabel, 0, 0, 0)
	greenLayout.AddWidget(greenSpinBox, 0, 1, 0)
	greenGroup.SetLayout(greenLayout)
	// Construct box for blue input
	blueGroup := widgets.NewQGroupBox(nil)
	blueLabel := widgets.NewQLabel2("b:", nil, 0)
	blueSpinBox := widgets.NewQDoubleSpinBox(nil)
	blueSpinBox.SetMaximum(1.0)
	blueSpinBox.SetMinimum(0.0)
	blueSpinBox.SetSingleStep(0.01)
	blueLayout := widgets.NewQGridLayout2()
	blueLayout.AddWidget(blueLabel, 0, 0, 0)
	blueLayout.AddWidget(blueSpinBox, 0, 1, 0)
	blueGroup.SetLayout(blueLayout)
	targetColorLayout.AddWidget(redGroup, 0, 0)
	targetColorLayout.AddWidget(greenGroup, 0, 0)
	targetColorLayout.AddWidget(blueGroup, 0, 0)
	toRGBButton := widgets.NewQPushButton2("Convert to RGB", nil)
	toRGBButton.ConnectClicked(func(_ bool) {hsbToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toLabButton := widgets.NewQPushButton2("Convert to L*a*b*", nil)
	toLabButton.ConnectClicked(func(_ bool) {hsbToLabDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toXYZButton := widgets.NewQPushButton2("Convert to XYZ", nil)
	toXYZButton.ConnectClicked(func(_ bool) {hsbToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYButton := widgets.NewQPushButton2("Convert to CMY", nil)
	toCMYButton.ConnectClicked(func(_ bool) {hsbToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	toCMYKButton := widgets.NewQPushButton2("Convert to CMYK", nil)
	toCMYKButton.ConnectClicked(func(_ bool) {hsbToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox)} )
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) {backToConverterFrom(hsbWindow)} )
	hsbLayout.AddWidget(targetColor, 0, 0)
	hsbLayout.AddWidget(toRGBButton, 0, 0)
	hsbLayout.AddWidget(toLabButton, 0, 0)
	hsbLayout.AddWidget(toXYZButton, 0, 0)
	hsbLayout.AddWidget(toCMYButton, 0, 0)
	hsbLayout.AddWidget(toCMYKButton, 0, 0)
	hsbLayout.AddWidget(backButton, 0, 0)
}

func hsbToCMYKDialog(redSpinBox *widgets.QSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	h := redSpinBox.Value()
	s := greenSpinBox.Value()
	b := blueSpinBox.Value()
	C, M, Y, K := ConvertHsbToCMYK(h, s, b)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	KString := strconv.FormatFloat(K, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString+", K: "+KString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func hsbToCMYDialog(redSpinBox *widgets.QSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	h := redSpinBox.Value()
	s := greenSpinBox.Value()
	b := blueSpinBox.Value()
	C, M, Y := ConvertHsbToCMY(h, s, b)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func hsbToXYZDialog(redSpinBox *widgets.QSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	h := redSpinBox.Value()
	s := greenSpinBox.Value()
	b := blueSpinBox.Value()
	X, Y, Z := ConvertHsbToXYZ(h, s, b)
	XString := strconv.FormatFloat(X, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	ZString := strconv.FormatFloat(Z, 'f', 4, 64)
	message := "X: "+XString+", Y: "+YString+", Z: "+ZString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func hsbToLabDialog(redSpinBox *widgets.QSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	h := redSpinBox.Value()
	s := greenSpinBox.Value()
	b := blueSpinBox.Value()
	L, A, B := ConvertHsbToLab(h, s, b)
	LString := strconv.FormatFloat(L, 'f', 0, 64)
	AString := strconv.FormatFloat(A, 'f', 0, 64)
	BString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "L: "+LString+", A: "+AString+", B: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func hsbToRGBDialog(redSpinBox *widgets.QSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	h := redSpinBox.Value()
	s := greenSpinBox.Value()
	b := blueSpinBox.Value()
	R, G, B := ConvertHsbToRGB(h, s, b)
	rString := strconv.FormatFloat(R, 'f', 0, 64)
	gString := strconv.FormatFloat(G, 'f', 0, 64)
	bString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "R: "+rString+", G: "+gString+", B: "+bString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmykToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	k := kSpinBox.Value()
	H, S, B := ConvertCMYKToHsb(c, m, y, k)
	HString := strconv.FormatFloat(H, 'f', 0, 64)
	SString := strconv.FormatFloat(S, 'f', 4, 64)
	BString := strconv.FormatFloat(B, 'f', 4, 64)
	message := "H: "+HString+", s: "+SString+", b: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmykToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	k := kSpinBox.Value()
	C, M, Y := ConvertCMYKToCMY(c, m, y, k)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmykToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	k := kSpinBox.Value()
	X, Y, Z := ConvertCMYKToXYZ(c, m, y, k)
	XString := strconv.FormatFloat(X, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	ZString := strconv.FormatFloat(Z, 'f', 4, 64)
	message := "X: "+XString+", Y: "+YString+", Z: "+ZString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmykToLabDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	k := kSpinBox.Value()
	L, A, B := ConvertCMYKToLab(c, m, y, k)
	LString := strconv.FormatFloat(L, 'f', 0, 64)
	AString := strconv.FormatFloat(A, 'f', 0, 64)
	BString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "L: "+LString+", A: "+AString+", B: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmykToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox, kSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	k := kSpinBox.Value()
	r, g, b := ConvertCMYKToRGB(c, m, y, k)
	rString := strconv.FormatFloat(r, 'f', 0, 64)
	gString := strconv.FormatFloat(g, 'f', 0, 64)
	bString := strconv.FormatFloat(b, 'f', 0, 64)
	message := "R: "+rString+", G: "+gString+", B: "+bString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmyToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	H, S, B := ConvertCMYToHsb(c, m, y)
	HString := strconv.FormatFloat(H, 'f', 0, 64)
	SString := strconv.FormatFloat(S, 'f', 4, 64)
	BString := strconv.FormatFloat(B, 'f', 4, 64)
	message := "H: "+HString+", s: "+SString+", b: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmyToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	C, M, Y, K := ConvertCMYToCMYK(c, m, y)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	KString := strconv.FormatFloat(K, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString+", K: "+KString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmyToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	X, Y, Z := ConvertCMYToXYZ(c, m, y)
	XString := strconv.FormatFloat(X, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	ZString := strconv.FormatFloat(Z, 'f', 4, 64)
	message := "X: "+XString+", Y: "+YString+", Z: "+ZString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmyToLabDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	L, A, B := ConvertCMYToLab(c, m, y)
	LString := strconv.FormatFloat(L, 'f', 0, 64)
	AString := strconv.FormatFloat(A, 'f', 0, 64)
	BString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "L: "+LString+", A: "+AString+", B: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func cmyToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	c := redSpinBox.Value()
	m := greenSpinBox.Value()
	y := blueSpinBox.Value()
	r, g, b := ConvertCMYToRGB(c, m, y)
	rString := strconv.FormatFloat(r, 'f', 0, 64)
	gString := strconv.FormatFloat(g, 'f', 0, 64)
	bString := strconv.FormatFloat(b, 'f', 0, 64)
	message := "R: "+rString+", G: "+gString+", B: "+bString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func xyzToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	x := redSpinBox.Value()
	y := greenSpinBox.Value()
	z := blueSpinBox.Value()
	H, S, B := ConvertXYZToHsb(x, y, z)
	HString := strconv.FormatFloat(H, 'f', 0, 64)
	SString := strconv.FormatFloat(S, 'f', 4, 64)
	BString := strconv.FormatFloat(B, 'f', 4, 64)
	message := "H: "+HString+", s: "+SString+", b: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func xyzToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	x := redSpinBox.Value()
	y := greenSpinBox.Value()
	z := blueSpinBox.Value()
	C, M, Y, K := ConvertXYZToCMYK(x, y, z)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	KString := strconv.FormatFloat(K, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString+", K: "+KString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func xyzToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	x := redSpinBox.Value()
	y := greenSpinBox.Value()
	z := blueSpinBox.Value()
	C, M, Y := ConvertXYZToCMY(x, y, z)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func xyzToLabDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	x := redSpinBox.Value()
	y := greenSpinBox.Value()
	z := blueSpinBox.Value()
	L, A, B := ConvertXYZToLab(x, y, z)
	LString := strconv.FormatFloat(L, 'f', 0, 64)
	AString := strconv.FormatFloat(A, 'f', 0, 64)
	BString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "L: "+LString+", A: "+AString+", B: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func xyzToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QDoubleSpinBox) {
	x := redSpinBox.Value()
	y := greenSpinBox.Value()
	z := blueSpinBox.Value()
	r, g, b := ConvertXYZToRGB(x, y, z)
	rString := strconv.FormatFloat(r, 'f', 0, 64)
	gString := strconv.FormatFloat(g, 'f', 0, 64)
	bString := strconv.FormatFloat(b, 'f', 0, 64)
	message := "R: "+rString+", G: "+gString+", B: "+bString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func labToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	l := redSpinBox.Value()
	a := greenSpinBox.Value()
	b := blueSpinBox.Value()
	H, S, B := ConvertLabToHsb(l, a, b)
	HString := strconv.FormatFloat(H, 'f', 0, 64)
	SString := strconv.FormatFloat(S, 'f', 4, 64)
	BString := strconv.FormatFloat(B, 'f', 4, 64)
	message := "H: "+HString+", s: "+SString+", b: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func labToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	l := redSpinBox.Value()
	a := greenSpinBox.Value()
	b := blueSpinBox.Value()
	C, M, Y, K := ConvertLabToCMYK(l, a, b)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	KString := strconv.FormatFloat(K, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString+", K: "+KString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func labToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	l := redSpinBox.Value()
	a := greenSpinBox.Value()
	b := blueSpinBox.Value()
	C, M, Y := ConvertLabToCMY(l, a, b)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func labToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	l := redSpinBox.Value()
	a := greenSpinBox.Value()
	b := blueSpinBox.Value()
	X, Y, Z := ConvertLabToXYZ(l, a, b)
	XString := strconv.FormatFloat(X, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	ZString := strconv.FormatFloat(Z, 'f', 4, 64)
	message := "X: "+XString+", Y: "+YString+", Z: "+ZString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func labToRGBDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	l := redSpinBox.Value()
	a := greenSpinBox.Value()
	b := blueSpinBox.Value()
	R, G, B := ConvertLabToRGB(l, a, b)
	RString := strconv.FormatFloat(R, 'f', 0, 64)
	GString := strconv.FormatFloat(G, 'f', 0, 64)
	BString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "R: "+RString+", G: "+GString+", B: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func backToConverterFrom(functionWindow *widgets.QMainWindow) {
	functionWindow.Hide()
	converterWindow.Show()
}

func openConverterFromConverter(functionWindow *widgets.QMainWindow) {
	functionWindow.Show()
	converterWindow.Hide()
}

func rgbToHsbDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	r := redSpinBox.Value()
	g := greenSpinBox.Value()
	b := blueSpinBox.Value()
	H, S, B := ConvertRGBToHsb(r, g, b)
	HString := strconv.FormatFloat(H, 'f', 0, 64)
	SString := strconv.FormatFloat(S, 'f', 4, 64)
	BString := strconv.FormatFloat(B, 'f', 4, 64)
	message := "H: "+HString+", s: "+SString+", b: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func rgbToCMYKDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	r := redSpinBox.Value()
	g := greenSpinBox.Value()
	b := blueSpinBox.Value()
	C, M, Y, K := ConvertRGBToCMYK(r, g, b)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	KString := strconv.FormatFloat(K, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString+", K: "+KString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func rgbToCMYDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	r := redSpinBox.Value()
	g := greenSpinBox.Value()
	b := blueSpinBox.Value()
	C, M, Y := ConvertRGBToCMY(r, g, b)
	CString := strconv.FormatFloat(C, 'f', 4, 64)
	MString := strconv.FormatFloat(M, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	message := "C: "+CString+", M: "+MString+", Y: "+YString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func rgbToXYZDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	r := redSpinBox.Value()
	g := greenSpinBox.Value()
	b := blueSpinBox.Value()
	X, Y, Z := ConvertRGBToXYZ(r, g, b)
	XString := strconv.FormatFloat(X, 'f', 4, 64)
	YString := strconv.FormatFloat(Y, 'f', 4, 64)
	ZString := strconv.FormatFloat(Z, 'f', 4, 64)
	message := "X: "+XString+", Y: "+YString+", Z: "+ZString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func rgbToLabDialog(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	r := redSpinBox.Value()
	g := greenSpinBox.Value()
	b := blueSpinBox.Value()
	L, A, B := ConvertRGBToLab(r, g, b)
	LString := strconv.FormatFloat(L, 'f', 0, 64)
	AString := strconv.FormatFloat(A, 'f', 0, 64)
	BString := strconv.FormatFloat(B, 'f', 0, 64)
	message := "L: "+LString+", A: "+AString+", B: "+BString
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func visualize(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox) {
	// Visualize a given RGB color using the DrawCar function. The png file is to be saved to the folder
	red := uint8(redSpinBox.Value())
	green := uint8(greenSpinBox.Value())
	blue := uint8(blueSpinBox.Value())
	carFileName := DrawCar(red, green, blue)
	// Report to the user the file name
	message := "Your color visualization picture is saved as " + carFileName+" in the folder."
	widgets.QMessageBox_Information(nil, "OK", message,
	widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
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
	// Open a function window from the index window
	functionWindow.Show()
	indexWindow.Hide()
}

func backToIndexFromFunction(functionWindow *widgets.QMainWindow) {
	// Go back to the index window from a function window
	functionWindow.Hide()
	indexWindow.Show()
}

func backToMixFromResults() {
	// Go back to the mix color window from the results window
	mixColorWindow.Show()
	mixColorOutputWindow.Hide()
}

func mixColor(redSpinBox, greenSpinBox, blueSpinBox *widgets.QSpinBox, quantitySpinBox *widgets.QDoubleSpinBox) {
	// Mix color. Get input from the QSpinBoxes and call the FindComponents function
	if !repositoryLoaded {
		// Check whether a valid repository file has been loaded
		noFileErrorMessage := "Sorry! You have not uploaded a valid repository file."
		widgets.QMessageBox_Information(nil, "OK", noFileErrorMessage,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else{
		// Get RGB and required quantity value from the boxes
		targetR := redSpinBox.Value()
		targetG := greenSpinBox.Value()
		targetB := blueSpinBox.Value()
		quantityRequired := quantitySpinBox.Value()
		exists, inStock, composition, _:= FindComponents(targetR, targetG, targetB, quantityRequired, repository)
		// Show the results dialog
		ShowMixColorDialog(exists, inStock, composition)
		//CreateMixColorOutputWindow(func1Exists, func1InStock, func1Composition)
		//mixColorOutputWindow.Show()
		//mixColorWindow.Hide()
	}
}

func CreateMixColorOutputWindow(inStock bool, composition []Component) {
	// Create output window for the mix color function
	mixColorOutputWindow = widgets.NewQMainWindow(nil, 0)
	mixColorOutputWindow.SetWindowTitle("Mix target color results")
	mixColorOutputWindow.SetMinimumSize2(500,500)
	mixColorOutputWindow.SetMaximumSize2(500,500)
	mixColorOutputMainWidget := widgets.NewQWidget(nil, 0)
	mixColorOutputLayout := widgets.NewQVBoxLayout()
	mixColorOutputMainWidget.SetLayout(mixColorOutputLayout)
	mixColorOutputWindow.SetCentralWidget(mixColorOutputMainWidget)
	// Create back button
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToMixFromResults() } )
	/*if !exists {
		resultTextContent := "We are very sorry. With the given repository, your target color cannot be obtained." +
		" Perhaps you could try another repository or some other target color."
		resultText := widgets.NewQLabel(nil, 0)
		resultText.SetWordWrap(true)
		resultText.SetText(resultTextContent)
		mixColorOutputLayout.AddWidget(resultText, 0, 0)*/
	if !inStock {
		// When ont or more pigments are out of stock
		resultTextContent := "Congratulations! Your target color can be mixed using the following pigments. But "+
		"note that one or more pigments are out of stock."
		resultText := widgets.NewQLabel(nil, 0)
		resultText.SetWordWrap(true)
		resultText.SetText(resultTextContent)
		componentLayout := widgets.NewQGridLayout2()
		componentNum := len(composition)
		// There are three colomns of the outpus: pigment, percentage and in strock or not
		pigmentLabel := widgets.NewQLabel2("Pigment", nil, 0)
		percentageLabel := widgets.NewQLabel2("Percentage", nil, 0)
		inStockLabel := widgets.NewQLabel2("In stock", nil, 0)
		componentLayout.AddWidget(pigmentLabel, 0, 0, 0)
		componentLayout.AddWidget(percentageLabel, 0, 1, 0)
		componentLayout.AddWidget(inStockLabel, 0, 2, 0)
		// List component one by one
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
		// Add the result text and result table to the layout
		mixColorOutputLayout.AddWidget(resultText, 0, 0)
		mixColorOutputLayout.AddLayout(componentLayout, 0)
	} else {
		// When all components are in stock
		resultTextContent := "Congratulations! Your target color can be mixed using the following pigments (with lowest price)."
		resultText := widgets.NewQLabel(nil, 0)
		resultText.SetWordWrap(true)
		resultText.SetText(resultTextContent)
		componentLayout := widgets.NewQGridLayout2()
		componentNum := len(composition)
		// There are two columns of the output: pigment and percentage
		pigmentLabel := widgets.NewQLabel2("Pigment", nil, 0)
		percentageLabel := widgets.NewQLabel2("Percentage", nil, 0)
		componentLayout.AddWidget(pigmentLabel, 0, 0, 0)
		componentLayout.AddWidget(percentageLabel, 0, 1, 0)
		// Add components to the table one by one
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
		// Add the output text and table to the layout
		mixColorOutputLayout.AddWidget(resultText, 0, 0)
		mixColorOutputLayout.AddLayout(componentLayout, 0)
	}
	// Add back button to the end of the layout
	mixColorOutputLayout.AddWidget(backButton, 0, 0)
}

func ShowMixColorDialog(exists, inStock bool, composition []Component) {
	// Show dialog windows for mix color results
	if !exists {
		// When there is no solution
		message :=" We are very sorry. With the given repository, your target color cannot be obtained." +
		" Perhaps you could try another repository or some other target color."
		widgets.QMessageBox_Information(nil, "OK", message,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else {
		// If there are solutions (whether all components are in stock or not), a new window is needed to display the results
		// Go to the new window created
		CreateMixColorOutputWindow(inStock, composition)
		mixColorOutputWindow.Show()
		mixColorWindow.Hide()
	}
}

func makeHits(redSpinBox, greenSpinBox, blueSpinBox, batchRedSpinBox, batchGreenSpinBox, batchBlueSpinBox, batchPriceSpinBox *widgets.QSpinBox) {
	// Make hits to the off-color batch. Take target RGB, batch RGB and batch price inputs from the spinBoxes
	if !repositoryLoaded {
		// Check whether a valid repository file has been loaded
		noFileErrorMessage := "Sorry! You have not uploaded a valid repository file."
		widgets.QMessageBox_Information(nil, "OK", noFileErrorMessage,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else {
		targetR := redSpinBox.Value()
		targetG := greenSpinBox.Value()
		targetB := blueSpinBox.Value()
		batchR := batchRedSpinBox.Value()
		batchG := batchGreenSpinBox.Value()
		batchB := batchBlueSpinBox.Value()
		batchPrice := batchPriceSpinBox.Value()
		// Call the OffColorHit function (in coreFunctions.go)
		exists, currentPercentage, composition, finalPrice := OffColorHit(batchR, batchG, batchB, targetR, targetG, targetB, batchPrice, repository)
		// Show dialog windows for the result
		ShowMakeHitsDialog(exists, currentPercentage, composition, finalPrice)
	}
}

func ShowMakeHitsDialog(exists bool, currentPercentage float64, composition []Component, finalPrice int) {
	// Show results dialog for make-hits function
	if !exists {
		// When there is no solution, a dialog window pops up
		message :=" We are very sorry. With the given repository, your target color cannot be obtained." +
		" Perhaps you could try another repository or some other target color."
		widgets.QMessageBox_Information(nil, "OK", message,
		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	} else {
		// When there is a solution, a new window is needed to display the results
		// Go to the new result window from the current window
		CreateOffColorOutputWindow(currentPercentage, composition, finalPrice)
		offColorOutputWindow.Show()
		offColorWindow.Hide()
	}
}

func CreateOffColorOutputWindow(currentPercentage float64, composition []Component, finalPrice int) {
	// Create window for output for the off-color hits function
	offColorOutputWindow = widgets.NewQMainWindow(nil, 0)
	offColorOutputWindow.SetWindowTitle("Off-color hits results")
	offColorOutputWindow.SetMinimumSize2(500,500)
	offColorOutputWindow.SetMaximumSize2(500,500)
	offColorOutputMainWidget := widgets.NewQWidget(nil, 0)
	offColorOutputLayout := widgets.NewQVBoxLayout()
	offColorOutputMainWidget.SetLayout(offColorOutputLayout)
	offColorOutputWindow.SetCentralWidget(offColorOutputMainWidget)
	// Create back button
	backButton := widgets.NewQPushButton2("Back", nil)
	backButton.ConnectClicked(func(_ bool) { backToOffFromResults() } )
	// Convert percentage from float64 to string to display a text box
	currentPercentageString := strconv.FormatFloat(currentPercentage, 'f', -1, 64)
	finalPriceString := strconv.Itoa(finalPrice)
	resultTextContent := "Congratulations! Your target color can be mixed using the following pigments. Final"+
	" price is "+finalPriceString+"."
	resultText := widgets.NewQLabel(nil, 0)
	resultText.SetWordWrap(true)
	resultText.SetText(resultTextContent)
	componentLayout := widgets.NewQGridLayout2()
	componentNum := len(composition)
	// There are two columns of the results: pigment and percentage
	pigmentLabel := widgets.NewQLabel2("Pigment", nil, 0)
	percentageLabel := widgets.NewQLabel2("Percentage", nil, 0)
	componentLayout.AddWidget(pigmentLabel, 0, 0, 0)
	componentLayout.AddWidget(percentageLabel, 0, 1, 0)
	// Results for the current off-color batch is displayed first
	batchLabel := widgets.NewQLabel2("Off-color batch", nil, 0)
	batchPercentageLabel := widgets.NewQLabel2(currentPercentageString, nil, 0)
	componentLayout.AddWidget(batchLabel, 1, 0, 0)
	componentLayout.AddWidget(batchPercentageLabel, 1, 1, 0)
	// Components needed are listed one by one
	for i:=0; i<componentNum; i++ {
		component := composition[i]
		pigment := component.pigment
		pigmentName := pigment.name
		pigmentPercentage := strconv.FormatFloat(component.percentage, 'f', -1, 64)
		pigmentNameLabel := widgets.NewQLabel2(pigmentName, nil, 0)
		pigmentPercentageLabel := widgets.NewQLabel2(pigmentPercentage, nil, 0)
		componentLayout.AddWidget(pigmentNameLabel, i+2, 0, 0)
		componentLayout.AddWidget(pigmentPercentageLabel, i+2, 1, 0)
	}
	// Add the result text and table and the back button to the layout.
	offColorOutputLayout.AddWidget(resultText, 0, 0)
	offColorOutputLayout.AddLayout(componentLayout, 0)
	offColorOutputLayout.AddWidget(backButton, 0, 0)
}

func backToOffFromResults() {
	// Go back to the off-color hits function window from the results window
	offColorOutputWindow.Hide()
	offColorWindow.Show()
}

func colorRange() {
	
}
