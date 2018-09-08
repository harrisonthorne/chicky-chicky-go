package characters

// Chicken is the main character of this game. we ain't callin
// it chicky chicky for nothing folks
// 
// Chicken has the following fields implemented from
// the Character struct:
// 		Killable
//			IKillable
//				Hit(power int)
//				Kill() []Item
//			Lifespan			int
//			HealthLeft 			float32
// 		Position				Point
// 		Velocity 				Velocity
type Chicken struct {
	SimpleCharacter
}

// RunLeft runs the chicken to the left
func (c *Chicken) RunLeft() {
	
}

// RunRight runs the chicken to the right
func (c *Chicken) RunRight() {

}

// Jump jumps the chicken
func (c *Chicken) Jump() {

}

// Squat squats the chicken
func (c *Chicken) Squat() {

}

