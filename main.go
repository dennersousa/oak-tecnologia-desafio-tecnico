package main

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	ForSale     string  `json:"for_sale"`
}

var db *sql.DB

func main() {
	// Configurar o modo Gin para "release mode"
	gin.SetMode(gin.ReleaseMode)

	// Inicializar o banco de dados
	if err := InitDB("database.db"); err != nil {
		fmt.Println("Erro ao inicializar o banco de dados:", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o banco de dados:", err)
		}
	}(db)

	// Configurar rotas da API
	r := gin.Default()

	// Adicionar middleware de CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.POST("/cadastrar-produto", cadastrarProdutoHandler)
	r.GET("/listagem-produtos", listaProdutosHandler)

	// Determinar a porta disponível automaticamente
	port := ":8080" // Porta padrão
	go func() {
		if err := http.ListenAndServe(port, r); err != nil {
			fmt.Println("Erro ao iniciar o servidor:", err)
			return
		}
	}()

	// Abrir o arquivo cadastro.html após um pequeno atraso
	go func() {
		time.Sleep(1 * time.Second) // Atraso de 1 segundos para abrir o arquivo
		openFile("view/cadastro.html")
	}()

	// Aguardar indefinidamente
	select {}
}

func InitDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (name TEXT, description TEXT, value REAL, for_sale TEXT)")
	return err
}

func CreateProduct(name, description string, value float64, forSale string) error {
	_, err := db.Exec("INSERT INTO products (name, description, value, for_sale) VALUES (?, ?, ?, ?)",
		name, description, value, forSale)
	return err
}

func GetAllProducts() ([]Product, error) {
	rows, err := db.Query("SELECT name, description, value, for_sale FROM products")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Erro ao fechar o cursor:", err)
		}
	}(rows)

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Name, &p.Description, &p.Value, &p.ForSale)
		if err != nil {
			return nil, err
		}
		// Formatar o valor para duas casas decimais
		p.Value = round(p.Value, 2)
		products = append(products, p)
	}
	return products, nil
}

func cadastrarProdutoHandler(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CreateProduct(product.Name, product.Description, product.Value, product.ForSale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Produto cadastrado com sucesso"})
}

func listaProdutosHandler(c *gin.Context) {
	products, err := GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// round arredonda o número 'n' para 'places' casas decimais.
func round(n float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(n*shift) / shift
}

func openFile(filepath string) {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", filepath).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", filepath).Start()
	case "linux":
		err = exec.Command("xdg-open", filepath).Start()
	default:
		err = fmt.Errorf("sistema operacional não suportado")
	}
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
	}
}
