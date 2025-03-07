package validators

import (
	"errors"
	"library-Notify/src/domain"
)

func Checknotify(notify domain.Notify) error {
	if notify.Id_reader <= 0 {
		return errors.New("La cantidad debe ser mayor a 0")
	}
	if notify.Return_date == "" {
		return errors.New("Verifique la fecha de devolucion")
	}

	return nil
}