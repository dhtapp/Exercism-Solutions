package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var canAttack bool
	if knightIsAwake != canAttack {
		return false
	} else {
		return true
	}
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var isAwake = true
	if knightIsAwake == isAwake && archerIsAwake == isAwake && prisonerIsAwake == isAwake {
		return true
	} else if knightIsAwake == isAwake || archerIsAwake == isAwake || prisonerIsAwake == isAwake {
		return true
	} else {
		return false
	}
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var isAwake = true
	if prisonerIsAwake == isAwake && archerIsAwake == isAwake {
		return false
	} else if prisonerIsAwake == isAwake {
		return true
	} else {
		return false
	}
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var isAwake = true
	var hasDog = true
	var isAsleep = false
	if knightIsAwake == isAsleep && archerIsAwake == isAsleep && prisonerIsAwake == isAwake {
		return true
	} else if archerIsAwake == isAsleep && petDogIsPresent == hasDog {
		return true
	} else {
		return false
	}
	panic("Please implement the CanFreePrisoner() function")
}
