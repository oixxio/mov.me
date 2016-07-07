package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// error response contains everything we need to use http.Error
type handlerError struct {
	Error   error
	Message string
	Code    int
}
//App resource model
type app struct{
	Id int 
	Type string 
	Name string 
	Title string 
	Description string 
	Url_demo string 
	Icon_path string 
	Header_path string 
	Screenshot_path_1 string 
	Screenshot_path_2 string 
	Screenshot_path_3 string 
	Url_video string 
	Version string 
	Version_notes string 
	State string 
	Keywords string 
}
//App_requirements resource model
type app_requirement struct{
	Id int 
	Id_app int 
	Type string 
	Description string 
	Rules string 
}
//user_features resource model
type app_user_feature struct{
	Id int 
	Id_app int 
	Screenshot_path string 
	Title string 
	Description string 
}
//client_features resource model
type app_client_feature struct{
	Id int 
	Id_app int 
	Screenshot_path string 
	Title string 
	Description string 
}
//plans resource model
type plan struct{
	Id int 
	Id_developer int 
	Id_client int 
	Id_app int 
	State string 
	Type int 
	Requirements int 
}
//client resource model
type client struct{
	Id int 
	Email string 
	Password string 
	Name string 
}
//developer resource model
type developer struct{
	Id int 
	Email string 
	Password string 
	Name string 
}

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to

	// call the actual handler
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}
	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}
/*App resource functions*/

//app CRUD functions

