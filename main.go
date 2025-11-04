// package main

// import (
// 	"fmt"
// )

// // На уровне пакетов глобальные переменные можно объявлять только используя ключевое слово var
// var GlobalVar = 2333
// var AnotherGlobalVar int = 22

// func main() {
// 	//На уровне ф-ии чаще всего используется сокращенное объявление, но допускается и использование var для создания переменных без начального значения, либо для явного указания типа.
// 	var anotherLocalVar float64 = float64(GlobalVar + AnotherGlobalVar)
// 	localVar := GlobalVar
// 	var localVarInteger int32 = int32(anotherLocalVar)
// 	fmt.Printf("First Global Var: %d \nAnother Global Var: %d \nLocal Vars:%d, %.3f, %d", GlobalVar, AnotherGlobalVar, localVar, anotherLocalVar, localVarInteger)
// 	//В go допускаются только явные преобразования типов данных
// }
