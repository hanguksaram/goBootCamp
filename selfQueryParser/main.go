package main

import (
"bufio"
"fmt"
"log"
"net"
"strings"
"strconv"
)
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			querymap := map[string]string{}
			var queryIndexStart int
			queryString := []rune(strings.Fields(ln)[1])
			for i := 0; i < len(queryString); i++ {
				if queryString[i] == '?' {
					queryIndexStart = i
					break
				}
			}
			queryString = queryString[queryIndexStart+1:]
			var accName, accValue string
			encodedAcc := []rune{}
			nameValueFlag, isEncoded := false, false
			for i := 0; i < len(queryString); i++ {

				if queryString[i] == '%' {
					isEncoded = true
					hexVal, _ := strconv.ParseInt((string(queryString[i + 1]) + string(queryString[i + 2])), 16, 32)
					fmt.Println(hexVal, string(queryString[i + 1]) + string(queryString[i + 2]))
					encodedAcc = append(encodedAcc, rune(hexVal))

				} else if queryString[i] != '=' && queryString[i] != '&' && !nameValueFlag {
					accName += string(queryString[i])
				} else if queryString[i] != '=' && queryString[i] != '&' && nameValueFlag {
					accValue += string(queryString[i])
				}
				if queryString[i] == '=' {
					nameValueFlag = true
					if isEncoded {
						fmt.Println(encodedAcc)
						accName = string(encodedAcc)
						isEncoded = false
						encodedAcc = encodedAcc[:1]
					}

				}
				if queryString[i] == '&' || i == len(queryString) - 1 {
					if isEncoded {
						fmt.Println(encodedAcc)
						accValue = string(encodedAcc)
						isEncoded = false
						encodedAcc = encodedAcc[:1]
					}
					querymap[accName] = accValue
					nameValueFlag = false
					accName, accValue = "", ""
				}
			}
			fmt.Println(querymap)

		}
		if i > 1 {
			break
		}
		i++
	}
}
