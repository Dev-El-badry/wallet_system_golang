package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/Dev-El-badry/wallet-system/api"
	db "github.com/Dev-El-badry/wallet-system/db/sqlc"
	"github.com/Dev-El-badry/wallet-system/gapi"
	"github.com/Dev-El-badry/wallet-system/pb"
	"github.com/Dev-El-badry/wallet-system/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/Dev-El-badry/wallet-system/doc/statik"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conf, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("can't load config file", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// run db migration
	runDBMigration(conf.MigrationURL, conf.DBSource)

	store := db.NewStore(conn)
	go runGatewayServer(conf, store)
	runGrpcServer(conf, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(
		migrationURL,
		dbSource,
	)
	if err != nil {
		log.Fatal("cannot create a new migration instance: ", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up: ", err)
	}

	log.Println("db migrated successfully")
}

func runGrpcServer(conf util.Config, store db.Store) {
	server, err := gapi.NewServer(conf, store)
	if err != nil {
		log.Fatal("cannot create a server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleWalletServer(grpcServer, server)
	reflection.Register(grpcServer) //what RPC are available on the server, and how to call them.

	listener, err := net.Listen("tcp", conf.GRPCServerAddress)
	if err != nil {
		log.Fatal("can not create listener")
	}

	log.Printf("Start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server")
	}
}

func runGatewayServer(conf util.Config, store db.Store) {
	server, err := gapi.NewServer(conf, store)
	if err != nil {
		log.Fatal("cannot create a server", err)
	}

	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOptions)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //to prevent to do unnecessary work

	err = pb.RegisterSimpleWalletHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("cannot create statik files: ", err)
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)
	// fs := http.FileServer(http.Dir("./doc/swagger"))
	// mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	listener, err := net.Listen("tcp", conf.HTTPServerAddress)
	if err != nil {
		log.Fatal("can not create listener")
	}

	log.Printf("Start HTTP Gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, grpcMux)
	if err != nil {
		log.Fatal("cannot start HTTP server")
	}
}

func runGinServer(conf util.Config, store db.Store) {
	server, err := api.NewServer(conf, store)
	if err != nil {
		log.Fatal("cannot create a server", err)
	}

	err = server.Start(conf.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
