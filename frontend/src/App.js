import logo from './logo.svg';
import './App.css';
import React, { useEffect, useState } from "react";


function App() {
  const [mensaje, setMensaje] = useState("Cargando...");

  useEffect(() => {
    fetch("http://localhost:8080/ping")
      .then((res) => res.text())
      .then((data) => setMensaje(data))
      .catch((err) => setMensaje("Error al conectar con la API"));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      <h1>Frontend React</h1>
      <p>Respuesta de la API: {mensaje}</p>
      </header>
    </div>
  );
}

export default App;
