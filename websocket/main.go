package websocket

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"origin-tender-backend/websocket/wsModels"
	"strings"
)

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

var addr = flag.String("addr", ":8081", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./websocket/home.html")
}

func SendNotification(writer http.ResponseWriter, request *http.Request) {

}

// return userId
func initRead(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}

	return conn, nil
}

func Run() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws/test", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	http.HandleFunc("/ws/notification/", func(writer http.ResponseWriter, request *http.Request) {

		routeNames := strings.Split(request.URL.Path, "/")
		userId := routeNames[len(routeNames)-1]

		conn, err := initRead(writer, request)
		if err != nil {
			fmt.Println(err)
		}

		if userId != "" {
			wsModels.UserConnections[userId] = wsModels.WsConnections{
				Bets:         wsModels.UserConnections[userId].Bets,
				Session:      wsModels.UserConnections[userId].Session,
				Notification: conn,
			}
			conn.WriteMessage(1, []byte("ok"))
		}

	})

	http.HandleFunc("/ws/bets/", func(writer http.ResponseWriter, request *http.Request) {

		routeNames := strings.Split(request.URL.Path, "/")
		userId := routeNames[len(routeNames)-1]

		conn, err := initRead(writer, request)
		if err != nil {
			fmt.Println(err)
		}

		if userId != "" {
			wsModels.UserConnections[userId] = wsModels.WsConnections{
				Bets:         conn,
				Session:      wsModels.UserConnections[userId].Session,
				Notification: wsModels.UserConnections[userId].Notification,
			}

			conn.WriteMessage(1, []byte("ok"))
		}

	})

	http.HandleFunc("/ws/session/", func(writer http.ResponseWriter, request *http.Request) {

		routeNames := strings.Split(request.URL.Path, "/")
		userId := routeNames[len(routeNames)-1]

		conn, err := initRead(writer, request)
		if err != nil {
			fmt.Println(err)
		}

		if userId != "" {
			wsModels.UserConnections[userId] = wsModels.WsConnections{
				Bets:         wsModels.UserConnections[userId].Bets,
				Session:      conn,
				Notification: wsModels.UserConnections[userId].Notification,
			}

			conn.WriteMessage(1, []byte("ok"))
		}

	})

	http.HandleFunc("/ws/", func(writer http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
		}

		myString := string(body[:])
		fmt.Println(myString)

	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
