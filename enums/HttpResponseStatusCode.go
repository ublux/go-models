package enums

type HttpResponseStatusCode string

const (
	HttpResponseStatusCode_Status100Continue                      HttpResponseStatusCode = "Status100Continue"
	HttpResponseStatusCode_Status101SwitchingProtocols            HttpResponseStatusCode = "Status101SwitchingProtocols"
	HttpResponseStatusCode_Status102Processing                    HttpResponseStatusCode = "Status102Processing"
	HttpResponseStatusCode_Status200OK                            HttpResponseStatusCode = "Status200OK"
	HttpResponseStatusCode_Status201Created                       HttpResponseStatusCode = "Status201Created"
	HttpResponseStatusCode_Status202Accepted                      HttpResponseStatusCode = "Status202Accepted"
	HttpResponseStatusCode_Status203NonAuthoritative              HttpResponseStatusCode = "Status203NonAuthoritative"
	HttpResponseStatusCode_Status204NoContent                     HttpResponseStatusCode = "Status204NoContent"
	HttpResponseStatusCode_Status205ResetContent                  HttpResponseStatusCode = "Status205ResetContent"
	HttpResponseStatusCode_Status206PartialContent                HttpResponseStatusCode = "Status206PartialContent"
	HttpResponseStatusCode_Status207MultiStatus                   HttpResponseStatusCode = "Status207MultiStatus"
	HttpResponseStatusCode_Status208AlreadyReported               HttpResponseStatusCode = "Status208AlreadyReported"
	HttpResponseStatusCode_Status226IMUsed                        HttpResponseStatusCode = "Status226IMUsed"
	HttpResponseStatusCode_Status300MultipleChoices               HttpResponseStatusCode = "Status300MultipleChoices"
	HttpResponseStatusCode_Status301MovedPermanently              HttpResponseStatusCode = "Status301MovedPermanently"
	HttpResponseStatusCode_Status302Found                         HttpResponseStatusCode = "Status302Found"
	HttpResponseStatusCode_Status303SeeOther                      HttpResponseStatusCode = "Status303SeeOther"
	HttpResponseStatusCode_Status304NotModified                   HttpResponseStatusCode = "Status304NotModified"
	HttpResponseStatusCode_Status305UseProxy                      HttpResponseStatusCode = "Status305UseProxy"
	HttpResponseStatusCode_Status306SwitchProxy                   HttpResponseStatusCode = "Status306SwitchProxy"
	HttpResponseStatusCode_Status307TemporaryRedirect             HttpResponseStatusCode = "Status307TemporaryRedirect"
	HttpResponseStatusCode_Status308PermanentRedirect             HttpResponseStatusCode = "Status308PermanentRedirect"
	HttpResponseStatusCode_Status400BadRequest                    HttpResponseStatusCode = "Status400BadRequest"
	HttpResponseStatusCode_Status401Unauthorized                  HttpResponseStatusCode = "Status401Unauthorized"
	HttpResponseStatusCode_Status402PaymentRequired               HttpResponseStatusCode = "Status402PaymentRequired"
	HttpResponseStatusCode_Status403Forbidden                     HttpResponseStatusCode = "Status403Forbidden"
	HttpResponseStatusCode_Status404NotFound                      HttpResponseStatusCode = "Status404NotFound"
	HttpResponseStatusCode_Status405MethodNotAllowed              HttpResponseStatusCode = "Status405MethodNotAllowed"
	HttpResponseStatusCode_Status406NotAcceptable                 HttpResponseStatusCode = "Status406NotAcceptable"
	HttpResponseStatusCode_Status407ProxyAuthenticationRequired   HttpResponseStatusCode = "Status407ProxyAuthenticationRequired"
	HttpResponseStatusCode_Status408RequestTimeout                HttpResponseStatusCode = "Status408RequestTimeout"
	HttpResponseStatusCode_Status409Conflict                      HttpResponseStatusCode = "Status409Conflict"
	HttpResponseStatusCode_Status410Gone                          HttpResponseStatusCode = "Status410Gone"
	HttpResponseStatusCode_Status411LengthRequired                HttpResponseStatusCode = "Status411LengthRequired"
	HttpResponseStatusCode_Status412PreconditionFailed            HttpResponseStatusCode = "Status412PreconditionFailed"
	HttpResponseStatusCode_Status413PayloadTooLarge               HttpResponseStatusCode = "Status413PayloadTooLarge"
	HttpResponseStatusCode_Status414RequestUriTooLong             HttpResponseStatusCode = "Status414RequestUriTooLong"
	HttpResponseStatusCode_Status415UnsupportedMediaType          HttpResponseStatusCode = "Status415UnsupportedMediaType"
	HttpResponseStatusCode_Status416RequestedRangeNotSatisfiable  HttpResponseStatusCode = "Status416RequestedRangeNotSatisfiable"
	HttpResponseStatusCode_Status417ExpectationFailed             HttpResponseStatusCode = "Status417ExpectationFailed"
	HttpResponseStatusCode_Status418ImATeapot                     HttpResponseStatusCode = "Status418ImATeapot"
	HttpResponseStatusCode_Status419AuthenticationTimeout         HttpResponseStatusCode = "Status419AuthenticationTimeout"
	HttpResponseStatusCode_Status421MisdirectedRequest            HttpResponseStatusCode = "Status421MisdirectedRequest"
	HttpResponseStatusCode_Status422UnprocessableEntity           HttpResponseStatusCode = "Status422UnprocessableEntity"
	HttpResponseStatusCode_Status423Locked                        HttpResponseStatusCode = "Status423Locked"
	HttpResponseStatusCode_Status424FailedDependency              HttpResponseStatusCode = "Status424FailedDependency"
	HttpResponseStatusCode_Status426UpgradeRequired               HttpResponseStatusCode = "Status426UpgradeRequired"
	HttpResponseStatusCode_Status428PreconditionRequired          HttpResponseStatusCode = "Status428PreconditionRequired"
	HttpResponseStatusCode_Status429TooManyRequests               HttpResponseStatusCode = "Status429TooManyRequests"
	HttpResponseStatusCode_Status431RequestHeaderFieldsTooLarge   HttpResponseStatusCode = "Status431RequestHeaderFieldsTooLarge"
	HttpResponseStatusCode_Status451UnavailableForLegalReasons    HttpResponseStatusCode = "Status451UnavailableForLegalReasons"
	HttpResponseStatusCode_Status500InternalServerError           HttpResponseStatusCode = "Status500InternalServerError"
	HttpResponseStatusCode_Status501NotImplemented                HttpResponseStatusCode = "Status501NotImplemented"
	HttpResponseStatusCode_Status502BadGateway                    HttpResponseStatusCode = "Status502BadGateway"
	HttpResponseStatusCode_Status503ServiceUnavailable            HttpResponseStatusCode = "Status503ServiceUnavailable"
	HttpResponseStatusCode_Status504GatewayTimeout                HttpResponseStatusCode = "Status504GatewayTimeout"
	HttpResponseStatusCode_Status505HttpVersionNotsupported       HttpResponseStatusCode = "Status505HttpVersionNotsupported"
	HttpResponseStatusCode_Status506VariantAlsoNegotiates         HttpResponseStatusCode = "Status506VariantAlsoNegotiates"
	HttpResponseStatusCode_Status507InsufficientStorage           HttpResponseStatusCode = "Status507InsufficientStorage"
	HttpResponseStatusCode_Status508LoopDetected                  HttpResponseStatusCode = "Status508LoopDetected"
	HttpResponseStatusCode_Status510NotExtended                   HttpResponseStatusCode = "Status510NotExtended"
	HttpResponseStatusCode_Status511NetworkAuthenticationRequired HttpResponseStatusCode = "Status511NetworkAuthenticationRequired"
)
