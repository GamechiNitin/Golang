package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "Url"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// Initilize with MongoBD
func init() {
	// It is a initilize method, run only one time

	// Client option
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	response, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connected successfully")

	collection = response.Database(dbName).Collection(colName)

	// Collection instance
	fmt.Println("Collection instance is ready")
}

// Mongo-DB Helpers - File
// Insert 1 record
func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one data in Database with ID: ", inserted.InsertedID)
}

// Update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkNilError(err)

	// Filter the mongodb database base on ID
	filter := bson.M{"_id": id}

	// Updating DB value
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkNilError(err)
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkNilError(err)

	// Filter the mongodb database base on ID
	filter := bson.M{"_id": id}

	// Delete one record
	result, err := collection.DeleteOne(context.Background(), filter)
	checkNilError(err)
	fmt.Println("Delete count: ", result)
}

// Delete all record
func deleteAllMovie() int64 {
	// Delete All record
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) // Delete Many Filter
	checkNilError(err)
	fmt.Println("Delete many count: ", result.DeletedCount) // result : key, values
	return result.DeletedCount
}

// GET ALL Movies with primitive return type
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkNilError(err)

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkNilError(err)

		// If No error Add data to movies
		movies = append(movies, movie)
	}
	// Always Close Cursor context
	defer cursor.Close(context.Background())
	return movies
}

// Actual Controller - File
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	//Return all movies
	allMovies := getAllMovies()

	// Response
	json.NewEncoder(w).Encode(allMovies)
}

// Create Movie Controller - File
func CreateMovies(w http.ResponseWriter, r *http.Request) {

	// Different Header
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	//Return all movies
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	checkNilError(err)

	// Add to Database
	insertOneMovie(movie)

	// Response
	json.NewEncoder(w).Encode(movie)
}

// Marked movie has watched - Controller
func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete one movie - Controller
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all movie - Controller
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
