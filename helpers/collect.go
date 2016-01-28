package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/rytsar-server/accessors"
	"github.com/jessemillar/rytsar-server/models"
)

// Collect marks certain loot as collected for the specified user
func Collect(userID string, latitude string, longitude string, ag *accessors.AccessorGroup) string {
	stock := models.CheckStock(symbol)
	user := ag.GetUser(userID)

	if stock.Name == "N/A" || stock.Price == 0 {
		return fmt.Sprintf("%s does not appear to be a valid stock...\n", strings.ToUpper(symbol)) // Return the price through the API endpoint
	}

	// Make sure they have enough turnips to buy
	if user.Turnips < stock.Price*quantity {
		return fmt.Sprintf("%s shares of %s costs %s turnips and you have %s turnips.\n", Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity), Comma(user.Turnips)) // Return information about a user's portfolio
	}

	ag.SubtractTurnips(userID, stock.Price*quantity)

	if quantity == 0 {
		quantity = 1
	}

	ag.AddShares(userID, symbol, quantity)

	Webhook(fmt.Sprintf("<@%s|%s> purchased %s share(s) of %s for %s turnips.", user.UserID, user.Username, Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity)))

	return fmt.Sprintf("%s turnips were spent to add %s share(s) of %s to your portfolio.\n", Comma(stock.Price*quantity), Comma(quantity), strings.ToUpper(symbol)) // Return information about a user's portfolio
}
