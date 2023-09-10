package main

import (
	"fmt"
	"io/ioutil"
	"github.com/iden3/go-circom-prover-verifier/parsers"
	"github.com/iden3/go-circom-prover-verifier/verifier"
)

func main() {
	// read proof & verificationKey & publicSignals
	proofJson, _ := ioutil.ReadFile("./proof.json")
	// fmt.Println(string(proofJson))
	vkJson, _ := ioutil.ReadFile("./verification_key.json")
	// fmt.Println(string(vkJson))
	publicJson, _ := ioutil.ReadFile("./public.json")
	// fmt.Println(string(publicJson))

	// parse proof & verificationKey & publicSignals
	public, _ := parsers.ParsePublicSignals(publicJson)
	// fmt.Println(public)
	proof, _ := parsers.ParseProof(proofJson)
	// fmt.Println(proof)
	// fmt.Println(vkJson)
	vk, err := parsers.ParseVk(vkJson)
	if err != nil {
		fmt.Printf("Error parsing verification key: %v\n", err)
		return
	}
	fmt.Println(vk)

	// verify the proof with the given verificationKey & publicSignals
	v := verifier.Verify(vk, proof, public)
	fmt.Println(v)
}
