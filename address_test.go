package apidq

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/nikitaksv/apidq-client-go/dto/address"

	"github.com/stretchr/testify/require"
)

func TestAddressClean(t *testing.T) {
	reqBs := []byte(`{"query":"москва спартаковская 10с12","countryCode":"RU"}`)
	rspBs := []byte(`{"original":"москва спартаковская 10с12","address":"г Москва, пл Спартаковская","postcodeIn":"","postcode":"105082","region":{"fullName":"г Москва","name":"Москва","type":"г","codes":{"fias":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5","ga":"RU0770000000000000000000000","osm":""}},"area":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"city":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"cityArea":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"settlement":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"planStructure":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"street":{"fullName":"пл Спартаковская","name":"Спартаковская","type":"пл","codes":{"fias":"cd6717bf-1b64-4004-a042-ff1164313e7c","ga":"RU0770000000000000000002733","osm":""}},"houseDetails":{"fullName":"дом 10, строение 12","floor":"","house":"10","case":"","build":"12","liter":"","lend":"","block":"","pav":"","flat":"","office":"","kab":"","abon":"","plot":"","sek":"","entr":"","room":"","hostel":"","munit":""},"coordinates":{"latitude":55.777322,"longitude":37.677688},"country":{"name":"Россия","alpha2":"RU","alpha3":"RUS","numeric":643},"valid":true,"quality":{"unique":0,"actuality":0,"undefined":0,"level":8,"house":3,"geo":8}}`)
	h := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		require.JSONEq(t, string(bs), string(reqBs))

		if _, err = w.Write(rspBs); err != nil {
			panic(err)
		}
	}
	client, tS := NewTestClient(h)
	defer tS.Close()

	cleanRsp, _, err := client.Address.Clean(context.Background(), &address.CleanRequest{
		Query:       "москва спартаковская 10с12",
		CountryCode: "RU",
	})
	if err != nil {
		t.Fatal(err)
	}
	bs, err := json.Marshal(cleanRsp)
	if err != nil {
		t.Fatal(err)
	}

	require.JSONEq(t, string(bs), string(rspBs))
}

