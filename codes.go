package main

type Codes struct {
	bits string
	i    int
	used bool
}

func initializeCodes() (cd []Codes) {
	for i := 0; i < 256; i++ {
		cd = append(cd, Codes{bits: "", i: 0, used: false})
	}
	return cd
}
