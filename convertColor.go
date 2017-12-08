/*==============================================================================

 * Convert color in one space into another

 * Function format: Convert___To___()
 
 * Color spaces to be selected from: RGB,CMYK,CMY,Hsb,Lab,XYZ
   Note: 
   1. Here the XYZ is CIE XYZ)
   2. ConvertColor uses sRGB color space. More imformation about sRGB vs AdobeRGB: https://fstoppers.com/pictures/adobergb-vs-srgb-3167
 
 * Output type: float64
 
 * Input type and range:
   RGB: int, 0-255
   CMY/CMYK: float64, 0-1
   Hsb: h, int, 0-360; s, b, float64, 0-1
   Lab: L, int, 0-100; a, b, int, -128-128
   XYZ: float64, 0-1
  
 * You can call the Round function to limit the number of digits after decimal point

 *============================================================================*/


package main

import(
  "math"
)

//ConvertRGBToXYZ takes integer R,G,B values
//and return the corresponding float64 X,Y,Z values
func ConvertRGBToXYZ(R,G,B int) (float64,float64,float64) {
  return convertRGBToXYZ(float64(R),float64(G),float64(B))
}

//convertRGBToXYZ takes float64 R,G,B values
//and return the corresponding float64 X,Y,Z values
//You can apply D65 white point by uncommenting the commented code.
func convertRGBToXYZ(R,G,B float64) (float64,float64,float64) {
  R = Gamma(R/255)
  G = Gamma(G/255)
  B = Gamma(B/255)
  //sRGB
  X := R * 0.433910 + G * 0.376220 + B * 0.189860
  Y := R * 0.212649 + G * 0.715169 + B * 0.072182
  Z := R * 0.017756 + G * 0.109478 + B * 0.872915
  //applying D65 white point
  /*
  X := R * 0.412453 + G * 0.357580 + B * 0.180423
  Y := R * 0.212671 + G * 0.715160 + B * 0.072169
  Z := R * 0.019334 + G * 0.119193 + B * 0.950227 */
  X = LimitInRange(X,0,1)
  Y = LimitInRange(Y,0,1)
  Z = LimitInRange(Z,0,1)
  return X,Y,Z
}


//ConvertRGBToLab takes integer R,G,B values
//and return the corresponding float64 L*,a*,b* values
func ConvertRGBToLab(R,G,B int) (float64,float64,float64) {
  return convertRGBToLab(float64(R),float64(G),float64(B))
}

//ConvertXYZToLab takes float64 X,Y,Z values
//and return the corresponding float64 L*,a*,b* values
//X,Y,Z range from 0-1
func ConvertXYZToLab(X,Y,Z float64)(float64,float64,float64) {
  L := 116 * FunctionLab(Y) - 16
  if L<0 {
    L = 0
  }
  a := 500 * (FunctionLab(X) - FunctionLab(Y))
  b := 200 * (FunctionLab(Y) - FunctionLab(Z))
  L = LimitInRange(L,0,100)
  a = LimitInRange(a,-128,128)
  b = LimitInRange(b,-128,128)
  return L,a,b
}

//convertRGBToLab takes float64 R,G,B values
//and return the corresponding float64 L*,a*,b* values
func convertRGBToLab(R,G,B float64) (float64,float64,float64) {
  X,Y,Z := convertRGBToXYZ(R,G,B)
  return ConvertXYZToLab(X,Y,Z)
}

//ConvertRGBToHsb takes the integer R,G,B values of a color
//and returns the float64 hsb(hsv) values of this color
func ConvertRGBToHsb(R,G,B int) (float64,float64,float64){
  return convertRGBToHsb(float64(R),float64(G),float64(B))
}

//convertRGBToHsb takes the float64 R,G,B values of a color
//and returns the float64 hsb(hsv) values of this color
func convertRGBToHsb(R,G,B float64) (float64,float64,float64){
  var h,s,b float64
  b = Max(R,G,B)
  if b != 0{
    s = (b-Min(R,G,B))/b
    switch b {
      case R: h = (G-B)*60/s/b
      case G: h = (B-R)*60/s/b + 180
      case B: h = (R-G)*60/s/b + 240
    }
  }else {
    s = 0
    h = 0
  }
  if h < 0{
    h += 360
  }
  h = LimitInRange(h,0,360)
  s = LimitInRange(s,0,1)
  b = LimitInRange(b/255,0,1)
  return h,s,b
}

//ConvertRGBToYUV takes the integer R,G,B values of a color
//and returns the float64 YUV values of this color
func ConvertRGBToYUV(R,G,B int)(float64,float64,float64){
  return convertRGBToYUV(float64(R),float64(G),float64(B))
}