func TestAddressSuggest(t *testing.T) {
	reqBs := []byte(`{"query": "москва варш","countryCode": "RU", "count": 5}`)
	rspBs := []byte(`{"suggestions":[{"address":"г Москва, Варшавское ш","postcode":"117105","region":{"fullName":"г Москва","name":"Москва","type":"г","codes":{"fias":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5","ga":"RU0770000000000000000000000","osm":""}},"area":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"city":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"cityArea":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"settlement":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"planStructure":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"street":{"fullName":"Варшавское ш","name":"Варшавское","type":"ш","codes":{"fias":"8fc06b0b-5de3-4a72-9e6f-9e0647a37a66","ga":"RU0770000000000000000000476","osm":""}},"houseDetails":{"fullName":"","floor":"","house":"","case":"","build":"","liter":"","lend":"","block":"","pav":"","flat":"","office":"","kab":"","abon":"","plot":"","sek":"","entr":"","room":"","hostel":"","munit":""},"coordinates":{"latitude":55.646,"longitude":37.6203},"country":{"name":"Россия","alpha2":"RU","alpha3":"RUS","numeric":643}},{"address":"г Москва, 2-й Варшавский проезд","postcode":"115201","region":{"fullName":"г Москва","name":"Москва","type":"г","codes":{"fias":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5","ga":"RU0770000000000000000000000","osm":""}},"area":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"city":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"cityArea":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"settlement":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"planStructure":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"street":{"fullName":"2-й Варшавский проезд","name":"2-й Варшавский","type":"проезд","codes":{"fias":"b89718e1-8b56-4ba8-8383-5c7b596aee6c","ga":"RU0770000000000000000000475","osm":""}},"houseDetails":{"fullName":"","floor":"","house":"","case":"","build":"","liter":"","lend":"","block":"","pav":"","flat":"","office":"","kab":"","abon":"","plot":"","sek":"","entr":"","room":"","hostel":"","munit":""},"coordinates":{"latitude":55.6442,"longitude":37.63},"country":{"name":"Россия","alpha2":"RU","alpha3":"RUS","numeric":643}},{"address":"г Москва, 1-й Варшавский проезд","postcode":"115201","region":{"fullName":"г Москва","name":"Москва","type":"г","codes":{"fias":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5","ga":"RU0770000000000000000000000","osm":""}},"area":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"city":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"cityArea":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"settlement":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"planStructure":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"street":{"fullName":"1-й Варшавский проезд","name":"1-й Варшавский","type":"проезд","codes":{"fias":"09ffd474-1ca8-42e1-8217-876300fd7c2c","ga":"RU0770000000000000000000474","osm":""}},"houseDetails":{"fullName":"","floor":"","house":"","case":"","build":"","liter":"","lend":"","block":"","pav":"","flat":"","office":"","kab":"","abon":"","plot":"","sek":"","entr":"","room":"","hostel":"","munit":""},"coordinates":{"latitude":55.6501,"longitude":37.6264},"country":{"name":"Россия","alpha2":"RU","alpha3":"RUS","numeric":643}},{"address":"г Москва, п Вороновское, Варшавское 64-й км ш","postcode":"108830","region":{"fullName":"г Москва","name":"Москва","type":"г","codes":{"fias":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5","ga":"RU0770000000000000000000000","osm":""}},"area":{"fullName":"п Вороновское","name":"Вороновское","type":"п","codes":{"fias":"10409e98-eb2d-4a52-acdd-7166ca7e0e48","ga":"RU0770020000000000000000000","osm":""}},"city":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"cityArea":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"settlement":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"planStructure":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"street":{"fullName":"Варшавское 64-й км ш","name":"Варшавское 64-й км","type":"ш","codes":{"fias":"dc6cb90e-fe77-44c7-93f7-ec39909489e1","ga":"RU0770020000000000000000015","osm":""}},"houseDetails":{"fullName":"","floor":"","house":"","case":"","build":"","liter":"","lend":"","block":"","pav":"","flat":"","office":"","kab":"","abon":"","plot":"","sek":"","entr":"","room":"","hostel":"","munit":""},"coordinates":{"latitude":55.2921,"longitude":37.1821},"country":{"name":"Россия","alpha2":"RU","alpha3":"RUS","numeric":643}},{"address":"г Москва, Варшавское шоссе 28-й км (п Воскресенско км","postcode":"117148","region":{"fullName":"г Москва","name":"Москва","type":"г","codes":{"fias":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5","ga":"RU0770000000000000000000000","osm":""}},"area":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"city":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"cityArea":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"settlement":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"planStructure":{"fullName":"","name":"","type":"","codes":{"fias":"","ga":"","osm":""}},"street":{"fullName":"Варшавское шоссе 28-й км (п Воскресенско км","name":"Варшавское шоссе 28-й км (п Воскресенско","type":"км","codes":{"fias":"b4a45703-7ca1-4dff-9f9d-8e34deadbf33","ga":"RU0770000000000000000007569","osm":""}},"houseDetails":{"fullName":"","floor":"","house":"","case":"","build":"","liter":"","lend":"","block":"","pav":"","flat":"","office":"","kab":"","abon":"","plot":"","sek":"","entr":"","room":"","hostel":"","munit":""},"coordinates":{"latitude":55.4926,"longitude":37.5928},"country":{"name":"Россия","alpha2":"RU","alpha3":"RUS","numeric":643}}]}`)
	h := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		require.JSONEq(t, string(bs), string(reqBs))

		if _, err = w.Write(rspBs); err != nil {
			panic(err)
		}
	}
	client, tS := NewTestClient(h)
	defer tS.Close()

	cleanRsp, _, err := client.Address.Suggest(context.Background(), &address.SuggestRequest{
		Query:       "москва варш",
		CountryCode: "RU",
		Count:       5,
	})
	if err != nil {
		t.Fatal(err)
	}

	bs, err := json.Marshal(cleanRsp)
	if err != nil {
		t.Fatal(err)
	}

	require.JSONEq(t, string(bs), string(rspBs))
}
