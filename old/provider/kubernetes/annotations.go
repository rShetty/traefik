package kubernetes

import (
	"strconv"

	"github.com/containous/traefik/old/provider/label"
)

const (
	annotationKubernetesIngressClass                    = "kubernetes.io/ingress.class"
	annotationKubernetesAuthRealm                       = "ingress.kubernetes.io/auth-realm"
	annotationKubernetesAuthType                        = "ingress.kubernetes.io/auth-type"
	annotationKubernetesAuthSecret                      = "ingress.kubernetes.io/auth-secret"
	annotationKubernetesAuthHeaderField                 = "ingress.kubernetes.io/auth-header-field"
	annotationKubernetesAuthForwardResponseHeaders      = "ingress.kubernetes.io/auth-response-headers"
	annotationKubernetesAuthRemoveHeader                = "ingress.kubernetes.io/auth-remove-header"
	annotationKubernetesAuthForwardURL                  = "ingress.kubernetes.io/auth-url"
	annotationKubernetesAuthForwardTrustHeaders         = "ingress.kubernetes.io/auth-trust-headers"
	annotationKubernetesAuthForwardTLSSecret            = "ingress.kubernetes.io/auth-tls-secret"
	annotationKubernetesAuthForwardTLSInsecure          = "ingress.kubernetes.io/auth-tls-insecure"
	annotationKubernetesRewriteTarget                   = "ingress.kubernetes.io/rewrite-target"
	annotationKubernetesWhiteListSourceRange            = "ingress.kubernetes.io/whitelist-source-range"
	annotationKubernetesWhiteListIPStrategy             = "ingress.kubernetes.io/whitelist-ipstrategy"
	annotationKubernetesWhiteListIPStrategyDepth        = "ingress.kubernetes.io/whitelist-ipstrategy-depth"
	annotationKubernetesWhiteListIPStrategyExcludedIPs  = "ingress.kubernetes.io/whitelist-ipstrategy-excluded-ips"
	annotationKubernetesPreserveHost                    = "ingress.kubernetes.io/preserve-host"
	annotationKubernetesPassTLSCert                     = "ingress.kubernetes.io/pass-tls-cert" // Deprecated
	annotationKubernetesPassTLSClientCert               = "ingress.kubernetes.io/pass-client-tls-cert"
	annotationKubernetesFrontendEntryPoints             = "ingress.kubernetes.io/frontend-entry-points"
	annotationKubernetesPriority                        = "ingress.kubernetes.io/priority"
	annotationKubernetesCircuitBreakerExpression        = "ingress.kubernetes.io/circuit-breaker-expression"
	annotationKubernetesLoadBalancerMethod              = "ingress.kubernetes.io/load-balancer-method"
	annotationKubernetesAffinity                        = "ingress.kubernetes.io/affinity"
	annotationKubernetesSessionCookieName               = "ingress.kubernetes.io/session-cookie-name"
	annotationKubernetesRuleType                        = "ingress.kubernetes.io/rule-type"
	annotationKubernetesRedirectEntryPoint              = "ingress.kubernetes.io/redirect-entry-point"
	annotationKubernetesRedirectPermanent               = "ingress.kubernetes.io/redirect-permanent"
	annotationKubernetesRedirectRegex                   = "ingress.kubernetes.io/redirect-regex"
	annotationKubernetesRedirectReplacement             = "ingress.kubernetes.io/redirect-replacement"
	annotationKubernetesMaxConnAmount                   = "ingress.kubernetes.io/max-conn-amount"
	annotationKubernetesMaxConnExtractorFunc            = "ingress.kubernetes.io/max-conn-extractor-func"
	annotationKubernetesRateLimit                       = "ingress.kubernetes.io/rate-limit"
	annotationKubernetesErrorPages                      = "ingress.kubernetes.io/error-pages"
	annotationKubernetesBuffering                       = "ingress.kubernetes.io/buffering"
	annotationKubernetesResponseForwardingFlushInterval = "ingress.kubernetes.io/responseforwarding-flushinterval"
	annotationKubernetesAppRoot                         = "ingress.kubernetes.io/app-root"
	annotationKubernetesServiceWeights                  = "ingress.kubernetes.io/service-weights"
	annotationKubernetesRequestModifier                 = "ingress.kubernetes.io/request-modifier"

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
	annotationKubernetesProtocol                = "ingress.kubernetes.io/protocol"
)

func getAnnotationName(annotations map[string]string, name string) string {
	if _, ok := annotations[name]; ok {
		return name
	}

	if _, ok := annotations[label.Prefix+name]; ok {
		return label.Prefix + name
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
