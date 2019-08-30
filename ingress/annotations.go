package ingress

import (
	"strconv"

	"github.com/containous/traefik-migration/label"
)

const (
	annotationKubernetesProtocol            = "ingress.kubernetes.io/protocol"
	annotationKubernetesRuleType            = "ingress.kubernetes.io/rule-type"
	annotationKubernetesIngressClass        = "kubernetes.io/ingress.class"
	annotationKubernetesFrontendEntryPoints = "ingress.kubernetes.io/frontend-entry-points"
	annotationKubernetesPriority            = "ingress.kubernetes.io/priority"

	// CustomHeadersMiddleware
	annotationKubernetesSSLForceHost            = "ingress.kubernetes.io/ssl-force-host"
	annotationKubernetesSSLRedirect             = "ingress.kubernetes.io/ssl-redirect"
	annotationKubernetesHSTSMaxAge              = "ingress.kubernetes.io/hsts-max-age"
	annotationKubernetesHSTSIncludeSubdomains   = "ingress.kubernetes.io/hsts-include-subdomains"
	annotationKubernetesCustomRequestHeaders    = "ingress.kubernetes.io/custom-request-headers"
	annotationKubernetesCustomResponseHeaders   = "ingress.kubernetes.io/custom-response-headers"
	annotationKubernetesAllowedHosts            = "ingress.kubernetes.io/allowed-hosts"
	annotationKubernetesProxyHeaders            = "ingress.kubernetes.io/proxy-headers"
	annotationKubernetesSSLTemporaryRedirect    = "ingress.kubernetes.io/ssl-temporary-redirect"
	annotationKubernetesSSLHost                 = "ingress.kubernetes.io/ssl-host"
	annotationKubernetesSSLProxyHeaders         = "ingress.kubernetes.io/ssl-proxy-headers"
	annotationKubernetesHSTSPreload             = "ingress.kubernetes.io/hsts-preload"
	annotationKubernetesForceHSTSHeader         = "ingress.kubernetes.io/force-hsts"
	annotationKubernetesFrameDeny               = "ingress.kubernetes.io/frame-deny"
	annotationKubernetesCustomFrameOptionsValue = "ingress.kubernetes.io/custom-frame-options-value"
	annotationKubernetesContentTypeNosniff      = "ingress.kubernetes.io/content-type-nosniff"
	annotationKubernetesBrowserXSSFilter        = "ingress.kubernetes.io/browser-xss-filter"
	annotationKubernetesCustomBrowserXSSValue   = "ingress.kubernetes.io/custom-browser-xss-value"
	annotationKubernetesContentSecurityPolicy   = "ingress.kubernetes.io/content-security-policy"
	annotationKubernetesPublicKey               = "ingress.kubernetes.io/public-key"
	annotationKubernetesReferrerPolicy          = "ingress.kubernetes.io/referrer-policy"
	annotationKubernetesIsDevelopment           = "ingress.kubernetes.io/is-development"

	// TODO AuthMiddleware
	// annotationKubernetesAuthRealm                       = "ingress.kubernetes.io/auth-realm"
	// annotationKubernetesAuthType                        = "ingress.kubernetes.io/auth-type"
	// annotationKubernetesAuthSecret                      = "ingress.kubernetes.io/auth-secret"
	// annotationKubernetesAuthHeaderField                 = "ingress.kubernetes.io/auth-header-field"
	// annotationKubernetesAuthForwardResponseHeaders      = "ingress.kubernetes.io/auth-response-headers"
	// annotationKubernetesAuthRemoveHeader                = "ingress.kubernetes.io/auth-remove-header"
	// annotationKubernetesAuthForwardURL                  = "ingress.kubernetes.io/auth-url"
	// annotationKubernetesAuthForwardTrustHeaders         = "ingress.kubernetes.io/auth-trust-headers"
	// annotationKubernetesAuthForwardTLSSecret            = "ingress.kubernetes.io/auth-tls-secret"
	// annotationKubernetesAuthForwardTLSInsecure          = "ingress.kubernetes.io/auth-tls-insecure"

	// TODO ReplacePathRegexMiddleware
	// annotationKubernetesRewriteTarget                   = "ingress.kubernetes.io/rewrite-target"

	// TODO whitelistMiddleware
	// annotationKubernetesWhiteListSourceRange            = "ingress.kubernetes.io/whitelist-source-range"
	// annotationKubernetesWhiteListUseXForwardedFor       = "ingress.kubernetes.io/whitelist-x-forwarded-for"

	// TODO PassTLSCertMiddleware
	annotationKubernetesPassTLSCert       = "ingress.kubernetes.io/pass-tls-cert" // Deprecated
	annotationKubernetesPassTLSClientCert = "ingress.kubernetes.io/pass-client-tls-cert"

	// TODO CircuitBreakMiddleware
	// annotationKubernetesCircuitBreakerExpression        = "ingress.kubernetes.io/circuit-breaker-expression"

	// TODO RedirectMiddleware
	// annotationKubernetesRedirectEntryPoint              = "ingress.kubernetes.io/redirect-entry-point"
	// annotationKubernetesRedirectPermanent               = "ingress.kubernetes.io/redirect-permanent"
	// annotationKubernetesRedirectRegex                   = "ingress.kubernetes.io/redirect-regex"
	// annotationKubernetesRedirectReplacement             = "ingress.kubernetes.io/redirect-replacement"

	// TODO InFlightReqMiddleware
	// annotationKubernetesMaxConnAmount                   = "ingress.kubernetes.io/max-conn-amount"
	// annotationKubernetesMaxConnExtractorFunc            = "ingress.kubernetes.io/max-conn-extractor-func"

	// TODO RateLimitMiddleware
	// annotationKubernetesRateLimit                       = "ingress.kubernetes.io/rate-limit"

	// TODO ErrorPagesMiddleware
	// annotationKubernetesErrorPages                      = "ingress.kubernetes.io/error-pages"

	// TODO BufferingMiddleware
	// annotationKubernetesBuffering                       = "ingress.kubernetes.io/buffering"

	// TODO Modifiers Middlewares
	annotationKubernetesRequestModifier = "ingress.kubernetes.io/request-modifier"

	// FIXME Not possible yet
	// annotationKubernetesPreserveHost                    = "ingress.kubernetes.io/preserve-host"
	// annotationKubernetesSessionCookieName               = "ingress.kubernetes.io/session-cookie-name"
	// annotationKubernetesAffinity                        = "ingress.kubernetes.io/affinity"
	// annotationKubernetesResponseForwardingFlushInterval = "ingress.kubernetes.io/responseforwarding-flushinterval"

	// TODO ??
	// annotationKubernetesLoadBalancerMethod              = "ingress.kubernetes.io/load-balancer-method"
	// annotationKubernetesServiceWeights                  = "ingress.kubernetes.io/service-weights"
	// annotationKubernetesAppRoot                         = "ingress.kubernetes.io/app-root"

	// FIXME global backend
)

