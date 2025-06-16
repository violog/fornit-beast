<?php
class Stemmer_RU {
	private $vowel = "аеёиоуыэюя";
	private $regexPerfectiveGerunds = array(
		"(в|вши|вшись)$",
		"(ив|ивши|ившись|ыв|ывши|ывшись)$"
	);
	private $regexAdjective = "(ее|ие|ые|ое|ими|ыми|ей|ий|ый|ой|ем|им|ым|ом|его|ого|ему|ому|их|ых|ую|юю|ая|яя|ою|ею)$";
	private $regexParticiple = array(
		"(ем|нн|вш|ющ|щ)",
		"(ивш|ывш|ующ)"
	);
	private $regexReflexives = "(ся|сь)$";
	private $regexVerb = array(
		"(ла|на|ете|йте|ли|й|л|ем|н|ло|но|ет|ют|ны|ть|ешь|нно)$",
		"(ила|ыла|ена|ейте|уйте|ите|или|ыли|ей|уй|ил|ыл|им|ым|ен|ило|ыло|ено|ят|ует|уют|ит|ыт|ены|ить|ыть|ишь|ую|ю)$"
	);
	private $regexNoun = "(а|ев|ов|ие|ье|е|иями|ями|ами|еи|ии|и|ией|ей|ой|ий|й|иям|ям|ием|ем|ам|ом|о|у|ах|иях|ях|ы|ь|ию|ью|ю|ия|ья|я)$";
	private $regexSuperlative = "(ейш|ейше)$";
	private $regexDerivational = "(ост|ость)$";
	private $regexI = "и$";
	private $regexNN = "нн$";
	private $regexSoftSign = "ь$";
	private $word = '';
	private $RV = 0;
	private $R2 = 0;
	private $cache = array();

	public function getWordBase($word){
		if( isset($this->cache[$word]) ) return $this->cache[$word];
		$this->word = $word;
		$this->findRegions();
		if( !$this->removeEndings($this->regexPerfectiveGerunds, $this->RV) ){
			$this->removeEndings($this->regexReflexives, $this->RV);
			if( !($this->removeEndings(array($this->regexParticiple[0] . $this->regexAdjective, $this->regexParticiple[1] . $this->regexAdjective), $this->RV ) || $this->removeEndings($this->regexAdjective, $this->RV)) ){
				if (!$this->removeEndings($this->regexVerb, $this->RV)) {
					$this->removeEndings($this->regexNoun, $this->RV);
				}
			}
		}
		$this->removeEndings($this->regexI, $this->RV);
		$this->removeEndings($this->regexDerivational, $this->R2);
		if ($this->removeEndings($this->regexNN, $this->RV)) $this->word .= 'н';
		$this->removeEndings($this->regexSuperlative, $this->RV);
		$this->removeEndings($this->regexSoftSign, $this->RV);
		$this->cache[$word] = $this->word;
		return $this->word;
	}

	public function removeEndings($regex, $region){
		$prefix = mb_substr($this->word, 0, $region, 'utf8');
		$word   = substr($this->word,strlen($prefix));
		if (is_array($regex)) {
			if (preg_match('/.+[а|я]' . $regex[0] . '/u', $word)) {
				$this->word = $prefix . preg_replace('/' . $regex[0] . '/u', '', $word);
				return true;
			}
			$regex = $regex[1];
		}
		if (preg_match('/.+' . $regex . '/u', $word)) {
			$this->word = $prefix . preg_replace('/' . $regex . '/u', '', $word);
			return true;
		}
		return false;
	}

	private function findRegions(){
		$state = 0;
		$wordLength = mb_strlen($this->word, 'utf8');
		for ($i = 1; $i < $wordLength; $i++) {
			$prevChar = mb_substr($this->word, $i - 1, 1, 'utf8');
			$char	 = mb_substr($this->word, $i, 1, 'utf8');
			switch ($state) {
				case 0:
					if ($this->isVowel($char)) {
						$this->RV = $i + 1;
						$state	= 1;
					}
				break;
				case 1:
					if ($this->isVowel($prevChar) && !$this->isVowel($char)) {
						$state	= 2;
					}
				break;
				case 2:
					if ($this->isVowel($prevChar) && !$this->isVowel($char)) {
						$this->R2 = $i + 1;
						return;
					}
				break;
			}
		}
	}

	private function isVowel($char){
		return (strpos($this->vowel, $char) !== false);
	}
}