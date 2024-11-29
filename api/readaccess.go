package handler



import (
	"encoding/json"
	"fmt"
	"log"
	"context"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v4/pgxpool"
)


type AccessSite struct {
	ID       int
	Site     string
	IP       string
	Hostname string
	Date     time.Time
	Provedor string
	City     string
	State    string
	Country  string
}


type AccessSiteJson struct {
	ID       int       `json:"id"`
	Site     string    `json:"site"`
	IP       string    `json:"ip"`
	Hostname string    `json:"hostname"`
	Date     time.Time `json:"date"`
	Provedor string    `json:"provedor"`
	City     string    `json:"city"`
	State    string    `json:"state"`
	Country  string    `json:"country"`
}

type ResponseModel struct {
	Success int
	Message string
}



func ReadAccess(w http.ResponseWriter,r*http.Request){
	pool := ConnectDB()
        defer pool.Close()
	infoAccess,err := ReadAccessSites(pool)
	if err == nil{
		w.Header().Set("Content-Type", "application/json")
		if errJson := json.NewEncoder(w).Encode(infoAccess); errJson != nil {
			http.Error(w, "Erro ao gerar o JSON", http.StatusInternalServerError)
			log.Printf("Erro ao codificar JSON: %v", errJson)
	}
	}else{
		fmt.Fprintf(w,"error: %v\n",err)
	}
}


func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}


func ReadAccessSites(pool *pgxpool.Pool) ([]AccessSiteJson, error) {
    rows, err := pool.Query(context.Background(), "SELECT * FROM table_access_sites")
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar registros: %v", err)
    }
    defer rows.Close()

    var sites []AccessSiteJson
    for rows.Next() {
        var site AccessSiteJson
        err := rows.Scan(&site.ID, &site.Site, &site.IP, &site.Hostname, &site.Date, &site.Provedor, &site.City, &site.State, &site.Country)
        if err != nil {
            return nil, fmt.Errorf("erro ao escanear registro: %v", err)
        }
        sites = append(sites, site)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("erro ao iterar registros: %v", err)
    }

    return sites, nil
}



// Conectar ao banco de dados
func ConnectDB() *pgxpool.Pool {
	LoadEnv()
	databaseURL := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Erro ao analisar a URL do banco de dados: ", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
	fmt.Println("Conexão com o banco de dados estabelecida.")
	return conn
}






