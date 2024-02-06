package main

import (
	e "github.com/alvarollopi/go/ejer_interfaces"
	m "github.com/alvarollopi/go/modelos"
)

func main() {
	/*
						estado, texto := variables.ConviertoaTexto(1588)
						fmt.Println(estado)
						fmt.Println(texto)

						if os := runtime.GOOS; os == "linux" || os == "Os X." {
							fmt.Println("Esto no es Windows, es ", os)
						} else {
							fmt.Println("Esto es Windows")
						}

						switch os := runtime.GOOS; os {
						case "linux":
							fmt.Println("Esto es Linux")
						case "darwin":
							fmt.Printf("Esto es Darwin")
						default:
							fmt.Printf("%s \n", os)
						}

					numero, texto := ejercicios.ConviertoaNumero("999")
					fmt.Println(numero)
					fmt.Println(texto)


				teclado.IngresoNumeros()


			iteraciones.Iterar()


		fmt.Println(ejercicios.NumeroTeclado())
	*/
	//files.Grabatabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	Pedro := new(m.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(m.Mujer)
	e.HumanosRespirando(Maria)
}