//convertRGBToYUV takes the float64 R,G,B values of a color
//and returns the float64 YUV values of this color
func convertRGBToYUV(R,G,B float64)(float64,float64,float64){
  Y := 0.299 * R + 0.587 * G + 0.114 * B
  U := -0.147 * R - 0.289 * G + 0.436 * B
  V := 0.615 * R - 0.515 * G - 0.100 * B
  return Y,U,V
}

//ConvertRGBToCMY takes the integer R,G,B values of a color
//and returns the float64 CMY values of this color
func ConvertRGBToCMY(R,G,B int)(float64,float64,float64){
  return float64(255-R)/255,float64(255-G)/255,float64(255-B)/255
}

//convertRGBToCMY takes the float64 R,G,B values of a color
//and returns the float64 CMY values of this color
func convertRGBToCMY(R,G,B float64)(float64,float64,float64){
  return (255-R)/255,(255-G)/255,(255-B)/255
}

//ConvertRGBToCMYK takes the integer R,G,B values of a color
//and returns the float64 CMYK values of this color
func ConvertRGBToCMYK(R,G,B int)(float64,float64,float64,float64){
  return convertRGBToCMYK(float64(R),float64(G),float64(B))
}

//convertRGBToCMYK takes the float64 R,G,B values of a color
//and returns the float64 CMYK values of this color
func convertRGBToCMYK(R,G,B float64)(float64,float64,float64,float64){
  C := R
  M := G
  Y := B
  K := Max(C,M,Y)
  if K == 0{
    return 0,0,0,1
  }
  C = (K-C)/K
  M = (K-M)/K
  Y = (K-Y)/K
  return C,M,Y,1-K/255
}

//ConvertCMYToRGB takes the float64 C,M,Y values of a color
//and returns the float64 R,G,B values of this color
//C,M,Y range from 0-1
func ConvertCMYToRGB(C,M,Y float64) (float64,float64,float64){
  return (255-255*C),(255-255*M),(255-255*Y)
}

//ConvertCMYKToRGB takes the float64 C,M,Y,K values of a color
//and returns the float64 R,G,B values of this color
//C,Y,M,K range from 0-1
func ConvertCMYKToRGB(C,M,Y,K float64) (float64,float64,float64){
  return 255*(1-C)*(1-K),255*(1-M)*(1-K),255*(1-Y)*(1-K)
}

//ConvertHsbToRGB takes the interger h and float64 s,b values of a color
//and returns the float64 R,G,B values of this color
//h ranges from 0-360, s and b range from 0-1
func ConvertHsbToRGB(h int,s,b float64) (float64,float64,float64) {
  h1 := (h/60)%6
  f := float64(h)/60 - float64(h1)
  p := b * (1-s)
  q := b * (1 - f * s)
  t := b * (1 - (1- f) * s)
  var R,G,B float64
  switch h1 {
    case 0: R,G,B = b,t,p
    case 1: R,G,B = q,b,p
    case 2: R,G,B = p,b,t
    case 3: R,G,B = p,q,b
    case 4: R,G,B = t,p,b
    case 5: R,G,B = b,p,q
  }
  R = LimitInRange(R,0,1)
  G = LimitInRange(G,0,1)
  B = LimitInRange(B,0,1)
  return R*255,G*255,B*255
}

//VerseGamma is the inverse function of Gamma
//which is used in ConvertXYZToRGB
func VerseGamma(x float64) float64{
  if x > 0.00313{
    return math.Pow(x,1/2.4)*1.055-0.055
  }else{
    return x*12.92
  }
}

//ConvertXYZToRGB takes the float64 X,Y,Z values of a color
//and returns the float64 R,G,B values of this color
//X,Y,Z range from 0-1
func ConvertXYZToRGB(X,Y,Z float64) (float64,float64,float64){
  R := 3.240479 * X - 1.537150 * Y - 0.498535 * Z
  G := -0.969256 * X + 1.875992 * Y + 0.041556 * Z
  B := 0.055648 * X - 0.204043 * Y + 1.057311 * Z
  R = LimitInRange(VerseGamma(R),0,1)
  G = LimitInRange(VerseGamma(G),0,1)
  B = LimitInRange(VerseGamma(B),0,1)
  return R*255,G*255,B*255
}

//ConvertLabToXYZ takes int L*,a*,b* values of a color
//and returns the float64 X,Y,Z values of this color
//L* ranges from 0-100, and a*,b* ranges from -128-128
func ConvertLabToXYZ(L,a,b int)(float64,float64,float64){
  y := float64(L+16)/116
  x := float64(a)/500 + y
  z := y - float64(b)/200
  Y := FunctionXYZ(y)
  X := FunctionXYZ(x)
  Z := FunctionXYZ(z)
  X = LimitInRange(X*0.95047,0,1)
  Y = LimitInRange(Y,0,1)
  Z = LimitInRange(Z*1.08883,0,1)
  return X,Y,Z
}

