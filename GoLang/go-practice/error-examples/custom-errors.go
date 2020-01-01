package error_examples

import "log"

type badError struct {
	message           string
	additionalContext string
}

func (e *badError) Error() string {
	return e.message
}

type worseError struct {
	message           string
	warpDriveUnstable bool
}

func (e *worseError) Error() string {
	return e.message
}

type riskLevel int

const (
	riskLevelLow riskLevel = iota
	riskLevelMedium
	riskLevelHigh
)

func doRiskyManeuver(rl riskLevel) (bool, error) {
	switch rl {
	case riskLevelMedium:
		return false, &badError{
			"we're running on empyt.",
			"Not enough dillitam crystals.",
		}
	case riskLevelHigh:
		return false, &worseError{
			message:           "Plasma chambes melting.",
			warpDriveUnstable: true,
		}
	}

	return true, nil
}

func CustomErrors() {
	//success, err := doRiskyManeuver(riskLevelLow)
	success, err := doRiskyManeuver(riskLevelHigh)
	if err != nil {
		switch err.(type) {
		case *badError:
			log.Printf("Is's bad Captain: %s %s\n", err, err.(*badError).additionalContext)
		case *worseError:
			log.Printf("It's really bad Captain: %s\n", err)
			if err.(*worseError).warpDriveUnstable {
				log.Printf("The warp drive is unstable Captain! What do we do?")
			}
		}
	}
	if success {
		log.Println("We made it to port, Captain!")
	}
}
