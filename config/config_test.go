package config

import (
	"testing"

	"baptiste.com/database"
	"github.com/stretchr/testify/assert"
)

func TestConstructorServer_Exitosa(t *testing.T) {
	// 1. Valores de prueba:
	//    a. Definimos el puerto en el que se espera que el servidor escuche.
	//    b. Definimos un DNS simulado que se utilizará para la conexión a la base de datos.
	port := ":8080"
	dns := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable search_path=baptiste"

	// 2. Mock de la base de datos:
	//    a. Creamos un mock de la base de datos que simula un objeto de base de datos real.
	//    b. Sobrescribimos la función dbConstructor para devolver nuestro mock en lugar de una conexión real.
	databaseMock := &database.Database{}
	mockDbConstructor := func(dns string) (*database.Database, error) {
		return databaseMock, nil
	}

	// 3. Llamada al constructor:
	//    a. Llamamos a la función ConstructorServer con los valores de prueba y el mock.
	//    b. Se espera que esto inicialice un servidor con la configuración dada.
	server, err := ConstructorServer(port, dns, mockDbConstructor)

	// 4. Verificaciones:
	//    a. Verifica que no exista un error al inicializar el servidor.
	assert.NoError(t, err)

	//    b. Verifica que el servidor no sea nil, lo que indica que se inicializó correctamente.
	assert.NotNil(t, server)

	//    c. Verifica que el puerto del servidor sea el mismo que el proporcionado en la prueba.
	assert.Equal(t, port, server.Port)

	//    d. Verifica que el router del servidor se haya inicializado correctamente (no sea nil).
	assert.NotNil(t, server.Router)
}
