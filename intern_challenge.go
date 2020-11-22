package main

import (
    "fmt"
    "net/http"
)

type Hello struct {
    lang string
    hello string
}

var hellos []Hello

// Initialize the list of implemented languages with fr and en languages
func Init() {
    hello_fr := Hello{"fr", "bonjour"}
    hello_en := Hello{"en", "hello"}
    hellos = []Hello{hello_fr, hello_en}
    return
}

// Add a language to the list of implemented languages
func Add_Hello(lang, hello string) {
    if Get_Hello(lang) == hello {
        fmt.Println("Already implemented language\n")
        return
    }
    new_hello := Hello{lang, hello}
    hellos = append(hellos, new_hello)
    return
}

// Return "hello" in the language "lang"
func Get_Hello(lang string) string {
    for i := 0; i < len(hellos) ; i++ {
        if hellos[i].lang == lang {
            return hellos[i].hello
        }
    }
    return "Language not implemented"
}

// Delete a language from the list of implemented languages
func Delete_Hello(lang string) {
    for i := 0; i < len(hellos) ; i++ {
        if hellos[i].lang == lang {
            hellos[i] = hellos[len(hellos)-1]
            hellos = hellos[:len(hellos)-1]
            return
        }
    }
    fmt.Println("This language is not implemented, it can't be removed\n")
    return
}

func main() {

    // Initialization
    Init()

    // GET method of "/"
    http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ) {
        fmt.Fprint(res, "Hello")
    } )

    // GET, POST and DELETE methods of "/hello"
    http.HandleFunc( "/hello", func( res http.ResponseWriter, req *http.Request ) {
        switch req.Method {
        case "GET":
            lang, ok := req.URL.Query()["lang"]	
            if !ok || len(lang[0]) < 1 {
                fmt.Fprint(res, "lang parameter is missing")
                return
            }
            hello := Get_Hello(lang[0])
            fmt.Fprint(res, hello)
            return
        case "POST":
            lang := req.FormValue("language")
            hello := req.FormValue("hello")
            Add_Hello(lang, hello)
            return
        case "DELETE":
            lang, ok := req.URL.Query()["lang"]	
            if !ok || len(lang[0]) < 1 {
                fmt.Fprint(res, "lang parameter is missing")
                return
            }
            Delete_Hello(lang[0])
            return
        default:
            fmt.Fprint(res, "Only GET, POST and DELETE methods are available")
            return
        }
    })


    // Listen and serve using `http.DefaultServeMux`
    http.ListenAndServe(":9000", nil)

}
