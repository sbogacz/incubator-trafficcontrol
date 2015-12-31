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
	"gopkg.in/guregu/null.v3"
	"time"
)

type Type struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	UseInTable  null.String `db:"use_in_table" json:"useInTable"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleType(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getType(id)
	} else if method == "POST" {
		return postType(payload)
	} else if method == "PUT" {
		return putType(id, payload)
	} else if method == "DELETE" {
		return delType(id)
	}
	return nil, nil
}

func getType(id int) (interface{}, error) {
	ret := []Type{}
	arg := Type{Id: int64(id)}
	if id >= 0 {
		nstmt, err := db.GlobalDB.PrepareNamed(`select * from type where id=:id`)
		err = nstmt.Select(&ret, arg)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		nstmt.Close()
	} else {
		queryStr := "select * from type"
		err := db.GlobalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postType(payload []byte) (interface{}, error) {
	var v Type
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO type("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ",use_in_table"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ",:use_in_table"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putType(id int, payload []byte) (interface{}, error) {
	var v Type
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE type SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",use_in_table = :use_in_table"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delType(id int) (interface{}, error) {
	arg := Type{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM type WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
