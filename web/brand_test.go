package web

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TecharoHQ/anubis"
	"github.com/TecharoHQ/anubis/lib/localization"
	"github.com/a-h/templ"
)

// renderBrandBase renders the base layout with the given corporate branding.
func renderBrandBase(t *testing.T, brandName, brandLink string) string {
	t.Helper()

	origName, origLink := anubis.BrandName, anubis.BrandLink
	defer func() { anubis.BrandName, anubis.BrandLink = origName, origLink }()
	anubis.BrandName, anubis.BrandLink = brandName, brandLink

	req := httptest.NewRequest("GET", "/", nil)
	localizer := &localization.SimpleLocalizer{}
	localizer.Localizer = localization.NewLocalizationService().GetLocalizerFromRequest(req)

	var buf strings.Builder
	component := base("Test title", templ.NopComponent, nil, nil, nil, nil, localizer)
	if err := component.Render(context.Background(), &buf); err != nil {
		t.Fatalf("failed to render template: %v", err)
	}
	return buf.String()
}

// The corporate fork must not leak upstream branding on challenge pages.
func TestCorporateFooter(t *testing.T) {
	out := renderBrandBase(t, "Medee", "")

	for _, banned := range []string{"Techaro", "CELPHASE", "celphase", "Anubis", "Made with"} {
		if strings.Contains(out, banned) {
			t.Errorf("footer leaks upstream branding %q", banned)
		}
	}
	if !strings.Contains(out, "Medee") {
		t.Errorf("brand name missing from footer")
	}
	if strings.Contains(out, `<a href="https://intranet.example.com">`) {
		t.Errorf("unexpected brand link")
	}
	if !strings.Contains(out, "css/brand.css") {
		t.Errorf("corporate stylesheet not linked")
	}
}

func TestCorporateFooterWithLink(t *testing.T) {
	out := renderBrandBase(t, "Medee", "https://intranet.example.com")

	if !strings.Contains(out, `<a href="https://intranet.example.com">Medee</a>`) {
		t.Errorf("brand link missing from footer")
	}
}
