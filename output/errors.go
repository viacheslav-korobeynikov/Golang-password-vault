package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
	}
	color.Red("Неизвестный тип ошибки")

	//switch t := value.(type) {
	//case string:
	//	color.Red(t)
	//case int:
	//	color.Red("Код ошибки: %d", t)
	//case error:
	//	color.Red(t.Error())
	//default:
	//	color.Red("Неизвестный тип ошибки")
	//}
}

//func sum[T int | float32 | float64 | int16 | string](a, b T) T {
//	return a + b
//}
//
//type List[T any] struct {
//	elements []T
//}
