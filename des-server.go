package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func logRequest(r *http.Request) {
	fmt.Println("REQUEST:", r.URL.Path)
	//fmt.Println(r.Form)
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

func serveError(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	logRequest(r)
	http.Error(w, "Unknown URL", http.StatusInternalServerError)
}

// serve the "Service Terminated" message
func serveSSInfoClosed(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	logRequest(r)

	resp := "<ss>9</ss>\n" +
		"<lang1>1. The Demon's Souls Online Service has been terminated.</lang1>\n" +
		"<lang2>2. The Demon's Souls Online Service has been terminated.</lang2>\n" +
		"<lang3>3. </lang3>\n" +
		"<lang5>5. La connessione ai servizi online di Demon's Souls è stata terminata.</lang5>\n" +
		"<lang6>6. El servicio online de Demon's Souls ha terminado.</lang6>\n" +
		"<lang7>7. Der Online-Dienst für Demon's Souls wurde eingestellt.</lang7>\n" +
		"<lang8>8. Le service en ligne de Demon's Souls a été interrompu.</lang8>\n"

	fmt.Fprintf(w, "%s", base64.StdEncoding.EncodeToString([]byte(resp)))
}

// serve the URL of the actual game server
func serveSSInfoOpen(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	logRequest(r)

	resp := "<ss>0</ss>\n" +
		"<lang1></lang1>\n" +
		"<lang2></lang2>\n" +
		"<lang5></lang5>\n" +
		"<lang6></lang6>\n" +
		"<lang7></lang7>\n" +
		"<lang8></lang8>\n" +
		"<gameurl1>http://ns1.demons-souls.com:18000/cgi-bin/</gameurl1>\n" +
		"<gameurl2>http://ns1.demons-souls.com:18000/cgi-bin/</gameurl2>\n" +
		"<gameurl5>http://ns1.demons-souls.com:18000/cgi-bin/</gameurl5>\n" +
		"<gameurl6>http://ns1.demons-souls.com:18000/cgi-bin/</gameurl6>\n" +
		"<gameurl7>http://ns1.demons-souls.com:18000/cgi-bin/</gameurl7>\n" +
		"<gameurl8>http://ns1.demons-souls.com:18000/cgi-bin/</gameurl8>\n" +
		"<interval1>120</interval1>\n" +
		"<interval2>120</interval2>\n" +
		"<interval5>120</interval5>\n" +
		"<interval6>120</interval6>\n" +
		"<interval7>120</interval7>\n" +
		"<interval8>120</interval8>\n" +
		"<getWanderingGhostInterval>20</getWanderingGhostInterval>\n" +
		"<setWanderingGhostInterval>20</setWanderingGhostInterval>\n" +
		"<getBloodMessageNum>80</getBloodMessageNum>\n" +
		"<getReplayListNum>80</getReplayListNum>\n" +
		"<enableWanderingGhost>1</enableWanderingGhost>\n"

	fmt.Fprintf(w, "%s", base64.StdEncoding.EncodeToString([]byte(resp)))
}

// initial login request
func serveCgiBinLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	logRequest(r)

	header := []byte{0x01, 0xF4, 0x02, 0x00, 0x00, 0x01, 0x01}
	footer := []byte{0x00}

	msg := "This is the announcement message.\r\n" +
		"We can write anything here!\r\n"

	data := []byte{}
	data = append(data, header...)
	data = append(data, []byte(msg)...)
	data = append(data, footer...)

	fmt.Fprintf(w, "%s\n", base64.StdEncoding.EncodeToString(data))
}

func main() {
	http.HandleFunc("/demons-souls-us/ss.info", serveSSInfoOpen)
	http.HandleFunc("/cgi-bin/login.spd", serveCgiBinLogin)
	http.HandleFunc("/", serveError)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":18000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
