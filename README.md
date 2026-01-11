<div align="center">
  <h1>FastDeploy CLI (fdc)</h1>
  <p><strong>Tu asistente personal para desplegar aplicaciones sin complicaciones.</strong></p>
  <p><i>Orquesta despliegues complejos con comandos sencillos.</i></p>
  
  <p>
    <a href="https://github.com/jairoprogramador/fastdeploy-client/releases">
      <img src="https://img.shields.io/github/v/release/jairoprogramador/fastdeploy-client?style=for-the-badge" alt="Latest Release">
    </a>
    <a href="https://github.com/jairoprogramador/fastdeploy-client/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/jairoprogramador/fastdeploy-client?style=for-the-badge" alt="License">
    </a>
  </p>
</div>

---

**`fdc`** es una herramienta de línea de comandos (CLI) que actúa como un cliente inteligente para `fd`. Su misión es simplificar al máximo el proceso de despliegue, permitiéndote inicializar y ejecutar el despliegue de tus proyectos en un entorno contenerizado con una configuración mínima y comandos intuitivos.

Olvídate de la complejidad de Docker y los detalles de bajo nivel, `fdc` es el puente que te conecta con un motor de despliegue potente, haciendo que el proceso sea simple y repetible.

## ✨ Características Principales

*   **🚀 Inicialización Rápida**: Con `fdc init`, la herramienta genera un archivo `fdconfig.yaml` adaptado a tus necesidades.
*   **📄 Configuración Declarativa**: Define tu configuracion de despliegue en un único archivo `fdconfig.yaml`. Fácil de leer, modificar y versionar.
*   **🐳 Abstracción de Docker**: `fdc [step] [environment]` se encarga de construir la imagen de Docker y ejecutar comando en el contenedor que aloja a `fd`. No necesitas ser un experto.
*   **🔌 Orquestación Transparente**: Actúa como un punto de entrada único para `fd`, pasándole tus instrucciones y gestionando el ciclo de vida del contenedor por ti.

## 🚀 Instalación

Instala `fdc` en segundos.

*(Nota: Las siguientes instrucciones son un ejemplo. Ajústalas según tu método de distribución final).*

### macOS (Homebrew)
```sh
# brew tap jairoprogramador/fastdeploy-client
# brew install fastdeploy-cliente
```

### Linux
Puedes descargar el paquete `.deb` o `.rpm` desde la [página de Releases](https://github.com/jairoprogramador/fastdeploy-client/releases) y usar tu gestor de paquetes.

```sh
# Para sistemas basados en Debian/Ubuntu
sudo dpkg -i fastdeploy-client_*.deb

# Para sistemas basados en Red Hat/Fedora
sudo rpm -i fastdeploy-client_*.rpm
```
Alternativamente, puedes descargar el binario directamente:
```sh
curl -sL https://github.com/jairoprogramador/fastdeploy-client/releases/latest/download/fastdeploy-client_linux_amd64.tar.gz | tar xz

sudo mv fdc /usr/local/bin/
```

### Windows
1.  Descarga el archivo `fastdeploy-client_*_windows_a*64.zip` desde la [página de Releases](https://github.com/jairoprogramador/fastdeploy-client/releases).
2.  Descomprime el archivo.
3.  Añade el ejecutable `fdc.exe` a tu variable de entorno `PATH`.

## 🏁 Guía de Inicio Rápido

Este es el flujo de trabajo típico con `fdc`.

### Paso 1: Inicializa tu Proyecto

Navega al directorio raíz de tu proyecto y ejecuta:
```sh
fdc init
```
La herramienta te guiará con unas sencillas preguntas para generar el archivo `fdconfig.yaml`, que conecta tu proyecto con la plantilla de despliegue de `fastdeploy`.

### Paso 2: Ejecuta los Pasos de Despliegue

Una vez configurado, usa el comando `fdc` para enviar instrucciones directamente a `fastdeploy`. Los `steps` como `test`, `supply`, `package`  o `deploy` son gestionados por el motor de `fastdeploy`, no por esta CLI.

Por ejemplo, para ejecutar las pruebas en el entorno de `sand`:
```sh
fdc test sand
```
Para desplegar en el mismo entorno:
```sh
fdc deploy sand
```
`fdc` se encargará de iniciar el contenedor con el core de `fd` y le pasará estos comandos para que los ejecute.

## 📚 Comandos Básicos

| Comando | Descripción |
| :--- | :--- |
| `fdc init` | Inicializa un proyecto creando el archivo de configuración `fdconfig.yaml`. |
| `fdc [step] [env]` | Ejecuta un comando en `fastdeploy`. Los `steps` (`test`, `supply`, `deploy`, etc.) dependen de la plantilla utilizada. |
| `fdc version` | Muestra la versión de la CLI. |

**Flags comunes:**
*   `--yes` o `-y`: Salta las confirmaciones interactivas para `fdc init`.

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si tienes ideas, sugerencias o encuentras un error, por favor abre un [issue](https://github.com/jairoprogramador/fastdeploy-client/issues) o envía un [pull request](https://github.com/jairoprogramador/fastdeploy-client/pulls).

## 📄 Licencia

`fdc` está distribuido bajo la [Apache License 2.0](https://github.com/jairoprogramador/fastdeploy-client/blob/main/LICENSE).
