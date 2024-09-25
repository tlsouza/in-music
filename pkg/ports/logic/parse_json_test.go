package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(testDescribe *testing.T) {

	testDescribe.Run("Should return result from unmarshall", func(test *testing.T) {

		buyerExpected := make(map[string]interface{})
		buyerExpected["accountId"] = "26651003"
		buyerExpected["userId"] = "29328490"
		buyerExpected["publicAccountId"] = "03ad2b52-564e-4d3f-b292-896d7d7ad3bd"
		buyerExpected["email"] = "vegha@mailinator.com"
		buyerExpected["fullName"] = "Vegha Multi marcas"
		buyerExpected["nickname"] = "Vegha Multimarcas"
		buyerExpected["company"] = true
		buyerExpected["gender"] = ""
		buyerExpected["creationDate"] = "2016-03-16T16:39:25.841Z"
		buyerExpected["status"] = "active"
		buyerExpected["documento"] = "24513769000171"
		buyerExpected["facebookId"] = "127208587687036"
		buyerExpected["googleId"] = ""
		buyerExpected["address"] = "Av Barao Homem de Melo 1281"
		buyerExpected["addressComplement"] = "Loja"
		buyerExpected["addressNumber"] = "1281"
		buyerExpected["neighbourhood"] = "43802"
		buyerExpected["city"] = "9085"
		buyerExpected["state"] = "2"
		buyerExpected["zipCode"] = "30431327"
		buyerExpected["region"] = "31"
		buyerExpected["emailVerified"] = true
		buyerExpected["verificationDate"] = "2016-03-16T16:47:58.945Z"
		buyerExpected["phone"] = "3125269900"
		buyerExpected["phoneVerified"] = true
		buyerExpected["phoneHidden"] = false
		buyerExpected["phoneVerifiedAt"] = "2017-05-08T16:16:36.285Z"
		buyerExpected["secondaryEmail"] = ""

		body := []byte(`{"accountId":"26651003","userId":"29328490","publicAccountId":"03ad2b52-564e-4d3f-b292-896d7d7ad3bd","email":"vegha@mailinator.com","fullName":"Vegha Multi marcas","nickname":"Vegha Multimarcas","company":true,"gender":"","creationDate":"2016-03-16T16:39:25.841Z","status":"active","documento":"24513769000171","facebookId":"127208587687036","googleId":"","address":"Av Barao Homem de Melo 1281","addressComplement":"Loja","addressNumber":"1281","neighbourhood":"43802","city":"9085","state":"2","zipCode":"30431327","region":"31","emailVerified":true,"verificationDate":"2016-03-16T16:47:58.945Z","phone":"3125269900","phoneVerified":true,"phoneHidden":false,"phoneVerifiedAt":"2017-05-08T16:16:36.285Z","secondaryEmail":""}`)

		buyerResult, err := Unmarshal[map[string]interface{}](body, context.TODO())

		assert.Nil(test, err)
		assert.Equal(test, buyerExpected, buyerResult)
	})
}

func TestParseJSON(testDescribe *testing.T) {

	testDescribe.Run("Should return result from unmarshall", func(test *testing.T) {

		buyerExpected := make(map[string]interface{})
		buyerExpected["accountId"] = "26651003"
		buyerExpected["userId"] = "29328490"

		type Buyer struct {
			AccountId string `json:"accountId"`
			UserId    string `json:"userId"`
		}

		input := &Buyer{AccountId: "26651003", UserId: "29328490"}

		buyerResult, err := ParseJSON[map[string]interface{}](input, context.TODO())

		assert.Nil(test, err)
		assert.Equal(test, buyerExpected, buyerResult)
	})
}
