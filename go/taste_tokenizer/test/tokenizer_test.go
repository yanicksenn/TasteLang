package tokenizer

import (
	"reflect"
	"testing"

	"github.com/yanicksenn/TasteLang/go/shared"
	"github.com/yanicksenn/TasteLang/go/taste_tokenizer"
)

func TestTokenize(t *testing.T) {
	actualTokens := taste_tokenizer.Tokenize(`

# this is a comment

namespace Tokenizer.Test;
# this is another comment
type Token { # comment
	string value;
	# bool is_valid;
	int counter; # comment
    # comment line 1 
    # comment line 2 
}
# comment
	
	`)

	expectedTokens := []shared.Token {
		{ 
			Token: "namespace", 
			Len: 9,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 0, 
				Y: 4, 
			},
		},
		{
			Token: "Tokenizer", 
			Len: 9,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 10, 
				Y: 4, 
			},

		},
		{ 
			Token: ".", 
			Len: 1,
			IsValidIdentifier: false,
			StartLocation: shared.Location{ 
				X: 19, 
				Y: 4, 
			},

		},
		{ 
			Token: "Test", 
			Len: 4,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 20, 
				Y: 4, 
			},

		},
		{ 
			Token: ";", 
			Len: 1,
			IsValidIdentifier: false,
			StartLocation: shared.Location{ 
				X: 24, 
				Y: 4, 
			},

		},
		{ 
			Token: "type", 
			Len: 4,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 0, 
				Y: 6, 
			},
 
		},
		{ 
			Token: "Token", 
			Len: 5,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 5, 
				Y: 6, 
			},
 
		},
		{ 
			Token: "{", 
			Len: 1,
			IsValidIdentifier: false,
			StartLocation: shared.Location{ 
				X: 11, 
				Y: 6, 
			},
 
		},
		{ 
			Token: "string", 
			Len: 6,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 1, 
				Y: 7, 
			},

		},
		{ 
			Token: "value", 
			Len: 5,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 8, 
				Y: 7, 
			},

		},
		{ 
			Token: ";", 
			Len: 1,
			IsValidIdentifier: false,
			StartLocation: shared.Location{ 
				X: 13, 
				Y: 7, 
			},

		},
		{ 

			Token: "int", 
			Len: 3,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 1, 
				Y: 9, 
			},

		},
		{ 
			Token: "counter", 
			Len: 7,
			IsValidIdentifier: true,
			StartLocation: shared.Location{ 
				X: 5, 
				Y: 9, 
			},
 
		},
		{ 
			Token: ";", 
			Len: 1,
			IsValidIdentifier: false,
			StartLocation: shared.Location{ 
				X: 12, 
				Y: 9, 
			},

		},
		{ 
			Token: "}", 
			Len: 1,
			IsValidIdentifier: false,
			StartLocation: shared.Location{ 
				X: 0, 
				Y: 12, 
			},

		},

	}

	if len(actualTokens) != len(expectedTokens) {
		t.Fatalf("Expected: %d but got: %d\n", len(expectedTokens), len(actualTokens))	
	}

	if !reflect.DeepEqual(actualTokens, expectedTokens) {
		t.Fatalf("Expected: %+v but got: %+v\n", expectedTokens, actualTokens)	
	}
}
