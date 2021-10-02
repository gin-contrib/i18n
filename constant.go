package i18n

import (
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	defaultFormatFile = "yaml"
)

var (
	defaultLanguage      = language.English
	defaultUnmarshalFunc = yaml.Unmarshal

	acceptLanguage = map[language.Tag]bool{
		defaultLanguage: true,
		language.German: true,
		language.French: true,
	}
)

// Key messages
var (
	CheckoutCannotProcessThisCard = KeyMessage{
		DefaultMessage: "Sorry, we are unable to process this card. Please try another, or refer to your bank for more details.",
		I18nKey:        "CheckoutCannotProcessThisCard",
	}

	CheckoutCannotProcessThisCardTryAnotherOneOrTryAgainLater = KeyMessage{
		DefaultMessage: "Sorry, we are unable to process this card. Please try another payment method, or try again later.",
		I18nKey:        "CheckoutCannotProcessThisCardTryAnotherOneOrTryAgainLater",
	}

	CheckoutCannotProcessThisCardAndTryAnotherOne = KeyMessage{
		DefaultMessage: "Sorry, we are unable to process payment with this card. Please try another.",
		I18nKey:        "CheckoutCannotProcessThisCardAndTryAnotherOne",
	}

	CheckoutCannotProcessThisCardAndContactBankOrTryAnotherOne = KeyMessage{
		DefaultMessage: "Sorry, we are unable to process this card. Please contact your bank or try another.",
		I18nKey:        "CheckoutCannotProcessThisCardAndContactBankOrTryAnotherOne",
	}

	CheckoutInsufficientFunds = KeyMessage{
		DefaultMessage: "Sorry, we are unable to process payment due to the spending limit on this payment method. Please try another method, or refer to your bank for more details.",
		I18nKey:        "CheckoutInsufficientFunds",
	}

	CheckoutInsufficientFundsViaPaypal = KeyMessage{
		DefaultMessage: "Sorry, we are unable to process payment due to the spending limit on this payment method. Please try another method, or follow up with Paypal and try again.",
		I18nKey:        "CheckoutInsufficientFundsViaPaypal",
	}

	CheckoutMakeSureCardCorrectly = KeyMessage{
		DefaultMessage: "Please make sure your card number was entered correctly, and try again.",
		I18nKey:        "CheckoutMakeSureCardCorrectly",
	}

	CheckoutMakeSureCardCVVCorrectly = KeyMessage{
		DefaultMessage: "Please make sure your CVC/CVV was entered correctly, and try again.",
		I18nKey:        "CheckoutMakeSureCardCVVCorrectly",
	}

	CheckoutAuthenticationRequired = KeyMessage{
		DefaultMessage: "This payment needs to be authenticated with your card provider. Please follow up with your bank, or try a different payment method.",
		I18nKey:        "CheckoutAuthenticationRequired",
	}

	CheckoutCardExpired = KeyMessage{
		DefaultMessage: "Your card seems to have expired. Please try another.",
		I18nKey:        "CheckoutCardExpired",
	}

	CheckoutHavingTroubleViaPaypal = KeyMessage{
		DefaultMessage: "We are having trouble processing your payment via Paypal. Please try another payment method, or contact Paypal.",
		I18nKey:        "CheckoutHavingTroubleViaPaypal",
	}

	CheckoutBillingAddressNotMatchCardHolderAndTryAgain = KeyMessage{
		DefaultMessage: "Please make sure your billing address matches your cardholder address and try again.",
		I18nKey:        "CheckoutBillingAddressNotMatchCardHolderAndTryAgain",
	}

	CheckoutBillingNameNotMatchCardHolder = KeyMessage{
		DefaultMessage: "Please make sure the name on your billing address matches the cardholder's.",
		I18nKey:        "CheckoutBillingNameNotMatchCardHolder",
	}

	CheckoutBillingAddressNotMatchCardHolderAddress = KeyMessage{
		DefaultMessage: "Please make sure that the billing address entered matches the cardholder address.",
		I18nKey:        "CheckoutBillingAddressNotMatchCardHolderAddress",
	}

	CheckoutBillingAddressAndZipCodeNotMatchCardHolder = KeyMessage{
		DefaultMessage: "Please make sure that both billing address and zip code match the cardholder's information.",
		I18nKey:        "CheckoutBillingAddressAndZipCodeNotMatchCardHolder",
	}

	CheckoutBillingZipCodeNotMatchCardHolder = KeyMessage{
		DefaultMessage: "Please make sure that the zip code entered matches the cardholder zip code.",
		I18nKey:        "CheckoutBillingZipCodeNotMatchCardHolder",
	}

	CheckoutBillingNameAndAddressNotMatchCardHolder = KeyMessage{
		DefaultMessage: "Please make sure that the address and name on your billing address matches the cardholder's.",
		I18nKey:        "CheckoutBillingNameAndAddressNotMatchCardHolder",
	}

	CheckoutAdyenAVSFailedDefault = KeyMessage{
		DefaultMessage: "Please make sure that your billing address matches the cardholder address.",
		I18nKey:        "CheckoutAdyenAVSFailedDefault",
	}

	CheckoutAdyenFraudMessage = KeyMessage{
		DefaultMessage: "Please enter a valid card and try again.",
		I18nKey:        "CheckoutAdyenFraudMessage",
	}
)

// error
var (
	ErrCheckoutCannotProcessThisCardAndContactBankOrTryAnotherOne = errHandler(CheckoutCannotProcessThisCardAndContactBankOrTryAnotherOne)
	ErrCheckoutAdyenAVSFailedDefault                              = errHandler(CheckoutAdyenAVSFailedDefault)
	ErrCheckoutAdyenFraud                                         = errHandler(CheckoutAdyenFraudMessage)
)
