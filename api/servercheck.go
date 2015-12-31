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

type Servercheck struct {
	Id          int64     `db:"id" json:"id"`
	Server      int64     `db:"server" json:"server"`
	Aa          null.Int  `db:"aa" json:"aa"`
	Ab          null.Int  `db:"ab" json:"ab"`
	Ac          null.Int  `db:"ac" json:"ac"`
	Ad          null.Int  `db:"ad" json:"ad"`
	Ae          null.Int  `db:"ae" json:"ae"`
	Af          null.Int  `db:"af" json:"af"`
	Ag          null.Int  `db:"ag" json:"ag"`
	Ah          null.Int  `db:"ah" json:"ah"`
	Ai          null.Int  `db:"ai" json:"ai"`
	Aj          null.Int  `db:"aj" json:"aj"`
	Ak          null.Int  `db:"ak" json:"ak"`
	Al          null.Int  `db:"al" json:"al"`
	Am          null.Int  `db:"am" json:"am"`
	An          null.Int  `db:"an" json:"an"`
	Ao          null.Int  `db:"ao" json:"ao"`
	Ap          null.Int  `db:"ap" json:"ap"`
	Aq          null.Int  `db:"aq" json:"aq"`
	Ar          null.Int  `db:"ar" json:"ar"`
	As          null.Int  `db:"as" json:"as"`
	At          null.Int  `db:"at" json:"at"`
	Au          null.Int  `db:"au" json:"au"`
	Av          null.Int  `db:"av" json:"av"`
	Aw          null.Int  `db:"aw" json:"aw"`
	Ax          null.Int  `db:"ax" json:"ax"`
	Ay          null.Int  `db:"ay" json:"ay"`
	Az          null.Int  `db:"az" json:"az"`
	Ba          null.Int  `db:"ba" json:"ba"`
	Bb          null.Int  `db:"bb" json:"bb"`
	Bc          null.Int  `db:"bc" json:"bc"`
	Bd          null.Int  `db:"bd" json:"bd"`
	Be          null.Int  `db:"be" json:"be"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleServercheck(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getServercheck(id)
	} else if method == "POST" {
		return postServercheck(payload)
	} else if method == "PUT" {
		return putServercheck(id, payload)
	} else if method == "DELETE" {
		return delServercheck(id)
	}
	return nil, nil
}

func getServercheck(id int) (interface{}, error) {
	ret := []Servercheck{}
	arg := Servercheck{Id: int64(id)}
	if id >= 0 {
		nstmt, err := db.GlobalDB.PrepareNamed(`select * from servercheck where id=:id`)
		err = nstmt.Select(&ret, arg)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		nstmt.Close()
	} else {
		queryStr := "select * from servercheck"
		err := db.GlobalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postServercheck(payload []byte) (interface{}, error) {
	var v Servercheck
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO servercheck("
	sqlString += "server"
	sqlString += ",aa"
	sqlString += ",ab"
	sqlString += ",ac"
	sqlString += ",ad"
	sqlString += ",ae"
	sqlString += ",af"
	sqlString += ",ag"
	sqlString += ",ah"
	sqlString += ",ai"
	sqlString += ",aj"
	sqlString += ",ak"
	sqlString += ",al"
	sqlString += ",am"
	sqlString += ",an"
	sqlString += ",ao"
	sqlString += ",ap"
	sqlString += ",aq"
	sqlString += ",ar"
	sqlString += ",as"
	sqlString += ",at"
	sqlString += ",au"
	sqlString += ",av"
	sqlString += ",aw"
	sqlString += ",ax"
	sqlString += ",ay"
	sqlString += ",az"
	sqlString += ",ba"
	sqlString += ",bb"
	sqlString += ",bc"
	sqlString += ",bd"
	sqlString += ",be"
	sqlString += ") VALUES ("
	sqlString += ":server"
	sqlString += ",:aa"
	sqlString += ",:ab"
	sqlString += ",:ac"
	sqlString += ",:ad"
	sqlString += ",:ae"
	sqlString += ",:af"
	sqlString += ",:ag"
	sqlString += ",:ah"
	sqlString += ",:ai"
	sqlString += ",:aj"
	sqlString += ",:ak"
	sqlString += ",:al"
	sqlString += ",:am"
	sqlString += ",:an"
	sqlString += ",:ao"
	sqlString += ",:ap"
	sqlString += ",:aq"
	sqlString += ",:ar"
	sqlString += ",:as"
	sqlString += ",:at"
	sqlString += ",:au"
	sqlString += ",:av"
	sqlString += ",:aw"
	sqlString += ",:ax"
	sqlString += ",:ay"
	sqlString += ",:az"
	sqlString += ",:ba"
	sqlString += ",:bb"
	sqlString += ",:bc"
	sqlString += ",:bd"
	sqlString += ",:be"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putServercheck(id int, payload []byte) (interface{}, error) {
	var v Servercheck
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE servercheck SET "
	sqlString += "server = :server"
	sqlString += ",aa = :aa"
	sqlString += ",ab = :ab"
	sqlString += ",ac = :ac"
	sqlString += ",ad = :ad"
	sqlString += ",ae = :ae"
	sqlString += ",af = :af"
	sqlString += ",ag = :ag"
	sqlString += ",ah = :ah"
	sqlString += ",ai = :ai"
	sqlString += ",aj = :aj"
	sqlString += ",ak = :ak"
	sqlString += ",al = :al"
	sqlString += ",am = :am"
	sqlString += ",an = :an"
	sqlString += ",ao = :ao"
	sqlString += ",ap = :ap"
	sqlString += ",aq = :aq"
	sqlString += ",ar = :ar"
	sqlString += ",as = :as"
	sqlString += ",at = :at"
	sqlString += ",au = :au"
	sqlString += ",av = :av"
	sqlString += ",aw = :aw"
	sqlString += ",ax = :ax"
	sqlString += ",ay = :ay"
	sqlString += ",az = :az"
	sqlString += ",ba = :ba"
	sqlString += ",bb = :bb"
	sqlString += ",bc = :bc"
	sqlString += ",bd = :bd"
	sqlString += ",be = :be"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delServercheck(id int) (interface{}, error) {
	arg := Servercheck{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM servercheck WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