//ConvertLabToRGB takes int L*,a*,b* values of a color
//and returns the float64 R,G,B values of this color
//L* ranges from 0-100, and a*,b* ranges from -128-128
func ConvertLabToRGB(L,a,b int)(float64,float64,float64) {
  X,Y,Z := ConvertLabToXYZ(L,a,b)
  return ConvertXYZToRGB(X,Y,Z)
}

//ConvertLabToHsb takes int L*,a*,b* values of a color
//and returns the float64 h,s,b values of this color
//L* ranges from 0-100, and a*,b* ranges from -128-128
func ConvertLabToHsb(L,a,b int)(float64,float64,float64){
  R,G,B := ConvertLabToRGB(L,a,b)
  return convertRGBToHsb(R,G,B)
}

//ConvertHsbToLab takes interger h, and float64 s,b values of a color
//and returns the float64 L*,a*,b* values of this color
//h ranges from 0-360, s and b range from 0-1
func ConvertHsbToLab(h int,s,b float64)(float64,float64,float64){
  R,G,B := ConvertHsbToRGB(h,s,b)
  return convertRGBToLab(R,G,B)
}

//ConvertCMYToLab takes float64 C,M,Y values of a color
//and returns the float64 L*,a*,b* values of this color
//C,Y,M range from 0-1
func ConvertCMYToLab(C,M,Y float64)(float64,float64,float64){
  R,G,B := ConvertCMYToRGB(C,M,Y)
  return convertRGBToLab(R,G,B)
}

//ConvertLabToCMY takes interger L*,a*,b* values of a color
//and returns the float64 L*,a*,b* values of this color
//L* ranges from 0-100, and a*,b* ranges from -128-128
func ConvertLabToCMY(L,a,b int)(float64,float64,float64){
  R,G,B := ConvertLabToRGB(L,a,b)
  return convertRGBToCMY(R,G,B)
}

//ConvertCMYKToLab takes float64 C,M,Y,K values of a color
//and returns the float64 L*,a*,b* values of this color
//C,Y,M,K range from 0-1
func ConvertCMYKToLab(C,M,Y,K float64)(float64,float64,float64){
  R,G,B := ConvertCMYKToRGB(C,M,Y,K)
  return convertRGBToLab(R,G,B)
}

//ConvertLabToCMYK takes integer L*,a*,b* values of a color
//and returns the float64 C,M,Y,K values of this color
//L* ranges from 0-100, and a*,b* ranges from -128-128
func ConvertLabToCMYK(L,a,b int)(float64,float64,float64,float64){
  R,G,B := ConvertLabToRGB(L,a,b)
  return convertRGBToCMYK(R,G,B)
}

//ConvertCMYToHsb takes float64 C,M,Y values of a color
//and returns the float64 h,s,b values of this color
//C,Y,M range from 0-1
func ConvertCMYToHsb(C,M,Y float64)(float64,float64,float64){
  R,G,B := ConvertCMYToRGB(C,M,Y)
  return convertRGBToHsb(R,G,B)
}

//ConvertHsbToCMY takes integer h and float64 s,b values of a color
//and returns the float64 C,M,Y values of this color
//h ranges from 0-360, s and b range from 0-1
func ConvertHsbToCMY(h int,s,b float64)(float64,float64,float64){
  R,G,B := ConvertHsbToRGB(h,s,b)
  return convertRGBToCMY(R,G,B)
}

//ConvertCMYKToHsb takes float64 C,M,Y,K values of a color
//and returns the float64 h,s,b values of this color
//C,Y,M,K range from 0-1
func ConvertCMYKToHsb(C,M,Y,K float64)(float64,float64,float64){
  R,G,B := ConvertCMYKToRGB(C,M,Y,K)
  return convertRGBToHsb(R,G,B)
}

//ConvertHsbToCMYK takes integer h and float64 s,b values of a color
//and returns the float64 C,M,Y,K values of this color
//h ranges from 0-360, s and b range from 0-1
func ConvertHsbToCMYK(h int,s,b float64)(float64,float64,float64,float64){
  R,G,B := ConvertHsbToRGB(h,s,b)
  return convertRGBToCMYK(R,G,B)
}