func listApps(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM apps")

	var list []app
	for rows.Next(){
		var A app
		err := rows.Scan(&A.Id,&A.Type,&A.Name,&A.Title,&A.Description,&A.Url_demo,&A.Icon_path,&A.Header_path,
			&A.Screenshot_path_1,&A.Screenshot_path_2,&A.Screenshot_path_3,&A.Url_video,&A.Version,&A.Version_notes,
			&A.State,&A.Keywords)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getApp(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM apps WHERE ID = " + param)
	var list []app
	for rows.Next(){
		var A app
		err := rows.Scan(&A.Id,&A.Type,&A.Name,&A.Title,&A.Description,&A.Url_demo,&A.Icon_path,&A.Header_path,
			&A.Screenshot_path_1,&A.Screenshot_path_2,&A.Screenshot_path_3,&A.Url_video,&A.Version,&A.Version_notes,
			&A.State,&A.Keywords)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createApp(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO apps (ID,type,name,title,description,url_demo,icon_path,header_path,screenshot_path_1,screenshot_path_2,screenshot_path_3,url_video,version,version_notes,state,keywords)VALUES(NULL,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	_,err := stmt.Exec(appData.Type,appData.Name,appData.Title,appData.Description,appData.Url_demo,appData.Icon_path,appData.Header_path,appData.Screenshot_path_1,appData.Screenshot_path_2,appData.Screenshot_path_3,appData.Url_video,appData.Version,appData.Version_notes,appData.State,appData.Keywords)
	if err != nil {
		return app{}, &handlerError{e, "Could not create app", http.StatusBadRequest}
	}

	return appData,nil
}
func deleteApp(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM apps WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete app", http.StatusBadRequest}
	}
	
	return list,nil
}

func updateApp(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE apps set type = ?, name = ?, title = ?, description = ?, url_demo = ?, icon_path = ?, header_path = ?, screenshot_path_1 = ?, screenshot_path_2 = ?, screenshot_path_3 = ?, url_video = ?, version = ?, version_notes = ?, state = ?, keywords = ? WHERE ID = "+strconv.Itoa(appData.Id))
	_,err := stmt.Exec(appData.Type,appData.Name,appData.Title,appData.Description,appData.Url_demo,appData.Icon_path,appData.Header_path,appData.Screenshot_path_1,appData.Screenshot_path_2,appData.Screenshot_path_3,appData.Url_video,appData.Version,appData.Version_notes,appData.State,appData.Keywords)
	if err != nil {
		return app{}, &handlerError{e, "Could not update app", http.StatusBadRequest}
	}

	return appData,nil
}
//app requirements CRUD functions
func listAppsRequirements(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM app_requirements")

	var list []app_requirement
	for rows.Next(){
		var A app_requirement
		err := rows.Scan(&A.Id,&A.Id_app,&A.Type,&A.Description,&A.Rules)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getAppRequirements(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM app_requirements WHERE id_app =" + param)
	var list []app_requirement
	for rows.Next(){
		var A app_requirement
		err := rows.Scan(&A.Id,&A.Id_app,&A.Type,&A.Description,&A.Rules)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}
	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createAppRequirement(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app_requirement{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app_requirement
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app_requirement{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO app_requirements (ID,id_app,type,description,rules)VALUES(NULL,?,?,?,?)")
	_,err := stmt.Exec(appData.Id_app,appData.Type,appData.Description,appData.Rules)
	if err != nil {
		return app_requirement{}, &handlerError{e, "Could not create app_requirement", http.StatusBadRequest}
	}

	return appData,nil
}
func deleteAppRequirement(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM app_requirements WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete requirements", http.StatusBadRequest}
	}
	
	return list,nil
}
func updateAppRequirement(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app_requirement{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app_requirement
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app_requirement{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE app_requirements SET type = ?,description = ?,rules = ? WHERE ID = "+ strconv.Itoa(appData.Id))
	_,err := stmt.Exec(appData.Type,appData.Description,appData.Rules)
	if err != nil {
		return app_requirement{}, &handlerError{e, "Could not update app_requirement", http.StatusBadRequest}
	}

	return appData,nil
}
//user features CRUD functions
func listAppsUserFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM user_features")

	var list []app_user_feature
	for rows.Next(){
		var A app_user_feature
		err := rows.Scan(&A.Id,&A.Id_app,&A.Screenshot_path,&A.Title,&A.Description)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getAppUserFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM user_features WHERE id_app =" + param)
	var list []app_user_feature
	for rows.Next(){
		var A app_user_feature
		err := rows.Scan(&A.Id,&A.Id_app,&A.Screenshot_path,&A.Title,&A.Description)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}
	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createAppUserFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app_user_feature{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app_user_feature
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app_user_feature{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO user_features (ID,id_app,screenshot_path,title,description)VALUES(NULL,?,?,?,?)")
	_,err := stmt.Exec(appData.Id_app,appData.Screenshot_path,appData.Title,appData.Description)
	if err != nil {
		return app_user_feature{}, &handlerError{e, "Could not create app_user_feature", http.StatusBadRequest}
	}

	return appData,nil
}
func deleteAppUserFeature(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM user_features WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete requirements", http.StatusBadRequest}
	}
	
	return list,nil
}
func updateAppUserFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app_user_feature{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app_user_feature
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app_user_feature{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE user_features SET screenshot_path = ?,title = ?,description = ? WHERE ID = "+ strconv.Itoa(appData.Id))
	_,err := stmt.Exec(appData.Screenshot_path,appData.Title,appData.Description)
	if err != nil {
		return app_user_feature{}, &handlerError{e, "Could not update app_user_feature", http.StatusBadRequest}
	}

	return appData,nil
}
//client features CRUD functions
func listAppsClientFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM client_features")

	var list []app_client_feature
	for rows.Next(){
		var A app_client_feature
		err := rows.Scan(&A.Id,&A.Id_app,&A.Screenshot_path,&A.Title,&A.Description)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getAppClientFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM client_features WHERE id_app =" + param)
	var list []app_client_feature
	for rows.Next(){
		var A app_client_feature
		err := rows.Scan(&A.Id,&A.Id_app,&A.Screenshot_path,&A.Title,&A.Description)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}
	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createClientFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app_client_feature{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app_client_feature
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app_client_feature{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO client_features (ID,id_app,screenshot_path,title,description)VALUES(NULL,?,?,?,?)")
	_,err := stmt.Exec(appData.Id_app,appData.Screenshot_path,appData.Title,appData.Description)
	if err != nil {
		return app_client_feature{}, &handlerError{e, "Could not create app_client_feature", http.StatusBadRequest}
	}

	return appData,nil
}
func deleteAppClientFeature(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM client_features WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete requirements", http.StatusBadRequest}
	}
	
	return list,nil
}
func updateClientFeatures(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return app_client_feature{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var appData app_client_feature
	e = json.Unmarshal(data, &appData)
	if e != nil {
		return app_client_feature{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE client_features SET screenshot_path = ?,title = ?,description = ? WHEN ID = " + strconv.Itoa(appData.Id))
	_,err := stmt.Exec(appData.Screenshot_path,appData.Title,appData.Description)
	if err != nil {
		return app_client_feature{}, &handlerError{e, "Could not update app_client_feature", http.StatusBadRequest}
	}

	return appData,nil
}
//plans CRUD functions
func listPlans(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM plans")

	var list []plan
	for rows.Next(){
		var A plan
		err := rows.Scan(&A.Id,&A.Id_developer,&A.Id_client,&A.Id_app,&A.State,&A.Type,&A.Requirements)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getPlan(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM plans WHERE ID =" +param)
	var list []plan
	for rows.Next(){
		var A plan
		err := rows.Scan(&A.Id,&A.Id_developer,&A.Id_client,&A.Id_app,&A.State,&A.Type,&A.Requirements)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}
	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createPlan(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return plan{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var planData plan
	e = json.Unmarshal(data, &planData)
	if e != nil {
		return plan{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO plans (ID,id_developer,id_client,id_app,state,type,requirements)VALUES(NULL,?,?,?,?,?,?)")
	_,err := stmt.Exec(planData.Id_developer,planData.Id_client,planData.Id_app,planData.State,planData.Type,planData.Requirements)
	if err != nil {
		return plan{}, &handlerError{e, "Could not create plan", http.StatusBadRequest}
	}

	return planData,nil
}
func deletePlan(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM plans WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete requirements", http.StatusBadRequest}
	}
	
	return list,nil
}
func updatePlan(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return plan{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var planData plan
	e = json.Unmarshal(data, &planData)
	if e != nil {
		return plan{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE plans state = ?,type = ?,requirements = ? WHEN ID = " + strconv.Itoa(planData.Id))
	_,err := stmt.Exec(planData.State,planData.Type,planData.Requirements)
	if err != nil {
		return plan{}, &handlerError{e, "Could not create plan", http.StatusBadRequest}
	}

	return planData,nil
}
//clients CRUD functions
func listClients(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM clients")

	var list []client
	for rows.Next(){
		var A client
		err := rows.Scan(&A.Id,&A.Email,&A.Password,&A.Name)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getClient(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM clients WHERE ID =" +param)
	var list []client
	for rows.Next(){
		var A client
		err := rows.Scan(&A.Id,&A.Email,&A.Password,&A.Name)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}
	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createClient(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return client{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var clientData client
	e = json.Unmarshal(data, &clientData)
	if e != nil {
		return client{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO clients (ID,email,password,name)VALUES(NULL,?,?,?)")
	_,err := stmt.Exec(clientData.Email,clientData.Password,clientData.Name)
	if err != nil {
		return client{}, &handlerError{e, "Could not create client", http.StatusBadRequest}
	}

	return clientData,nil
}
func deleteClient(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM clients WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete requirements", http.StatusBadRequest}
	}
	
	return list,nil
}
func updateClient(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return client{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var clientData client
	e = json.Unmarshal(data, &clientData)
	if e != nil {
		return client{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE clients email = ?,password = ?,name = ? WHEN = ID" + strconv.Itoa(clientData.Id))
	_,err := stmt.Exec(clientData.Email,clientData.Password,clientData.Name)
	if err != nil {
		return client{}, &handlerError{e, "Could not update client", http.StatusBadRequest}
	}

	return clientData,nil
}
// developer CRUD functions
func listDevelopers(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	rows,_ := db.Query("SELECT * FROM developers")

	var list []developer
	for rows.Next(){
		var A developer
		err := rows.Scan(&A.Id,&A.Email,&A.Password,&A.Name)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}

	}
	return list,nil
}
func getDeveloper(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	rows,_ := db.Query("SELECT * FROM developers WHERE ID =" +param)
	var list []developer
	for rows.Next(){
		var A developer
		err := rows.Scan(&A.Id,&A.Email,&A.Password,&A.Name)
        if err != nil {
			log.Fatal(err)
		}else{
			list = append(list,A)
		}
	}
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	
	return list,nil
}
func createDeveloper(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return developer{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var developerData developer
	e = json.Unmarshal(data, &developerData)
	if e != nil {
		return developer{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("INSERT INTO developers (ID,email,password,name)VALUES(NULL,?,?,?)")
	_,err := stmt.Exec(developerData.Email,developerData.Password,developerData.Name)
	if err != nil {
		return developer{}, &handlerError{e, "Could not create developer", http.StatusBadRequest}
	}

	return developerData,nil
}
func deleteDeveloper(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	param := mux.Vars(r)["id"]
	_, e := strconv.Atoi(param)
	_,err := db.Exec("DELETE FROM developers WHERE ID = " + param)
	var list []app
	if err != nil {
		return developer{}, &handlerError{e, "Could not delete requirements", http.StatusBadRequest}
	}
	
	return list,nil
}
func updateDeveloper(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError){

	db,_ := sql.Open("mysql", "root:@/movilizame?charset=utf8")
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return developer{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}
	var developerData developer
	e = json.Unmarshal(data, &developerData)
	if e != nil {
		return developer{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	} 
	stmt,_ := db.Prepare("UPDATE developers email = ?,password = ?,name = ? WHEN = ID" + strconv.Itoa(developerData.Id))
	_,err := stmt.Exec(developerData.Email,developerData.Password,developerData.Name)
	if err != nil {
		return developer{}, &handlerError{e, "Could not update client", http.StatusBadRequest}
	}

	return developerData,nil
}
/***********************/

func main() {
	// command line flags
	 port := flag.Int("port", 3000, "port to serve on")
	 dir := flag.String("directory", "web/dashboard/", "directory of web files")
	 flag.Parse()
	 // handle all requests by serving a file of the same name
	 fs := http.Dir(*dir)
	 fileHandler := http.FileServer(fs)

	// setup routes
	//Habria que reorganizar el orden para que sea mas entendible
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/dashboard/", 302))
	router.PathPrefix("/dashboard/").Handler(http.StripPrefix("/dashboard", fileHandler))
	http.Handle("/", router)
	// CRUD apps
	router.Handle("/apps", handler(listApps)).Methods("GET")
	router.Handle("/apps/{id}", handler(deleteApp)).Methods("DELETE")
	router.Handle("/apps/{id}", handler(getApp)).Methods("GET")
	router.Handle("/apps", handler(createApp)).Methods("POST")
	router.Handle("/apps", handler(updateApp)).Methods("PUT")
	//CRUD app requirements
	router.Handle("/apps/requirements", handler(listAppsRequirements)).Methods("GET")
	router.Handle("/apps/requirements/{id}", handler(getAppRequirements)).Methods("GET")
	router.Handle("/apps/requirements/{id}", handler(deleteAppRequirement)).Methods("DELETE")
	router.Handle("/requirements", handler(createAppRequirement)).Methods("POST")
	router.Handle("/requirements", handler(createAppRequirement)).Methods("PUT")
	//CRUD client features
	router.Handle("/apps/clientFeatures", handler(listAppsClientFeatures)).Methods("GET")
	router.Handle("/apps/clientFeatures/{id}", handler(getAppClientFeatures)).Methods("GET")
	router.Handle("/apps/clientFeatures/{id}", handler(deleteAppClientFeature)).Methods("DELETE")
	router.Handle("/clientFeature", handler(createClientFeatures)).Methods("POST")
	router.Handle("/clientFeature", handler(updateClientFeatures)).Methods("PUT")
	//CRUD user features
	router.Handle("/apps/userFeatures", handler(listAppsUserFeatures)).Methods("GET")
	router.Handle("/apps/userFeatures/{id}", handler(getAppUserFeatures)).Methods("GET")
	router.Handle("/apps/userFeatures/{id}", handler(deleteAppUserFeature)).Methods("DELETE")
	router.Handle("/userFeature", handler(createAppUserFeatures)).Methods("POST")
	router.Handle("/userFeature", handler(updateAppUserFeatures)).Methods("PUT")
	//CRUD plans
	router.Handle("/plans", handler(listPlans)).Methods("GET")
	router.Handle("/plans/{id}", handler(getPlan)).Methods("GET")
	router.Handle("/plans/{id}", handler(deletePlan)).Methods("DELETE")
	router.Handle("/plans", handler(createPlan)).Methods("POST")
	router.Handle("/plans", handler(updatePlan)).Methods("PUT")
	//CRUD clients
	router.Handle("/clients", handler(listClients)).Methods("GET")
	router.Handle("/clients/{id}", handler(getClient)).Methods("GET")
	router.Handle("/clients/{id}", handler(deleteClient)).Methods("DELETE")
	router.Handle("/clients", handler(createClient)).Methods("POST")
	router.Handle("/clients", handler(updateClient)).Methods("PUT")
	//CRUD developers
	router.Handle("/developers", handler(listDevelopers)).Methods("GET")
	router.Handle("/developers/{id}", handler(getDeveloper)).Methods("GET")
	router.Handle("/developers/{id}", handler(deleteDeveloper)).Methods("DELETE")
	router.Handle("/developers", handler(createDeveloper)).Methods("POST")
	router.Handle("/developers", handler(updateDeveloper)).Methods("PUT")

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// // this call blocks -- the progam runs here forever
	 err := http.ListenAndServe(addr, nil)
	 fmt.Println(err.Error())
}
