// Copyright 2019 by Teradata Corporation. All Rights Reserved.

package main

import (
	"database/sql"
	"fmt"
	"os"
	"plugin"
	"reflect"
)

func main () {

	if len (os.Args) < 3 {
		fmt.Fprintf (os.Stderr, "Parameters: PluginFileName ConnParams [Query]...\n")
		return
	}

	sFileName   := os.Args [1]
	sConnParams := os.Args [2]
	asQueries   := os.Args [3 : ]

	_, err := plugin.Open (sFileName)
	if err != nil {
		fmt.Fprintf (os.Stderr, "plugin.Open failed: %v\n", err)
		return
	}

	pool, err := sql.Open ("teradata", sConnParams)
	if err != nil {
		fmt.Fprintf (os.Stderr, "sql.Open %v failed: %v\n", sConnParams, err)
		return
	}

	defer pool.Close ()

	for iQuery, sQuery := range asQueries {

		fmt.Printf ("  Executing Query %v of %v: %v\n", iQuery + 1, len (asQueries), sQuery)
		rows, err := pool.Query (sQuery)
		if err != nil {
			fmt.Fprintf (os.Stderr, "pool.Query failed: %v\n", err)
			break
		}

		for nResult := 1 ; ; nResult++ {

			asColumnNames, err := rows.Columns ()
			if err != nil {
				fmt.Fprintf (os.Stderr, "rows.Columns failed for Query %v Result %v: %v\n", iQuery + 1, nResult, err)
				break
			}

			nColumnCount := len (asColumnNames)

			apColumnTypes, err := rows.ColumnTypes ()
			if err != nil {
				fmt.Fprintf (os.Stderr, "rows.ColumnTypes failed for Query %v Result %v: %v\n", iQuery + 1, nResult, err)
				break
			}

			if len (apColumnTypes) > 0 {

				for iColumn, pColumnType := range apColumnTypes {
					fmt.Printf ("    Result %v Metadata Column %v %+v\n", nResult, iColumn + 1, *pColumnType)
				}
			}

			nRowCount := 0
			for rows.Next () {

				nRowCount++

				aoColumnValues := make ([] interface {}, nColumnCount)
				aoColumnValuePointers := make ([] interface {}, nColumnCount)

				for i := range aoColumnValues {
					aoColumnValuePointers [i] = &aoColumnValues [i]
				}

				err = rows.Scan (aoColumnValuePointers ...)
				if err != nil {
					fmt.Fprintf (os.Stderr, "rows.Scan failed for Query %v Result %v Row %v: %v\n", iQuery + 1, nResult, nRowCount, err)
					break
				}

				for iColumn, oValue := range aoColumnValues {

					sTypeDesc := func () reflect.Type { if oValue != nil { return reflect.TypeOf (oValue) } ; return apColumnTypes [iColumn].ScanType () } ().String ()

					fmt.Printf ("    Result %v Row %v Column %v \"%v\" %v = %v\n", nResult, nRowCount, iColumn + 1, asColumnNames [iColumn], sTypeDesc, oValue)

				} // end for iColumn
			} // end for rows.Next

			err = rows.Err ()
			if err != nil {
				fmt.Fprintf (os.Stderr, "rows.Next failed for Query %v Result %v Row %v: %v\n", iQuery + 1, nResult, nRowCount, err)
				break
			}

			if ! rows.NextResultSet () {
				break
			}

			err = rows.Err ()
			if err != nil {
				fmt.Fprintf (os.Stderr, "rows.NextResultSet failed for Query %v Result %v: %v\n", iQuery + 1, nResult, err)
				break
			}
		} // end for nResult

	} // end for range asQueries

} // end main
