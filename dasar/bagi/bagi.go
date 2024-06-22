package bagi

import "errors"

func Bagi(nilai, pembagi int) (int, error) {
    if pembagi == 0 {
        return 0, errors.New("Error pokoknya")
    } else {
        return nilai / pembagi, nil
    }
}