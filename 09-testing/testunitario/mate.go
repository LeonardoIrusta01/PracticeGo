package testunitario

/* Testeos
* Con el comando -> go test -cover <- nos permite saber que tan testeado esta nuestro código
* Con el comando -> go test -coverprofile={nombre del archivo} <- creamos un archivo que nos indica que tan testeado esta nuestro código
* Con el comando -> go tool cover -func={nombre del arhivo} <- leemos el archivo que creamos con los testeos hechos y no hechos
* Con el comando -> go tool cover -html={nombre del arhivo} <- podemos leer mejor el archivo que creamos de los testeos hechos o no hechos.
 */

/*
 * Con el comando -> go test -cpuprofile=cou <- generamos un archivo que para poder detectar donde se esta demorando nuestro archivo
 * Con el comando -> go tool pprof cou <- accedemos a una configuracion para manipual la lentitud de dicha demora
 */

func Suma(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Fibonacci(n int) int {
	if n >= 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
