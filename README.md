# GoProject
Group project for 02-601 <br>
This is a simple color calculator for automotive paints.

# Prerequisites for this package:
1. https://github.com/draffensperger/golp
2. https://github.com/therecipe/qt
3. qt5

# Installation
Follow the installation instructions of the two github repositories to intall them. <br>
To install qt5, you can use homebrew. <br>
You have to have these files from this package in your go directory:
1. main.go
2. coreFunc.go
3. drawCar.go
4. convertColor.go

Run go build and then you can execute the executable file!

# Input file
data.csv is an example input file of this program. You can also write your own, simply use the same colunms of the given file.
BadData is a bad example, which cannot be successfully converted to a slice of Pigments.

# How to use this calculator?
Run the executable file after building it. The GUI should pop up. Click on buttons to use the functions you need. The color range function is not available at this point. <br>
There is a help button on the main window where you can find instructions of the functions.
