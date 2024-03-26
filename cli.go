package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// Iniciar o servidor
	go func() {
		if err := startServer(); err != nil {
			fmt.Println("Erro ao iniciar o servidor:", err)
			return
		}
	}()

	// Abrir o arquivo HTML após um pequeno atraso
	go func() {
		time.Sleep(1 * time.Second) // Atraso de 1 segundo para abrir o arquivo
		if err := openFile("./view/cadastro.html"); err != nil {
			fmt.Println("Erro ao abrir o arquivo HTML:", err)
			return
		}
	}()

	// Aguardar indefinidamente
	select {}
}

func startServer() error {
	fmt.Println("Iniciando o servidor...")
	cmd := exec.Command("./src/main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao iniciar o servidor: %v", err)
	}
	return nil
}

func openFile(filepath string) error {
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
		return fmt.Errorf("erro ao abrir o arquivo %s: %v", filepath, err)
	}
	return nil
}
