package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const dit = 200 * time.Millisecond
const dah = 3 * dit
const pause = 200 * time.Millisecond

func sendmorse(t *net.Conn, lang bool, color string) {
	trans := ""
	transoff := ""
	for _, letter := range color {
		trans += string(letter) + "1"
		transoff += string(letter) + "0"
	}
	//sendln(t, trans)
	fmt.Fprintf(*t, "%s\n", trans)
	if lang == true {
		time.Sleep(dah)
	} else {
		time.Sleep(dit)
	}
	fmt.Fprintf(*t, "%s\n", transoff)
	time.Sleep(dah)
}

func sendreset(t *net.Conn) {
	fmt.Fprintf(*t, "%s\n", "r0g0y0")
	time.Sleep(pause)
}

func populatetable(m *map[rune]string) {
	words := *m
	words['A'] = "01"
	words['B'] = "100"
	words['C'] = "101"
	words['D'] = "100"
	words['E'] = "0"
	words['F'] = "0010"
	words['G'] = "110"
	words['H'] = "0000"
	words['I'] = "00"
	words['J'] = "0111"
	words['K'] = "101"
	words['L'] = "0100"
	words['M'] = "11"
	words['N'] = "10"
	words['O'] = "111"
	words['P'] = "0110"
	words['Q'] = "1101"
	words['R'] = "010"
	words['S'] = "000"
	words['T'] = "1"
	words['U'] = "001"
	words['V'] = "0001"
	words['W'] = "011"
	words['X'] = "1001"
	words['Y'] = "1011"
	words['Z'] = "1100"
	words['1'] = "01111"
	words['2'] = "00111"
	words['3'] = "00011"
	words['4'] = "00001"
	words['5'] = "00000"
	words['6'] = "10000"
	words['7'] = "11000"
	words['8'] = "11100"
	words['9'] = "11110"
	words['0'] = "11111"
	*m = words
}

func sendword(t *net.Conn, m *map[rune]string, w string, c string) {
	w = strings.ToUpper(w)
	for _, letter := range w {
		words := *m
		//fmt.Printf("%c", letter)
		code := words[letter]
		for _, symbol := range code {
			//fmt.Printf("%c", symbol)
			if symbol == '1' {
				sendmorse(t, true, c)
			} else {
				sendmorse(t, false, c)
			}
		}
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s IP:Port WORD COLOR [-r]\n", os.Args[0])
		return
	}
	ip := ""
	wort := ""
	color := ""
	repeat := ""
	if len(os.Args) > 4 {
		ip, wort, color, repeat = os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	} else {
		ip, wort, color = os.Args[1], os.Args[2], os.Args[3]
	}
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[rune]string)
	ma := &m
	populatetable(ma)
	time.Sleep(100)
	sendreset(&conn)

	check := 0
	if strings.TrimSpace(repeat) != "" {
		if repeat == "-r" {
			check = 1
		}
	}

	if check == 1 {
		for {
			sendword(&conn, ma, wort, color)
		}
	} else {
		//expect(t, "Escape character is '^]'.")
		sendword(&conn, ma, wort, color)
		//ls, err := t.ReadBytes('$')
		//checkErr(err)
	}
	//os.Stdout.Write(ls)
}
