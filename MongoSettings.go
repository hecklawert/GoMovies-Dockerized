/*
*    @author      Hëck Lawert
*    @githuh      https://github.com/hecklawert
*    @date        11/09/2019
*    @description Here we have our settings for MongoDB connection
 */

package main

import "gopkg.in/mgo.v2"

// Stupid global variable
var collection = getSession().DB("GoMovies").C("movies")

// Return connection to MongoDB
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://simpledb-mongodb.default.svc.cluster.local")
	if err != nil {
		panic(err)
	}

	return session
}
