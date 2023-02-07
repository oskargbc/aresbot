package constants

import (
	"fmt"
)

var ErrWrongTaskModeSelected = fmt.Errorf("wrong task mode selected")
var ErrGeneratingCookiesFailed = fmt.Errorf("generating cookies failed")
var ErrLoginFailed = fmt.Errorf("account login failed")
var ErrProductNotFound = fmt.Errorf("product not found")
var ErrErrorResponseDecodeFailed = fmt.Errorf("could not decode error response")
var ErrSizeNotFound = fmt.Errorf("size not found")
var ErrProductOos = fmt.Errorf("product out of stock")
var ErrAtcStepFailed = fmt.Errorf("adding product to cart failed")
var ErrRateLimited = fmt.Errorf("task/proxy is rate limitied")
var ErrForbiddenResponse = fmt.Errorf("ip/proxy blocked")
var ErrUnexpectedStatusCode = fmt.Errorf("unexpected status code returned")
var ErrQueueHandlingFailed = fmt.Errorf("failed to handle checkout queue")
var ErrPreviousResponseError = fmt.Errorf("previous response error")
var ErrAddressStepFailed = fmt.Errorf("submitting address failed")
var ErrShippingStepFailed = fmt.Errorf("submitting shippig failed")
var ErrPickupSelection = fmt.Errorf("failed selecting pickup location")
var ErrPaymentStepFailed = fmt.Errorf("submitting payment failed")
var ErrCheckoutStepFailed = fmt.Errorf("checkout failed")
var ErrCreditCardFailed = fmt.Errorf("credit card payment failed")
var ErrOosOnCheckout = fmt.Errorf("oos on checkout")
var ErrGoogleRecaptcha = fmt.Errorf("blocked by google recaptcha")
var ErrGeneralError = fmt.Errorf("general error occured")
var ErrKeywordsNotSupportedOnThisShop = fmt.Errorf("keyword matching not supported")
var ErrProxySelectionFailed = fmt.Errorf("proxy selection failed")
var ErrBadLicence = fmt.Errorf("bad licence")
var ErrConnectionToLicenceServer = fmt.Errorf("can't connect to licence server")
var ErrNoQuickTaskSettings = fmt.Errorf("no quicktask settings found")
var ErrNoCaptchaProvider = fmt.Errorf("no captcha provider given")

type GeneralError struct {
	msg string
}

func NewGeneralError(msg string) error {
	return &GeneralError{
		msg: msg,
	}
}

func (g *GeneralError) Error() string {
	return g.msg
}
