

# gitscribe (gs)

```
           /$$   /$$                                  /$$ /$$                
          |__/  | $$                                 |__/| $$                
  /$$$$$$  /$$ /$$$$$$   /$$$$$$$  /$$$$$$$  /$$$$$$  /$$| $$$$$$$   /$$$$$$ 
 /$$__  $$| $$|_  $$_/  /$$_____/ /$$_____/ /$$__  $$| $$| $$__  $$ /$$__  $$
| $$  \ $$| $$  | $$   |  $$$$$$ | $$      | $$  \__/| $$| $$  \ $$| $$$$$$$$
| $$  | $$| $$  | $$ /$$\____  $$| $$      | $$      | $$| $$  | $$| $$_____/
|  $$$$$$$| $$  |  $$$$//$$$$$$$/|  $$$$$$$| $$      | $$| $$$$$$$/|  $$$$$$$
 \____  $$|__/   \___/ |_______/  \_______/|__/      |__/|_______/  \_______/
 /$$  \ $$                                                                    
|  $$$$$$/                                                                    
 \______/                                                                    
```

**Tu asistente git multi-agente impulsado por IA con gestión inteligente de contexto.**

gitscribe analiza tus cambios preparados y genera mensajes de commit convencionales utilizando tu modelo de IA preferido. Gestiona múltiples proveedores, mantiene tus credenciales seguras usando la bóveda de claves (keyring) de tu sistema y ahora incluye un potente sistema de contexto para instrucciones específicas del proyecto.

---

## Tabla de Contenidos