//ConvertCMYToXYZ takes float64 C,M,Y values of a color
//and returns the float64 X,Y,Z values of this color
//C,Y,M range from 0-1
func ConvertCMYToXYZ(C,M,Y float64)(float64,float64,float64){
  R,G,B := ConvertCMYToRGB(C,M,Y)
  return convertRGBToXYZ(R,G,B)
}

//ConvertXYZToCMY takes float64 X,Y,Z values of a color
//and returns the float64 C,M,Y values of this color
//X,Y,Z range from 0-1
func ConvertXYZToCMY(X,Y,Z float64)(float64,float64,float64){
  R,G,B := ConvertXYZToRGB(X,Y,Z)
  return convertRGBToCMY(R,G,B)
}

//ConvertCMYKToXYZ takes float64 C,M,Y,K values of a color
//and returns the float64 X,Y,Z values of this color
//C,Y,M,K range from 0-1
func ConvertCMYKToXYZ(C,M,Y,K float64)(float64,float64,float64){
  R,G,B := ConvertCMYKToRGB(C,M,Y,K)
  return convertRGBToXYZ(R,G,B)
}

//ConvertXYZToCMYK takes float64 X,Y,Z values of a color
//and returns the float64 C,M,Y,K values of this color
//X,Y,Z range from 0-1
func ConvertXYZToCMYK(X,Y,Z float64)(float64,float64,float64,float64){
  R,G,B := ConvertXYZToRGB(X,Y,Z)
  return convertRGBToCMYK(R,G,B)
}

//ConvertXYZToHsb takes float64 X,Y,Z values of a color
//and returns the float64 h,s,b values of this color
//X,Y,Z range from 0-1
func ConvertXYZToHsb(X,Y,Z float64)(float64,float64,float64){
  R,G,B := ConvertXYZToRGB(X,Y,Z)
  return convertRGBToHsb(R,G,B)
}

//ConvertHsbToXYZ takes integer h and float64 s,b values of a color
//and returns the float64 x,y,z values of this color
//h ranges from 0-360, s and b range from 0-1
func ConvertHsbToXYZ(h int,s,b float64)(float64,float64,float64){
  R,G,B := ConvertHsbToRGB(h,s,b)
  return convertRGBToXYZ(R,G,B)
}

//ConvertCMYToCMYK takes float64 C,M,Y values of a color
//and returns the float64 C,M,Y,K values of this color
//C,Y,M range from 0-1
func ConvertCMYToCMYK(C,M,Y float64)(float64,float64,float64,float64){
  K := 1 - Min(C,M,Y)
  C = (K-(1-C))/K
  M = (K-(1-M))/K
  Y = (K-(1-Y))/K
  return C,M,Y,1-K
}

//ConvertCMYKToCMY takes float64 C,M,Y,K values of a color
//and returns the float64 C,M,Y values of this color
//C,Y,M,K range from 0-1
func ConvertCMYKToCMY(C,M,Y,K float64)(float64,float64,float64){
  C = 1-(1-K)*(1-C)
  M = 1-(1-K)*(1-M)
  Y = 1-(1-K)*(1-Y)
  return C,M,Y
}

//Round takes a float64 number x and an interger n
//and returns the number rournded to n digits after the decimal point
func Round(x float64, n int) float64{
  pow := math.Pow(10,float64(n))
  return math.Trunc((x+0.5/pow)*pow)/pow
}

//Max takes a slice of numbers and returns the maximum
func Max(a ... float64) float64{
  maximum := a[1]
  for i := range a{
    if a[i] > maximum{
      maximum = a[i]
    }
  }
  return maximum
}

func Min(a ... float64) float64 {
  minimum := a[1]
  for i := range a{
    if a[i] < minimum {
      minimum = a[i]
    }
  }
  return minimum
}

//FunctionLab is the function used when convert XYZ space into L*a*b* space
func FunctionLab(t float64) float64{
  if t > 0.00856 {
    return math.Pow(t,1.00/3.00)
  }else{
    return 7.787*t + 4.00/29.00
  }
}

//FunctionLab is the function used when convert L*a*b* space into XYZ space
func FunctionXYZ(t float64) float64{
  if t > 0.206893 {
    return math.Pow(t,3)
  }else{
    return (t-4.00/29.00)/7.787
  }
}

//Gamma is the function used when convert RGB space into XYZ space
func Gamma(x float64) float64{
  if x > 0.04045{
    return math.Pow((x+0.055)/1.055,2.4)
  }else{
    return x/12.92
  }
}

//LimitInRange force float64 x to be limited from lowerLimit to upperLimit
func LimitInRange(x,lowerLimit,upperLimit float64) float64{
  if x<lowerLimit{
    x = lowerLimit
  }else if x>upperLimit{
    x = upperLimit
  }
  return x
}