var compatibilityMapping = map[string]string{
	// annotationKubernetesPreserveHost:             "traefik.frontend.passHostHeader",
	annotationKubernetesPassTLSCert:         "traefik.frontend.passTLSCert",
	annotationKubernetesFrontendEntryPoints: "traefik.frontend.entryPoints",
	annotationKubernetesPriority:            "traefik.frontend.priority",
	// annotationKubernetesCircuitBreakerExpression: "traefik.backend.circuitbreaker",
	// annotationKubernetesLoadBalancerMethod:       "traefik.backend.loadbalancer.method",
	// annotationKubernetesAffinity:                 "traefik.backend.loadbalancer.stickiness",
	// annotationKubernetesSessionCookieName:        "traefik.backend.loadbalancer.stickiness.cookieName",
	annotationKubernetesRuleType: "traefik.frontend.rule.type",
	// annotationKubernetesRedirectEntryPoint:       "traefik.frontend.redirect.entrypoint",
	// annotationKubernetesRedirectRegex:            "traefik.frontend.redirect.regex",
	// annotationKubernetesRedirectReplacement:      "traefik.frontend.redirect.replacement",
}

func getAnnotationName(annotations map[string]string, name string) string {
	if _, ok := annotations[name]; ok {
		return name
	}

	if _, ok := annotations[label.Prefix+name]; ok {
		return label.Prefix + name
	}

	if lbl, compat := compatibilityMapping[name]; compat {
		if _, ok := annotations[lbl]; ok {
			return lbl
		}
	}

	return name
}

func getStringValue(annotations map[string]string, annotation string, defaultValue string) string {
	annotationName := getAnnotationName(annotations, annotation)
	return label.GetStringValue(annotations, annotationName, defaultValue)
}

func getStringSafeValue(annotations map[string]string, annotation string, defaultValue string) (string, error) {
	annotationName := getAnnotationName(annotations, annotation)
	value := label.GetStringValue(annotations, annotationName, defaultValue)
	_, err := strconv.Unquote(`"` + value + `"`)
	return value, err
}

func getBoolValue(annotations map[string]string, annotation string, defaultValue bool) bool {
	annotationName := getAnnotationName(annotations, annotation)
	return label.GetBoolValue(annotations, annotationName, defaultValue)
}

func getIntValue(annotations map[string]string, annotation string, defaultValue int) int {
	annotationName := getAnnotationName(annotations, annotation)
	return label.GetIntValue(annotations, annotationName, defaultValue)
}

func getInt64Value(annotations map[string]string, annotation string, defaultValue int64) int64 {
	annotationName := getAnnotationName(annotations, annotation)
	return label.GetInt64Value(annotations, annotationName, defaultValue)
}

func getSliceStringValue(annotations map[string]string, annotation string) []string {
	annotationName := getAnnotationName(annotations, annotation)
	return label.GetSliceStringValue(annotations, annotationName)
}

func getMapValue(annotations map[string]string, annotation string) map[string]string {
	annotationName := getAnnotationName(annotations, annotation)
	return label.GetMapValue(annotations, annotationName)
}
