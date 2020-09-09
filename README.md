## Teradata GoSQL Driver Plugin

The Teradata GoSQL Driver is available as a Go plugin.

This Go plugin requires 64-bit Go 1.14.6, and runs on macOS and Linux. Windows is not supported.

For community support, please visit the [Teradata Community forums](https://community.teradata.com/).

For Teradata customer support, please visit [Teradata Access](https://access.teradata.com/).

Please note, this driver may contain beta/preview features ("Beta Features"). As such, by downloading and/or using the driver, in addition to agreeing to the licensing terms below, you acknowledge that the Beta Features are experimental in nature and that the Beta Features are provided "AS IS" and may not be functional on any machine or in any environment.

Copyright 2020 Teradata. All Rights Reserved.

### Table of Contents

* [Download](#Download)
* [Sample Program](#SampleProgram)
* [Connection Parameters](#ConnectionParameters)
* [Data Types](#DataTypes)

<a name="Download"></a>

### Download

Download the shared library for your platform.

* For Mac, download `teradatasql-go1.14.6.dylib`

* For Linux, download `teradatasql-go1.14.6.so`

<a name="SampleProgram"></a>

### Sample Program

The Teradata GoSQL Driver accepts a JSON string to specify connection parameters, [documented below](#ConnectionParameters).

You must properly quote the `UseGoSQLPlugin.go` command line arguments for the JSON string and the SQL query text parameters according to what your shell requires.

In the following example, single-quotes are used to enclose the JSON string, which contains double-quotes. This is the typical quoting style on Mac and Linux. Specify your platform's shared library filename as the first argument after the `UseGoSQLPlugin.go` filename argument.
```
go run UseGoSQLPlugin.go teradatasql-go1.14.6.dylib '{"host":"whomooz","user":"guest","password":"please"}' "select 123"
```

The `QueryRows.go` program can execute multiple SQL requests using the same database connection.
```
go run UseGoSQLPlugin.go teradatasql-go1.14.6.dylib '{"host":"whomooz","user":"guest","password":"please"}' "create volatile table tomtab1 (c1 integer) on commit preserve rows" "insert into tomtab1 values(123)" "select * from tomtab1"
```

This is different from executing a multi-statement request in which multiple DML statements are separated by semicolons.
```
go run UseGoSQLPlugin.go teradatasql-go1.14.6.dylib '{"host":"whomooz","user":"guest","password":"please"}' "help session ; select current_timestamp ; select * from dbc.dbcinfo"
```

<a name="ConnectionParameters"></a>

### Teradata GoSQL Driver Connection Parameters

The Teradata GoSQL Driver accepts a JSON string to specify connection parameters.

The connection parameter string must be formatted as a JSON object with open and close curly braces containing key-value pairs, with both the keys and the values being JSON strings.

Example: `{"host":"dbserver.foo.com","user":"guest","password":"please"}`

We want to provide consistency for Teradata JDBC Driver and Teradata GoSQL Driver connection parameters, in terms of both connection parameter names and functionality. We want the Teradata GoSQL Driver to offer all the same connection parameters that the Teradata JDBC Driver offers, for those connection parameters that make sense outside of Java.

Here is the current list of the Teradata GoSQL Driver's JSON string connection parameters, and their default values if omitted.

Parameter          | Default     | Type           | Description
------------------ | ----------- | -------------- | ---
`account`          |             | string         | Specifies the Teradata Database account. Equivalent to the Teradata JDBC Driver `ACCOUNT` connection parameter.
`column_name`      | `"false"`   | quoted boolean | Controls the `name` column returned by `DBI::dbColumnInfo`. Equivalent to the Teradata JDBC Driver `COLUMN_NAME` connection parameter. False specifies that the returned `name` column provides the AS-clause name if available, or the column name if available, or the column title. True specifies that the returned `name` column provides the column name if available, but has no effect when StatementInfo parcel support is unavailable.
`cop`              | `"true"`    | quoted boolean | Specifies whether COP Discovery is performed. Equivalent to the Teradata JDBC Driver `COP` connection parameter.
`coplast`          | `"false"`   | quoted boolean | Specifies how COP Discovery determines the last COP hostname. Equivalent to the Teradata JDBC Driver `COPLAST` connection parameter. When `coplast` is `false` or omitted, or COP Discovery is turned off, then no DNS lookup occurs for the coplast hostname. When `coplast` is `true`, and COP Discovery is turned on, then a DNS lookup occurs for a coplast hostname.
`database`         |             | string         | Specifies the initial database to use after logon, instead of the user's default database. Equivalent to the Teradata JDBC Driver `DATABASE` connection parameter.
`dbs_port`         | `"1025"`    | quoted integer | Specifies the Teradata Database port number. Equivalent to the Teradata JDBC Driver `DBS_PORT` connection parameter.
`encryptdata`      | `"false"`   | quoted boolean | Controls encryption of data exchanged with the Teradata Database. Equivalent to the Teradata JDBC Driver `ENCRYPTDATA` connection parameter.
`fake_result_sets` | `"false"`   | quoted boolean | Controls whether a fake result set containing statement metadata precedes each real result set.
`host`             |             | string         | Specifies the Teradata Database hostname.
`lob_support`      | `"true"`    | quoted boolean | Controls LOB support. Equivalent to the Teradata JDBC Driver `LOB_SUPPORT` connection parameter.
`log`              | `"0"`       | quoted integer | Controls debug logging. Somewhat equivalent to the Teradata JDBC Driver `LOG` connection parameter. This parameter's behavior is subject to change in the future. This parameter's value is currently defined as an integer in which the 1-bit governs function and method tracing, the 2-bit governs debug logging, the 4-bit governs transmit and receive message hex dumps, and the 8-bit governs timing. Compose the value by adding together 1, 2, 4, and/or 8.
`logdata`          |             | string         | Specifies extra data for the chosen logon authentication method. Equivalent to the Teradata JDBC Driver `LOGDATA` connection parameter.
`logmech`          | `"TD2"`     | string         | Specifies the logon authentication method. Equivalent to the Teradata JDBC Driver `LOGMECH` connection parameter. Possible values are `TD2` (the default), `JWT`, `LDAP`, `KRB5` for Kerberos, or `TDNEGO`.
`max_message_body` | `"2097000"` | quoted integer | Not fully implemented yet and intended for future usage. Equivalent to the Teradata JDBC Driver `MAX_MESSAGE_BODY` connection parameter.
`partition`        | `"DBC/SQL"` | string         | Specifies the Teradata Database Partition. Equivalent to the Teradata JDBC Driver `PARTITION` connection parameter.
`password`         |             | string         | Specifies the Teradata Database password. Equivalent to the Teradata JDBC Driver `PASSWORD` connection parameter.
`sip_support`      | `"true"`    | quoted boolean | Controls whether StatementInfo parcel is used. Equivalent to the Teradata JDBC Driver `SIP_SUPPORT` connection parameter.
`teradata_values`  | `"false"`   | quoted boolean | Controls whether result set column values are provided as built-in Go data types or type `teradatasql.TeradataValue`. Refer to the [table below](#DataTypes) for details.
`tmode`            | `"DEFAULT"` | string         | Specifies the transaction mode. Equivalent to the Teradata JDBC Driver `TMODE` connection parameter. Possible values are `DEFAULT` (the default), `ANSI`, or `TERA`.
`user`             |             | string         | Specifies the Teradata Database username. Equivalent to the Teradata JDBC Driver `USER` connection parameter.

<a name="DataTypes"></a>

### Data Types

Teradata Database data type | Go data type
--------------------------- | ---
`BYTEINT`                   | `int8`
`SMALLINT`                  | `int16`
`INTEGER`                   | `int32`
`BIGINT`                    | `int64`
`FLOAT`                     | `float64`
All others                  | `string`
