package repositories

import (
	"quotes-api/models"
	"testing"

	"github.com/google/uuid"
)

func TestSave_AssignsIDAndStoresQuote(t *testing.T) {
	repo := NewQuoteRepository()

	input := models.Quote{Text: "Stay hungry, stay foolish", Author: "Steve Jobs"}
	saved := repo.Save(input)

	if saved.ID == "" {
		t.Fatalf("expected saved quote to have a generated ID")
	}

	if _, err := uuid.Parse(saved.ID); err != nil {
		t.Fatalf("expected saved quote ID to be a valid UUID, got %q", saved.ID)
	}

	if saved.Text != input.Text || saved.Author != input.Author {
		t.Fatalf("expected saved quote fields to match input, got %+v", saved)
	}

	all := repo.FindAll()
	if len(all) != 1 {
		t.Fatalf("expected 1 quote after save, got %d", len(all))
	}

	if all[0] != saved {
		t.Fatalf("expected stored quote to match saved quote, got %+v", all[0])
	}
}

func TestSaveMany_SavesAllQuotesWithUniqueIDs(t *testing.T) {
	repo := NewQuoteRepository()

	input := []models.Quote{
		{Text: "Do or do not. There is no try.", Author: "Yoda"},
		{Text: "Simplicity is the ultimate sophistication.", Author: "Leonardo da Vinci"},
		{Text: "Knowledge is power.", Author: "Francis Bacon"},
	}

	saved := repo.SaveMany(input)

	if len(saved) != len(input) {
		t.Fatalf("expected %d saved quotes, got %d", len(input), len(saved))
	}

	seen := map[string]bool{}
	for i, quote := range saved {
		if quote.ID == "" {
			t.Fatalf("expected saved quote at index %d to have ID", i)
		}

		if _, err := uuid.Parse(quote.ID); err != nil {
			t.Fatalf("expected valid UUID for quote at index %d, got %q", i, quote.ID)
		}

		if seen[quote.ID] {
			t.Fatalf("expected unique IDs, duplicate found: %q", quote.ID)
		}
		seen[quote.ID] = true

		if quote.Text != input[i].Text || quote.Author != input[i].Author {
			t.Fatalf("expected saved quote fields to match input at index %d, got %+v", i, quote)
		}
	}

	all := repo.FindAll()
	if len(all) != len(input) {
		t.Fatalf("expected repository to contain %d quotes, got %d", len(input), len(all))
	}
}

func TestFindByID_ReturnsQuoteWhenFound(t *testing.T) {
	repo := NewQuoteRepository()
	saved := repo.Save(models.Quote{Text: "The unexamined life is not worth living.", Author: "Socrates"})

	got, found := repo.FindByID(saved.ID)
	if !found {
		t.Fatalf("expected quote to be found by ID")
	}

	if got != saved {
		t.Fatalf("expected found quote to match saved quote, got %+v", got)
	}
}

func TestFindByID_ReturnsFalseWhenMissing(t *testing.T) {
	repo := NewQuoteRepository()

	got, found := repo.FindByID("missing-id")
	if found {
		t.Fatalf("expected no quote to be found for missing ID")
	}

	if got != (models.Quote{}) {
		t.Fatalf("expected zero-value quote when not found, got %+v", got)
	}
}

func TestUpdate_UpdatesExistingQuoteAndPreservesID(t *testing.T) {
	repo := NewQuoteRepository()
	saved := repo.Save(models.Quote{Text: "Original", Author: "Original Author"})

	updatedInput := models.Quote{Text: "Updated", Author: "Updated Author"}
	updated, ok := repo.Update(saved.ID, updatedInput)
	if !ok {
		t.Fatalf("expected update to succeed")
	}

	if updated.ID != saved.ID {
		t.Fatalf("expected updated quote to preserve ID %q, got %q", saved.ID, updated.ID)
	}

	if updated.Text != updatedInput.Text || updated.Author != updatedInput.Author {
		t.Fatalf("expected updated fields to match input, got %+v", updated)
	}

	fromRepo, found := repo.FindByID(saved.ID)
	if !found {
		t.Fatalf("expected updated quote to still exist in repository")
	}

	if fromRepo != updated {
		t.Fatalf("expected repository to store updated quote, got %+v", fromRepo)
	}
}

func TestUpdate_ReturnsFalseWhenQuoteMissing(t *testing.T) {
	repo := NewQuoteRepository()

	updated, ok := repo.Update("missing-id", models.Quote{Text: "Updated", Author: "Someone"})
	if ok {
		t.Fatalf("expected update to fail for missing quote")
	}

	if updated != (models.Quote{}) {
		t.Fatalf("expected zero-value quote on failed update, got %+v", updated)
	}
}

func TestDelete_RemovesExistingQuote(t *testing.T) {
	repo := NewQuoteRepository()
	saved := repo.Save(models.Quote{Text: "Delete me", Author: "Unknown"})

	deleted := repo.Delete(saved.ID)
	if !deleted {
		t.Fatalf("expected delete to return true for existing quote")
	}

	if len(repo.FindAll()) != 0 {
		t.Fatalf("expected repository to be empty after delete")
	}

	if _, found := repo.FindByID(saved.ID); found {
		t.Fatalf("expected deleted quote to be absent")
	}
}

func TestDelete_ReturnsFalseWhenQuoteMissing(t *testing.T) {
	repo := NewQuoteRepository()

	deleted := repo.Delete("missing-id")
	if deleted {
		t.Fatalf("expected delete to return false for missing quote")
	}
}