- [Características](#features)
- [Instalación](#installation)
- [Inicio Rápido](#quick-start)
- [Comandos](#commands)
  - [`gs commit` - Commits impulsados por IA](#gs-commit)
  - [`gs pr` - Creación de Pull Requests](#gs-pr)
  - [`gs context` - Gestión de Contexto](#gs-context)
  - [`gs agent` - Gestión de Agentes](#gs-agent)
  - [`gs models` - Navegador de Modelos](#gs-models)
  - [Otros Comandos](#other-commands)
- [Sistema de Contexto](#context-system)
- [Seguridad](#security)
- [Configuración](#configuration)
- [Ejemplos](#examples)
- [Solución de Problemas](#troubleshooting)

---

## Características

- **Soporte Multi-Agente**: Elige entre OpenAI (GPT-4o), Anthropic (Claude), Groq (Llama), OpenCode (Kimi) y más
- **Sistema de Contexto**: Añade instrucciones específicas del proyecto (hasta 3 por proyecto) para guiar la generación de commits por IA
- **Creación de PR**: Detecta automáticamente GitHub/GitLab y crea PR con títulos y descripciones generados por IA
- **Almacenamiento Seguro de Claves**: Todas las claves de API están cifradas en tu **Keyring del SO** (Keychain, GNOME Keyring, etc.)
- **Interfaz Interactiva**: Edita mensajes, visualiza contextos, cancela o confirma, todo con atajos de teclado
- **Retroalimentación Visual**: Spinners elegantes y salida estilizada usando herramientas de Charmbracelet
- **Flujo de Trabajo Todo en Uno**: Prepara, confirma, envía y crea PR con una sola herramienta

---

## Instalación

### Desde GitHub Releases (Recomendado)

#### Linux
```shell
# Descargar la última versión
curl -L -o gs.tar.gz https://github.com/albuquerquesz/gitscribe/releases/latest/download/gs_linux_amd64.tar.gz

# Extraer e instalar
tar -xzf gs.tar.gz
sudo mv gs /usr/local/bin/
rm gs.tar.gz
```

#### macOS
```shell
# Descargar la última versión
curl -L -o gs.tar.gz https://github.com/albuquerquesz/gitscribe/releases/latest/download/gs_darwin_amd64.tar.gz

# Extraer e instalar
tar -xzf gs.tar.gz
sudo mv gs /usr/local/bin/
rm gs.tar.gz
```

#### Windows
1. Descarga `gs_windows_amd64.tar.gz` desde la [página de releases](https://github.com/albuquerquesz/gitscribe/releases/latest)
2. Extrae `gs.exe` usando 7-Zip o similar
3. Añádelo a tu `PATH` del sistema

### Usando `go install`
```shell
go install github.com/albuquerquesz/gitscribe@latest
```

### Usando Homebrew (macOS/Linux)
```shell
brew tap albuquerquesz/gitscribe
brew install gs
```

---

## Inicio Rápido

### 1. Inicialización
```shell
gs init
```
Esto verificará configuraciones existentes y te guiará a través de la configuración.

### 2. Configura Tu Primer Modelo de IA
```shell
gs models
```
- Selecciona un proveedor (Anthropic, OpenAI, Groq, etc.)
- Elige un modelo
- Introduce tu clave de API (se almacena de forma segura en el keyring del SO)

### 3. Realiza Tu Primer Commit
```shell
gs commit
```
O usa el alias más corto:
```shell
gs cmt
```

### 4. (Opcional) Añade Contexto del Proyecto
```shell
gs ctx add "Este es un proyecto de React TypeScript que usa Redux"
gs ctx add "Sigue commits convencionales con prefijos de emoji"
```

---

## Comandos

### `gs commit` (alias: `gs cmt`)

Flujo de trabajo git add, commit y push impulsado por IA.

**Uso:**
```shell
# Preparar todos los cambios, generar commit con IA, confirmar y enviar
gs commit

# Preparar solo archivos específicos
gs commit main.go internal/auth/

# Usar mensaje personalizado (omitir generación por IA)
gs commit -m "feat: add secure key storage"

# Usar un agente específico solo para este commit
gs commit --agent claude-sonnet

# Enviar a una rama específica
gs commit -b feature-branch
```

**Prompt Interactivo:**
Después de generar el mensaje de commit, verás:
```
feat: implement user authentication

- Add JWT token validation
- Create auth middleware
- Update user model

[E] Edit  [C] Contexts  [ESC] Cancel  [↵] Continue
```

**Atajos de Teclado:**
- **E** - Editar el mensaje de commit en línea
- **C** - Ver contextos activos para este proyecto
- **ESC** - Cancelar el commit
- **Enter** - Proceder con el commit y el push

---

### `gs pr`

Crea pull requests en GitHub o GitLab con título y descripción generados por IA.

**Características:**
- Detecta automáticamente el proveedor (GitHub/GitLab) desde la URL remota
- Genera título y cuerpo desde el historial de commits
- Envía la rama automáticamente si es necesario
- Valida que existan commits entre ramas

**Requisitos:**
- [CLI de GitHub (`gh`)](https://cli.github.com/) para repos de GitHub
- [CLI de GitLab (`glab`)](https://glab.readthedocs.io/) para repos de GitLab

**Uso:**
```shell
# Generar título/cuerpo de PR con IA y crear PR
gs pr

# Crear PR con título personalizado (cuerpo generado por IA)
gs pr -t "feat: implement new feature"

# Crear PR con título y cuerpo personalizados
gs pr -t "feat: auth" -b "Detailed description here"

# Crear PR en borrador apuntando a una rama específica
gs pr --target develop --draft

# Apuntar a una rama diferente
gs pr --target staging
```

**Flujo de Trabajo:**
1. Verifica cambios sin confirmar (avisa si los hay)
2. Detecta el proveedor de Git desde la URL remota
3. Verifica que la herramienta CLI esté instalada
4. Envía la rama actual
5. Genera título/cuerpo de PR desde los últimos 20 commits
6. Muestra prompt interactivo (similar al commit)
7. Crea la PR usando la CLI del proveedor

---

### `gs context` (alias: `gs ctx`)

Gestiona contextos específicos del proyecto para guiar la generación de commits por IA.

**¿Qué son los Contextos?**
Los contextos son instrucciones o información sobre tu proyecto que ayudan a la IA a generar mejores mensajes de commit. Se almacenan por repositorio git y se incluyen en cada prompt de IA.

**Límite:** Máximo 3 contextos por proyecto (FIFO - Primero en Entrar, Primero en Salir)

#### `gs ctx add [contexto]`

Añade una cadena de contexto para el proyecto actual.

```shell
# Añadir contexto de stack tecnológico
gs ctx add "This is a Go project using Chi router and PostgreSQL"

# Añadir contexto de convenciones
gs ctx add "Use conventional commits with scope: feat(api):, fix(auth):, etc."

# Añadir directrices del equipo
gs ctx add "Always reference issue numbers: Fixes #123"
```

**Verificación de Límite:**
```shell
gs ctx add "Fourth context"
# Error: Limite de 3 contextos atingido para este projeto
# Use 'gs ctx remove' para remover um contexto existente
```

#### `gs ctx list`

Lista todos los contextos del proyecto actual.

```shell
gs ctx list
```

**Salida:**
```
Contextos para /home/user/myproject (2/3):
1. This is a Go project using Chi router and PostgreSQL
2. Use conventional commits with scope: feat(api):, fix(auth):, etc.
```

#### `gs ctx remove` (alias: `gs ctx rm`)

Elimina un contexto de forma interactiva.

```shell
gs ctx remove
```

**Flujo Interactivo:**
```
Contextos disponíveis:
1. This is a Go project using Chi router and PostgreSQL
2. Use conventional commits with scope

Qual deseja remover? (1-2): 1
✓ Contexto removido
```

Presiona `ESC` o introduce un número inválido para cancelar.

#### Cómo Afectan los Contextos a los Prompts de IA

Cuando tienes contextos configurados, la IA recibe:

```
Contextos adicionais do projeto:
- This is a Go project using Chi router and PostgreSQL
- Use conventional commits with scope: feat(api):, fix(auth):, etc.

Analise o diff abaixo considerando os contextos acima:

[git diff content here]
```

Esto resulta en mensajes de commit más relevantes y apropiados para el proyecto.

---

### `gs agent`

Gestiona perfiles de agentes de IA.

#### `gs agent list`

Lista todos los agentes configurados.

```shell
gs agent list
```

**Salida:**
```
★ 🟢 claude-sonnet       Provider: anthropic    Model: claude-3-5-sonnet    API Key: ✅
  🟢 groq-llama          Provider: groq          Model: llama-3.3-70b       API Key: ✅
  🔴 old-openai          Provider: openai        Model: gpt-4                API Key: ❌
```

**Leyenda:**
- `★` = Agente predeterminado
- `🟢`/`🔴` = Habilitado/Deshabilitado
- `✅`/`❌` = Clave de API configurada/no configurada

#### `gs agent add`

Añade un nuevo perfil de agente.

```shell
# Añadir agente Groq
gs agent add -n my-groq -p groq -m llama-3.3-70b-versatile

# Añadir con clave de API en línea
gs agent add -n production -p anthropic -m claude-3-opus -k sk-ant-xxx

# Añadir endpoint personalizado compatible con OpenAI
gs agent add -n local-ollama -p ollama -m llama2 --base-url http://localhost:11434/v1
```

**Banderas:**
- `-n, --name`: Nombre del perfil del agente (requerido)
- `-p, --provider`: Nombre del proveedor (requerido)
  - Opciones: `anthropic`, `openai`, `groq`, `opencode`, `gemini`, `openrouter`, `ollama`
- `-m, --model`: Nombre del modelo (requerido)
- `-k, --key`: Clave de API (opcional, solicita de forma segura si no se proporciona)
- `--base-url`: URL base personalizada para endpoints personalizados

#### `gs agent remove [name]`

Elimina un perfil de agente.

```shell
gs agent remove old-agent
```

#### `gs agent set-key [name]`

Actualiza la clave de API de un agente existente.

```shell
gs agent set-key my-agent
# Introduce la nueva clave de API de forma segura
```

---

### `gs models`

Navega y habilita modelos de IA de forma interactiva.

```shell
gs models
```

**TUI Interactiva:**
1. Selecciona proveedor de la lista
2. Elige modelo de las opciones disponibles
3. Introduce clave de API (entrada oculta)
4. El modelo se habilita y se establece como predeterminado

**Características:**
- Navegador visual de modelos con descripciones
- Validación automática de claves de API
- Almacenamiento seguro de claves

---

### Otros Comandos

#### `gs init`

Inicializa la configuración de GitScribe.

```shell
gs init
```

Verifica:
- Configuración existente
- Autenticación de OpenCode (`~/.local/share/opencode/auth.json`)
- Configuración del repositorio Git

#### `gs update`

Auto-actualización a la última versión.

```shell
gs update
```

Verifica releases de GitHub, muestra el registro de cambios y actualiza el binario automáticamente.

---

## Sistema de Contexto

El Sistema de Contexto es una de las características más poderosas de GitScribe. Te permite proporcionar instrucciones específicas del proyecto que guían a la IA para generar mejores mensajes de commit.

### Casos de Uso

**1. Información del Stack Tecnológico:**
```shell
gs ctx add "React 18 + TypeScript + Redux Toolkit project"
```

**2. Convenciones de Commit:**
```shell
gs ctx add "Use Angular commit convention: type(scope): subject"
gs ctx add "Available scopes: auth, api, ui, db, deps"
```

**3. Directrices del Equipo:**
```shell
gs ctx add "Reference Jira tickets in commits: PROJ-123"
gs ctx add "Mark breaking changes with BREAKING CHANGE: in body"
```

**4. Especificidades del Proyecto:**
```shell
gs ctx add "This is a microservices architecture - mention service name"
gs ctx add "Database migrations should be explicitly mentioned"
```

### Almacenamiento

Los contextos se almacenan en:
```
~/.config/gitscribe/contexts.json
```

Formato:
```json
{
  "contexts": {
    "/home/user/projects/myapp": [
      {
        "text": "React TypeScript project",
        "created_at": "2026-02-04T10:00:00Z"
      }
    ]
  }
}
```

### Límites

- **Máximo 3 contextos** por repositorio git
- **Orden FIFO** - el primero en añadirse es el primero en la lista
- **Basado en ruta** - vinculado a la raíz del repositorio git (`git rev-parse --show-toplevel`)

### Visualización Durante el Commit

Durante el flujo de commit, presiona `C` para ver los contextos activos:
```
feat: implement user authentication

[E] Edit  [C] Contexts  [ESC] Cancel  [↵] Continue

[Press C]

Contextos ativos:
  • React 18 + TypeScript + Redux Toolkit project
  • Use Angular commit convention
```

---

## Seguridad

### Almacenamiento de Claves de API

GitScribe utiliza tu **Keyring Nativo del SO** para el almacenamiento seguro de claves de API:

- **macOS**: Keychain
- **Linux**: Secret Service API / GNOME Keyring / KWallet
- **Windows**: Administrador de Credenciales de Windows

Las claves son:
- ✅ Cifradas en reposo por el SO
- ✅ Nunca almacenadas en archivos de texto plano
- ✅ Borradas de la memoria después de su uso
- ✅ Accesibles solo para tu cuenta de usuario

### Prioridad de Resolución de Claves

Cuando se necesita una clave de API, GitScribe intenta en orden:

1. **Entrada específica del agente en el keyring** (más específica)
2. **Entrada genérica del keyring** (desde la configuración)
3. **Variables de entorno**:
   - `OPENAI_API_KEY`
   - `ANTHROPIC_API_KEY`
   - `GROQ_API_KEY`
   - `OPENCODE_API_KEY`
   - `GOOGLE_API_KEY`
   - `OPENROUTER_API_KEY`
4. **Archivo de auth de OpenCode** (`~/.local/share/opencode/auth.json`)

### Entrada Segura

- Los prompts de claves de API usan enmascaramiento de contraseña
- Sin datos sensibles en los registros
- Borrado seguro de memoria después del uso

### Permisos de Archivo

```
~/.config/gitscribe/     (drwxr-xr-x)
├── config.yaml          (-rw-------)  # Solo propietario
├── contexts.json        (-rw-r--r--)  # Escritura propietario, lectura todos
└── ...
```

---

## Configuración

### Ubicación del Archivo de Configuración

```
~/.config/gitscribe/config.yaml
```

### Ejemplo de Configuración

```yaml
version: "1.0"
global:
  default_agent: "claude-sonnet"
  auto_select: true
  request_timeout_seconds: 30
  max_retries: 3

agents:
  - name: "claude-sonnet"
    provider: "anthropic"
    model: "claude-3-5-sonnet-20241022"
    temperature: 0.7
    max_tokens: 4096
    timeout_seconds: 30
    enabled: true
    priority: 1
    keyring_key: "agent:claude-sonnet:api-key"

  - name: "groq-fast"
    provider: "groq"
    model: "llama-3.1-8b-instant"
    temperature: 0.5
    max_tokens: 2048
    enabled: true
    priority: 2
    keyring_key: "agent:groq-fast:api-key"

routing:
  - name: "quick-commits"
    agent_profile: "groq-fast"
    conditions: ["token_count < 1000"]
    priority: 1
```

### Proveedores Soportados

| Proveedor | Modelos | Autenticación |
|----------|--------|----------------|
| **Anthropic** | Claude 3.5 Sonnet, Claude 3.5 Haiku | Clave de API |
| **OpenAI** | GPT-4o, GPT-4o Mini, GPT-4 | Token Bearer |
| **Groq** | Llama 3.3 70B, Llama 3.1 8B | Token Bearer |
| **OpenCode** | Kimi 2.5, Mini Pickle, GLM | Clave de API |
| **Gemini** | Gemini 1.5 Pro, Gemini 1.5 Flash | Clave de API |
| **OpenRouter** | Varios modelos | Token Bearer |
| **Ollama** | Modelos locales (Llama2, Mistral, etc.) | Ninguno (local) |

---

## Ejemplos

### Ejemplo 1: Flujo Básico

```shell
# Inicializar
gs init

# Configurar primer modelo
gs models
# Seleccionar: Groq → Llama 3.3 70B → Introducir clave de API

# Realizar cambios
echo "console.log('hello')" > app.js

# Preparar, confirmar con IA, enviar
gs commit
# Revisar mensaje → Presionar Enter → ¡Listo!
```

### Ejemplo 2: Proyecto con Contextos

```shell
# Configurar proyecto
cd my-react-project
git init

# Añadir contextos
gs ctx add "React 18 + TypeScript project with Redux Toolkit"
gs ctx add "Use conventional commits: feat:, fix:, docs:, style:, refactor:, test:, chore:"
gs ctx add "Available scopes: components, hooks, store, api, auth"

# Trabajar en función
gs commit -b feature/new-component
# IA genera: "feat(components): add UserProfile card with avatar"

# Crear PR
gs pr
# IA genera título y descripción de PR desde los commits
```

### Ejemplo 3: Configuración Multi-Agente

```shell
# Añadir múltiples agentes
gs agent add -n claude -p anthropic -m claude-3-5-sonnet-20241022
gs agent add -n fast -p groq -m llama-3.1-8b-instant

# Usar agente específico para refactorización compleja
gs commit --agent claude

# Usar agente rápido para correcciones rápidas
gs commit --agent fast -m "fix: typo in readme"

# Listar agentes
gs agent list
```

### Ejemplo 4: Flujo de PR

```shell
# Crear rama de característica
git checkout -b feature/auth-improvements

# Realizar varios commits
gs commit -m "feat(auth): add JWT refresh token logic"
gs commit -m "feat(auth): implement token rotation"
gs commit -m "test(auth): add tests for token refresh"

# Enviar y crear PR
gs pr --target main

# O con título personalizado
gs pr -t "feat: implement secure token refresh mechanism"
```

### Ejemplo 5: Flujo de Gestión de Contextos

```shell
# Iniciar nuevo proyecto
cd ~/projects/api-service
gs init

# Añadir contextos con el tiempo
gs ctx add "Go microservice using Chi router and PostgreSQL"
gs ctx add "Follow DDD patterns: repository, service, handler layers"
gs ctx add "Always include database migration notes in commits"

# Verificar contextos
gs ctx list
# Salida: 3/3 contextos

# Intentar añadir cuarto (falla)
gs ctx add "Use structured logging with zap"
# Error: Limite de 3 contextos atingido

# Eliminar uno
gs ctx remove
# Seleccionar: 2 (Follow DDD patterns...)

# Ahora se puede añadir uno nuevo
gs ctx add "Use structured logging with zap"
```

---

## Solución de Problemas

### "No commits between main and <branch>"

**Causa:** La rama no se ha enviado o no tiene commits

**Solución:**
```shell
git push origin <branch>
# Luego reintenta
gs pr
```

### "gh CLI is not installed"

**Causa:** CLI de GitHub no encontrado en PATH

**Solución:**
```shell
# macOS
brew install gh

# Linux
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
sudo apt update && sudo apt install gh

# O descargar desde https://cli.github.com/
```

### "Limite de 3 contextos atingido"

**Causa:** Límite máximo de contextos por proyecto alcanzado

**Solución:**
```shell
gs ctx remove
# Elimina un contexto antes de añadir uno nuevo
```

### Clave de API no encontrada

**Causa:** Clave no en keyring o entorno

**Solución:**
```shell
# Opción 1: Establecer vía comando de agente
gs agent set-key my-agent

# Opción 2: Usar variable de entorno
export GROQ_API_KEY="gsk_xxx"

# Opción 3: Reconfigurar modelo
gs models
```

### Commit cancelado muestra error

**Comportamiento:** Al presionar ESC durante commit, anteriormente mostraba "Error: commit cancelled"

**Comportamiento Actual:** Ahora muestra un mensaje limpio sin error:
```
ℹ Commit cancelled
```

### Spinner no se muestra

**Causa:** El terminal podría no soportar códigos ANSI

**Solución:** Prueba con la bandera `--accessible` (si está disponible) o verifica la compatibilidad del terminal.

### Contextos no aparecen en la salida de IA

**Verifica:**
1. ¿Están añadidos contextos para el proyecto actual? `gs ctx list`
2. ¿El directorio actual está dentro de un repositorio git? `git rev-parse --show-toplevel`
3. Prueba presionar `C` durante commit para verificar que los contextos se cargan

---

## Contribuir

¡Las contribuciones son bienvenidas! No dudes en enviar un Pull Request.

1. Haz un fork del repositorio
2. Crea tu rama de características (`git checkout -b feature/AmazingFeature`)
3. Confirma tus cambios (`gs commit -m 'feat: add some amazing feature'`)
4. Envía a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request (`gs pr`)

## Licencia

Licencia MIT - consulte el archivo LICENSE para más detalles.

## Agradecimientos

- Construido con herramientas de [Charmbracelet](https://charmbracelet.com/) (Bubble Tea, Lipgloss, Huh)
- Arquitectura multi-agente inspirada en patrones modernos de enrutamiento de LLM
- ¡Gracias a todos los contribuyentes y usuarios!

---

**Construido con ❤️ y mucho ☕**
