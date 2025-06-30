package helper

type Msg struct {
	Body string
}

var (
	MsgErrorMenu  = &Msg{Body: "❗ Entrada inválida. Digite *menu* ou escolha uma opção abaixo:"}
	MsgWellcome   = &Msg{Body: "🍔 Que bom te ver por aqui! Qual vai ser seu pedido?"}
	MsgSeller     = &Msg{Body: "👨‍🍳 Aguardando um atendente. Ele falará com você em breve!"}
	MsgMenu       = &Msg{Body: "📋 Cardápio: X-Burger, X-Bacon, X-Egg..."}
	MsgPromos     = &Msg{Body: "🔥 Promoções do dia: 2x1 no X-Burger até às 21h!"}
)
