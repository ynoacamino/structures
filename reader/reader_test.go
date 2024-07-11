package reader

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	start := time.Now()
	arr := []byte(`"hola","mun,do",como,"est,asadwwddwawawaawdaaaaaaaaaaaaaaaaaaaaaw"`)
	linea := *SplitCSVLineForBytes(&arr)

	print("Linea: ", linea[0][0])

	end := time.Now()

	duration := end.Sub(start)

	fmt.Println("Elapsed time: ", duration.Microseconds())
}

/*
	start := time.Now()
	arr := []byte(`"hola","mun,do",como,"est,asadwwddawdwawawddwadaw"`)
	linea := SplitCSVLine(&arr)

	println("Linea: ", linea)

	end := time.Now()

	duration := end.Sub(start)

	fmt.Println("Elapsed time: ", duration.Microseconds())
*/
