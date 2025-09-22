package content_based_router

import "testing"

func TestNewContentBasedRouter(t *testing.T) {
	router := NewContentBasedRouter()

	if router.WidgetQueue == nil {
		t.Error("expected WidgetQueue to be initialized, got nil")
	}
	if router.GadgetQueue == nil {
		t.Error("expected GadgetQueue to be initialized, got nil")
	}
	if router.DunnoQueue == nil {
		t.Error("expected DunnoQueue to be initialized, got nil")
	}

}
