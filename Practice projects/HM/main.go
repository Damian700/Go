package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/* iit receives the name of the file to be read on the terminal */
/* Usa la variable os.Args para leer los argumentos que se pasan a un programa
para abrir archivos, existe la funcion "Open" en el package "os"
Hay que hacer un type File?
File puede aplicar la Reader interface
io.Copy function
*/
func main() {
	fmt.Println("Welcome to files in Golang!")
	content := "This needs to go in a File: Cristo Reina!"
	file, err := os.Create(os.Args[1]) /* "./ BLA BLA BLA" significa que lo creamos en el mismo directorio doned estamos */
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	/* para escribir en un archivo se puede usar WriteString, ByteWritter, CopyBuffer, etc
	La funcion WriteString te devuelve el Length (lo dice la documentacion)
	*/checkNilErr(err)
	fmt.Println("Length is", length)
	defer file.Close() /* ACORDARSE DE CERRAR EL ARCHIVO DESPUES DE TRABAJAR SOBRE EL */
	readFile("./myText.txt")
}

/* PARA OBTENER LA REFERENCIA DE UN ARCHIVO EN LA MEMORIA,
UTILIZAR:
f, err := os.Open(os.Args[1])
O SI QUERES RECIBIR LA ESTRUCTURA EN SI...
*f, err := os.Open(os.Args[1])*/

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) /* ioutil tiene multiples herramientas del estilo de io*/
	/* la información que llega, está en formato de data BYTE */
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n as a string: \n", string(databyte), "\n and in byte lenguage: \n", databyte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) /* cierra la ejecucion del programa */
	}
	/* another option:
	fmt.Println("Error:", err)
	os.Exit(1) */
}

/* USE THE READER INTERFACE

io.Copy(os.Stdout, f)
FUNCION READ:
func(f *File*) Read(b []byte) (n int, err error)*/
