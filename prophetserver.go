package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func createProphetAppData(w http.ResponseWriter, r *http.Request) {
	var global = Global{
		ProjectName: "TMS",
		Communication: `graph LR;EMS-->CMS;QMSa-->CMS;`,
		ContextMap: `graph TD;User-->Exam;Exam-->Question;Exam-->Answer;`,
	}
	var ms1 = Ms{
		Name: "EMS",
		BoundedContext: `graph LR;
        User-->Exam;
        Exam-->Choice;
        Choice-->Answer;`,
	}
	var ms2 = Ms{
		Name: "CMS",
		BoundedContext: `graph LR;
        Question-->Language;
        Question-->Code;
        Question-->Choice;
        Exam-->Question;`,
	}
	var pad = ProphetAppData{
		Global: global,
		Ms:    []Ms{ms1, ms2},
	}
	js, err := json.Marshal(pad)
	if err != nil {
		//ToDo
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/app", createProphetAppData)
	log.Println("Starting server on :8081...")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
