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

type Hwinfo struct {
	Id          int64     `db:"id" json:"id"`
	Serverid    int64     `db:"serverid" json:"serverid"`
	Description string    `db:"description" json:"description"`
	Val         string    `db:"val" json:"val"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleHwinfo(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getHwinfo(id)
	} else if method == "POST" {
		return postHwinfo(payload)
	} else if method == "PUT" {
		return putHwinfo(id, payload)
	} else if method == "DELETE" {
		return delHwinfo(id)
	}
	return nil, nil
}

func getHwinfo(id int) (interface{}, error) {
	ret := []Hwinfo{}
	arg := Hwinfo{Id: int64(id)}
	if id >= 0 {
		nstmt, err := db.GlobalDB.PrepareNamed(`select * from hwinfo where id=:id`)
		err = nstmt.Select(&ret, arg)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		nstmt.Close()
	} else {
		queryStr := "select * from hwinfo"
		err := db.GlobalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postHwinfo(payload []byte) (interface{}, error) {
	var v Hwinfo
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO hwinfo("
	sqlString += "serverid"
	sqlString += ",description"
	sqlString += ",val"
	sqlString += ") VALUES ("
	sqlString += ":serverid"
	sqlString += ",:description"
	sqlString += ",:val"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putHwinfo(id int, payload []byte) (interface{}, error) {
	var v Hwinfo
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE hwinfo SET "
	sqlString += "serverid = :serverid"
	sqlString += ",description = :description"
	sqlString += ",val = :val"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delHwinfo(id int) (interface{}, error) {
	arg := Hwinfo{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM hwinfo WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
