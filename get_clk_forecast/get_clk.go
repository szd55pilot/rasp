package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Ref for RASP querying http://rasp.stratus.org.uk/page.php?42
// Challock airfield 51 12.508N,   000 49.840E
// or Longitude:	0.8306 (000 49.840E), Latitude:	51.208 (51 12.508N)

func main() {

	var baseurl = "http://rasp.mrsap.org/cgi-bin/get_rasp_blipspot.cgi?"

	// region=Monday&grid=d2&day=0&i=1332&k=1547&width=2000&height=2000&linfo=1&%20param="

	const (
		amp       = "&"
		clkGrid   = "d2"
		clkWidth  = "2000"
		clkHeight = "2000"
		clkLinfo  = "1"
		clkRegion = "UK12"
		format    = "json"
	)

	var fullurl string

	var (
		clkDay   = "0"
		clkI     = "1332"
		clkK     = "1547"
		clkParam = "param=wstar"
		//		clkParam = "param=\"BL Top\""
	)

	fullurl = baseurl +
		"region=" + clkRegion + amp +
		"grid=" + clkGrid + amp +
		"day=" + clkDay + amp +
		"i=" + clkI + amp +
		"k=" + clkK + amp +
		"width=" + clkWidth + amp +
		"height=" + clkHeight + amp +
		"linfo=" + clkLinfo + amp +
		clkParam + amp +
		"format=" + format

	// fmt.Printf("%s", fullurl)

	resp, err := http.Get(fullurl)

	if err != nil {
		fmt.Println("HTTP Status Code:", resp.StatusCode)
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("HTTP Code was not 200: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", fullurl, err)
		os.Exit(1)
	}
	// fmt.Printf("%s", body)

	var parsed downloaded_data

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&parsed); err != nil {
		log.Println(err)
	}

	fmt.Printf("parsed => %+v", parsed.GetRaspBlipspotResults)

	//	spew.Dump(parsed.GetRaspBlipspotResults)

	// pseudo

	// type SomeValue struct {
	// 	Time string
	// 	Speed string
	// }

	// _, v := range w.Values {

	// 	if v.One900Lst != nil {
	// 		t :=
	// 		s = SomeValue{Time: v.One900Lst, }
	// 	}
	// }

	// pseudo end

}

type downloaded_data struct {
	GetRaspBlipspotResults Header `json:"get_rasp_blipspot_results"`
	//	MapInfo                string                 `json:"map_info"`
}

type Values struct {
	// Zero600Lst json.Number `json:"0600lst,omitempty"`
	FcstPd json.Number `json:"Fcst Pd"`
	// Zero630Lst json.Number `json:"0630lst,omitempty"`
	// Zero700Lst json.Number `json:"0700lst,omitempty"`
	// Zero730Lst json.Number `json:"0730lst,omitempty"`
	// Zero800Lst json.Number `json:"0800lst,omitempty"`
	// Zero830Lst json.Number `json:"0830lst,omitempty"`
	Zero900Lst json.Number `json:"0900lst,omitempty"`
	Zero930Lst json.Number `json:"0930lst,omitempty"`
	One000Lst  json.Number `json:"1000lst,omitempty"`
	One030Lst  json.Number `json:"1030lst,omitempty"`
	One100Lst  json.Number `json:"1100lst,omitempty"`
	One130Lst  json.Number `json:"1130lst,omitempty"`
	One200Lst  json.Number `json:"1200lst,omitempty"`
	One230Lst  json.Number `json:"1230lst,omitempty"`
	One300Lst  json.Number `json:"1300lst,omitempty"`
	One330Lst  json.Number `json:"1330lst,omitempty"`
	One400Lst  json.Number `json:"1400lst,omitempty"`
	One430Lst  json.Number `json:"1430lst,omitempty"`
	One500Lst  json.Number `json:"1500lst,omitempty"`
	One530Lst  json.Number `json:"1530lst,omitempty"`
	One600Lst  json.Number `json:"1600lst,omitempty"`
	One630Lst  json.Number `json:"1630lst,omitempty"`
	One700Lst  json.Number `json:"1700lst,omitempty"`
	One730Lst  json.Number `json:"1730lst,omitempty"`
	One800Lst  json.Number `json:"1800lst,omitempty"`
	One830Lst  json.Number `json:"1830lst,omitempty"`
	One900Lst  json.Number `json:"1900lst,omitempty"`
}

type W struct {
	Values []Values `json:"values"`
}

// type BLTop struct {
// 	Values []Values `json:"values"`
// }
// type ThmlHt struct {
// 	Values []Values `json:"values"`
// }
// type Hcrit175 struct {
// 	Values []Values `json:"values"`
// }
// type SfcSun struct {
// 	Values []Values `json:"values"`
// }
// type Temp2M struct {
// 	Values []Values `json:"values"`
// }
// type DewPt2M struct {
// 	Values []Values `json:"values"`
// }
// type MSLPress struct {
// 	Values []Values `json:"values"`
// }
// type SfcWDir struct {
// 	Values []Values `json:"values"`
// }
// type SfcWSpd struct {
// 	Values []Values `json:"values"`
// }
// type BLWindSpd struct {
// 	Values []Values `json:"values"`
// }
// type BlWindDir struct {
// 	Values []Values `json:"values"`
// }
// type MaxConverg struct {
// 	Values []Values `json:"values"`
// }
// type CUpot struct {
// 	Values []Values `json:"values"`
// }
// type OneHrRain struct {
// 	Values []Values `json:"values"`
// }
// type Stars struct {
// 	Values []Values `json:"values"`
// }
// type Hcrit225 struct {
// 	Values []Values `json:"values"`
// }

type Results struct {
	W W `json:"W*"`
	// BLTop      BLTop      `json:"BL Top"`
	// ThmlHt     ThmlHt     `json:"Thml Ht"`
	// Hcrit175   Hcrit175   `json:"Hcrit(175)"`
	// SfcSun     SfcSun     `json:"Sfc. Sun %"`
	// Temp2M     Temp2M     `json:"Temp@2m"`
	// DewPt2M    DewPt2M    `json:"DewPt@2m"`
	// MSLPress   MSLPress   `json:"MSL Press"`
	// SfcWDir    SfcWDir    `json:"Sfc. W.Dir"`
	// SfcWSpd    SfcWSpd    `json:"Sfc. W.Spd"`
	// BLWindSpd  BLWindSpd  `json:"BL Wind Spd"`
	// BlWindDir  BlWindDir  `json:"Bl Wind Dir"`
	// MaxConverg MaxConverg `json:"Max.Converg"`
	// CUpot      CUpot      `json:"CUpot"`
	// OneHrRain  OneHrRain  `json:"1hr Rain"`
	// Stars      Stars      `json:"Stars"`
	// Hcrit225   Hcrit225   `json:"Hcrit(225)"`
}

type Header struct {
	Mapinfo string  `json:"mapinfo"`
	Region  string  `json:"region"`
	Grid    string  `json:"grid"`
	GridI   int     `json:"grid-i"`
	GridJ   int     `json:"grid-j"`
	Lat     float64 `json:"Lat"`
	Lon     float64 `json:"Lon"`
	Created string  `json:"Created"`
	Results Results `json:"Results"`
}
