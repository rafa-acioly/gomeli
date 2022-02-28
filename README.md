# Golang Mercado Livre SDK (non-oficial)

> Obs: essa lib foi construída pela comunidade, se vc deseja utilizar a biblioteca oficial mantida pelo Mercado Livre - [clique aqui](https://github.com/mercadolibre/golang-sdk)

# Funcionalidades

- [x] Autenticação e Autorização
- [ ] Consulta dos dados do usuário
- [ ] Consulta de categorias
- [ ] Consulta de moedas
- [ ] Consulta e publicação de anúncios
- [ ] Consulta de pedidos
- [ ] Consulta de pagamentos
- [ ] Envio de mensagem de pós venda para cliente
- [ ] Publicando feedback para um pedido

# Requisitos

- Golang >= 1.16.x

# Instalação

Para instalar o módulo execute:

```sh
go get -u github.com/rafa-acioly/meli
```

# Executando os testes

```sh
go test /tests
```

# Utilização

Para utilizar este modulo, primeiramente você deve ter uma [aplicação](http://applications.mercadolibre.com/) configurada no Mercado Livre. 
Caso não esteja habituado com estas funcionalidades, você pode conferir no [Getting Started](http://developers.mercadolibre.com/getting-started/) 
da página do manual do desenvolvedor.  

Após a criação da aplicação você terá as informações do **App ID (client-id)** e **Secret Key (client-secret)** disponibilizados pelo Mercado Livre. Estas informações
serão utilizadas quando você acessar algum recurso que necessita de autorização. 

Atualmente o Mercado Livre não possui um ambiente de `Sandbox` para realização de testes. Todas as publicações serão executadas na sua conta real, conforme
descrito no [manual](http://developers.mercadolibre.com/start-testing/).

Você também pode criar um [usuário de teste](http://developers.mercadolibre.com/start-testing/#Create-a-test-user) se achar necessário.
Com o usuário de teste criado, é possível configurar outra aplicação e ter a **App ID (client-id)** e **Secret Key (client-secret)** para o usuário de teste.
Lembrando novamente que mesmo com o usuário de teste, os dados aparecerão no ambiente de **produção** do Mercado Livre.

# Exemplo de autenticação e autorização

A forma de conseguir o AccessToken é realizando a consulta via `client_credentials`. Esta forma, é recomendada para scripts que rodam em rotinas automáticas (via cron, ou tarefas agendadas). **OBS:** para conseguir utilizar, você precisa ter configurado em sua APP, o **Scope offline access** marcado.

```go
import "github.com/rafa-acioly/meli"

func main() {
    client := meli.NewClient("client-id", "client-secret")
    service := meli.NewAuthService(client)

    // seu script...
}
```


**Importante:** o modulo irá armazenar o access_token e o refresh_token para utilizar nas requisiçōes que necessitarão de autenticação. Ou seja, quando o access_token estiver expirado, ele será atualizado automaticamente pelo modulo, utilizando o refresh_token.

Com o usuário autenticado já podemos publicar nosso primeiro anúncio.

# Publicando um anúncio

Com aplicação configurada e o usuário autenticado, será possível realizar a publicação de um anúncio no Mercado Livre, 
portanto, você precisa ter as informações da sua **App ID** e **Secret Key** criada na aplicação.

```go
import "github.com/rafa-acioly/meli"

func main() {
    item := meli.NewItem()

    item.
        SetTitle("Test item - no offer").
        SetCategoryId("MLB46585").
        SetPrice(100).
        SetCurrencyId("BRL").
        SetAvailableQuantity(10).
        SetBuyingMode("buy_it_now").
        SetListingTypeId("gold_especial").
        SetCondition("new").
        SetDescription("Test item - no offer").
        SetWarranty("12 months")
    
    picture := meli.Picture{Source: "https://placehold.it/200x200"}
    item.AddPicture(picture)

    cli := meli.Newclient("app-id", "secret-id")
    announcement := meli.NewAnnouncement(cli)
    response, err := announcement.Create(item)
    if err != nil {
        panic(err.Error())
    }

    fmt.Println(response.Permalink)
}
```