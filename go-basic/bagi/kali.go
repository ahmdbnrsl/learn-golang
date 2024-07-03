package bagi

type validationError struct{
    Mess string
}

func (v *validationError) Error() string {
    return v.Mess
}

func Kali(num, numm int) (int, error) {
    if num == 0 {
        return 0, &validationError{Mess: "pokonya error"}
    } else {
        return num * numm, nil
    }
}