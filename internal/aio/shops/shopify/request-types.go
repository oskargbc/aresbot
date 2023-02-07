package shopify

import (
	"aresbot/internal/aio/models"
	"fmt"
	"time"
)

type ProductList struct {
	Products []Product `json:"products"`
}

type SingleProduct struct {
	Product Product `json:"product"`
}

type Product struct {
	Images []struct {
		ID         int64         `json:"id"`
		CreatedAt  time.Time     `json:"created_at"`
		Position   int           `json:"position"`
		UpdatedAt  time.Time     `json:"updated_at"`
		ProductID  int64         `json:"product_id"`
		VariantIds []interface{} `json:"variant_ids"`
		Src        string        `json:"src"`
		Width      int           `json:"width"`
		Height     int           `json:"height"`
	} `json:"images"`
	Options []struct {
		Name     string   `json:"name"`
		Position int      `json:"position"`
		Values   []string `json:"values"`
	} `json:"options"`
	PublishedAt time.Time        `json:"published_at"`
	Variants    []ProductVariant `json:"variants"`
	BodyHTML    string           `json:"body_html"`
	Handle      string           `json:"handle"`
	ID          int64            `json:"id"`
	ProductType string           `json:"product_type"`
	Title       string           `json:"title"`
	Vendor      string           `json:"vendor"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type ProductVariant struct {
	FeaturedImage    interface{} `json:"featured_image"`
	Option2          interface{} `json:"option2"`
	Option3          interface{} `json:"option3"`
	Available        bool        `json:"available"`
	CompareAtPrice   string      `json:"compare_at_price"`
	Grams            int         `json:"grams"`
	ID               int64       `json:"id"`
	Option1          string      `json:"option1"`
	Position         int         `json:"position"`
	Price            string      `json:"price"`
	ProductID        int64       `json:"product_id"`
	RequiresShipping bool        `json:"requires_shipping"`
	Sku              string      `json:"sku"`
	Taxable          bool        `json:"taxable"`
	Title            string      `json:"title"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

func BuildShippingDataMap(task models.Task) map[string]string {
	return map[string]string{
		"email_or_phone": task.Profile.Phone,
		"customer_email": task.Profile.Email,
		"first_name":     task.Profile.FirstName,
		"last_name":      task.Profile.LastName,
		"company":        "",
		"address1":       fmt.Sprintf("%s+%s", task.Profile.Address, task.Profile.Address2),
		"address2":       task.Profile.Address2,
		"zip":            task.Profile.Zip,
		"city":           task.Profile.City,
		"country":        task.Profile.Country,
		"phone":          task.Profile.Phone,
	}
}

var StaticFields = map[string]string{
	"checkout[remember_me]":                "0",
	"checkout[buyer_accepts_marketing]":    "0",
	"checkout[pick_up_in_store][selected]": "0",
	"checkout[id]":                         "delivery-shipping",
}

var StaticClientFields = map[string]string{
	"checkout[client_details][browser_width]":      "1200",
	"checkout[client_details][browser_height]":     "851",
	"checkout[client_details][javascript_enabled]": "1",
	"checkout[client_details][color_depth]":        "24",
	"checkout[client_details][java_enabled]":       "false",
	"checkout[client_details][browser_tz]":         "-60",
}

var AdditionalStaticFields = map[string]string{
	"checkout[remember_me]": "false",
}

type CreditCardRequest struct {
	CreditCard          ShopifyCreditCard `json:"credit_card"`
	PaymentSessionScope string            `json:"payment_session_scope"`
}

type ShopifyCreditCard struct {
	IssueNumber       string `json:"issue_number"`
	Month             int    `json:"month"`
	Name              string `json:"name"`
	Number            string `json:"number"`
	StartMonth        int    `json:"start_month"`
	StartYear         int    `json:"start_year"`
	VerificationValue string `json:"verification_value"`
	Year              int    `json:"year"`
}

type CreditCardResponse struct {
	ID string `json:"id"`
}
