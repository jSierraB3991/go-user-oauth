# Go User OAuth

Un módulo de Golang para la centralización de usuarios con soporte para autenticación, gestión de usuarios y autenticación de dos factores.

## Instalación

Para instalar la versión más reciente del módulo:

```go
go get github.com/jSierraB3991/go-user-oauth
```

Para instalar una versión específica:

```go
go get github.com/jSierraB3991/go-user-oauth@v0.3.6
```

## Inicialización

La inicialización del servicio debe realizarse en el archivo `main.go` de tu aplicación:

```go
package main

import (
    "time"
    "gorm.io/gorm"
    
    gooauthservice "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-service"
    gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
)

func main() {
    // Configuración previa
    var database *gorm.DB // Tu conexión a la base de datos
    jwtSecret := "tu-clave-secreta-para-jwt"
    aesEncryptKey := "tu-clave-para-encriptar"
    appName := "Mi Aplicación"
    appImageUrl := "https://miapp.com/logo.png"
    
    // Tiempo máximo para usar el QR de doble autenticación
    timeToQrToTwoFactorInMinutes := time.Duration(10)
    
    // Inicializar el servicio OAuth
    oauthService := gooauthservice.NewGoOauthService(
        database, 
        jwtSecret, 
        aesEncryptKey, 
        gooauthmodel.ServiceModeParam{
            AppName: appName,
            UrlImagenApp: appImageUrl,
        }, 
        timeToQrToTwoFactorInMinutes,
    )
    
    // Continuar con la inicialización de otros servicios...
}
```

### Parámetros de inicialización

- `database`: Conexión a la base de datos (`*gorm.DB`)
- `jwtSecret`: Clave secreta para crear tokens JWT
- `aesEncryptKey`: Clave para encriptar datos sensibles en la base de datos
- `appName` y `appImageUrl`: Información para configurar la autenticación de dos factores
- `timeToQrToTwoFactorInMinutes`: Tiempo máximo para que los usuarios puedan usar el QR de doble autenticación

## Uso en Servicios

Para utilizar el servicio OAuth en tus propios servicios (por ejemplo, en un servicio de usuarios):

```go
package services

import (
    gooauthinterface "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-interface"
    repositoryinterface "miaplicacion/domain/repository-interface"
)

type UserService struct {
    oauthService gooauthinterface.GoOauthInterface
    repo repositoryinterface.UserRepositoryInterface
}

func NewUserService(
    oauthService gooauthinterface.GoOauthInterface,
    repo repositoryinterface.UserRepositoryInterface,
) *UserService {
    return &UserService{
        oauthService: oauthService,
        repo: repo,
    }
}

// Implementa tus métodos del servicio aquí...
```

## Interfaz GoOauthInterface

La interfaz `GoOauthInterface` proporciona métodos para:

- Autenticación y autorización de usuarios
- Creación y actualización de usuarios
- Gestión de autenticación de dos factores
- Cambio y recuperación de contraseñas
- Validación de correos electrónicos

```go
type GoOauthInterface interface {
    CheckoutMiddleware(requets *http.Request) bool
    GetSecretByClient(ctx context.Context) error
    CreateUser(ctx context.Context, userParam gooauthrequest.CreateUser, roleUser string, attributes *map[string][]string) (string, error)
    UpdateUser(ctx context.Context, keyCloakUserId string, attributes *map[string][]string, req gooauthrequest.UpdateUserRequest) error
    ErrorHandler() error
    GetUserByRole(ctx context.Context, role string) ([]*gooauthrest.User, error)
    LoginUser(ctx context.Context, userName, password, ip, userAgent string) (*gooauthrest.JWT, error)
    LoginWithTwoFactor(ctx context.Context, userName, codeTwoFactor, ip, userAgent string) (*gooauthrest.JWT, error)
    GetUserByUserId(ctx context.Context, keycloakId string) (*gooauthrest.User, error)
    GetUsersByUsersId(ctx context.Context, keycloakUsersId []string) ([]gooauthrest.User, error)
    ChangePassword(ctx context.Context, req gooauthrequest.ChangePasswordRequest) error
    ValidateMailByUserId(ctx context.Context, userId string) error
    GenerateQrForDobleOuath(userName string) (*gooauthrest.QrTwoOauthRest, error)
    ValidateCodeOtp(req gooauthrequest.ValidateOauthCodeRequest) (bool, error)
    GeneratetokenToValidate(userId, keyToGenerateToken string, limitInHours time.Duration) (*string, error)
    RemenberPassword(token, newPassword, codeTwoFactor string) error
    IsActiveTwoFactorOauth(token string) (bool, error)
    IsActiveTwoFactor(user string) (bool, error)
    DisAvailableTwoFactorAuth(userEmail, codeTwoFactor string) error
}
```

## Modelos Principales

### CreateUser
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

### UpdateUserRequest
```go
type UpdateUserRequest struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}
```

