package middleware

import (
	// "context"
	"encoding/json"
	"fmt"

	// "io/ioutil"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nboaldin/pento-time-tracker/models"
	"github.com/satori/go.uuid"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb+srv://root:SERD6whew.lah.runk@pento-time-tracker.zn3fi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func init() {
	err := mgm.SetDefaultConfig(nil, "pento_tt", options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Could not connect to database")
	}
	log.Println("connected to mongo atlas")
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Pento"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	encoder := json.NewEncoder(w)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Error decoding json", err)
	}

	foundUser, err := findFirstUser(user)
	fmt.Println("!!!!!!!!!!!!! FOUND USER", foundUser)

	if err != nil {
		log.Println("Error finding user", err)
		createdUser, err := createOneUser(user)
		if err != nil {
			log.Print("Error creating user", err)
		}
		log.Println("!!!!!!!!!!!!! USER CREATED", createdUser)
		encoder.Encode(createdUser)
	}
	if foundUser != nil {
		encoder.Encode(foundUser)
	}

}

func findFirstUser(user models.User) (*models.User, error) {
	// foundUser := &models.User{}
	coll := mgm.Coll(&user)
	fmt.Printf("First and last name: %s %s \n", user.FirstName, user.LastName)
	err := coll.First(bson.M{"firstname": user.FirstName, "lastname": user.LastName}, &user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func createOneUser(user models.User) (*models.User, error) {
	newUser := models.NewUser(user.FirstName, user.LastName)
	err := mgm.Coll(newUser).Create(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func SessionStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	var session models.Session

	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		log.Fatal("Error decoding json", err)
	}

	fmt.Println("!!!!!!!!!!!!!!!!DECODED SESSIONSTART JSON:", session)

	foundUser, err := startSession(params, session)
	fmt.Println("!!!!!!!!!!!!! FOUND USER", foundUser)
	if err != nil {
		log.Println("Error starting session")
	}

	json.NewEncoder(w).Encode(foundUser)
}

func startSession(params map[string]string, session models.Session) (*models.User, error) {
	foundUser := &models.User{}
	coll := mgm.Coll(foundUser)

	userID := params["id"]

	fmt.Printf("ID of the user: %s", userID)
	err := coll.FindByID(userID, foundUser)
	if err != nil {
		return nil, err
	}

	newSession := models.Session{
		UUID:  uuid.NewV4().String(),
		Name:  session.Name,
		Start: time.Now(),
		End:   time.Time{},
	}

	foundUser.TimeTrackingSessions = append(foundUser.TimeTrackingSessions, newSession)
	err = coll.Update(foundUser)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func SessionStop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	var session models.Session

	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		log.Fatal("Error decoding json", err)
	}

	fmt.Println("!!!!!!!!!!!!!!!!DECODED SESSIONSTOP JSON:", session)

	foundUser, err := stopSession(params, session)
	fmt.Println("!!!!!!!!!!!!! FOUND USER", foundUser)
	if err != nil {
		log.Println("Error starting session")
	}

	json.NewEncoder(w).Encode(foundUser)

}

func stopSession(params map[string]string, session models.Session) (*models.User, error) {
	foundUser := &models.User{}
	coll := mgm.Coll(foundUser)

	userID := params["id"]

	fmt.Printf("ID of the user: %s", userID)
	err := coll.FindByID(userID, foundUser)

	if err != nil {
		return nil, err
	}

	for i := range foundUser.TimeTrackingSessions {
		if foundUser.TimeTrackingSessions[i].UUID == session.UUID {
			foundUser.TimeTrackingSessions[i].End = time.Now()
			break
		}
	}

	err = coll.Update(foundUser)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
