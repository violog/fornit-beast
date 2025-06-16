package word_sensor

////////////////////////////////////////////////

// вызов из words_sensor.go
func afetrInitPhraseTree() {
	wordRecognizerInit() //
	// VerbalDetection("привет новая абзаца",1) // текст с пульта
	// PhraseSeparator("привет") // распознавание фразы
	// WordDetection("привет") // распознавание слова
	isReadyWordSensorLevel = 1
	initWordPult()
	initPrasePult()

	// проверка
	// GetFirstSymbolFromWordID(556)
}