### User
```go
type User struct {
    Id         string              `json:"id"`
    Email      string              `json:"email"`
    Name       string              `json:"name"`
    SubName    string              `json:"sub_name"`
    Enabled    bool                `json:"enabled"`
    Password   string              `json:"password"`
    Role       string              `json:"role"`
    Attributes *map[string][]string `json:"attributtes"`
}
```

### JWT
```go
type JWT struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiredIn    int    `json:"expired_in"`
    Role         string `json:"role"`
    IsTwoFactor  bool   `json:"is_two_factor"`
}
```

## Atributos Personalizados

Los atributos son extensiones del modelo de usuario que pueden utilizarse para almacenar información adicional:

```go
func GetAttributtesSingUp(req request.SingUpRequest *map[string][]string {
    attributes := make(map[string][]string)
    attributes["gender"] = []string{req.Gender}
    return &attributes
}
```

## Ejemplo Completo

A continuación se muestra un ejemplo completo de cómo inicializar el servicio en `main.go` y usarlo en un servicio de usuario:

```go
// main.go
package main

import (
    "log"
    "time"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    
    gooauthservice "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-service"
    gooauthmodel "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-model"
    
    "miapp/infrastructure/repositories"
    "miapp/domain/services"
)

func main() {
    // Conexión a la base de datos
    dsn := "host=localhost user=postgres password=postgres dbname=miapp port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }
    
    // Configuración del servicio OAuth
    jwtSecret := "mi-clave-super-secreta-jwt"
    aesEncryptKey := "clave-para-encriptar-16bytes"
    appName := "Mi Aplicacion"
    appImageUrl := "https://miapp.com/logo.png"
    timeToQrToTwoFactorInMinutes := time.Duration(10)
    
    // Inicializar el servicio OAuth
    oauthService := gooauthservice.NewGoOauthService(
        database, 
        jwtSecret, 
        aesEncryptKey, 
        gooauthmodel.ServiceModeParam{
            AppName: appName,
            UrlImagenApp: appImageUrl,
        }, 
        timeToQrToTwoFactorInMinutes,
    )
    
    // Inicializar repositorio de usuarios
    userRepository := repositories.NewUserRepository(database)
    
    // Inicializar servicio de usuarios
    userService := services.NewUserService(oauthService, userRepository)
    
    // Continuar con la configuración del servidor, rutas, etc.
    // ...
}
```

```go
// services/user_service.go
package services

import (
    "context"
    "time"
    
    gooauthinterface "github.com/jSierraB3991/go-user-oauth/domain/go-oauth-interface"
    gooauthrequest "github.com/jSierraB3991/go-user-oauth/infrastructure/go-oauth-request"
    
    "miapp/domain/models"
    "miapp/domain/repositories"
    "miapp/domain/request"
)

type UserService struct {
    oauthService gooauthinterface.GoOauthInterface
    repo        repositories.UserRepositoryInterface
}

func NewUserService(
    oauthService gooauthinterface.GoOauthInterface,
    repo repositories.UserRepositoryInterface,
) *UserService {
    return &UserService{
        oauthService: oauthService,
        repo: repo,
    }
}

// Ejemplo de método para registrar un usuario
func (s *UserService) RegisterUser(ctx context.Context, req request.SingUpRequest) (string, error) {
    // Preparar los atributos personalizados
    attributes := getAttributesSignUp(req)
    
    // Crear usuario en el servicio OAuth
    createUserReq := gooauthrequest.CreateUser{
        Email:      req.Email,
        UserName:   req.Email, // Usamos el email como nombre de usuario
        Password:   req.Password,
        FirstName:  req.FirstName,
        LastName:   req.LastName,
        Emailverify: true,
    }
    
    // Crear el usuario con el rol "USER"
    userId, err := s.oauthService.CreateUser(ctx, createUserReq, "USER", attributes)
    if err != nil {
        return "", err
    }
    
    // Guardar información adicional en nuestra base de datos
    user := models.User{
        ID:        userId,
        Email:     req.Email,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        CreatedAt: time.Now(),
    }
    
    err = s.repo.Create(ctx, &user)
    if err != nil {
        return "", err
    }
    
    return userId, nil
}

// Función para generar atributos para el registro
func getAttributesSignUp(req request.SingUpRequest) *map[string][]string {
    attributes := make(map[string][]string)
    attributes["gender"] = []string{req.Gender}
    return &attributes
}

// Añade más métodos según sea necesario...
```

## Contribución

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request para sugerir cambios o mejoras.

## Licencia

Este proyecto está licenciado bajo los términos de la Licencia Pública General GNU, versión 3 (GPLv3).

Resumen de la licencia:
* Puedes copiar, distribuir y modificar este proyecto siempre que mantengas la misma licencia.
* Cualquier modificación realizada y distribuida debe estar bajo los mismos términos.
* No hay garantías sobre el uso del software.

Puedes encontrar una copia completa de la licencia en el archivo LICENSE en el repositorio de este proyecto o en [https://www.gnu.org/licenses/gpl-3.0.html](https://www.gnu.org/licenses/gpl-3.0.html).