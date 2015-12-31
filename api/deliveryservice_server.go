// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"../db"
	"encoding/json"
	"fmt"
	"time"
)

type DeliveryserviceServer struct {
	Deliveryservice int64     `db:"deliveryservice" json:"deliveryservice"`
	Server          int64     `db:"server" json:"server"`
	LastUpdated     time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleDeliveryserviceServer(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getDeliveryserviceServer(id)
	} else if method == "POST" {
		return postDeliveryserviceServer(payload)
	} else if method == "PUT" {
		return putDeliveryserviceServer(id, payload)
	} else if method == "DELETE" {
		return delDeliveryserviceServer(id)
	}
	return nil, nil
}

func getDeliveryserviceServer(id int) (interface{}, error) {
	ret := []DeliveryserviceServer{}
	arg := DeliveryserviceServer{Deliveryservice: int64(id)}
	if id >= 0 {
		nstmt, err := db.GlobalDB.PrepareNamed(`select * from deliveryservice_server where deliveryservice=:deliveryservice`)
		err = nstmt.Select(&ret, arg)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		nstmt.Close()
	} else {
		queryStr := "select * from deliveryservice_server"
		err := db.GlobalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postDeliveryserviceServer(payload []byte) (interface{}, error) {
	var v DeliveryserviceServer
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO deliveryservice_server("
	sqlString += "deliveryservice"
	sqlString += ",server"
	sqlString += ") VALUES ("
	sqlString += ":deliveryservice"
	sqlString += ",:server"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putDeliveryserviceServer(id int, payload []byte) (interface{}, error) {
	var v DeliveryserviceServer
	err := json.Unmarshal(payload, &v)
	v.Deliveryservice = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE deliveryservice_server SET "
	sqlString += "deliveryservice = :deliveryservice"
	sqlString += ",server = :server"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE deliveryservice=:deliveryservice"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delDeliveryserviceServer(id int) (interface{}, error) {
	arg := DeliveryserviceServer{Deliveryservice: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM deliveryservice_server WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
