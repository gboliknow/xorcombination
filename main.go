package main

import(
	"encoding/hex"
	"fmt"
)


func XorBuffers(buf1, buf2 [] byte)([] byte , error){
	if len(buf1) != len(buf2){
		return nil, fmt.Errorf("buffers must be of equal length")
	}

	result := make([]byte, len(buf1))
	for i := range buf1{
		result[i] = buf1[i] ^ buf2[i]
	}

	return result, nil
}

func main(){
	hexStr1 := "1c0111001f010100061a024b53535009181c"
	hexStr2 := "686974207468652062756c6c277320657965"

	buf1 , err := hex.DecodeString(hexStr1)
	if err != nil {
		fmt.Println("Error decoding hex1: ", err)
		return
	}

	buf2, err := hex.DecodeString(hexStr2)
	if err != nil {
		fmt.Println("Error decoding hex2: ", err)
		return
	}

	result , err := XorBuffers(buf1, buf2)
	if err != nil {
		fmt.Println("Error computing Xor", err)
		return
	}
	fmt.Println("XOR result:", hex.EncodeToString(result))
}
