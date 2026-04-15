<div align="center">
  <h1>🛡️ Go User OAuth</h1>
  <p><strong>Un módulo de Golang robusto para la centralización de usuarios con soporte nativo para autenticación, gestión avanzada, sesiones y autenticación de dos factores (2FA).</strong></p>
  
  [![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-00ADD8?style=flat&logo=go)](https://golang.org/)
  [![License](https://img.shields.io/badge/License-GPLv3-green.svg)](https://www.gnu.org/licenses/gpl-3.0)
</div>

---

## 📖 Tabla de Contenidos

- [✨ Características Principales](#-características-principales)
- [⚙️ Instalación](#️-instalación)
- [🚀 Inicialización](#-inicialización)
- [💻 Uso en Servicios Arquitectura Limpia](#-uso-en-servicios-arquitectura-limpia)
- [🔌 Interfaz `GoOauthInterface`](#-interfaz-gooauthinterface)
- [📦 Modelos Principales](#-modelos-principales)
- [🎯 Ejemplos Extendidos](#-ejemplos-extendidos)
- [🤝 Contribución](#-contribución)
- [📄 Licencia](#-licencia)

---

## ✨ Características Principales

- 🔐 **Autenticación Completa**: Login seguro con JWT y gestión mediante *Refresh Tokens*.
- 🛡️ **Autenticación de Dos Factores (2FA)**: Soporte completo de TOTP y códigos QR.
- 👥 **Gestión de Usuarios Avanzada**: CRUD de usuarios, roles, búsqueda paginada y atributos extensibles.
- 📱 **Gestión de Sesiones**: Control y auditoría de sesiones activas, cierre de sesión remoto.
- 📧 **Comunicaciones**: Validación de emails, restablecimiento de contraseñas (con y sin 2FA).
- 🛠️ **Administración**: Creación y validación de usuarios administradores.

---

## ⚙️ Instalación

Para integrar el módulo a tu proyecto, instala la última versión:

```go
go get github.com/jSierraB3991/go-user-oauth
```

O si necesitas anclar una versión específica:

```go
go get github.com/jSierraB3991/go-user-oauth@v0.3.6
```

---

## 🚀 Inicialización

El núcleo de este módulo debe ser instanciado en tu archivo de arranque (p.ej. `main.go`).

```go
package main

import (
	"time"

	"gorm.io/gorm"
	
	gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
	gooauthservice "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-service"
)

func main() {
	// ... Configuración de Base de Datos y entono
	var database *gorm.DB // Instancia viva de *gorm.DB (postgres, mysql, etc.)
	
	// Variables de entorno o secrectos
	jwtSecret := "MI_SUPER_SECRETO_PARA_JWT"
	aesEncryptKey := "CLAVE_DE_ENCRIPTACION_32_BYTES" // 16, 24, o 32 bytes
	appName := "Mi App Increíble"
	appImageUrl := "https://miapp.com/assets/logo.png"
	
	// Tiempo máximo de expiración para el QR del 2FA
	timeToQrToTwoFactorInMinutes := time.Duration(10)
	
	oauthService := gooauthservice.NewGoOauthService(
		database, 
		jwtSecret, 
		aesEncryptKey, 
		gooauthmodel.ServiceModeParam{
			AppName:      appName,
			UrlImagenApp: appImageUrl,
		}, 
		timeToQrToTwoFactorInMinutes,
	)
	
	// Inyectar en tus servicios (Ver la siguiente sección)
}
```

### 📋 Parámetros de inicialización

| Parámetro | Tipo | Descripción |
| :--- | :--- | :--- |
| `database` | `*gorm.DB` | Conexión GORM activa. |
| `jwtSecret` | `string` | Secreto para firmar tokens JWT para las sesiones. |
| `aesEncryptKey` | `string` | Llave AES simétrica para encriptación de datos sensibles. |
| `serviceModeParam` | `struct` | Configuración estética (Nombre de la App, Logo) para correos y 2FA. |
| `timeToQr...` | `time.Duration` | Tiempo de vida para vincular un dispositivo TOTP mediante QR. |

---

## 💻 Uso en Servicios Arquitectura Limpia

Inyecta la interfaz `GoOauthInterface` en cualquier controlador o servicio que requiera acciones de autenticación.

```go
package services

import (
	gooauthinterface "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-interface"
)

type AuthService struct {
	oauthService gooauthinterface.GoOauthInterface
	// Otras dependencias...
}

func NewAuthService(oauthService gooauthinterface.GoOauthInterface) *AuthService {
	return &AuthService{
		oauthService: oauthService,
	}
}
```

---

## 🔌 Interfaz `GoOauthInterface`

La interfaz ha evolucionado bastante, soportando de manera estructurada la mayoría de casos de uso requeridos en plataformas modernas. **Esta es la declaración completa obligatoria:**

```go
type GoOauthInterface interface {
	// --- Gestión de Request / Errores ---
	CheckoutMiddleware(requets *http.Request) bool
	GetSecretByClient(ctx context.Context) error
	ErrorHandler() error

	// --- Gestión de Autenticación y Tokens ---
	LoginUser(ctx context.Context, req gooauthrequest.GoLoginRequest) (*gooauthrest.JWT, error)
	LoginWithTwoFactor(ctx context.Context, req gooauthrequest.GoLoginRequestTwoFactor) (*gooauthrest.JWT, error)
	ValidateTokenIsValidSession(ctx context.Context, tokenStr string) error
	RefreshToken(ctx context.Context, refreshToken string) (*gooauthrest.JWT, error)

	// --- Gestión de Usuarios (CRUD y Consultas) ---
	CreateUser(ctx context.Context, userParam gooauthrequest.CreateUser, roleUser string, attributes *map[string][]string) (string, error)
	UpdateUser(ctx context.Context, keyCloakUserId string, attributes *map[string][]string, req gooauthrequest.UpdateUserRequest) error
	UpdateOneAttr(ctx context.Context, keyCloakUserId string, attribute string, value string) error
	
	GetUserByRole(ctx context.Context, role string) ([]*gooauthrest.User, error)
	GetUserByUserId(ctx context.Context, keycloakId string) (*gooauthrest.User, error)
	GetUserByEmail(ctx context.Context, email string) (*gooauthrest.User, error)
	GetUserByName(ctx context.Context, name string, page *eliotlibs.Paggination) ([]string, error)
	GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error)
	GetUsersByEmail(ctx context.Context, emails []string) ([]gooauthrest.User, error)

	// --- Credenciales, Contraseñas y Validaciones ---
	ChangePassword(ctx context.Context, req gooauthrequest.ChangePasswordRequest) error
	RemenberPassword(ctx context.Context, token, newPassword, codeTwoFactor string) error
	ValidateMailByUserId(ctx context.Context, userId string) error
	GeneratetokenToValidate(ctx context.Context, userId, keyToGenerateToken string, limitInHours time.Duration) (*string, error)
	GenerateValidateMail(ctx context.Context, mailSend, keyToGenerateToken string) (string, error)

	// --- Autenticación Dos Factores (2FA) ---
	GenerateQrForDobleOuath(ctx context.Context, userName, appName, imageUrl string) (*gooauthrest.QrTwoOauthRest, error)
	ValidateCodeOtp(ctx context.Context, req gooauthrequest.ValidateOauthCodeRequest) (bool, error)
	IsActiveTwoFactorOauth(ctx context.Context, token string) (bool, error)
	IsActiveTwoFactor(ctx context.Context, user string) (bool, error)
	DisAvailableTwoFactorAuth(ctx context.Context, userEmail, codeTwoFactor string) error

	// --- Administración y Mantenimiento de Usuarios ---
	CreateUserAdministrator(ctx context.Context, userEmail, userpassword, appName string, attributes *map[string][]string) (string, error)
	ExistsUserAdministrator(ctx context.Context) (bool, error)
	ChangeEmailByAdmin(ctx context.Context, kUserId, newEmail string) error
	ChangePasswordToGeneric(ctx context.Context, kUserId string) error
	RemoveUserTwoMonthsNoValidate(ctx context.Context, usersNoRemove []string) ([]string, error)

	// --- Sesiones e Inicios de Sesión y Auditoría ---
	GetInvalidLogins(ctx context.Context, page *eliotlibs.Paggination) (*gooauthrest.InvalidLoginRestPagg, error)
	GetActiveSessions(ctx context.Context, email, refreshToken string, page, limit int) (*gooauthrest.LoginSessionRestPagination, error)
	RemoveSessionByRefreshToken(ctx context.Context, email, refreshToken string) error
	RemoveSessionById(ctx context.Context, sessionId uint) error
}
```

---

## 📦 Modelos Principales

A continuación se listan las estructuras de datos vitales para usar el módulo:

### `CreateUser`
```go
type CreateUser struct {
	Email       string `json:"email"`
	UserName    string `json:"-"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Emailverify bool   `json:"-"`
}
```

### `UpdateUserRequest`
```go
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
```

### `User` (Rest)
```go
type User struct {
	Id         string               `json:"id"`
	Email      string               `json:"email"`
	Name       string               `json:"name"`
	SubName    string               `json:"sub_name"`
	Enabled    bool                 `json:"enabled"`
	Password   string               `json:"password"`
	Role       string               `json:"role"`
	Attributes *map[string][]string `json:"attributtes"` // Metadatos extras
}
```

### `JWT`
```go
type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int    `json:"expired_in"`
	Role         string `json:"role"`
	IsTwoFactor  bool   `json:"is_two_factor"`
}
```

---

## 🎯 Ejemplos Extendidos

### 1. Iniciar Sesión (Login) y Uso de 2FA
Un flujo clásico para autenticar un usuario verificando si requiere segundo factor de autenticación.

```go
func (s *AuthService) SignIn(ctx context.Context, email, password string) (*gooauthrest.JWT, error) {
	req := gooauthrequest.GoLoginRequest{
		Email:    email,
		Password: password,
	}

	// El servicio internamente guarda el historial y valida intentos fallidos
	return s.oauthService.LoginUser(ctx, req)
}

func (s *AuthService) Verify2FA(ctx context.Context, email, password, otpCode string) (*gooauthrest.JWT, error) {
	req := gooauthrequest.GoLoginRequestTwoFactor{
		Email:     email,
		Password:  password,
		CodeOauth: otpCode,
	}

	return s.oauthService.LoginWithTwoFactor(ctx, req)
}
```

### 2. Gestión de Sesiones Activas
Ideal para mostrar la pantalla de "Dispositivos conectados" permitiendo cerrar sesiones activas remótamente.

```go
// Listado de sesiones activas paginadas
func (s *AuthService) GetSessions(ctx context.Context, email, refreshToken string) (*gooauthrest.LoginSessionRestPagination, error) {
	page, limit := 1, 10
	return s.oauthService.GetActiveSessions(ctx, email, refreshToken, page, limit)
}

// Cerrar sesión / revocación
func (s *AuthService) Logout(ctx context.Context, email, refreshToken string) error {
	return s.oauthService.RemoveSessionByRefreshToken(ctx, email, refreshToken)
}
```

### 3. Registro de Usuario con Atributos Flexibles

```go
func (s *AuthService) Register(ctx context.Context, payload RegisterPayload) (string, error) {
	createUserReq := gooauthrequest.CreateUser{
		Email:       payload.Email,
		UserName:    payload.Email,
		Password:    payload.Password,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Emailverify: false,
	}

	// Atributos extendidos que puedes usar para lo que requieras
	attributes := map[string][]string{
		"gender":           {payload.Gender},
		"terms_accepted":   {"true"},
		"registration_ip":  {"192.168.1.1"},
	}

	userId, err := s.oauthService.CreateUser(ctx, createUserReq, "USER", &attributes)
	if err != nil {
		return "", err
	}

	return userId, nil
}
```

### 4. Renovar Token de Acceso (Refresh Token)

```go
func (s *AuthService) RenewToken(ctx context.Context, refreshToken string) (*gooauthrest.JWT, error) {
	return s.oauthService.RefreshToken(ctx, refreshToken)
}
```

---

## 🤝 Contribución

¡Tus pull requests y reportajes de problemas son bienvenidos! 
Si tienes alguna idea para mejorar el código o ampliar la documentación, por favor abre un _issue_ o un _pull request_. Para cambios mayores, abre primero un _issue_ para discutir lo que te gustaría cambiar.

---

## 📄 Licencia

Este proyecto está licenciado bajo los términos de la [Licencia Pública General GNU, versión 3 (GPLv3)](https://www.gnu.org/licenses/gpl-3.0.html).

### Resumen Rápido:
- ✅ Puedes copiar, modificar y usar este componente en todos los entornos que desees.
- ❗️ Cualquier modificación al código fuente del módulo debe distribuirse bajo la misma licencia.
- 🚫 Se entrega tal cual como está sin ningún tipo de garantías legales.

<br>
<p align="center">Construido con ❤️ para ecosistemas distribuidos en <strong>Golang</strong>.</p>