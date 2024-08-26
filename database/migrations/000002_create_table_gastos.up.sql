-- Crear la tabla gastos:
CREATE TABLE IF NOT EXISTS baptiste.gastos (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL,
    costo_del_gasto INT NOT NULL,
    fecha_del_gasto TIMESTAMPTZ NOT NULL,
    categoria TEXT NOT NULL,
    lugar TEXT NOT NULL
);