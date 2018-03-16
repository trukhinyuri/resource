package validation

import (
	"regexp"

	"fmt"

	rstypes "git.containerum.net/ch/json-types/resource-service"
	kubtypes "git.containerum.net/ch/kube-client/pkg/model"
	"git.containerum.net/ch/resource-service/pkg/server"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

func StandardResourceValidator(uni *ut.UniversalTranslator) (ret *validator.Validate) {
	ret = validator.New()
	ret.SetTagName("binding")

	enTranslator, _ := uni.GetTranslator(en.New().Locale())
	enUSTranslator, _ := uni.GetTranslator(en_US.New().Locale())

	enTranslations.RegisterDefaultTranslations(ret, enTranslator)
	enTranslations.RegisterDefaultTranslations(ret, enUSTranslator)

	ret.RegisterValidation("dns", dnsValidationFunc)
	ret.RegisterValidation("docker_image", dockerImageValidationFunc)

	ret.RegisterStructValidation(createIngressRequestValidate, rstypes.CreateIngressRequest{})
	ret.RegisterStructValidation(serviceValidate, kubtypes.Service{})
	ret.RegisterStructValidation(updateServiceValidate, server.UpdateServiceRequest{})

	return
}

var (
	dnsLabel    = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	dockerImage = regexp.MustCompile(`(?:.+/)?([^:]+)(?::.+)?`)
)

func dnsValidationFunc(fl validator.FieldLevel) bool {
	return dnsLabel.MatchString(fl.Field().String())
}

func dockerImageValidationFunc(fl validator.FieldLevel) bool {
	return dockerImage.MatchString(fl.Field().String())
}

func createIngressRequestValidate(structLevel validator.StructLevel) {
	req := structLevel.Current().Interface().(rstypes.CreateIngressRequest)

	if req.Type == rstypes.IngressCustomHTTPS {
		if req.TLS == nil {
			structLevel.ReportError(req.TLS, "TLS", "tls", "exists", "")
		}
	}
}

func serviceValidate(structLevel validator.StructLevel) {
	req := structLevel.Current().Interface().(kubtypes.Service)

	v := structLevel.Validator()

	if err := v.Var(req.Name, "dns"); err != nil {
		structLevel.ReportValidationErrors("Name", "", err.(validator.ValidationErrors))
	}

	if err := v.Var(req.Deploy, "dns"); err != nil {
		structLevel.ReportValidationErrors("Deploy", "", err.(validator.ValidationErrors))
	}

	for i, port := range req.Ports {
		if err := v.Var(port.Protocol, "eq=TCP|eq=UDP"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("Ports[%d].Protocol", i), "", err.(validator.ValidationErrors))
		}

		if err := v.Var(port.Port, "omitempty,min=1,max=65535"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("Ports[%d].Port", i), "", err.(validator.ValidationErrors))
		}

		if err := v.Var(port.TargetPort, "min=1,max=65535"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("Ports[%d].TargetPort", i), "", err.(validator.ValidationErrors))
		}
	}

	for i, ip := range req.IPs {
		if err := v.Var(ip, "ip"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("IPs[%d]", i), "", err.(validator.ValidationErrors))
		}
	}
}

func updateServiceValidate(structLevel validator.StructLevel) {
	req := structLevel.Current().Interface().(server.UpdateServiceRequest)

	v := structLevel.Validator()

	for i, port := range req.Ports {
		if err := v.Var(port.Protocol, "eq=TCP|eq=UDP"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("Ports[%d].Protocol", i), "", err.(validator.ValidationErrors))
		}

		if err := v.Var(port.Port, "min=1,max=65535"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("Ports[%d].Port", i), "", err.(validator.ValidationErrors))
		}

		if err := v.Var(port.TargetPort, "omitempty,min=1,max=65535"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("Ports[%d].TargetPort", i), "", err.(validator.ValidationErrors))
		}

	}

	for i, ip := range req.IPs {
		if err := v.Var(ip, "ip"); err != nil {
			structLevel.ReportValidationErrors(fmt.Sprintf("IPs[%d]", i), "", err.(validator.ValidationErrors))
		}
	}
}
