package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"library-Notify/src/domain"
	"net/http"
	"strconv"
)

type FetchAPI struct{}

func NewFetchAPI() *FetchAPI {
	return &FetchAPI{}
}

func (f *FetchAPI) FetchAPI(id_reader int) (domain.Response, error) {
    URL := "http://100.26.65.193/readers/" + strconv.Itoa(id_reader) 

    // Realizar la petición GET
    resp, err := http.Get(URL)
    if err != nil {
        return domain.Response{}, fmt.Errorf("error realizando la petición GET: %w", err)
    }
    defer resp.Body.Close()

    // Verificar si la respuesta fue exitosa
    if resp.StatusCode != http.StatusOK {
        return domain.Response{}, fmt.Errorf("error: código de estado %d", resp.StatusCode)
    }

    // Leer el cuerpo de la respuesta para depuración
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return domain.Response{}, fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
    }

    // Decodificar el JSON en una estructura de libro
    var responseReader domain.Response
    err = json.Unmarshal(bodyBytes, &responseReader)
    if err != nil {
        return domain.Response{}, fmt.Errorf("error decodificando la respuesta JSON: %w", err)
    }

    return responseReader, nil
}