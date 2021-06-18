package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"luis/webservice/models"
	"net/http"
	"os"
	"strconv"
	"time"
	"bytes"
	"image"
	"image/jpeg"
)

type PokemonModel struct {
	Name string `json:"name"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type ResponsePokemon struct {
   Name string `json:"name"`

}

type ResponsePlaceHolder struct {
	Title string `json:"title"`

}

//VARIABLES
 var contador = 0;
 var consultaJsonPlaceHolder = "https://jsonplaceholder.typicode.com/photos/";
 var consultaPokeApi = "https://pokeapi.co/api/v2/pokemon/";

//var response  models.ResponsePlaceHolder
var response  ResponsePlaceHolder
func serveFrames(imgByte []byte) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}
//nil is the zero value for pointers, interfaces, maps, slices, channels and function types
//uninitialized value.
	out, _ := os.Create("./Pictures/img.jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}
 
// ESTA RUTA DEBERIA IR APARTE YA QUE HACE LA FUNCION DE UN CONTROLLER
func getDatosApi(num int) {
	//tr := &http.Transport{
	//	MaxIdleConns:       5000,
	//	IdleConnTimeout:    120* time.Second,
	//	DisableCompression: true,}

	//client:= &http.Client{Transport: tr}
	res, err := http.Get(consultaJsonPlaceHolder + strconv.Itoa(num))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//req.Header.Add("Accept","application/json")
	//req.Header.Add("Content-Type", "application/json")
	//res, err:=client.Do(req)

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
   

	//json Unmarshal vacia el contenido de bodybytes en el puntero de referencia
	json.Unmarshal(responseData, &response)

	//fmt.Println(strconv.Itoa(num) + " " + responseObject.Title)
    //return responseObject
    contador = contador + 1
	fmt.Println(contador , response)
	time.Sleep(1000 * time.Millisecond)
      
}

func (t ResponsePlaceHolder) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)

}

func getPhotos() []ResponsePlaceHolder {
	datos := make([]ResponsePlaceHolder, 0)
	raw, err := ioutil.ReadFile("./photos.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		
	}
	json.Unmarshal(raw, &datos)
	return datos
}

func funcion() {
	fotos := getPhotos()
	//fmt.Println(fotos)          // Traigo todo el JSON
	for _, pic := range fotos {   //  Itero Todo el JSON
		fmt.Println(pic.toString())
		time.Sleep(1000 * time.Millisecond)
		}
}
func ConsultaAPiJson(){

	for i := 1; i < 5000; i++{
		getDatosApi(i)
	  }
}

func main() {
	
	for i := 1; i < 5000; i++{
	   //getDatosApi(i)
	  funcion()
	  }
	
	var s string
	fmt.Scan(&s)
}
