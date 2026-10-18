package parser


import(
	"testing"
	"Checkpoint1/internal/domain/task"
)

func TestParseLog(t *testing.T){

	tests := []struct { 
		name string 
		input string 
		expectedResult task.Result
		expectedError error
		
	}{ 
		{
			name:  "успешный парсинг",
			input: "[ERROR] Database crashed",
			expectedResult: task.Result{
					Level:   "ERROR",
					Message: "Database crashed",
			},
			expectedError: nil, 
		},

		{
			name:  "пустая строка",
			input: "",
			expectedResult: task.Result{},
			expectedError: ErrEmptyLog, 
		},

		{
			name:  "неверный формат лога",
			input: "INFO system started",
			expectedResult: task.Result{},
			expectedError: ErrInvalidFormat, 
		},
			
	}

	for _, tc := range tests {
		t.Run(tc.name,func(t *testing.T){
			result, err := ParseLog(tc.input)

			if err != tc.expectedError {
				t.Errorf("для '%s' ожидали ошибку %v, но получили %v", tc.input, tc.expectedError, err)
			}

			if result != tc.expectedResult {
				t.Errorf("для '%s' ожидали результат %v, но получили %v", tc.input, tc.expectedResult, err)
			}





			
		})
	}
}