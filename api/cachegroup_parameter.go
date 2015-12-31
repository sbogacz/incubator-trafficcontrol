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

type CachegroupParameter struct {
	Cachegroup  int64     `db:"cachegroup" json:"cachegroup"`
	Parameter   int64     `db:"parameter" json:"parameter"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleCachegroupParameter(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getCachegroupParameter(id)
	} else if method == "POST" {
		return postCachegroupParameter(payload)
	} else if method == "PUT" {
		return putCachegroupParameter(id, payload)
	} else if method == "DELETE" {
		return delCachegroupParameter(id)
	}
	return nil, nil
}

func getCachegroupParameter(id int) (interface{}, error) {
	ret := []CachegroupParameter{}
	arg := CachegroupParameter{Cachegroup: int64(id)}
	if id >= 0 {
		nstmt, err := db.GlobalDB.PrepareNamed(`select * from cachegroup_parameter where cachegroup=:cachegroup`)
		err = nstmt.Select(&ret, arg)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		nstmt.Close()
	} else {
		queryStr := "select * from cachegroup_parameter"
		err := db.GlobalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postCachegroupParameter(payload []byte) (interface{}, error) {
	var v CachegroupParameter
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO cachegroup_parameter("
	sqlString += "cachegroup"
	sqlString += ",parameter"
	sqlString += ") VALUES ("
	sqlString += ":cachegroup"
	sqlString += ",:parameter"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putCachegroupParameter(id int, payload []byte) (interface{}, error) {
	var v CachegroupParameter
	err := json.Unmarshal(payload, &v)
	v.Cachegroup = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE cachegroup_parameter SET "
	sqlString += "cachegroup = :cachegroup"
	sqlString += ",parameter = :parameter"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE cachegroup=:cachegroup"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delCachegroupParameter(id int) (interface{}, error) {
	arg := CachegroupParameter{Cachegroup: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM cachegroup_parameter WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
