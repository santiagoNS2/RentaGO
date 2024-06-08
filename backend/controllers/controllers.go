package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/santiagoNS2/go_docker_compose/models"
	repositorio "github.com/santiagoNS2/go_docker_compose/repository"
)

var (
	updateQuery = "UPDATE usuarios SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM usuarios WHERE id=$1;"
	selectQuery = "SELECT id, name, email, password FROM usuarios WHERE id=$1;"
	listQuery   = "SELECT id, name, email, password FROM usuarios limit $1 offset $2;"
	createQuery = "INSERT INTO usuarios (name, email, password) VALUES (:name, :email, :password) returning id;"
)

type Controller struct {
	repo repositorio.Repository[models.Usuario]
}

func NewController(repo repositorio.Repository[models.Usuario]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) ActualizarUnComentario(reqBody []byte, id string) error {
	nuevosValoresUsuario := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresUsuario)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}

	if len(nuevosValoresUsuario) == 0 {
		log.Printf("fallo al actualizar un usuario, con error: no hay datos")
		return fmt.Errorf("fallo al actualizar un usuario, con error: no hay datos")
	}

	query := construirUpdateQuery(nuevosValoresUsuario)
	nuevosValoresUsuario["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresUsuario)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columnas := []string{}
	for key := range nuevosValores {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return fmt.Sprintf(updateQuery, columnasString)
}

func (c *Controller) EliminarUnComentario(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un usuario, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) LeerUnComentario(id string) ([]byte, error) {
	usuario, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un usuario, con error: %s", err.Error())
	}

	usuarioJson, err := json.Marshal(usuario)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un usuario, con error: %s", err.Error())
	}
	return usuarioJson, nil
}

func (c *Controller) ListarComentarios(limit, offset int) ([]byte, error) {
	usuarios, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("[Controller] Fallo al leer usuarios , con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer usuarios, con error: %s", err.Error())
	}

	jsonUsuarios, err := json.Marshal(usuarios)
	if err != nil {
		log.Printf("fallo al leer usuarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer usuarios, con error: %s", err.Error())
	}
	return jsonUsuarios, nil
}

func (c *Controller) CrearComentario(reqBody []byte) (int64, error) {
	nuevoUsuario := &models.Usuario{}
	err := json.Unmarshal(reqBody, nuevoUsuario)
	if err != nil {
		log.Printf("fallo al crear un nuevo usuario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo usuario, con error: %s", err.Error())
	}

	valoresColumnasNuevoUsuario := map[string]any{
		"name":     nuevoUsuario.Name,
		"email":    nuevoUsuario.Email,
		"password": nuevoUsuario.Password,
	}

	nuevoId, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoUsuario)
	if err != nil {
		log.Printf("fallo al crear un nuevo usuario, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo usuario, con error: %s", err.Error())
	}
	return nuevoId, nil
}
