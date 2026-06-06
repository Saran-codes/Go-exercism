package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	canfastAttack:= !knightIsAwake
    return canfastAttack
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	canSpy:= knightIsAwake || archerIsAwake || prisonerIsAwake
    return canSpy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	canSignalPrisoner:= prisonerIsAwake && (!archerIsAwake)
    return canSignalPrisoner
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	b1:= petDogIsPresent && (!archerIsAwake)
    b2:= prisonerIsAwake && (!knightIsAwake) && (!archerIsAwake)
    canFreePrisoner:= b1 || b2
    return canFreePrisoner
}
