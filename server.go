package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Az3z3l/GQL-SERVER/graph"
	"github.com/Az3z3l/GQL-SERVER/graph/generated"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = "80"

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// RequestLogger goserv.log
func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func main() {
	// create logs for the go server
	fileName := "goserv.log"
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	r := mux.NewRouter()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = false
		return next(ctx)
	})

	// srv := handler.NewDefaultServer(
	// 	generated.NewExecutableSchema(
	// 		generated.Config{Resolvers: &graph.Resolver{}},
	// 		handler.IntrospectionEnabled(false)))

	// r.HandleFunc("/gql", playground.Handler("GraphQL playground", "/qry"))
	r.Handle("/qry", (graph.Middleware((srv)))).Methods("POST")

	r.Handle("/admincli", graph.AdminMiddleware(playground.Handler("GraphQL playground", "/query")))
	r.Handle("/query", graph.Middleware(srv)).Methods("POST")
	r.HandleFunc("/api/login", graph.LoginHandler).Methods("POST")
	r.HandleFunc("/api/register", graph.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/isset", graph.CookieChecker).Methods("POST")
	r.Handle("/api/challenge/add", graph.AdminMiddleware(http.HandlerFunc(graph.AddChallengeFile))).Methods("POST")
	r.Handle("/api/photo/add", graph.Middleware(http.HandlerFunc(graph.UserPhoto))).Methods("POST")
	r.HandleFunc("/api/scoreboard", graph.Pubscore).Methods("GET")

	challengeServer := http.FileServer(http.Dir("challenges"))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files", challengeServer))

	spa := spaHandler{staticPath: "static/frontend/build/", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	fmt.Printf("Platform runs at http://localhost:%s/\n", port)

	// production

	// srvr := &http.Server{
	// 	Handler: RequestLogger(r),
	// 	Addr:    ":" + port,
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatal(srvr.ListenAndServe())

	// developement  - add  "github.com/gorilla/handlers" in import

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowCredentials(), handlers.AllowedOrigins([]string{"http://localhost:3000"}))(RequestLogger(r))))
}
